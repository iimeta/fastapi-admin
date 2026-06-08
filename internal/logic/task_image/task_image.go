package task_image

import (
	"context"
	"encoding/base64"
	"net/url"
	"regexp"
	"time"

	"github.com/go-redsync/redsync/v4"
	"github.com/go-redsync/redsync/v4/redis/goredis/v9"
	"github.com/gogf/gf/v2/encoding/gjson"
	"github.com/gogf/gf/v2/os/gfile"
	"github.com/gogf/gf/v2/os/gtime"
	"github.com/gogf/gf/v2/text/gstr"
	"github.com/iimeta/fastapi-admin/v2/internal/config"
	"github.com/iimeta/fastapi-admin/v2/internal/consts"
	"github.com/iimeta/fastapi-admin/v2/internal/dao"
	"github.com/iimeta/fastapi-admin/v2/internal/errors"
	"github.com/iimeta/fastapi-admin/v2/internal/logic/common"
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

type sTaskImage struct {
	imageRedsync *redsync.Redsync
}

func init() {
	service.RegisterTaskImage(New())
}

func New() service.ITaskImage {
	return &sTaskImage{
		imageRedsync: redsync.New(goredis.NewPool(redis.UniversalClient)),
	}
}

// 绘图任务详情
func (s *sTaskImage) Detail(ctx context.Context, id string) (*model.TaskImage, error) {

	taskImage, err := dao.TaskImage.FindById(ctx, id)
	if err != nil {
		logger.Error(ctx, err)
		return nil, err
	}

	detail := &model.TaskImage{
		Id:             taskImage.Id,
		TraceId:        taskImage.TraceId,
		UserId:         taskImage.UserId,
		AppId:          taskImage.AppId,
		Model:          taskImage.Model,
		ImageId:        taskImage.ImageId,
		Width:          taskImage.Width,
		Height:         taskImage.Height,
		N:              taskImage.N,
		Quality:        taskImage.Quality,
		Size:           taskImage.Size,
		ResponseFormat: taskImage.ResponseFormat,
		Prompt:         taskImage.Prompt,
		Progress:       taskImage.Progress,
		Status:         taskImage.Status,
		CompletedAt:    util.FormatDateTime(taskImage.CompletedAt),
		ExpiresAt:      util.FormatDateTime(taskImage.ExpiresAt),
		Error:          taskImage.Error,
		Creator:        util.Desensitize(taskImage.Creator),
		CreatedAt:      util.FormatDateTime(taskImage.CreatedAt),
		UpdatedAt:      util.FormatDateTime(taskImage.UpdatedAt),
	}

	if config.Cfg.ImageTask.IsEnableStorage && taskImage.ImageUrl != "" {

		if config.Cfg.ImageTask.StorageBaseUrl != "" {
			if gstr.HasSuffix(config.Cfg.ImageTask.StorageBaseUrl, "/") {
				taskImage.ImageUrl = gstr.TrimLeft(taskImage.ImageUrl, "/")
			} else if !gstr.HasPrefix(taskImage.ImageUrl, "/") {
				taskImage.ImageUrl = "/" + taskImage.ImageUrl
			}
		}

		detail.ImageUrl = config.Cfg.ImageTask.StorageBaseUrl + taskImage.ImageUrl
	}

	if service.Session().IsAdminRole(ctx) {
		detail.FileName = taskImage.FileName
		detail.FilePath = taskImage.FilePath
	}

	return detail, nil
}

// 绘图任务分页列表
func (s *sTaskImage) Page(ctx context.Context, params model.TaskImagePageReq) (*model.TaskImagePageRes, error) {

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

	if params.ImageId != "" {
		filter["image_id"] = params.ImageId
	}

	if params.ImageUrl != "" {

		if gstr.HasPrefix(params.ImageUrl, "http") {
			if parse, err := url.Parse(params.ImageUrl); err == nil {
				params.ImageUrl = parse.Path
			}
		}

		filter["image_url"] = bson.M{
			"$regex": regexp.QuoteMeta(params.ImageUrl),
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

	results, err := dao.TaskImage.FindByPage(ctx, paging, filter, &dao.FindOptions{SortFields: []string{"-created_at"}})
	if err != nil {
		logger.Error(ctx, err)
		return nil, err
	}

	items := make([]*model.TaskImage, 0)
	for _, result := range results {

		image := &model.TaskImage{
			Id:        result.Id,
			TraceId:   result.TraceId,
			UserId:    result.UserId,
			AppId:     result.AppId,
			Model:     result.Model,
			ImageId:   result.ImageId,
			Width:     result.Width,
			Height:    result.Height,
			N:         result.N,
			Quality:   result.Quality,
			Size:      result.Size,
			Prompt:    result.Prompt,
			Status:    result.Status,
			CreatedAt: util.FormatDateTimeMonth(result.CreatedAt),
		}

		if config.Cfg.ImageTask.IsEnableStorage && result.ImageUrl != "" {

			if config.Cfg.ImageTask.StorageBaseUrl != "" {
				if gstr.HasSuffix(config.Cfg.ImageTask.StorageBaseUrl, "/") {
					result.ImageUrl = gstr.TrimLeft(result.ImageUrl, "/")
				} else if !gstr.HasPrefix(result.ImageUrl, "/") {
					result.ImageUrl = "/" + result.ImageUrl
				}
			}

			image.ImageUrl = config.Cfg.ImageTask.StorageBaseUrl + result.ImageUrl
		}

		items = append(items, image)
	}

	return &model.TaskImagePageRes{
		Items: items,
		Paging: &model.Paging{
			Page:     paging.Page,
			PageSize: paging.PageSize,
			Total:    paging.Total,
		},
	}, nil
}

// 绘图任务详情复制字段值
func (s *sTaskImage) CopyField(ctx context.Context, params model.TaskImageCopyFieldReq) (string, error) {

	result, err := dao.TaskImage.FindById(ctx, params.Id)
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

// 绘图任务
func (s *sTaskImage) Task(ctx context.Context) {

	logger.Info(ctx, "sTaskImage Task start")

	now := gtime.TimestampMilli()

	mutex := s.imageRedsync.NewMutex(consts.TASK_IMAGE_LOCK_KEY, redsync.WithExpiry(config.Cfg.ImageTask.LockMinutes*time.Minute))
	if err := mutex.LockContext(ctx); err != nil {
		logger.Info(ctx, "sTaskImage Task", err)
		logger.Debugf(ctx, "sTaskImage Task end time: %d", gtime.TimestampMilli()-now)
		return
	}
	logger.Debug(ctx, "sTaskImage Task lock")

	defer func() {
		if ok, err := mutex.UnlockContext(ctx); !ok || err != nil {
			logger.Error(ctx, err)
		} else {
			logger.Debug(ctx, "sTaskImage Task unlock")
		}
		logger.Debugf(ctx, "sTaskImage Task end time: %d", gtime.TimestampMilli()-now)
	}()

	taskImages, err := dao.TaskImage.Find(ctx, bson.M{"status": bson.M{"$in": []string{"queued", "completed"}}}, &dao.FindOptions{SortFields: []string{"created_at"}})
	if err != nil {
		logger.Error(ctx, err)
		return
	}

	providerMap := make(map[string]*entity.Provider)
	for _, taskImage := range taskImages {

		if taskImage.Status == "completed" {
			if taskImage.ExpiresAt <= now/1000 {

				update := bson.M{"status": "expired"}

				if config.Cfg.ImageTask.StorageExpiredDelete && taskImage.FilePath != "" {
					update["image_url"] = ""
					update["file_name"] = ""
					update["file_path"] = ""
					if err := gfile.RemoveFile(taskImage.FilePath); err != nil {
						logger.Error(ctx, err)
					}
				}

				if err = dao.TaskImage.UpdateById(ctx, taskImage.Id, update); err != nil {
					logger.Error(ctx, err)
				}
			}
			continue
		}

		logImage, err := dao.LogImage.FindOne(ctx, bson.M{"trace_id": taskImage.TraceId, "status": bson.M{"$in": []int{1, 2}}})
		if err != nil {
			logger.Error(ctx, err)
			continue
		}

		provider := providerMap[logImage.ModelAgent.ProviderId]
		if provider == nil {
			provider, err = dao.Provider.FindById(ctx, logImage.ModelAgent.ProviderId)
			if err != nil {
				logger.Error(ctx, err)
				continue
			}
			providerMap[logImage.ModelAgent.ProviderId] = provider
		}

		if err = dao.TaskImage.UpdateById(ctx, taskImage.Id, bson.M{"status": "in_progress"}); err != nil {
			logger.Error(ctx, err)
			continue
		}

		adapter := sdk.NewAdapter(ctx, &options.AdapterOptions{
			Provider: provider.Code,
			Model:    logImage.Model,
			Key:      logImage.Key,
			BaseUrl:  logImage.ModelAgent.BaseUrl,
			Path:     logImage.ModelAgent.Path,
			Timeout:  config.Cfg.Base.LongTimeout * time.Second,
			ProxyUrl: config.Cfg.Http.ProxyUrl,
		})

		requestBytes, err := gjson.Encode(taskImage.RequestData)
		if err != nil {
			logger.Error(ctx, err)

			if err = dao.TaskImage.UpdateById(ctx, taskImage.Id, bson.M{
				"status": "failed",
				"error":  &smodel.ImageError{Code: "request_encode_error", Message: err.Error()},
			}); err != nil {
				logger.Error(ctx, err)
			}

			continue
		}

		response, err := adapter.ImageGenerations(ctx, requestBytes)
		if err != nil {
			logger.Error(ctx, err)

			if err = dao.TaskImage.UpdateById(ctx, taskImage.Id, bson.M{
				"status": "failed",
				"error":  &smodel.ImageError{Code: "generation_error", Message: err.Error()},
			}); err != nil {
				logger.Error(ctx, err)
			}

			continue
		}

		// 计算花费
		common.Billing(ctx, response.Usage, &logImage.Spend)

		// 记录花费
		if err = common.RecordSpend(ctx, logImage.UserId, logImage.AppId, logImage.Creator, logImage.Rid, logImage.Key, logImage.Spend); err != nil {
			logger.Error(ctx, err)
			continue
		}

		if err = dao.LogImage.UpdateById(ctx, logImage.Id, bson.M{"spend": logImage.Spend}); err != nil {
			logger.Error(ctx, err)
			continue
		}

		var (
			imageUrl string
			fileName string
			filePath string
		)

		completedAt := gtime.TimestampMilli() / 1000
		var expiresAt int64

		if config.Cfg.ImageTask.IsEnableStorage && len(response.Data) > 0 {

			filePath = config.Cfg.ImageTask.StorageDir

			if filePath == "" {
				filePath = "./resource/public/image/"
			} else if !gstr.HasSuffix(filePath, "/") {
				filePath = filePath + "/"
			}

			fileName = taskImage.ImageId + "_image_0.png"

			imageData := response.Data[0]
			var imageBytes []byte

			if len(imageData.B64Json) > 0 {
				if decoded, err := base64.StdEncoding.DecodeString(imageData.B64Json); err == nil {
					imageBytes = decoded
				} else {
					logger.Error(ctx, err)
				}
			}

			if imageBytes != nil {
				if err = gfile.PutBytes(filePath+fileName, imageBytes); err != nil {
					logger.Error(ctx, err)
				} else {

					if gstr.HasPrefix(filePath, "./resource/public/") {
						imageUrl = "/public/" + gstr.TrimLeft(filePath, "./resource/public/") + fileName
					} else if config.Cfg.ImageTask.StorageBaseUrl == "" {
						imageUrl = "/open/image/" + fileName
					} else {
						imageUrl = fileName
					}

					if config.Cfg.ImageTask.StorageExpiresAt > 0 {
						expiresAt = gtime.NewFromTimeStamp(completedAt).Add(config.Cfg.ImageTask.StorageExpiresAt * time.Minute).Unix()
					}
				}
			}
		}

		if err = dao.TaskImage.UpdateById(ctx, taskImage.Id, bson.M{
			"progress":      100,
			"status":        "completed",
			"completed_at":  completedAt,
			"expires_at":    expiresAt,
			"image_url":     imageUrl,
			"file_name":     fileName,
			"file_path":     filePath + fileName,
			"response_data": util.ConvToMap(response),
			"error":         nil,
		}); err != nil {
			logger.Error(ctx, err)
		}
	}

	if _, err := redis.Set(ctx, consts.TASK_IMAGE_END_TIME_KEY, gtime.TimestampMilli()); err != nil {
		logger.Error(ctx, err)
	}
}
