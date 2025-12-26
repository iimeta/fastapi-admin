package task_file

import (
	"context"
	"net/url"
	"regexp"
	"time"

	"github.com/go-redsync/redsync/v4"
	"github.com/go-redsync/redsync/v4/redis/goredis/v9"
	"github.com/gogf/gf/v2/os/gfile"
	"github.com/gogf/gf/v2/os/gtime"
	"github.com/gogf/gf/v2/text/gstr"
	"github.com/iimeta/fastapi-admin/internal/config"
	"github.com/iimeta/fastapi-admin/internal/consts"
	"github.com/iimeta/fastapi-admin/internal/dao"
	"github.com/iimeta/fastapi-admin/internal/errors"
	"github.com/iimeta/fastapi-admin/internal/model"
	"github.com/iimeta/fastapi-admin/internal/model/entity"
	"github.com/iimeta/fastapi-admin/internal/service"
	"github.com/iimeta/fastapi-admin/utility/db"
	"github.com/iimeta/fastapi-admin/utility/logger"
	"github.com/iimeta/fastapi-admin/utility/redis"
	"github.com/iimeta/fastapi-admin/utility/util"
	sdk "github.com/iimeta/fastapi-sdk"
	smodel "github.com/iimeta/fastapi-sdk/model"
	"github.com/iimeta/fastapi-sdk/options"
	"go.mongodb.org/mongo-driver/bson"
)

type sTaskFile struct {
	fileRedsync *redsync.Redsync
}

func init() {
	service.RegisterTaskFile(New())
}

func New() service.ITaskFile {
	return &sTaskFile{
		fileRedsync: redsync.New(goredis.NewPool(redis.UniversalClient)),
	}
}

// 文件任务详情
func (s *sTaskFile) Detail(ctx context.Context, id string) (*model.TaskFile, error) {

	taskFile, err := dao.TaskFile.FindById(ctx, id)
	if err != nil {
		logger.Error(ctx, err)
		return nil, err
	}

	detail := &model.TaskFile{
		Id:        taskFile.Id,
		TraceId:   taskFile.TraceId,
		UserId:    taskFile.UserId,
		AppId:     taskFile.AppId,
		Model:     taskFile.Model,
		FileId:    taskFile.FileId,
		Status:    taskFile.Status,
		ExpiresAt: util.FormatDateTime(taskFile.ExpiresAt),
		Error:     taskFile.Error,
		Creator:   util.Desensitize(taskFile.Creator),
		CreatedAt: util.FormatDateTime(taskFile.CreatedAt),
		UpdatedAt: util.FormatDateTime(taskFile.UpdatedAt),
	}

	if config.Cfg.FileTask.IsEnableStorage && taskFile.FileUrl != "" {

		if config.Cfg.FileTask.StorageBaseUrl != "" {
			if gstr.HasSuffix(config.Cfg.FileTask.StorageBaseUrl, "/") {
				taskFile.FileUrl = gstr.TrimLeft(taskFile.FileUrl, "/")
			} else if !gstr.HasPrefix(taskFile.FileUrl, "/") {
				taskFile.FileUrl = "/" + taskFile.FileUrl
			}
		}

		detail.FileUrl = config.Cfg.FileTask.StorageBaseUrl + taskFile.FileUrl
	}

	if service.Session().IsAdminRole(ctx) {
		detail.FileName = taskFile.FileName
		detail.FilePath = taskFile.FilePath
	}

	return detail, nil
}

