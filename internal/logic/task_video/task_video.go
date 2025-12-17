package task_video

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

type sTaskVideo struct {
	videoRedsync *redsync.Redsync
}

func init() {
	service.RegisterTaskVideo(New())
}

func New() service.ITaskVideo {
	return &sTaskVideo{
		videoRedsync: redsync.New(goredis.NewPool(redis.UniversalClient)),
	}
}

// 视频任务详情
func (s *sTaskVideo) Detail(ctx context.Context, id string) (*model.TaskVideo, error) {

	taskVideo, err := dao.TaskVideo.FindById(ctx, id)
	if err != nil {
		logger.Error(ctx, err)
		return nil, err
	}

	detail := &model.TaskVideo{
		Id:                 taskVideo.Id,
		TraceId:            taskVideo.TraceId,
		UserId:             taskVideo.UserId,
		AppId:              taskVideo.AppId,
		Model:              taskVideo.Model,
		VideoId:            taskVideo.VideoId,
		Width:              taskVideo.Width,
		Height:             taskVideo.Height,
		Seconds:            taskVideo.Seconds,
		Prompt:             taskVideo.Prompt,
		Progress:           taskVideo.Progress,
		RemixedFromVideoId: taskVideo.RemixedFromVideoId,
		Status:             taskVideo.Status,
		CompletedAt:        util.FormatDateTime(taskVideo.CompletedAt),
		ExpiresAt:          util.FormatDateTime(taskVideo.ExpiresAt),
		Error:              taskVideo.Error,
		Creator:            util.Desensitize(taskVideo.Creator),
		CreatedAt:          util.FormatDateTime(taskVideo.CreatedAt),
		UpdatedAt:          util.FormatDateTime(taskVideo.UpdatedAt),
	}

	if config.Cfg.VideoTask.IsEnableStorage && taskVideo.VideoUrl != "" {
		detail.VideoUrl = config.Cfg.VideoTask.StorageBaseUrl + taskVideo.VideoUrl
	}

	if service.Session().IsAdminRole(ctx) {
		detail.FileName = taskVideo.FileName
		detail.FilePath = taskVideo.FilePath
	}

	return detail, nil
}

