package task_image

import (
	"bytes"
	"context"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"net/textproto"
	"net/url"
	"path"
	"regexp"
	"time"

	"github.com/go-redsync/redsync/v4"
	"github.com/go-redsync/redsync/v4/redis/goredis/v9"
	"github.com/gogf/gf/v2/encoding/gjson"
	"github.com/gogf/gf/v2/net/gtrace"
	"github.com/gogf/gf/v2/os/gctx"
	"github.com/gogf/gf/v2/os/gfile"
	"github.com/gogf/gf/v2/os/grpool"
	"github.com/gogf/gf/v2/os/gtime"
	"github.com/gogf/gf/v2/text/gstr"
	"github.com/gogf/gf/v2/util/gconv"
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
	"go.mongodb.org/mongo-driver/v2/mongo"
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
		Action:         taskImage.Action,
		ImageId:        taskImage.ImageId,
		Width:          taskImage.Width,
		Height:         taskImage.Height,
		N:              taskImage.N,
		Quality:        taskImage.Quality,
		Size:           taskImage.Size,
		OutputFormat:   taskImage.OutputFormat,
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
				taskImage.ImageUrl = gstr.TrimLeftStr(taskImage.ImageUrl, "/")
			} else if !gstr.HasPrefix(taskImage.ImageUrl, "/") {
				taskImage.ImageUrl = "/" + taskImage.ImageUrl
			}
		}

		detail.ImageUrl = config.Cfg.ImageTask.StorageBaseUrl + taskImage.ImageUrl
	}

	if service.Session().IsAdminRole(ctx) {
		detail.JobId = taskImage.JobId
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
			Action:    result.Action,
			ImageId:   result.ImageId,
			Width:     result.Width,
			Height:    result.Height,
			N:         result.N,
			Quality:   result.Quality,
			Size:      result.Size,
			Prompt:    result.Prompt,
			Progress:  result.Progress,
			Status:    result.Status,
			CreatedAt: util.FormatDateTimeMonth(result.CreatedAt),
		}

		if config.Cfg.ImageTask.IsEnableStorage && result.ImageUrl != "" {

			if config.Cfg.ImageTask.StorageBaseUrl != "" {
				if gstr.HasSuffix(config.Cfg.ImageTask.StorageBaseUrl, "/") {
					result.ImageUrl = gstr.TrimLeftStr(result.ImageUrl, "/")
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

// 绘图任务重新生成
func (s *sTaskImage) Regenerate(ctx context.Context, id string) error {

	if _, err := dao.TaskImage.FindOneAndUpdate(ctx, bson.M{
		"_id": id,
		"status": bson.M{
			"$in": []string{"in_progress", "failed"},
		},
	}, bson.M{
		"status":   "queued",
		"progress": 0,
		"error":    nil,
	}); err != nil {

		if errors.Is(err, mongo.ErrNoDocuments) {
			return errors.New("任务不在进行中或已失败状态, 无法重新生成")
		}

		logger.Error(ctx, err)
		return err
	}

	return nil
}

// 绘图任务批量操作
func (s *sTaskImage) BatchOperate(ctx context.Context, params model.TaskImageBatchOperateReq) error {

	switch params.Action {
	case consts.ACTION_REGENERATE:
		for _, id := range params.Ids {
			// 跳过不可重新生成的任务(如已完成、已过期等), 不中断整个批次
			if err := s.Regenerate(ctx, id); err != nil {
				logger.Error(ctx, err)
			}
		}
	}

	return nil
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

	// 计算僵死判定阈值: 服务重启或worker协程异常退出会导致任务一直停留在in_progress, 既不重试也不过期
	// 这里不单独写库, 而是把in_progress一并查出, 在内存里判断: updated_at超过最大处理时长仍未变化的视为僵死, 重新提升处理; 否则视为正在运行
	reclaimMillis := (config.Cfg.ImageTask.Reclaim * time.Second).Milliseconds()
	if reclaimMillis <= 0 {
		// 自动按单次超时 ×(重试次数+1)推算最大处理时长
		timeout := config.Cfg.ImageTask.Timeout
		if timeout <= 0 {
			timeout = config.Cfg.Base.LongTimeout
		}
		retryCount := config.Cfg.ImageTask.RetryCount
		if retryCount < 0 {
			retryCount = 0
		}
		reclaimMillis = (timeout * time.Duration(retryCount+1) * time.Second).Milliseconds()
	}

	// updated_at早于staleBefore的in_progress任务视为僵死; reclaimMillis<=0时不回收(staleBefore为0, 不会有任务命中)
	var staleBefore int64
	if reclaimMillis > 0 {
		staleBefore = now - reclaimMillis
	}

	taskImages, err := dao.TaskImage.Find(ctx, bson.M{"status": bson.M{"$in": []string{"queued", "in_progress", "completed"}}}, &dao.FindOptions{SortFields: []string{"created_at"}})
	if err != nil {
		logger.Error(ctx, err)
		return
	}

	// 进行中数量限制: 0为不限制, 大于0则限制同时进行中的任务数量
	// availableSlots为本轮还可(重新)提升为进行中的任务数量, 小于0表示不限制
	availableSlots := -1
	if config.Cfg.ImageTask.ConcurrencyLimit > 0 {
		// 统计真正在运行(未僵死)的in_progress数量, 僵死的不计入(本轮会被重新提升)
		liveInProgress := 0
		for _, taskImage := range taskImages {
			if taskImage.Status == "in_progress" && taskImage.UpdatedAt >= staleBefore {
				liveInProgress++
			}
		}
		if availableSlots = config.Cfg.ImageTask.ConcurrencyLimit - liveInProgress; availableSlots < 0 {
			availableSlots = 0
		}
	}

	var queuedTasks []*entity.TaskImage

	for _, taskImage := range taskImages {

		if taskImage.Status == "completed" {

			// expires_at为0表示永不过期(未开存储、永不过期配置或本地存储失败), 仅在已设置过期时间且已到期时才回收
			if taskImage.ExpiresAt > 0 && taskImage.ExpiresAt <= now/1000 {

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

		// 正在运行(未僵死)的in_progress任务跳过; 僵死的往下走重新提升处理
		if taskImage.Status == "in_progress" && taskImage.UpdatedAt >= staleBefore {
			// 巡检顺带推进伪进度(画图无真实进度), 由Task每轮按已耗时驱动, 重启后仍能续推, 无需常驻协程
			s.advanceFakeProgress(ctx, now, taskImage)
			continue
		}

		// 已达到进行中数量上限, 本轮不再(重新)提升任务为进行中
		if availableSlots == 0 {
			continue
		}

		// 提升为进行中时不重置progress: queued任务创建时progress已为0, 超时续轮询/回收的任务保留已有进度避免进度条倒退, 仅Regenerate(全新生图)才显式归零
		if err = dao.TaskImage.UpdateById(ctx, taskImage.Id, bson.M{"status": "in_progress", "error": nil}); err != nil {
			logger.Error(ctx, err)
			continue
		}

		queuedTasks = append(queuedTasks, taskImage)

		if availableSlots > 0 {
			availableSlots--
		}
	}

	for _, taskImage := range queuedTasks {
		if err := grpool.AddWithRecover(gctx.NeverDone(ctx), func(ctx context.Context) {

			ctx, err = gtrace.WithTraceID(ctx, taskImage.TraceId)

			s.processImageTask(ctx, taskImage)

		}, nil); err != nil {
			logger.Error(ctx, err)
		}
	}

	if _, err := redis.Set(ctx, consts.TASK_IMAGE_END_TIME_KEY, gtime.TimestampMilli()); err != nil {
		logger.Error(ctx, err)
	}
}

func (s *sTaskImage) processImageTask(ctx context.Context, taskImage *entity.TaskImage) {

	logImage, err := dao.LogImage.FindOne(ctx, bson.M{"trace_id": taskImage.TraceId, "status": bson.M{"$in": []int{1, 2}}})
	if err != nil {
		logger.Error(ctx, err)
		s.failTask(ctx, taskImage.Id, "log_not_found", err.Error())
		return
	}

	provider, err := dao.Provider.FindById(ctx, logImage.ModelAgent.ProviderId)
	if err != nil {
		logger.Error(ctx, err)
		s.failTask(ctx, taskImage.Id, "provider_not_found", err.Error())
		return
	}

	timeout := config.Cfg.ImageTask.Timeout * time.Second
	if timeout <= 0 {
		timeout = config.Cfg.Base.LongTimeout * time.Second
	}

	retryCount := config.Cfg.ImageTask.RetryCount
	if retryCount < 0 {
		retryCount = 0
	}

	var (
		response smodel.ImageResponse
		errCode  string
	)

	for attempt := 0; ; attempt++ {

		taskCtx, cancel := context.WithTimeout(ctx, timeout)

		if config.Cfg.ImageTask.SubmitMode == 2 {
			response, errCode, err = s.requestImageAsync(taskCtx, taskImage, logImage, provider, timeout)
		} else {
			response, errCode, err = s.requestImageSync(taskCtx, taskImage, logImage, provider, timeout)
		}

		cancel()

		if err == nil {
			break
		}

		// 仅异步: 轮询时钟到点, 上游任务可能仍在进行, 置回queued并保留job_id, 交由下一轮cron凭job_id续轮询, 避免重新提交导致上游重复出图
		if config.Cfg.ImageTask.SubmitMode == 2 && errCode == "timeout" {
			logger.Infof(ctx, "sTaskImage processImageTask task: %s async poll timeout, requeue to resume next round", taskImage.Id)
			s.requeueTask(ctx, taskImage.Id)
			return
		}

		if attempt < retryCount {

			// 重试前回查任务状态, 若已不在进行中(已被其它进程完成、被管理员重置或删除等), 则无需重试, 直接退出
			if latest, e := dao.TaskImage.FindById(ctx, taskImage.Id); e != nil {
				logger.Error(ctx, e)
			} else if latest.Status != "in_progress" {
				logger.Infof(ctx, "sTaskImage processImageTask task: %s status is %s, no need to retry, skip", taskImage.Id, latest.Status)
				return
			}

			logger.Errorf(ctx, "sTaskImage processImageTask task: %s failed: %s, retry: %d/%d", taskImage.Id, errCode, attempt+1, retryCount)
			continue
		}

		logger.Error(ctx, err)
		s.failTask(ctx, taskImage.Id, errCode, err.Error())
		return
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

		outputFormat := taskImage.OutputFormat
		if outputFormat == "" {
			outputFormat = "png"
		}

		fileName = taskImage.ImageId + "_image." + outputFormat

		imageData := response.Data[0]
		var imageBytes []byte

		if len(imageData.B64Json) > 0 {
			if decoded, err := base64.StdEncoding.DecodeString(imageData.B64Json); err == nil {
				imageBytes = decoded
			} else {
				logger.Error(ctx, err)
			}
		} else if len(imageData.Url) > 0 {

			if gstr.HasPrefix(imageData.Url, "data:image/png;base64,") {

				if decoded, err := base64.StdEncoding.DecodeString(gstr.TrimLeftStr(imageData.Url, "data:image/png;base64,")); err == nil {
					imageBytes = decoded
				} else {
					logger.Error(ctx, err)
				}

			} else if gstr.HasPrefix(imageData.Url, "http") {

				client := &http.Client{Timeout: timeout}

				resp, err := client.Get(imageData.Url)
				if err != nil {
					logger.Error(ctx, err)
					if resp.Body != nil {
						_ = resp.Body.Close()
					}
				} else {
					imageBytes, err = io.ReadAll(resp.Body)
					_ = resp.Body.Close()
					if err != nil {
						logger.Error(ctx, err)
					}
				}
			}
		}

		if imageBytes != nil {
			if err = gfile.PutBytes(filePath+fileName, imageBytes); err != nil {
				logger.Error(ctx, err)
			} else {

				if gstr.HasPrefix(filePath, "./resource/public/") {
					imageUrl = "/public/" + gstr.TrimLeftStr(filePath, "./resource/public/") + fileName
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

	responseData := make(map[string]any)
	if response.ResponseBytes != nil {
		if err := json.Unmarshal(response.ResponseBytes, &responseData); err != nil {
			logger.Error(ctx, err)
		} else {
			if data, ok := responseData["data"].([]any); ok {
				for _, d := range data {
					if v, ok := d.(map[string]any); ok {
						v["b64_json"] = ""
						if gstr.HasPrefix(gconv.String(v["url"]), "data:image/png;base64,") {
							v["url"] = ""
						}
					}
				}
			}
		}
	}

	// 通过CAS抢占任务完成: 仅当任务仍为进行中时才落库, 避免被管理员重新生成后多个进程重复完成、重复覆盖、重复计费
	if _, err = dao.TaskImage.FindOneAndUpdate(ctx, bson.M{
		"_id":    taskImage.Id,
		"status": "in_progress",
	}, bson.M{
		"progress":      100,
		"status":        "completed",
		"completed_at":  completedAt,
		"expires_at":    expiresAt,
		"image_url":     imageUrl,
		"file_name":     fileName,
		"file_path":     filePath + fileName,
		"response_data": responseData,
		"error":         nil,
	}); err != nil {

		// 任务已不在进行中, 说明已被其它进程完成或被重置/删除, 跳过且不计费
		if errors.Is(err, mongo.ErrNoDocuments) {
			logger.Infof(ctx, "sTaskImage processImageTask task: %s already handled by another worker, skip", taskImage.Id)
			// 抢占失败(任务已被删除/重置), 清理本次已落盘的孤儿文件, 避免无人回收
			if fileName != "" {
				if e := gfile.RemoveFile(filePath + fileName); e != nil {
					logger.Error(ctx, e)
				}
			}
			return
		}

		logger.Error(ctx, err)
		return
	}

	// 抢占成功, 计算并记录花费
	common.Billing(ctx, response.Usage, &logImage.Spend)

	if err = common.RecordSpend(ctx, logImage.UserId, logImage.AppId, logImage.Creator, logImage.Rid, logImage.Key, logImage.Spend); err != nil {
		logger.Error(ctx, err)
		return
	}

	if err = dao.LogImage.UpdateById(ctx, logImage.Id, bson.M{"spend": logImage.Spend}); err != nil {
		logger.Error(ctx, err)
		return
	}
}

// 按进行中已耗时(秒)计算伪进度档位: 30s→20, 60s→40, 90s→60, 120s→80, 150s→90, 180s→95, 210s→99, 此后维持99, 不足30s为0
func progressForElapsed(elapsedSec int64) int {
	switch {
	case elapsedSec >= 210:
		return 99
	case elapsedSec >= 180:
		return 95
	case elapsedSec >= 150:
		return 90
	case elapsedSec >= 120:
		return 80
	case elapsedSec >= 90:
		return 60
	case elapsedSec >= 60:
		return 40
	case elapsedSec >= 30:
		return 20
	default:
		return 0
	}
}

// 在Task()巡检中推进伪进度: 依据进行中已耗时(now-updated_at)计算应达档位, 仅在更高时CAS推进
// 进度写库时显式保留updated_at, 既不影响僵死回收判定, 也不会因写进度而刷新存活时间
func (s *sTaskImage) advanceFakeProgress(ctx context.Context, now int64, taskImage *entity.TaskImage) {

	target := progressForElapsed((now - taskImage.UpdatedAt) / 1000)
	if target <= taskImage.Progress {
		return
	}

	// CAS: 仍为进行中且进度未被更高值(如完成的100)覆盖时才更新; 用$not兼容progress字段缺失($lt不匹配缺失字段)
	if _, err := dao.TaskImage.FindOneAndUpdate(ctx, bson.M{
		"_id":      taskImage.Id,
		"status":   "in_progress",
		"progress": bson.M{"$not": bson.M{"$gte": target}},
	}, bson.M{
		"progress":   target,
		"updated_at": taskImage.UpdatedAt,
	}); err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return
		}
		logger.Error(ctx, err)
	}
}

func (s *sTaskImage) requeueTask(ctx context.Context, taskId string) {

	// 仅当任务仍为进行中时才置回排队中, 避免覆盖已被重新生成、重置或删除的任务; 保留job_id以便下一轮续轮询
	if _, err := dao.TaskImage.FindOneAndUpdate(ctx, bson.M{
		"_id":    taskId,
		"status": "in_progress",
	}, bson.M{
		"status": "queued",
	}); err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return
		}
		logger.Error(ctx, err)
	}
}

func (s *sTaskImage) failTask(ctx context.Context, taskId, code, message string) {
	// 仅当任务仍为进行中时才置为失败, 避免旧任务的失败覆盖已被重新生成并完成的新结果
	if _, err := dao.TaskImage.FindOneAndUpdate(ctx, bson.M{
		"_id":    taskId,
		"status": "in_progress",
	}, bson.M{
		"status": "failed",
		"error":  &smodel.ImageError{Code: code, Message: message},
	}); err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return
		}
		logger.Error(ctx, err)
	}
}

// 同步提交绘图任务, 阻塞等待上游返回结果
func (s *sTaskImage) requestImageSync(ctx context.Context, taskImage *entity.TaskImage, logImage *entity.LogImage, provider *entity.Provider, timeout time.Duration) (smodel.ImageResponse, string, error) {

	var response smodel.ImageResponse

	adapter := sdk.NewAdapter(ctx, &options.AdapterOptions{
		Provider: provider.Code,
		Model:    logImage.Model,
		Key:      logImage.Key,
		BaseUrl:  logImage.ModelAgent.BaseUrl,
		Path:     logImage.ModelAgent.Path,
		Timeout:  timeout,
		ProxyUrl: config.Cfg.Http.ProxyUrl,
	})

	if taskImage.Action == "edits" {

		var (
			imageEditReq smodel.ImageEditRequest
			err          error
		)

		if config.Cfg.ImageTask.DataFormat == 2 {
			imageEditReq, err = buildImageEditRequestByURL(ctx, taskImage)
		} else {
			imageEditReq, err = buildImageEditRequest(ctx, taskImage)
		}

		if err != nil {
			return response, "build_edit_request_error", err
		}

		if response, err = adapter.ImageEdits(ctx, imageEditReq); err != nil {
			errCode := "edit_error"
			if ctx.Err() != nil {
				errCode = "timeout"
			}
			return response, errCode, err
		}

	} else {

		requestBytes, err := gjson.Encode(taskImage.RequestData)
		if err != nil {
			return response, "request_encode_error", err
		}

		if response, err = adapter.ImageGenerations(ctx, requestBytes); err != nil {
			errCode := "generation_error"
			if ctx.Err() != nil {
				errCode = "timeout"
			}
			return response, errCode, err
		}
	}

	// 上游返回成功但无图像数据视为失败, 不能落库为已完成更不能计费
	if len(response.Data) == 0 {
		return response, "no_image", errors.New("no image in response")
	}

	return response, "", nil
}

// 异步提交绘图任务, 复用适配器提交后轮询上游任务状态
func (s *sTaskImage) requestImageAsync(ctx context.Context, taskImage *entity.TaskImage, logImage *entity.LogImage, provider *entity.Provider, timeout time.Duration) (smodel.ImageResponse, string, error) {

	var response smodel.ImageResponse

	jobId := taskImage.JobId

	// 已有上游句柄(进程内重试或重启后reclaim恢复)直接续轮询, 跳过提交, 避免上游重复出图与重复计费
	if jobId == "" {

		adapter := sdk.NewAdapter(ctx, &options.AdapterOptions{
			Provider: provider.Code,
			Model:    logImage.Model,
			Key:      logImage.Key,
			BaseUrl:  logImage.ModelAgent.BaseUrl,
			Path:     logImage.ModelAgent.Path,
			Timeout:  timeout,
			ProxyUrl: config.Cfg.Http.ProxyUrl,
		})

		var submitResponse smodel.ImageResponse

		if taskImage.Action == "edits" {

			// 异步编辑仅支持图像URL或file_id, 统一走URL方式提交
			imageEditReq, err := buildImageEditRequestByURL(ctx, taskImage)
			if err != nil {
				return response, "build_edit_request_error", err
			}

			imageEditReq.Async = true

			if submitResponse, err = adapter.ImageEdits(ctx, imageEditReq); err != nil {
				errCode := "edit_error"
				if ctx.Err() != nil {
					errCode = "timeout"
				}
				return response, errCode, err
			}

		} else {

			taskImage.RequestData["async"] = true

			requestBytes, err := gjson.Encode(taskImage.RequestData)
			if err != nil {
				return response, "request_encode_error", err
			}

			if submitResponse, err = adapter.ImageGenerations(ctx, requestBytes); err != nil {
				errCode := "generation_error"
				if ctx.Err() != nil {
					errCode = "timeout"
				}
				return response, errCode, err
			}
		}

		var jobResponse smodel.ImageJobResponse
		if submitResponse.ResponseBytes != nil {
			if err := json.Unmarshal(submitResponse.ResponseBytes, &jobResponse); err != nil {
				return response, "submit_response_parse_error", err
			}
		}

		if jobResponse.Id == "" {
			return response, "submit_response_invalid", errors.New("missing image id in async submit response")
		}

		jobId = jobResponse.Id

		// 立刻落库上游句柄, 必须先于轮询: 重启后reclaim可凭此续轮询而非重新提交
		if err := dao.TaskImage.UpdateById(ctx, taskImage.Id, bson.M{"job_id": jobId}); err != nil {
			logger.Error(ctx, err)
		}
		taskImage.JobId = jobId
	}

	// 轮询上游任务状态, 直到完成或超时
	job, errCode, err := s.pollImageJob(ctx, logImage, jobId, timeout)
	if err != nil {
		// timeout时上游任务可能仍在进行, 保留句柄交由上层置回queued续轮询; 其余(上游已失败/过期/删除)清空句柄, 由重试重新提交
		if errCode != "timeout" {
			taskImage.JobId = ""
		}
		return response, errCode, err
	}

	if job.Usage != nil {
		response.Usage = *job.Usage
	}

	// 上游标记已完成但无图像(无data且无image_url)视为失败, 清空句柄交由重试重新提交, 不能落库为已完成更不能计费
	if len(job.Data) == 0 && job.ImageUrl == "" {
		taskImage.JobId = ""
		return response, "no_image", errors.New("no image in completed job")
	}

	// 获取图像数据: 优先通过URL下载, 失败则调用content接口兜底
	if config.Cfg.ImageTask.IsEnableStorage {

		var imageBytes []byte

		if job.ImageUrl != "" {
			if imageBytes, err = s.downloadImage(ctx, job.ImageUrl, timeout); err != nil {
				logger.Errorf(ctx, "sTaskImage requestImageAsync download imageUrl: %s, error: %v", job.ImageUrl, err)
				imageBytes = nil
			}
		}

		if imageBytes == nil {
			if imageBytes, err = s.fetchImageContent(ctx, logImage, jobId, timeout); err != nil {
				logger.Errorf(ctx, "sTaskImage requestImageAsync fetchImageContent imageId: %s, error: %v", jobId, err)
				imageBytes = nil
			}
		}

		if len(imageBytes) > 0 {
			response.Data = []smodel.ImageResponseData{{B64Json: base64.StdEncoding.EncodeToString(imageBytes)}}
		}
	}

	if responseBytes, err := json.Marshal(job); err == nil {
		response.ResponseBytes = responseBytes
	}

	return response, "", nil
}

// 每5秒轮询一次上游任务状态, 直到任务完成、失败或超时
func (s *sTaskImage) pollImageJob(ctx context.Context, logImage *entity.LogImage, imageId string, timeout time.Duration) (smodel.ImageJobResponse, string, error) {

	var jobResponse smodel.ImageJobResponse

	ticker := time.NewTicker(5 * time.Second)
	defer ticker.Stop()

	for {

		job, err := s.retrieveImageJob(ctx, logImage, imageId, timeout)
		if err != nil {
			if ctx.Err() != nil {
				return jobResponse, "timeout", ctx.Err()
			}
			logger.Errorf(ctx, "sTaskImage pollImageJob retrieve imageId: %s, error: %v", imageId, err)
		} else {
			switch job.Status {
			case "completed":
				return job, "", nil
			case "failed", "expired", "deleted":
				message := job.Status
				if job.Error != nil {
					message = job.Error.Message
				}
				return job, "async_" + job.Status, errors.New(message)
			}
		}

		select {
		case <-ctx.Done():
			return jobResponse, "timeout", ctx.Err()
		case <-ticker.C:
		}
	}
}

// 查询上游绘图任务状态
func (s *sTaskImage) retrieveImageJob(ctx context.Context, logImage *entity.LogImage, imageId string, timeout time.Duration) (smodel.ImageJobResponse, error) {

	var jobResponse smodel.ImageJobResponse

	reqUrl := gstr.TrimRight(logImage.ModelAgent.BaseUrl, "/") + "/images/" + imageId

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, reqUrl, nil)
	if err != nil {
		return jobResponse, err
	}

	req.Header.Set("Authorization", "Bearer "+logImage.Key)

	client := &http.Client{Timeout: timeout}

	resp, err := client.Do(req)
	if err != nil {
		return jobResponse, err
	}
	defer func() {
		_ = resp.Body.Close()
	}()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return jobResponse, err
	}

	if resp.StatusCode != http.StatusOK {
		return jobResponse, errors.Newf("retrieve image job failed, status: %d, body: %s", resp.StatusCode, body)
	}

	if err = json.Unmarshal(body, &jobResponse); err != nil {
		return jobResponse, err
	}

	return jobResponse, nil
}

// 调用上游content接口下载图像字节数据
func (s *sTaskImage) fetchImageContent(ctx context.Context, logImage *entity.LogImage, imageId string, timeout time.Duration) ([]byte, error) {

	reqUrl := gstr.TrimRight(logImage.ModelAgent.BaseUrl, "/") + "/images/" + imageId + "/content"

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, reqUrl, nil)
	if err != nil {
		return nil, err
	}

	req.Header.Set("Authorization", "Bearer "+logImage.Key)

	client := &http.Client{Timeout: timeout}

	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer func() {
		_ = resp.Body.Close()
	}()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != http.StatusOK {
		return nil, errors.Newf("fetch image content failed, status: %d, body: %s", resp.StatusCode, body)
	}

	return body, nil
}

