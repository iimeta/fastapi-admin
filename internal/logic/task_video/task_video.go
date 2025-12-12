package task_video

import (
	"context"
	"time"

	"github.com/go-redsync/redsync/v4"
	"github.com/go-redsync/redsync/v4/redis/goredis/v9"
	"github.com/gogf/gf/v2/os/gfile"
	"github.com/gogf/gf/v2/os/gtime"
	"github.com/gogf/gf/v2/text/gstr"
	"github.com/iimeta/fastapi-admin/internal/config"
	"github.com/iimeta/fastapi-admin/internal/consts"
	"github.com/iimeta/fastapi-admin/internal/dao"
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
		filter["video_url"] = params.VideoUrl
	}

	if params.Status != "" {
		filter["status"] = params.Status
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
			VideoUrl:  result.VideoUrl,
			Status:    result.Status,
			CreatedAt: util.FormatDateTimeMonth(result.CreatedAt),
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

	taskVideos, err := dao.TaskVideo.Find(ctx, bson.M{"status": bson.M{"$in": []string{"queued", "in_progress"}}}, &dao.FindOptions{SortFields: []string{"created_at"}})
	if err != nil {
		logger.Error(ctx, err)
		return
	}

	providerMap := make(map[string]*entity.Provider)
	for _, taskVideo := range taskVideos {

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
			path     string
		)

		if retrieve.Status == "completed" {

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

				path = config.Cfg.VideoTask.StorageDir
				if path == "" {
					path = "./resource/public/video/"
				} else if !gstr.HasSuffix(path, "/") {
					path = path + "/"
				}

				videoUrl = taskVideo.VideoId + "_video.mp4"
				path += videoUrl

				if err = gfile.PutBytes(path, content.Data); err != nil {
					logger.Error(ctx, err)
				}
			}
		}

		if err = dao.TaskVideo.UpdateById(ctx, taskVideo.Id, bson.M{
			"progress":     retrieve.Progress,
			"status":       retrieve.Status,
			"completed_at": retrieve.CompletedAt,
			"expires_at":   retrieve.ExpiresAt,
			"video_url":    videoUrl,
			"path":         path,
			"error":        retrieve.Error,
		}); err != nil {
			logger.Error(ctx, err)
		}
	}

	if _, err := redis.Set(ctx, consts.TASK_VIDEO_END_TIME_KEY, gtime.TimestampMilli()); err != nil {
		logger.Error(ctx, err)
	}
}