// 视频任务分页列表
func (s *sTaskVideo) Page(ctx context.Context, params model.TaskVideoPageReq) (*model.TaskVideoPageRes, error) {

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

	if params.VideoId != "" {
		filter["video_id"] = params.VideoId
	}

	if params.VideoUrl != "" {

		if gstr.HasPrefix(params.VideoUrl, "http") {
			if parse, err := url.Parse(params.VideoUrl); err == nil {
				params.VideoUrl = parse.Path
			}
		}

		filter["video_url"] = bson.M{
			"$regex": regexp.QuoteMeta(params.VideoUrl),
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

	results, err := dao.TaskVideo.FindByPage(ctx, paging, filter, &dao.FindOptions{SortFields: []string{"-created_at"}})
	if err != nil {
		logger.Error(ctx, err)
		return nil, err
	}

	items := make([]*model.TaskVideo, 0)
	for _, result := range results {

		audio := &model.TaskVideo{
			Id:        result.Id,
			TraceId:   result.TraceId,
			UserId:    result.UserId,
			AppId:     result.AppId,
			Model:     result.Model,
			VideoId:   result.VideoId,
			Width:     result.Width,
			Height:    result.Height,
			Seconds:   result.Seconds,
			Prompt:    result.Prompt,
			Status:    result.Status,
			CreatedAt: util.FormatDateTimeMonth(result.CreatedAt),
		}

		if config.Cfg.VideoTask.IsEnableStorage && result.VideoUrl != "" {
			audio.VideoUrl = config.Cfg.VideoTask.StorageBaseUrl + result.VideoUrl
		}

		items = append(items, audio)
	}

	return &model.TaskVideoPageRes{
		Items: items,
		Paging: &model.Paging{
			Page:     paging.Page,
			PageSize: paging.PageSize,
			Total:    paging.Total,
		},
	}, nil
}

// 视频任务详情复制字段值
func (s *sTaskVideo) CopyField(ctx context.Context, params model.TaskVideoCopyFieldReq) (string, error) {

	result, err := dao.TaskVideo.FindById(ctx, params.Id)
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

// 视频文件
func (s *sTaskVideo) Video(ctx context.Context, fileName string) (string, error) {

	taskVideo, err := dao.TaskVideo.FindOne(ctx, bson.M{"file_name": fileName})
	if err != nil {
		logger.Error(ctx, err)
		return "", errors.New("视频文件未找到")
	}

	if taskVideo == nil || taskVideo.FilePath == "" {
		return "", errors.New("视频文件未找到")
	}

	return taskVideo.FilePath, nil
}

// 定时任务
func (s *sTaskVideo) Task(ctx context.Context) {

	logger.Info(ctx, "sTaskVideo Task start")

	now := gtime.TimestampMilli()

	mutex := s.videoRedsync.NewMutex(consts.TASK_VIDEO_LOCK_KEY, redsync.WithExpiry(config.Cfg.VideoTask.LockMinutes*time.Minute))
	if err := mutex.LockContext(ctx); err != nil {
		logger.Info(ctx, "sTaskVideo Task", err)
		logger.Debugf(ctx, "sTaskVideo Task end time: %d", gtime.TimestampMilli()-now)
		return
	}
	logger.Debug(ctx, "sTaskVideo Task lock")

	defer func() {
		if ok, err := mutex.UnlockContext(ctx); !ok || err != nil {
			logger.Error(ctx, err)
		} else {
			logger.Debug(ctx, "sTaskVideo Task unlock")
		}
		logger.Debugf(ctx, "sTaskVideo Task end time: %d", gtime.TimestampMilli()-now)
	}()

	taskVideos, err := dao.TaskVideo.Find(ctx, bson.M{"status": bson.M{"$in": []string{"queued", "in_progress", "completed"}}}, &dao.FindOptions{SortFields: []string{"created_at"}})
	if err != nil {
		logger.Error(ctx, err)
		return
	}

	providerMap := make(map[string]*entity.Provider)
	for _, taskVideo := range taskVideos {

		if taskVideo.Status == "completed" {
			if taskVideo.ExpiresAt <= now/1000 {
				if err = dao.TaskVideo.UpdateById(ctx, taskVideo.Id, bson.M{"status": "expired"}); err != nil {
					logger.Error(ctx, err)
				}
			}
			continue
		}

		logVideo, err := dao.LogVideo.FindOne(ctx, bson.M{"trace_id": taskVideo.TraceId})
		if err != nil {
			logger.Error(ctx, err)
			continue
		}

		provider := providerMap[logVideo.ModelAgent.ProviderId]
		if provider == nil {
			provider, err = dao.Provider.FindById(ctx, logVideo.ModelAgent.ProviderId)
			if err != nil {
				logger.Error(ctx, err)
				continue
			}
			providerMap[logVideo.ModelAgent.ProviderId] = provider
		}

		adapter := sdk.NewAdapter(ctx, &options.AdapterOptions{
			Provider: provider.Code,
			Model:    logVideo.Model,
			Key:      logVideo.Key,
			BaseUrl:  logVideo.ModelAgent.BaseUrl,
			Path:     logVideo.ModelAgent.Path,
			Timeout:  config.Cfg.Base.ShortTimeout * time.Second,
			ProxyUrl: config.Cfg.Http.ProxyUrl,
		})

		retrieve, err := adapter.VideoRetrieve(ctx, smodel.VideoRetrieveRequest{VideoId: taskVideo.VideoId})
		if err != nil {
			logger.Error(ctx, err)

			if err = dao.TaskVideo.UpdateById(ctx, taskVideo.Id, bson.M{
				"status": "failed",
				"error":  err,
			}); err != nil {
				logger.Error(ctx, err)
			}

			continue
		}

		var (
			videoUrl string
			fileName string
			filePath string
		)

		if retrieve.Status == "completed" && config.Cfg.VideoTask.IsEnableStorage {

			adapter := sdk.NewAdapter(ctx, &options.AdapterOptions{
				Provider: provider.Code,
				Model:    logVideo.Model,
				Key:      logVideo.Key,
				BaseUrl:  logVideo.ModelAgent.BaseUrl,
				Path:     logVideo.ModelAgent.Path,
				Timeout:  config.Cfg.Base.ShortTimeout * time.Second,
				ProxyUrl: config.Cfg.Http.ProxyUrl,
			})

			if content, err := adapter.VideoContent(ctx, smodel.VideoContentRequest{VideoId: taskVideo.VideoId}); err != nil {
				logger.Error(ctx, err)
			} else {

				filePath = config.Cfg.VideoTask.StorageDir

				if filePath == "" {
					filePath = "./resource/public/video/"
				} else if !gstr.HasSuffix(filePath, "/") {
					filePath = filePath + "/"
				}

				fileName = taskVideo.VideoId + "_video.mp4"

				if err = gfile.PutBytes(filePath+fileName, content.Data); err != nil {
					logger.Error(ctx, err)
				} else {
					if gstr.HasPrefix(filePath, "./resource/public/") {
						videoUrl = "/public/" + gstr.TrimLeft(filePath, "./resource/public/") + fileName
					} else if config.Cfg.VideoTask.StorageBaseUrl == "" {
						videoUrl = "/open/video/" + fileName
					} else {
						videoUrl = fileName
					}
				}
			}
		}

		if err = dao.TaskVideo.UpdateById(ctx, taskVideo.Id, bson.M{
			"progress":              retrieve.Progress,
			"status":                retrieve.Status,
			"completed_at":          retrieve.CompletedAt,
			"expires_at":            retrieve.ExpiresAt,
			"video_url":             videoUrl,
			"file_name":             fileName,
			"file_path":             filePath + fileName,
			"remixed_from_video_id": retrieve.RemixedFromVideoId,
			"error":                 retrieve.Error,
		}); err != nil {
			logger.Error(ctx, err)
		}
	}

	if _, err := redis.Set(ctx, consts.TASK_VIDEO_END_TIME_KEY, gtime.TimestampMilli()); err != nil {
		logger.Error(ctx, err)
	}
}