// 通过URL下载图像字节数据
func (s *sTaskImage) downloadImage(ctx context.Context, imageUrl string, timeout time.Duration) ([]byte, error) {

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, imageUrl, nil)
	if err != nil {
		return nil, err
	}

	client := &http.Client{Timeout: timeout}

	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer func() {
		_ = resp.Body.Close()
	}()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != http.StatusOK {
		return nil, errors.Newf("download image failed, status: %d", resp.StatusCode)
	}

	return body, nil
}

func buildImageEditRequest(ctx context.Context, taskImage *entity.TaskImage) (smodel.ImageEditRequest, error) {

	req := smodel.ImageEditRequest{
		Model: taskImage.Model,
	}

	if v, ok := taskImage.RequestData["prompt"]; ok {
		req.Prompt, _ = v.(string)
	}
	if v, ok := taskImage.RequestData["n"]; ok {
		if n, ok := v.(float64); ok {
			req.N = int(n)
		}
	}
	if v, ok := taskImage.RequestData["quality"]; ok {
		req.Quality, _ = v.(string)
	}
	if v, ok := taskImage.RequestData["size"]; ok {
		req.Size, _ = v.(string)
	}
	if v, ok := taskImage.RequestData["response_format"]; ok {
		req.ResponseFormat, _ = v.(string)
	}
	if v, ok := taskImage.RequestData["background"]; ok {
		req.Background, _ = v.(string)
	}

	imageVal, ok := taskImage.RequestData["images"]
	if !ok {
		imageVal, ok = taskImage.RequestData["image"]
		if !ok {
			return req, errors.New("missing image parameter in request data")
		}
	}

	var imageUrls []string
	switch v := imageVal.(type) {
	case string:
		imageUrls = append(imageUrls, v)
	case []any:
		for _, item := range v {
			if s, ok := item.(string); ok {
				imageUrls = append(imageUrls, s)
			}
		}
	case bson.A:
		for _, item := range v {
			if s, ok := item.(string); ok {
				imageUrls = append(imageUrls, s)
			} else if urls, ok := item.(bson.D); ok {
				for _, u := range urls {
					imageUrls = append(imageUrls, u.Value.(string))
				}
			}
		}
	case any:
		if url, ok := v.(string); ok {
			imageUrls = append(imageUrls, url)
		} else if urls, ok := v.([]any); ok {
			for _, item := range urls {
				if s, ok := item.(string); ok {
					imageUrls = append(imageUrls, s)
				}
			}
		} else {
			return req, errors.New("invalid image parameter type")
		}
	default:
		return req, errors.New("invalid image parameter type")
	}

	if len(imageUrls) == 0 {
		return req, errors.New("empty image urls")
	}

	timeout := config.Cfg.ImageTask.Timeout * time.Second
	if timeout <= 0 {
		timeout = config.Cfg.Base.LongTimeout * time.Second
	}

	client := &http.Client{Timeout: timeout}

	fileHeaders := make([]*multipart.FileHeader, 0, len(imageUrls))

	for _, imageUrl := range imageUrls {

		resp, err := client.Get(imageUrl)
		if err != nil {
			if resp.Body != nil {
				_ = resp.Body.Close()
			}
			return req, errors.Newf("download image failed: %s, error: %v", imageUrl, err)
		}

		imageBytes, err := io.ReadAll(resp.Body)
		_ = resp.Body.Close()
		if err != nil {
			return req, errors.Newf("read image failed: %s, error: %v", imageUrl, err)
		}

		fileHeader, err := bytesToFileHeader(imageUrl, imageBytes, resp.Header.Get("Content-Type"))
		if err != nil {
			return req, err
		}

		fileHeaders = append(fileHeaders, fileHeader)
	}

	req.Image = fileHeaders

	return req, nil
}

