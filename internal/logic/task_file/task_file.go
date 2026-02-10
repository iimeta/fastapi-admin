package task_file

import (
	"context"
	"regexp"
	"time"

	"github.com/go-redsync/redsync/v4"
	"github.com/go-redsync/redsync/v4/redis/goredis/v9"
	"github.com/gogf/gf/v2/os/gfile"
	"github.com/gogf/gf/v2/os/gtime"
	"github.com/gogf/gf/v2/text/gstr"
	"github.com/iimeta/fastapi-admin/v2/internal/config"
	"github.com/iimeta/fastapi-admin/v2/internal/consts"
	"github.com/iimeta/fastapi-admin/v2/internal/dao"
	"github.com/iimeta/fastapi-admin/v2/internal/errors"
	"github.com/iimeta/fastapi-admin/v2/internal/model"
	"github.com/iimeta/fastapi-admin/v2/internal/model/entity"
	"github.com/iimeta/fastapi-admin/v2/internal/service"
	"github.com/iimeta/fastapi-admin/v2/utility/db"
	"github.com/iimeta/fastapi-admin/v2/utility/logger"
	"github.com/iimeta/fastapi-admin/v2/utility/redis"
	"github.com/iimeta/fastapi-admin/v2/utility/util"
	sdk "github.com/iimeta/fastapi-sdk/v2"
	smodel "github.com/iimeta/fastapi-sdk/v2/model"
	"github.com/iimeta/fastapi-sdk/v2/options"
	"go.mongodb.org/mongo-driver/v2/bson"
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
		Id:           taskFile.Id,
		TraceId:      taskFile.TraceId,
		UserId:       taskFile.UserId,
		AppId:        taskFile.AppId,
		Model:        taskFile.Model,
		Purpose:      taskFile.Purpose,
		FileId:       taskFile.FileId,
		FileName:     taskFile.FileName,
		Bytes:        taskFile.Bytes,
		Status:       taskFile.Status,
		ExpiresAt:    util.FormatDateTime(taskFile.ExpiresAt),
		ResponseData: taskFile.ResponseData,
		Error:        taskFile.Error,
		BatchTraceId: taskFile.BatchTraceId,
		Creator:      util.Desensitize(taskFile.Creator),
		CreatedAt:    util.FormatDateTime(taskFile.CreatedAt),
		UpdatedAt:    util.FormatDateTime(taskFile.UpdatedAt),
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
		filter["$or"] = bson.A{
			bson.M{"trace_id": gstr.Trim(params.TraceId)},
			bson.M{"batch_trace_id": gstr.Trim(params.TraceId)},
		}
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

	if params.FileName != "" {
		filter["file_name"] = bson.M{
			"$regex": regexp.QuoteMeta(params.FileName),
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

		file := &model.TaskFile{
			Id:        result.Id,
			TraceId:   result.TraceId,
			UserId:    result.UserId,
			AppId:     result.AppId,
			Model:     result.Model,
			Purpose:   result.Purpose,
			FileId:    result.FileId,
			FileName:  result.FileName,
			Bytes:     result.Bytes,
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

			file.FileUrl = config.Cfg.FileTask.StorageBaseUrl + result.FileUrl
		}

		items = append(items, file)
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

// 文件
func (s *sTaskFile) File(ctx context.Context, fileName string) (string, error) {

	taskFile, err := dao.TaskFile.FindOne(ctx, bson.M{"file_id": gfile.Name(fileName)})
	if err != nil {
		logger.Error(ctx, err)
		return "", errors.New("文件未找到")
	}

	if taskFile == nil || taskFile.FilePath == "" {
		return "", errors.New("文件未找到")
	}

	return taskFile.FilePath, nil
}

// 文件任务
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

	taskFiles, err := dao.TaskFile.Find(ctx, bson.M{"status": bson.M{"$in": []string{"uploaded", "processing", "processed"}}}, &dao.FindOptions{SortFields: []string{"created_at"}})
	if err != nil {
		logger.Error(ctx, err)
		return
	}

	providerMap := make(map[string]*entity.Provider)
	for _, taskFile := range taskFiles {

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

			continue
		}

		if taskFile.Status == "processing" {

			var adapter sdk.AdapterGroup

			if taskFile.Purpose != "batch_output" {

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

				adapter = sdk.NewAdapter(ctx, &options.AdapterOptions{
					Provider: provider.Code,
					Model:    logFile.Model,
					Key:      logFile.Key,
					BaseUrl:  logFile.ModelAgent.BaseUrl,
					Path:     logFile.ModelAgent.Path,
					Timeout:  config.Cfg.Base.ShortTimeout * time.Second,
					ProxyUrl: config.Cfg.Http.ProxyUrl,
				})

			} else {

				logBatch, err := dao.LogBatch.FindOne(ctx, bson.M{"trace_id": taskFile.TraceId, "status": 1})
				if err != nil {
					logger.Error(ctx, err)
					continue
				}

				provider := providerMap[logBatch.ModelAgent.ProviderId]
				if provider == nil {
					provider, err = dao.Provider.FindById(ctx, logBatch.ModelAgent.ProviderId)
					if err != nil {
						logger.Error(ctx, err)
						continue
					}
					providerMap[logBatch.ModelAgent.ProviderId] = provider
				}

				adapter = sdk.NewAdapter(ctx, &options.AdapterOptions{
					Provider: provider.Code,
					Model:    logBatch.Model,
					Key:      logBatch.Key,
					BaseUrl:  logBatch.ModelAgent.BaseUrl,
					Path:     logBatch.ModelAgent.Path,
					Timeout:  config.Cfg.Base.ShortTimeout * time.Second,
					ProxyUrl: config.Cfg.Http.ProxyUrl,
				})
			}

			if retrieve, err := adapter.FileRetrieve(ctx, smodel.FileRetrieveRequest{FileId: taskFile.FileId}); err != nil {
				logger.Error(ctx, err)
			} else {
				if err = dao.TaskFile.UpdateById(ctx, taskFile.Id, bson.M{
					"status":        retrieve.Status,
					"expires_at":    retrieve.ExpiresAt,
					"response_data": util.ConvToMap(retrieve.ResponseBytes),
				}); err != nil {
					logger.Error(ctx, err)
				}
			}

		} else if config.Cfg.FileTask.IsEnableStorage && taskFile.FileUrl == "" && taskFile.FilePath == "" {

			var (
				fileUrl  string
				fileName string
				filePath string
				adapter  sdk.AdapterGroup
			)

			if taskFile.Purpose != "batch_output" {

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

				adapter = sdk.NewAdapter(ctx, &options.AdapterOptions{
					Provider: provider.Code,
					Model:    logFile.Model,
					Key:      logFile.Key,
					BaseUrl:  logFile.ModelAgent.BaseUrl,
					Path:     logFile.ModelAgent.Path,
					Timeout:  config.Cfg.Base.ShortTimeout * time.Second,
					ProxyUrl: config.Cfg.Http.ProxyUrl,
				})

			} else {

				logBatch, err := dao.LogBatch.FindOne(ctx, bson.M{"trace_id": taskFile.TraceId, "status": 1})
				if err != nil {
					logger.Error(ctx, err)
					continue
				}

				provider := providerMap[logBatch.ModelAgent.ProviderId]
				if provider == nil {
					provider, err = dao.Provider.FindById(ctx, logBatch.ModelAgent.ProviderId)
					if err != nil {
						logger.Error(ctx, err)
						continue
					}
					providerMap[logBatch.ModelAgent.ProviderId] = provider
				}

				adapter = sdk.NewAdapter(ctx, &options.AdapterOptions{
					Provider: provider.Code,
					Model:    logBatch.Model,
					Key:      logBatch.Key,
					BaseUrl:  logBatch.ModelAgent.BaseUrl,
					Path:     logBatch.ModelAgent.Path,
					Timeout:  config.Cfg.Base.ShortTimeout * time.Second,
					ProxyUrl: config.Cfg.Http.ProxyUrl,
				})
			}

			if content, err := adapter.FileContent(ctx, smodel.FileContentRequest{FileId: taskFile.FileId}); err != nil {
				logger.Error(ctx, err)
			} else {

				filePath = config.Cfg.FileTask.StorageDir

				if filePath == "" {
					filePath = "./resource/public/file/"
				} else if !gstr.HasSuffix(filePath, "/") {
					filePath = filePath + "/"
				}

				fileName = taskFile.FileId + gfile.Ext(taskFile.FileName)

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
						taskFile.ExpiresAt = gtime.NewFromTimeStamp(now / 1000).Add(config.Cfg.FileTask.StorageExpiresAt * time.Minute).Unix()
					}
				}
			}

			if err = dao.TaskFile.UpdateById(ctx, taskFile.Id, bson.M{
				"expires_at": taskFile.ExpiresAt,
				"file_url":   fileUrl,
				"file_path":  filePath + fileName,
			}); err != nil {
				logger.Error(ctx, err)
			}
		}
	}

	if _, err := redis.Set(ctx, consts.TASK_FILE_END_TIME_KEY, gtime.TimestampMilli()); err != nil {
		logger.Error(ctx, err)
	}
}