// 文件任务分页列表
func (s *sTaskFile) Page(ctx context.Context, params model.TaskFilePageReq) (*model.TaskFilePageRes, error) {

	paging := &db.Paging{
		Page:     params.Page,
		PageSize: params.PageSize,
	}

	filter := bson.M{}

	if params.TraceId != "" {
		filter["trace_id"] = gstr.Trim(params.TraceId)
	}

	if service.Session().IsResellerRole(ctx) {
		filter["rid"] = service.Session().GetRid(ctx)
	}

	if service.Session().IsUserRole(ctx) {
		filter["user_id"] = service.Session().GetUserId(ctx)
	} else if params.UserId != 0 {
		filter["user_id"] = params.UserId
	}

	if params.AppId != 0 {
		filter["app_id"] = params.AppId
	}

	if params.FileId != "" {
		filter["file_id"] = params.FileId
	}

	if params.FileUrl != "" {

		if gstr.HasPrefix(params.FileUrl, "http") {
			if parse, err := url.Parse(params.FileUrl); err == nil {
				params.FileUrl = parse.Path
			}
		}

		filter["file_url"] = bson.M{
			"$regex": regexp.QuoteMeta(params.FileUrl),
		}
	}

	if params.Status != "" {
		filter["status"] = params.Status
	} else if !service.Session().IsAdminRole(ctx) {
		filter["status"] = bson.M{"$ne": "deleted"}
	}

	if len(params.CreatedAt) > 0 {
		gte := gtime.NewFromStrFormat(params.CreatedAt[0], time.DateTime).TimestampMilli()
		lte := gtime.NewFromStrLayout(params.CreatedAt[1], time.DateTime).TimestampMilli() + 999
		filter["created_at"] = bson.M{
			"$gte": gte,
			"$lte": lte,
		}
	}

	results, err := dao.TaskFile.FindByPage(ctx, paging, filter, &dao.FindOptions{SortFields: []string{"-created_at"}})
	if err != nil {
		logger.Error(ctx, err)
		return nil, err
	}

	items := make([]*model.TaskFile, 0)
	for _, result := range results {

		audio := &model.TaskFile{
			Id:        result.Id,
			TraceId:   result.TraceId,
			UserId:    result.UserId,
			AppId:     result.AppId,
			Model:     result.Model,
			FileId:    result.FileId,
			Status:    result.Status,
			CreatedAt: util.FormatDateTimeMonth(result.CreatedAt),
		}

		if config.Cfg.FileTask.IsEnableStorage && result.FileUrl != "" {

			if config.Cfg.FileTask.StorageBaseUrl != "" {
				if gstr.HasSuffix(config.Cfg.FileTask.StorageBaseUrl, "/") {
					result.FileUrl = gstr.TrimLeft(result.FileUrl, "/")
				} else if !gstr.HasPrefix(result.FileUrl, "/") {
					result.FileUrl = "/" + result.FileUrl
				}
			}

			audio.FileUrl = config.Cfg.FileTask.StorageBaseUrl + result.FileUrl
		}

		items = append(items, audio)
	}

	return &model.TaskFilePageRes{
		Items: items,
		Paging: &model.Paging{
			Page:     paging.Page,
			PageSize: paging.PageSize,
			Total:    paging.Total,
		},
	}, nil
}

// 文件任务详情复制字段值
func (s *sTaskFile) CopyField(ctx context.Context, params model.TaskFileCopyFieldReq) (string, error) {

	result, err := dao.TaskFile.FindById(ctx, params.Id)
	if err != nil {
		logger.Error(ctx, err)
		return "", err
	}

	if service.Session().IsResellerRole(ctx) && result.Rid != service.Session().GetRid(ctx) {
		return "", errors.ERR_UNAUTHORIZED
	}

	if service.Session().IsUserRole(ctx) && result.UserId != service.Session().GetUserId(ctx) {
		return "", errors.ERR_UNAUTHORIZED
	}

	switch params.Field {
	case "creator":
		return result.Creator, nil
	}

	return "", nil
}

// 文件文件
func (s *sTaskFile) File(ctx context.Context, fileName string) (string, error) {

	taskFile, err := dao.TaskFile.FindOne(ctx, bson.M{"file_name": fileName})
	if err != nil {
		logger.Error(ctx, err)
		return "", errors.New("文件文件未找到")
	}

	if taskFile == nil || taskFile.FilePath == "" {
		return "", errors.New("文件文件未找到")
	}

	return taskFile.FilePath, nil
}