func buildImageEditRequestByURL(ctx context.Context, taskImage *entity.TaskImage) (smodel.ImageEditRequest, error) {

	var req smodel.ImageEditRequest

	req.Model = taskImage.RequestData["model"].(string)

	if v, ok := taskImage.RequestData["prompt"]; ok {
		req.Prompt, _ = v.(string)
	}
	if v, ok := taskImage.RequestData["n"]; ok {
		req.N, _ = v.(int)
	}
	if v, ok := taskImage.RequestData["quality"]; ok {
		req.Quality, _ = v.(string)
	}
	if v, ok := taskImage.RequestData["size"]; ok {
		req.Size, _ = v.(string)
	}
	if v, ok := taskImage.RequestData["response_format"]; ok {
		req.ResponseFormat, _ = v.(string)
	}
	if v, ok := taskImage.RequestData["background"]; ok {
		req.Background, _ = v.(string)
	}

	imageVal, ok := taskImage.RequestData["images"]
	if !ok {
		imageVal, ok = taskImage.RequestData["image"]
		if !ok {
			return req, errors.New("missing image parameter in request data")
		}
	}

	var imageUrls []string
	switch v := imageVal.(type) {
	case string:
		imageUrls = append(imageUrls, v)
	case []any:
		for _, item := range v {
			if s, ok := item.(string); ok {
				imageUrls = append(imageUrls, s)
			}
		}
	case bson.A:
		for _, item := range v {
			if s, ok := item.(string); ok {
				imageUrls = append(imageUrls, s)
			} else if urls, ok := item.(bson.D); ok {
				for _, u := range urls {
					imageUrls = append(imageUrls, u.Value.(string))
				}
			}
		}
	case any:
		if url, ok := v.(string); ok {
			imageUrls = append(imageUrls, url)
		} else if urls, ok := v.([]any); ok {
			for _, item := range urls {
				if s, ok := item.(string); ok {
					imageUrls = append(imageUrls, s)
				}
			}
		} else {
			return req, errors.New("invalid image parameter type")
		}
	default:
		return req, errors.New("invalid image parameter type")
	}

	if len(imageUrls) == 0 {
		return req, errors.New("empty image urls")
	}

	req.Images = make([]smodel.ImageEditImage, 0, len(imageUrls))
	for _, imageUrl := range imageUrls {
		req.Images = append(req.Images, smodel.ImageEditImage{ImageUrl: imageUrl})
	}

	return req, nil
}

func bytesToFileHeader(fileUrl string, data []byte, contentType string) (*multipart.FileHeader, error) {

	parsed, _ := url.Parse(fileUrl)
	fileName := path.Base(parsed.Path)
	if fileName == "" || fileName == "." || fileName == "/" {
		fileName = "image.png"
	}

	if contentType == "" {
		contentType = http.DetectContentType(data)
	}

	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)

	h := make(textproto.MIMEHeader)
	h.Set("Content-Disposition", fmt.Sprintf(`form-data; name="image"; filename="%s"`, fileName))
	h.Set("Content-Type", contentType)

	part, err := writer.CreatePart(h)
	if err != nil {
		return nil, err
	}

	if _, err = part.Write(data); err != nil {
		return nil, err
	}

	if err = writer.Close(); err != nil {
		return nil, err
	}

	reader := multipart.NewReader(body, writer.Boundary())
	form, err := reader.ReadForm(int64(len(data)) + 1024)
	if err != nil {
		return nil, err
	}

	files := form.File["image"]
	if len(files) == 0 {
		return nil, errors.Newf("failed to create file header for %s", fileUrl)
	}

	return files[0], nil
}
