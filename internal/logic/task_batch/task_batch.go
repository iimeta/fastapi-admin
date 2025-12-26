package task_batch

import (
	"context"
	"net/url"
	"regexp"
	"time"

	"github.com/go-redsync/redsync/v4"
	"github.com/go-redsync/redsync/v4/redis/goredis/v9"
	"github.com/gogf/gf/v2/os/gtime"
	"github.com/gogf/gf/v2/text/gstr"
	"github.com/gogf/gf/v2/util/gconv"
	"github.com/iimeta/fastapi-admin/internal/config"
	"github.com/iimeta/fastapi-admin/internal/consts"
	"github.com/iimeta/fastapi-admin/internal/dao"
	"github.com/iimeta/fastapi-admin/internal/errors"
	"github.com/iimeta/fastapi-admin/internal/model"
	"github.com/iimeta/fastapi-admin/internal/model/do"
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

type sTaskBatch struct {
	batchRedsync *redsync.Redsync
}

func init() {
	service.RegisterTaskBatch(New())
}

func New() service.ITaskBatch {
	return &sTaskBatch{
		batchRedsync: redsync.New(goredis.NewPool(redis.UniversalClient)),
	}
}

// 批处理任务详情
func (s *sTaskBatch) Detail(ctx context.Context, id string) (*model.TaskBatch, error) {

	taskBatch, err := dao.TaskBatch.FindById(ctx, id)
	if err != nil {
		logger.Error(ctx, err)
		return nil, err
	}

	detail := &model.TaskBatch{
		Id:          taskBatch.Id,
		TraceId:     taskBatch.TraceId,
		UserId:      taskBatch.UserId,
		AppId:       taskBatch.AppId,
		Model:       taskBatch.Model,
		BatchId:     taskBatch.BatchId,
		Status:      taskBatch.Status,
		CompletedAt: util.FormatDateTime(taskBatch.CompletedAt),
		ExpiresAt:   util.FormatDateTime(taskBatch.ExpiresAt),
		Creator:     util.Desensitize(taskBatch.Creator),
		CreatedAt:   util.FormatDateTime(taskBatch.CreatedAt),
		UpdatedAt:   util.FormatDateTime(taskBatch.UpdatedAt),
	}

	return detail, nil
}

// 批处理任务分页列表
func (s *sTaskBatch) Page(ctx context.Context, params model.TaskBatchPageReq) (*model.TaskBatchPageRes, error) {

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

	if params.BatchId != "" {
		filter["batch_id"] = params.BatchId
	}

	if params.BatchUrl != "" {

		if gstr.HasPrefix(params.BatchUrl, "http") {
			if parse, err := url.Parse(params.BatchUrl); err == nil {
				params.BatchUrl = parse.Path
			}
		}

		filter["batch_url"] = bson.M{
			"$regex": regexp.QuoteMeta(params.BatchUrl),
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

	results, err := dao.TaskBatch.FindByPage(ctx, paging, filter, &dao.FindOptions{SortFields: []string{"-created_at"}})
	if err != nil {
		logger.Error(ctx, err)
		return nil, err
	}

	items := make([]*model.TaskBatch, 0)
	for _, result := range results {

		audio := &model.TaskBatch{
			Id:        result.Id,
			TraceId:   result.TraceId,
			UserId:    result.UserId,
			AppId:     result.AppId,
			Model:     result.Model,
			BatchId:   result.BatchId,
			Status:    result.Status,
			CreatedAt: util.FormatDateTimeMonth(result.CreatedAt),
		}

		items = append(items, audio)
	}

	return &model.TaskBatchPageRes{
		Items: items,
		Paging: &model.Paging{
			Page:     paging.Page,
			PageSize: paging.PageSize,
			Total:    paging.Total,
		},
	}, nil
}

// 批处理任务详情复制字段值
func (s *sTaskBatch) CopyField(ctx context.Context, params model.TaskBatchCopyFieldReq) (string, error) {

	result, err := dao.TaskBatch.FindById(ctx, params.Id)
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

// 批处理定时任务
func (s *sTaskBatch) Task(ctx context.Context) {

	logger.Info(ctx, "sTaskBatch Task start")

	now := gtime.TimestampMilli()

	mutex := s.batchRedsync.NewMutex(consts.TASK_BATCH_LOCK_KEY, redsync.WithExpiry(config.Cfg.BatchTask.LockMinutes*time.Minute))
	if err := mutex.LockContext(ctx); err != nil {
		logger.Info(ctx, "sTaskBatch Task", err)
		logger.Debugf(ctx, "sTaskBatch Task end time: %d", gtime.TimestampMilli()-now)
		return
	}
	logger.Debug(ctx, "sTaskBatch Task lock")

	defer func() {
		if ok, err := mutex.UnlockContext(ctx); !ok || err != nil {
			logger.Error(ctx, err)
		} else {
			logger.Debug(ctx, "sTaskBatch Task unlock")
		}
		logger.Debugf(ctx, "sTaskBatch Task end time: %d", gtime.TimestampMilli()-now)
	}()

	taskBatchs, err := dao.TaskBatch.Find(ctx, bson.M{"status": bson.M{"$in": []string{"validating", "in_progress", "finalizing", "completed", "cancelling"}}}, &dao.FindOptions{SortFields: []string{"created_at"}})
	if err != nil {
		logger.Error(ctx, err)
		return
	}

	providerMap := make(map[string]*entity.Provider)
	for _, taskBatch := range taskBatchs {

		if taskBatch.Status == "completed" {
			if taskBatch.ExpiresAt <= now/1000 {
				if err = dao.TaskBatch.UpdateById(ctx, taskBatch.Id, bson.M{"status": "expired"}); err != nil {
					logger.Error(ctx, err)
				}
			}
			continue
		}

		logBatch, err := dao.LogBatch.FindOne(ctx, bson.M{"trace_id": taskBatch.TraceId, "status": 1})
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

		adapter := sdk.NewAdapter(ctx, &options.AdapterOptions{
			Provider: provider.Code,
			Model:    logBatch.Model,
			Key:      logBatch.Key,
			BaseUrl:  logBatch.ModelAgent.BaseUrl,
			Path:     logBatch.ModelAgent.Path,
			Timeout:  config.Cfg.Base.ShortTimeout * time.Second,
			ProxyUrl: config.Cfg.Http.ProxyUrl,
		})

		retrieve, err := adapter.BatchRetrieve(ctx, smodel.BatchRetrieveRequest{BatchId: taskBatch.BatchId})
		if err != nil {
			logger.Error(ctx, err)

			if err = dao.TaskBatch.UpdateById(ctx, taskBatch.Id, bson.M{
				"status": "failed",
				"error":  err,
			}); err != nil {
				logger.Error(ctx, err)
			}

			continue
		}

		if retrieve.Status == "completed" {

			if retrieve.OutputFileId != "" {

				taskFile := do.TaskFile{
					TraceId: taskBatch.TraceId,
					UserId:  taskBatch.UserId,
					AppId:   taskBatch.AppId,
					Model:   taskBatch.Model,
					FileId:  retrieve.OutputFileId,
					Rid:     taskBatch.Rid,
					Creator: taskBatch.Creator,
				}

				adapter := sdk.NewAdapter(ctx, &options.AdapterOptions{
					Provider: provider.Code,
					Model:    logBatch.Model,
					Key:      logBatch.Key,
					BaseUrl:  logBatch.ModelAgent.BaseUrl,
					Path:     logBatch.ModelAgent.Path,
					Timeout:  config.Cfg.Base.ShortTimeout * time.Second,
					ProxyUrl: config.Cfg.Http.ProxyUrl,
				})

				retrieve, err := adapter.FileRetrieve(ctx, smodel.FileRetrieveRequest{FileId: retrieve.OutputFileId})
				if err != nil {
					logger.Error(ctx, err)

					if id, err := dao.TaskFile.Insert(ctx, taskFile); err != nil {
						logger.Error(ctx, err)
					} else {
						if err = dao.TaskFile.UpdateById(ctx, id, bson.M{
							"status": "failed",
							"error":  err,
						}); err != nil {
							logger.Error(ctx, err)
						}
					}

					continue
				}

				taskFile.Purpose = retrieve.Purpose
				taskFile.FileName = retrieve.Filename
				taskFile.Bytes = retrieve.Bytes
				taskFile.ExpiresAt = retrieve.ExpiresAt
				taskFile.Status = retrieve.Status

				if _, err := dao.TaskFile.Insert(ctx, taskFile); err != nil {
					logger.Error(ctx, err)
				}
			}

			if retrieve.ErrorFileId != "" {

				taskFile := do.TaskFile{
					TraceId: taskBatch.TraceId,
					UserId:  taskBatch.UserId,
					AppId:   taskBatch.AppId,
					Model:   taskBatch.Model,
					FileId:  retrieve.ErrorFileId,
					Rid:     taskBatch.Rid,
					Creator: taskBatch.Creator,
				}

				adapter := sdk.NewAdapter(ctx, &options.AdapterOptions{
					Provider: provider.Code,
					Model:    logBatch.Model,
					Key:      logBatch.Key,
					BaseUrl:  logBatch.ModelAgent.BaseUrl,
					Path:     logBatch.ModelAgent.Path,
					Timeout:  config.Cfg.Base.ShortTimeout * time.Second,
					ProxyUrl: config.Cfg.Http.ProxyUrl,
				})

				retrieve, err := adapter.FileRetrieve(ctx, smodel.FileRetrieveRequest{FileId: retrieve.ErrorFileId})
				if err != nil {
					logger.Error(ctx, err)

					if id, err := dao.TaskFile.Insert(ctx, taskFile); err != nil {
						logger.Error(ctx, err)
					} else {
						if err = dao.TaskFile.UpdateById(ctx, id, bson.M{
							"status": "failed",
							"error":  err,
						}); err != nil {
							logger.Error(ctx, err)
						}
					}

					continue
				}

				taskFile.Purpose = retrieve.Purpose
				taskFile.FileName = retrieve.Filename
				taskFile.Bytes = retrieve.Bytes
				taskFile.ExpiresAt = retrieve.ExpiresAt
				taskFile.Status = retrieve.Status

				if _, err := dao.TaskFile.Insert(ctx, taskFile); err != nil {
					logger.Error(ctx, err)
				}
			}
		}

		if err = dao.TaskBatch.UpdateById(ctx, taskBatch.Id, bson.M{
			"output_file_id": retrieve.OutputFileId,
			"error_file_id":  retrieve.ErrorFileId,
			"status":         retrieve.Status,
			"in_progress_at": retrieve.InProgressAt,
			"finalizing_at":  retrieve.FinalizingAt,
			"completed_at":   retrieve.CompletedAt,
			"expires_at":     retrieve.ExpiresAt,
			"cancelling_at":  retrieve.CancellingAt,
			"cancelled_at":   retrieve.CancelledAt,
			"failed_at":      retrieve.FailedAt,
			"response_data":  gconv.Map(retrieve.ResponseBytes),
		}); err != nil {
			logger.Error(ctx, err)
		}
	}

	if _, err := redis.Set(ctx, consts.TASK_BATCH_END_TIME_KEY, gtime.TimestampMilli()); err != nil {
		logger.Error(ctx, err)
	}
}