// 文件定时任务
func (s *sTaskFile) Task(ctx context.Context) {

	logger.Info(ctx, "sTaskFile Task start")

	now := gtime.TimestampMilli()

	mutex := s.fileRedsync.NewMutex(consts.TASK_FILE_LOCK_KEY, redsync.WithExpiry(config.Cfg.FileTask.LockMinutes*time.Minute))
	if err := mutex.LockContext(ctx); err != nil {
		logger.Info(ctx, "sTaskFile Task", err)
		logger.Debugf(ctx, "sTaskFile Task end time: %d", gtime.TimestampMilli()-now)
		return
	}
	logger.Debug(ctx, "sTaskFile Task lock")

	defer func() {
		if ok, err := mutex.UnlockContext(ctx); !ok || err != nil {
			logger.Error(ctx, err)
		} else {
			logger.Debug(ctx, "sTaskFile Task unlock")
		}
		logger.Debugf(ctx, "sTaskFile Task end time: %d", gtime.TimestampMilli()-now)
	}()

	taskFiles, err := dao.TaskFile.Find(ctx, bson.M{"status": bson.M{"$in": []string{"queued", "in_progress", "completed"}}}, &dao.FindOptions{SortFields: []string{"created_at"}})
	if err != nil {
		logger.Error(ctx, err)
		return
	}

	providerMap := make(map[string]*entity.Provider)
	for _, taskFile := range taskFiles {

		if taskFile.Status == "completed" {
			if taskFile.ExpiresAt <= now/1000 {

				update := bson.M{"status": "expired"}

				if config.Cfg.FileTask.StorageExpiredDelete && taskFile.FilePath != "" {
					update["file_url"] = ""
					update["file_name"] = ""
					update["file_path"] = ""
					if err := gfile.RemoveFile(taskFile.FilePath); err != nil {
						logger.Error(ctx, err)
					}
				}

				if err = dao.TaskFile.UpdateById(ctx, taskFile.Id, update); err != nil {
					logger.Error(ctx, err)
				}
			}
			continue
		}

		logFile, err := dao.LogFile.FindOne(ctx, bson.M{"trace_id": taskFile.TraceId, "status": 1})
		if err != nil {
			logger.Error(ctx, err)
			continue
		}

		provider := providerMap[logFile.ModelAgent.ProviderId]
		if provider == nil {
			provider, err = dao.Provider.FindById(ctx, logFile.ModelAgent.ProviderId)
			if err != nil {
				logger.Error(ctx, err)
				continue
			}
			providerMap[logFile.ModelAgent.ProviderId] = provider
		}

		adapter := sdk.NewAdapter(ctx, &options.AdapterOptions{
			Provider: provider.Code,
			Model:    logFile.Model,
			Key:      logFile.Key,
			BaseUrl:  logFile.ModelAgent.BaseUrl,
			Path:     logFile.ModelAgent.Path,
			Timeout:  config.Cfg.Base.ShortTimeout * time.Second,
			ProxyUrl: config.Cfg.Http.ProxyUrl,
		})

		retrieve, err := adapter.FileRetrieve(ctx, smodel.FileRetrieveRequest{FileId: taskFile.FileId})
		if err != nil {
			logger.Error(ctx, err)

			if err = dao.TaskFile.UpdateById(ctx, taskFile.Id, bson.M{
				"status": "failed",
				"error":  err,
			}); err != nil {
				logger.Error(ctx, err)
			}

			continue
		}

		var (
			fileUrl  string
			fileName string
			filePath string
		)

		if retrieve.Status == "completed" && config.Cfg.FileTask.IsEnableStorage {

			adapter := sdk.NewAdapter(ctx, &options.AdapterOptions{
				Provider: provider.Code,
				Model:    logFile.Model,
				Key:      logFile.Key,
				BaseUrl:  logFile.ModelAgent.BaseUrl,
				Path:     logFile.ModelAgent.Path,
				Timeout:  config.Cfg.Base.ShortTimeout * time.Second,
				ProxyUrl: config.Cfg.Http.ProxyUrl,
			})

			if content, err := adapter.FileContent(ctx, smodel.FileContentRequest{FileId: taskFile.FileId}); err != nil {
				logger.Error(ctx, err)
			} else {

				filePath = config.Cfg.FileTask.StorageDir

				if filePath == "" {
					filePath = "./resource/public/file/"
				} else if !gstr.HasSuffix(filePath, "/") {
					filePath = filePath + "/"
				}

				fileName = retrieve.Filename

				if err = gfile.PutBytes(filePath+fileName, content.Data); err != nil {
					logger.Error(ctx, err)
				} else {

					if gstr.HasPrefix(filePath, "./resource/public/") {
						fileUrl = "/public/" + gstr.TrimLeft(filePath, "./resource/public/") + fileName
					} else if config.Cfg.FileTask.StorageBaseUrl == "" {
						fileUrl = "/open/file/" + fileName
					} else {
						fileUrl = fileName
					}

					if config.Cfg.FileTask.StorageExpiresAt > 0 {
						retrieve.ExpiresAt = gtime.NewFromTimeStamp(now / 1000).Add(config.Cfg.FileTask.StorageExpiresAt * time.Minute).Unix()
					}
				}
			}
		}

		if err = dao.TaskFile.UpdateById(ctx, taskFile.Id, bson.M{
			"status":     retrieve.Status,
			"expires_at": retrieve.ExpiresAt,
			"file_url":   fileUrl,
			"file_name":  fileName,
			"file_path":  filePath + fileName,
		}); err != nil {
			logger.Error(ctx, err)
		}
	}

	if _, err := redis.Set(ctx, consts.TASK_FILE_END_TIME_KEY, gtime.TimestampMilli()); err != nil {
		logger.Error(ctx, err)
	}
}
