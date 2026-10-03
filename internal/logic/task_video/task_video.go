package task_video

import (
	"context"
	"encoding/json"
	"net/url"
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
	sconsts "github.com/iimeta/fastapi-sdk/v2/consts"
	smodel "github.com/iimeta/fastapi-sdk/v2/model"
	"github.com/iimeta/fastapi-sdk/v2/options"
	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
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
		Action:             taskVideo.Action,
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

	if detail.Error != nil {
		common.ShieldTaskError(ctx, detail.Error)
	}

	if taskVideo.VideoUrl != "" {
		detail.VideoUrl = resolveVideoUrl(taskVideo.VideoUrl)
	}

	if service.Session().IsAdminRole(ctx) {
		if isVideoJobHandle(taskVideo.JobId) {
			detail.JobId = taskVideo.JobId
		}
		detail.FileName = taskVideo.FileName
		detail.FilePath = taskVideo.FilePath
		detail.ModelAgentId = taskVideo.ModelAgentId

		if taskVideo.ModelAgent != nil {

			providerName := taskVideo.ModelAgent.ProviderId
			if provider, err := dao.Provider.FindById(ctx, taskVideo.ModelAgent.ProviderId); err == nil && provider != nil {
				providerName = provider.Name
			}

			detail.ModelAgent = &model.ModelAgent{
				ProviderId:   taskVideo.ModelAgent.ProviderId,
				ProviderName: providerName,
				Name:         taskVideo.ModelAgent.Name,
				BaseUrl:      taskVideo.ModelAgent.BaseUrl,
				Path:         taskVideo.ModelAgent.Path,
				Weight:       taskVideo.ModelAgent.Weight,
				Remark:       taskVideo.ModelAgent.Remark,
			}
		}
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

	if params.Prompt != "" {
		filter["prompt"] = bson.M{
			"$regex": regexp.QuoteMeta(gstr.Trim(params.Prompt)),
		}
	}

	if !service.Session().IsAdminRole(ctx) && len(params.Models) > 0 {
		names := make([]string, 0, len(params.Models))
		if list, err := dao.Model.FindByIds(ctx, params.Models); err != nil {
			logger.Error(ctx, err)
		} else {
			for _, m := range list {
				if m.Name != "" {
					names = append(names, m.Name)
				}
			}
		}
		filter["model"] = bson.M{"$in": names}
	}

	if service.Session().IsAdminRole(ctx) && len(params.ModelAgents) > 0 {
		filter["model_agent_id"] = bson.M{"$in": params.ModelAgents}
	}

	if params.Status != "" {
		filter["status"] = params.Status
	} else if !service.Session().IsAdminRole(ctx) {
		filter["status"] = bson.M{"$ne": "deleted"}
	}

	if len(params.CreatedAt) > 0 && params.TraceId == "" {
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

		video := &model.TaskVideo{
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
			Progress:  result.Progress,
			Status:    result.Status,
			CreatedAt: util.FormatDateTimeMonth(result.CreatedAt),
		}

		if result.Status == "failed" && result.Error != nil {
			video.ErrMsg = common.ConvTaskErrMsg(ctx, result.Error)
		}

		if result.VideoUrl != "" {
			video.VideoUrl = resolveVideoUrl(result.VideoUrl)
		}

		if service.Session().IsAdminRole(ctx) {
			video.TotalTime = taskVideoTotalTime(result)
		}

		items = append(items, video)
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

// 视频任务重新生成
func (s *sTaskVideo) Regenerate(ctx context.Context, id string) error {

	taskVideo, err := dao.TaskVideo.FindOneAndUpdate(ctx, bson.M{
		"_id": id,
		"status": bson.M{
			"$in": []string{"in_progress", "failed"},
		},
	}, bson.M{
		"status":   "queued",
		"progress": 0,
		"error":    nil,
	})
	if err != nil {

		if errors.Is(err, mongo.ErrNoDocuments) {
			return errors.New("任务不在进行中或已失败状态, 无法重新生成")
		}

		logger.Error(ctx, err)
		return err
	}

	// 有原请求体则放弃旧句柄, 下一轮重新提交; 旧任务缺 RequestData 时保留句柄, 仍按原 video_id 轮询
	if taskVideo != nil && len(taskVideo.RequestData) > 0 {
		if err = dao.TaskVideo.UpdateById(ctx, taskVideo.Id, bson.M{"job_id": videoJobResubmit}); err != nil {
			logger.Error(ctx, err)
		}
	}

	if taskVideo != nil && taskVideo.TraceId != "" {
		if err = dao.LogVideo.UpdateOne(ctx, bson.M{"trace_id": taskVideo.TraceId}, bson.M{
			"status":  1,
			"err_msg": "",
		}); err != nil {
			logger.Error(ctx, err)
			return err
		}
	}

	return nil
}

// 视频任务批量操作
func (s *sTaskVideo) BatchOperate(ctx context.Context, params model.TaskVideoBatchOperateReq) error {

	switch params.Action {
	case consts.ACTION_REGENERATE:
		for _, id := range params.Ids {
			if err := s.Regenerate(ctx, id); err != nil {
				logger.Error(ctx, err)
			}
		}
	}

	return nil
}

// 视频任务
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

	reclaimMillis := (config.Cfg.VideoTask.Reclaim * time.Second).Milliseconds()
	if reclaimMillis <= 0 {
		timeout := config.Cfg.VideoTask.Timeout
		if timeout <= 0 {
			timeout = config.Cfg.Base.LongTimeout
		}
		retryCount := config.Cfg.VideoTask.RetryCount
		if retryCount < 0 {
			retryCount = 0
		}
		reclaimMillis = (timeout * time.Duration(retryCount+1) * time.Second).Milliseconds()
	}

	var staleBefore int64
	if reclaimMillis > 0 {
		staleBefore = now - reclaimMillis
	}

	taskVideos, err := dao.TaskVideo.Find(ctx, bson.M{"status": bson.M{"$in": []string{"queued", "in_progress"}}}, &dao.FindOptions{SortFields: []string{"created_at"}})
	if err != nil {
		logger.Error(ctx, err)
		return
	}

	availableSlots := -1
	if config.Cfg.VideoTask.ConcurrencyLimit > 0 {
		liveInProgress := 0
		for _, taskVideo := range taskVideos {
			if taskVideo.Status == "in_progress" && taskVideo.UpdatedAt >= staleBefore {
				liveInProgress++
			}
		}
		if availableSlots = config.Cfg.VideoTask.ConcurrencyLimit - liveInProgress; availableSlots < 0 {
			availableSlots = 0
		}
	}

	var queuedTasks []*entity.TaskVideo

	for _, taskVideo := range taskVideos {

		if taskVideo.Status == "in_progress" && taskVideo.UpdatedAt >= staleBefore {
			s.advanceFakeProgress(ctx, now, taskVideo)
			continue
		}

		if availableSlots == 0 {
			continue
		}

		if err = dao.TaskVideo.UpdateById(ctx, taskVideo.Id, bson.M{"status": "in_progress", "error": nil}); err != nil {
			logger.Error(ctx, err)
			continue
		}

		queuedTasks = append(queuedTasks, taskVideo)

		if availableSlots > 0 {
			availableSlots--
		}
	}

	for _, taskVideo := range queuedTasks {
		if err := grpool.AddWithRecover(gctx.NeverDone(ctx), func(ctx context.Context) {

			ctx, err = gtrace.WithTraceID(ctx, taskVideo.TraceId)

			s.processVideoTask(ctx, taskVideo)

		}, nil); err != nil {
			logger.Error(ctx, err)
		}
	}

	s.cleanExpiredAndFiles(ctx, now)

	if _, err := redis.Set(ctx, consts.TASK_VIDEO_END_TIME_KEY, gtime.TimestampMilli()); err != nil {
		logger.Error(ctx, err)
	}
}

func (s *sTaskVideo) processVideoTask(ctx context.Context, taskVideo *entity.TaskVideo) {

	logVideos, err := dao.LogVideo.Find(ctx, bson.M{
		"trace_id": taskVideo.TraceId,
		"action":   bson.M{"$in": []string{"create", "remix"}},
		"status":   bson.M{"$in": []int{1, -1}},
	}, &dao.FindOptions{SortFields: []string{"-created_at"}})
	if err != nil {
		logger.Error(ctx, err)
		s.failTask(ctx, taskVideo.Id, "log_not_found", err.Error(), nil)
		return
	}

	if len(logVideos) == 0 {
		logVideos, err = dao.LogVideo.Find(ctx, bson.M{
			"trace_id": taskVideo.TraceId,
			"status":   bson.M{"$in": []int{1, -1}},
		}, &dao.FindOptions{SortFields: []string{"-created_at"}})
		if err != nil {
			logger.Error(ctx, err)
			s.failTask(ctx, taskVideo.Id, "log_not_found", err.Error(), nil)
			return
		}
	}

	logVideo := pickBestLogVideo(logVideos)

	if logVideo == nil {
		if err = dao.TaskVideo.UpdateById(ctx, taskVideo.Id, bson.M{"status": "queued", "error": nil}); err != nil {
			logger.Error(ctx, err)
		}
		return
	}

	if logVideo.Status == -1 {
		logger.Infof(ctx, "sTaskVideo processVideoTask task: %s all log_video failed, mark task failed directly", taskVideo.Id)
		s.failTask(ctx, taskVideo.Id, "log_failed", logVideo.ErrMsg, nil, logVideo.Id)
		return
	}

	if logVideo.ModelAgent == nil {
		logger.Errorf(ctx, "sTaskVideo processVideoTask task: %s log_video: %s has no model_agent", taskVideo.Id, logVideo.Id)
		s.failTask(ctx, taskVideo.Id, "model_agent_not_found", "log_video has no model_agent", nil, logVideo.Id)
		return
	}

	provider, err := dao.Provider.FindById(ctx, logVideo.ModelAgent.ProviderId)
	if err != nil {
		logger.Error(ctx, err)
		s.failTask(ctx, taskVideo.Id, "provider_not_found", err.Error(), nil, logVideo.Id)
		return
	}

	timeout := config.Cfg.VideoTask.Timeout * time.Second
	if timeout <= 0 {
		timeout = config.Cfg.Base.LongTimeout * time.Second
	}

	retryCount := config.Cfg.VideoTask.RetryCount
	if retryCount < 0 {
		retryCount = 0
	}

	var (
		retrieve smodel.VideoJobResponse
		errCode  string
		videoUrl string
		fileName string
		filePath string
	)

	errorAgentIds := make([]string, 0)
	errorKeys := make([]string, 0)

	for attempt := 0; ; attempt++ {

		// 需要重新提交时(无上游句柄)才换代理; 已有 job_id 的续轮询必须留在原代理
		if attempt > 0 && videoPollHandle(taskVideo) == "" {
			agentChanged, e := s.switchRetryAgent(ctx, taskVideo, logVideo, errorAgentIds, errorKeys, attempt)
			if e != nil {
				logger.Errorf(ctx, "sTaskVideo processVideoTask task: %s switch agent failed: %v, retry current agent", taskVideo.Id, e)
			} else if agentChanged && logVideo.ModelAgent != nil {
				provider, err = dao.Provider.FindById(ctx, logVideo.ModelAgent.ProviderId)
				if err != nil {
					logger.Error(ctx, err)
					s.failTask(ctx, taskVideo.Id, "provider_not_found", err.Error(), nil, logVideo.Id)
					return
				}
			}
		}

		taskCtx, cancel := context.WithTimeout(ctx, timeout)
		retrieve, errCode, err = s.requestVideo(taskCtx, taskVideo, logVideo, provider, timeout)
		cancel()

		videoUrl, fileName, filePath = "", "", ""

		if err == nil && retrieve.Status == "completed" {

			if config.Cfg.VideoTask.IsEnableStorage {

				if content, e := s.downloadVideo(ctx, logVideo, provider, taskVideo, timeout); e != nil {
					logger.Error(ctx, e)
					err = e
					errCode = "storage_failed"
				} else if len(content) == 0 {
					err = errors.New("downloaded video content is empty")
					errCode = "storage_failed"
				} else {

					storageDir := config.Cfg.VideoTask.StorageDir
					if storageDir == "" {
						storageDir = "./resource/public/video/"
					} else if !gstr.HasSuffix(storageDir, "/") {
						storageDir = storageDir + "/"
					}

					fileName = pollVideoId(taskVideo) + "_video.mp4"
					if e := gfile.PutBytes(storageDir+fileName, content); e != nil {
						logger.Error(ctx, e)
						err = e
						errCode = "storage_failed"
					} else {
						filePath = storageDir + fileName
						if gstr.HasPrefix(storageDir, "./resource/public/") {
							videoUrl = "/public/" + gstr.TrimLeftStr(storageDir, "./resource/public/") + fileName
						} else if config.Cfg.VideoTask.StorageBaseUrl == "" {
							videoUrl = "/open/video/" + fileName
						} else {
							videoUrl = fileName
						}
					}
				}

			} else if retrieve.VideoUrl != "" {
				videoUrl = retrieve.VideoUrl
			} else {
				err = errors.New("no video in completed job")
				errCode = "no_video"
			}
		}

		if err == nil {
			break
		}

		if errCode == "retrieve_error" {
			logger.Infof(ctx, "sTaskVideo processVideoTask task: %s upstream api abnormal, requeue to resume next round", taskVideo.Id)
			s.requeueTask(ctx, taskVideo.Id)
			return
		}

		isRetry, isDisabled := common.IsNeedRetry(err)

		if isDisabled {
			if e := grpool.AddWithRecover(gctx.NeverDone(ctx), func(ctx context.Context) {

				service.Key().AutoDisabled(ctx, logVideo.Key, err.Error())

			}, nil); e != nil {
				logger.Error(ctx, e)
			}
		}

		if errCode != "timeout" && !isRetry {
			logger.Errorf(ctx, "sTaskVideo processVideoTask task: %s failed: %s, error: %v, no need to retry by config", taskVideo.Id, errCode, err)
			s.failTask(ctx, taskVideo.Id, errCode, err.Error(), util.ConvToMap(retrieve.ResponseBytes), logVideo.Id)
			return
		}

		if attempt < retryCount {

			if latest, e := dao.TaskVideo.FindById(ctx, taskVideo.Id); e != nil {
				logger.Error(ctx, e)
			} else if latest.Status != "in_progress" {
				logger.Infof(ctx, "sTaskVideo processVideoTask task: %s status is %s, no need to retry, skip", taskVideo.Id, latest.Status)
				return
			}

			errorAgentIds = appendUnique(errorAgentIds, logVideo.ModelAgentId)
			errorKeys = appendUnique(errorKeys, logVideo.Key)

			logger.Errorf(ctx, "sTaskVideo processVideoTask task: %s failed: %s, retry: %d/%d", taskVideo.Id, errCode, attempt+1, retryCount)
			continue
		}

		logger.Error(ctx, err)
		s.failTask(ctx, taskVideo.Id, errCode, err.Error(), util.ConvToMap(retrieve.ResponseBytes), logVideo.Id)
		return
	}

	completedAt := gtime.TimestampMilli() / 1000
	var expiresAt int64

	if config.Cfg.VideoTask.IsEnableStorage {
		if videoUrl != "" && config.Cfg.VideoTask.StorageExpiresAt > 0 {
			expiresAt = gtime.NewFromTimeStamp(completedAt).Add(config.Cfg.VideoTask.StorageExpiresAt * time.Minute).Unix()
		}
	} else if retrieve.ExpiresAt != nil {
		expiresAt = *retrieve.ExpiresAt
	}

	if retrieve.CompletedAt != nil {
		completedAt = *retrieve.CompletedAt
	}

	responseData := make(map[string]any)
	if retrieve.ResponseBytes != nil {
		if e := json.Unmarshal(retrieve.ResponseBytes, &responseData); e != nil {
			logger.Error(ctx, e)
		}
	}

	update := bson.M{
		"progress":              100,
		"status":                "completed",
		"completed_at":          completedAt,
		"expires_at":            expiresAt,
		"video_url":             videoUrl,
		"file_name":             fileName,
		"file_path":             filePath,
		"remixed_from_video_id": retrieve.RemixedFromVideoId,
		"response_data":         responseData,
		"model_agent_id":        taskVideo.ModelAgentId,
		"model_agent":           taskVideo.ModelAgent,
		"error":                 nil,
	}
	if retrieve.Id != "" {
		update["job_id"] = retrieve.Id
	}

	if _, err = dao.TaskVideo.FindOneAndUpdate(ctx, bson.M{
		"_id":    taskVideo.Id,
		"status": "in_progress",
	}, update); err != nil {

		if errors.Is(err, mongo.ErrNoDocuments) {
			logger.Infof(ctx, "sTaskVideo processVideoTask task: %s already handled by another worker, skip", taskVideo.Id)
			if filePath != "" {
				if e := gfile.RemoveFile(filePath); e != nil {
					logger.Error(ctx, e)
				}
			}
			return
		}

		logger.Error(ctx, err)
		return
	}

	if provider.Code == sconsts.PROVIDER_VOLC_ENGINE && retrieve.Usage != nil {

		common.Billing(ctx, *retrieve.Usage, &logVideo.Spend)

		if err = common.RecordSpend(ctx, logVideo.UserId, logVideo.AppId, logVideo.Creator, logVideo.Rid, logVideo.Key, logVideo.Spend); err != nil {
			logger.Error(ctx, err)
			return
		}
	}

	if provider.Code == sconsts.PROVIDER_BAILIAN && logVideo.Spend.VideoGeneration != nil && logVideo.Spend.VideoGeneration.Seconds == 0 {
		if seconds := gconv.Int(retrieve.Seconds); seconds > 0 {

			common.RecalcVideoSecondsSpend(&logVideo.Spend, seconds)

			if err = common.RecordSpend(ctx, logVideo.UserId, logVideo.AppId, logVideo.Creator, logVideo.Rid, logVideo.Key, logVideo.Spend); err != nil {
				logger.Error(ctx, err)
				return
			}

			if err = dao.TaskVideo.UpdateById(ctx, taskVideo.Id, bson.M{"seconds": seconds}); err != nil {
				logger.Error(ctx, err)
			}
		}
	}

	if err = dao.LogVideo.UpdateById(ctx, logVideo.Id, bson.M{
		"spend":          logVideo.Spend,
		"model_agent_id": logVideo.ModelAgentId,
		"model_agent":    logVideo.ModelAgent,
		"key":            logVideo.Key,
	}); err != nil {
		logger.Error(ctx, err)
	}
}

func (s *sTaskVideo) requestVideo(ctx context.Context, taskVideo *entity.TaskVideo, logVideo *entity.LogVideo, provider *entity.Provider, timeout time.Duration) (smodel.VideoJobResponse, string, error) {

	var retrieve smodel.VideoJobResponse

	jobId := videoPollHandle(taskVideo)

	if jobId == "" {

		if len(taskVideo.RequestData) == 0 {
			return retrieve, "missing_request_data", errors.New("missing request data")
		}

		upstreamModel := s.prepareUpstreamModel(ctx, taskVideo, logVideo)

		adapter := sdk.NewAdapterOfficial(ctx, &options.AdapterOptions{
			Provider: provider.Code,
			Model:    upstreamModel,
			Key:      logVideo.Key,
			BaseUrl:  logVideo.ModelAgent.BaseUrl,
			Path:     logVideo.ModelAgent.Path,
			Timeout:  timeout,
			ProxyUrl: config.Cfg.Http.ProxyUrl,
		})

		body := gjson.MustEncode(taskVideo.RequestData)
		bytes, _, err := adapter.VideoCreateOfficial(ctx, body)
		if err != nil {
			errCode := "generation_error"
			if ctx.Err() != nil {
				errCode = "timeout"
			}
			return retrieve, errCode, err
		}

		jobId = parseVideoCreateTaskId(bytes)
		if jobId == "" {
			return retrieve, "submit_response_invalid", errors.New("create video task failed: missing task_id")
		}

		// video_id 是客户端查询句柄, 重提只更新上游 job_id
		if err := dao.TaskVideo.UpdateById(ctx, taskVideo.Id, bson.M{"job_id": jobId}); err != nil {
			logger.Error(ctx, err)
		}

		taskVideo.JobId = jobId

	} else if !isVideoJobHandle(taskVideo.JobId) {

		// 旧任务只有 video_id: 收编为 job_id, 后续只认 job_id
		taskVideo.JobId = jobId
		if err := dao.TaskVideo.UpdateById(ctx, taskVideo.Id, bson.M{"job_id": jobId}); err != nil {
			logger.Error(ctx, err)
		}
	}

	job, errCode, err := s.pollVideoJob(ctx, logVideo, provider, jobId, timeout)
	if err != nil {
		if errCode != "retrieve_error" {
			abandonVideoJob(taskVideo)
		}
		return job, errCode, err
	}

	if job.Status == "completed" && job.VideoUrl == "" && !config.Cfg.VideoTask.IsEnableStorage {
		abandonVideoJob(taskVideo)
		return job, "no_video", errors.New("no video in completed job")
	}

	return job, "", nil
}

func (s *sTaskVideo) pollVideoJob(ctx context.Context, logVideo *entity.LogVideo, provider *entity.Provider, videoId string, timeout time.Duration) (smodel.VideoJobResponse, string, error) {

	var jobResponse smodel.VideoJobResponse

	retryCount := config.Cfg.VideoTask.RetryCount
	if retryCount < 0 {
		retryCount = 0
	}

	consecutiveFailures := 0
	lastPollInProgress := false

	ticker := time.NewTicker(5 * time.Second)
	defer ticker.Stop()

	for {

		job, err := s.retrieveVideoJob(ctx, logVideo, provider, videoId, timeout)
		if err != nil {

			lastPollInProgress = false

			if ctx.Err() != nil {
				return jobResponse, "retrieve_error", err
			}

			logger.Errorf(ctx, "sTaskVideo pollVideoJob retrieve videoId: %s, error: %v", videoId, err)
			if consecutiveFailures++; consecutiveFailures > retryCount {
				return jobResponse, "retrieve_error", err
			}

		} else {

			switch job.Status {
			case "completed":
				return job, "", nil
			case "failed", "expired", "deleted", "cancelled":
				message := job.Status
				if job.Error != nil {
					message = job.Error.Message
				}
				return job, "async_" + job.Status, errors.New(message)
			case "queued", "in_progress", "running":
				consecutiveFailures = 0
				lastPollInProgress = true
				s.updateRunningProgress(ctx, videoId, job.Progress)
			default:
				lastPollInProgress = false
				logger.Errorf(ctx, "sTaskVideo pollVideoJob videoId: %s, unknown status: %s", videoId, job.Status)
				if consecutiveFailures++; consecutiveFailures > retryCount {
					message := "unknown status: " + job.Status
					if job.Error != nil {
						message = job.Error.Message
					}
					return job, "async_unknown_status", errors.New(message)
				}
			}
		}

		select {
		case <-ctx.Done():
			if lastPollInProgress {
				return jobResponse, "timeout", ctx.Err()
			}
			return jobResponse, "retrieve_error", ctx.Err()
		case <-ticker.C:
		}
	}
}

func (s *sTaskVideo) retrieveVideoJob(ctx context.Context, logVideo *entity.LogVideo, provider *entity.Provider, videoId string, timeout time.Duration) (smodel.VideoJobResponse, error) {

	adapter := sdk.NewAdapter(ctx, &options.AdapterOptions{
		Provider: provider.Code,
		Model:    s.resolveUpstreamModel(ctx, logVideo, ""),
		Key:      logVideo.Key,
		BaseUrl:  logVideo.ModelAgent.BaseUrl,
		Path:     logVideo.ModelAgent.Path,
		Timeout:  timeout,
		ProxyUrl: config.Cfg.Http.ProxyUrl,
	})

	return adapter.VideoRetrieve(ctx, smodel.VideoRetrieveRequest{VideoId: videoId})
}

func (s *sTaskVideo) downloadVideo(ctx context.Context, logVideo *entity.LogVideo, provider *entity.Provider, taskVideo *entity.TaskVideo, timeout time.Duration, retry ...int) ([]byte, error) {

	adapter := sdk.NewAdapter(ctx, &options.AdapterOptions{
		Provider: provider.Code,
		Model:    s.resolveUpstreamModel(ctx, logVideo, ""),
		Key:      logVideo.Key,
		BaseUrl:  logVideo.ModelAgent.BaseUrl,
		Path:     logVideo.ModelAgent.Path,
		Timeout:  timeout,
		ProxyUrl: config.Cfg.Http.ProxyUrl,
	})

	content, err := adapter.VideoContent(ctx, smodel.VideoContentRequest{VideoId: pollVideoId(taskVideo)})
	if err == nil && len(content.Data) > 0 {
		return content.Data, nil
	}

	if err != nil {
		logger.Errorf(ctx, "sTaskVideo downloadVideo videoId: %s, error: %v", pollVideoId(taskVideo), err)
	} else {
		err = errors.New("downloaded video content is empty")
		logger.Errorf(ctx, "sTaskVideo downloadVideo videoId: %s, empty content", pollVideoId(taskVideo))
	}

	retryCount := config.Cfg.VideoTask.RetryCount
	if retryCount < 0 {
		retryCount = 0
	}

	if len(retry) == retryCount {
		return nil, err
	}

	retry = append(retry, 1)
	time.Sleep(time.Duration(len(retry)*5) * time.Second)
	logger.Infof(ctx, "sTaskVideo downloadVideo retry: %d, videoId: %s", len(retry), pollVideoId(taskVideo))

	return s.downloadVideo(ctx, logVideo, provider, taskVideo, timeout, retry...)
}

func (s *sTaskVideo) updateRunningProgress(ctx context.Context, videoId string, progress int) {

	if progress <= 0 {
		return
	}
	if progress > 0 && progress <= 1 {
		progress = progress * 100
	}
	if progress >= 100 {
		progress = 99
	}

	if _, err := dao.TaskVideo.FindOneAndUpdate(ctx, bson.M{
		"status":   "in_progress",
		"$or":      bson.A{bson.M{"job_id": videoId}, bson.M{"video_id": videoId}},
		"progress": bson.M{"$not": bson.M{"$gte": progress}},
	}, bson.M{
		"progress": progress,
	}); err != nil && !errors.Is(err, mongo.ErrNoDocuments) {
		logger.Error(ctx, err)
	}
}

func (s *sTaskVideo) cleanExpiredAndFiles(ctx context.Context, now int64) {

	expiredTasks, err := dao.TaskVideo.Find(ctx, bson.M{
		"status":     "completed",
		"expires_at": bson.M{"$gt": 0, "$lte": now / 1000},
	})
	if err != nil {
		logger.Error(ctx, err)
	}

	for _, taskVideo := range expiredTasks {

		update := bson.M{"status": "expired"}

		if config.Cfg.VideoTask.StorageExpiredDelete {

			if taskVideo.FilePath != "" {
				if err := gfile.RemoveFile(taskVideo.FilePath); err != nil {
					logger.Error(ctx, err)
				}
			}

			if taskVideo.VideoUrl != "" || taskVideo.FilePath != "" {
				update["video_url"] = ""
				update["file_name"] = ""
				update["file_path"] = ""
			}
		}

		if err := dao.TaskVideo.UpdateById(ctx, taskVideo.Id, update); err != nil {
			logger.Error(ctx, err)
		}
	}

	staleTasks, err := dao.TaskVideo.Find(ctx, bson.M{
		"status":    bson.M{"$in": []string{"failed", "deleted"}},
		"file_path": bson.M{"$exists": true, "$ne": ""},
	})
	if err != nil {
		logger.Error(ctx, err)
	}

	for _, taskVideo := range staleTasks {

		if config.Cfg.VideoTask.StorageExpiresAt <= 0 {
			continue
		}

		expiresAtMillis := taskVideo.CreatedAt + (config.Cfg.VideoTask.StorageExpiresAt * time.Minute).Milliseconds()
		if expiresAtMillis > now {
			continue
		}

		if taskVideo.FilePath != "" {
			if err := gfile.RemoveFile(taskVideo.FilePath); err != nil {
				logger.Error(ctx, err)
			}
		}

		if err := dao.TaskVideo.UpdateById(ctx, taskVideo.Id, bson.M{
			"video_url": "",
			"file_name": "",
			"file_path": "",
		}); err != nil {
			logger.Error(ctx, err)
		}
	}
}

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

func (s *sTaskVideo) advanceFakeProgress(ctx context.Context, now int64, taskVideo *entity.TaskVideo) {

	target := progressForElapsed((now - taskVideo.UpdatedAt) / 1000)
	if target <= taskVideo.Progress {
		return
	}

	if _, err := dao.TaskVideo.FindOneAndUpdate(ctx, bson.M{
		"_id":      taskVideo.Id,
		"status":   "in_progress",
		"progress": bson.M{"$not": bson.M{"$gte": target}},
	}, bson.M{
		"progress":   target,
		"updated_at": taskVideo.UpdatedAt,
	}); err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return
		}
		logger.Error(ctx, err)
	}
}

func (s *sTaskVideo) requeueTask(ctx context.Context, taskId string) {

	if _, err := dao.TaskVideo.FindOneAndUpdate(ctx, bson.M{
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

func (s *sTaskVideo) failTask(ctx context.Context, taskId, code, message string, responseData map[string]any, logVideoId ...string) {

	// 失败时也记录响应数据: 有上游响应则落上游原文, 否则按错误信息构造,
	// 便于任务查询接口直接返回落库结果, 无需再实时查询上游
	if len(responseData) == 0 {
		responseData = map[string]any{
			"status": "failed",
			"error": map[string]any{
				"code":    code,
				"message": message,
			},
		}
	}

	if _, err := dao.TaskVideo.FindOneAndUpdate(ctx, bson.M{
		"_id":    taskId,
		"status": "in_progress",
	}, bson.M{
		"status":        "failed",
		"error":         &smodel.VideoError{Code: code, Message: message},
		"response_data": responseData,
	}); err != nil {
		if !errors.Is(err, mongo.ErrNoDocuments) {
			logger.Error(ctx, err)
		}
		return
	}

	if len(logVideoId) > 0 && logVideoId[0] != "" {
		if err := dao.LogVideo.UpdateById(ctx, logVideoId[0], bson.M{
			"status":  -1,
			"err_msg": message,
		}); err != nil {
			logger.Error(ctx, err)
		}
	}
}

func (s *sTaskVideo) prepareUpstreamModel(ctx context.Context, taskVideo *entity.TaskVideo, logVideo *entity.LogVideo) string {

	requestModel := ""
	if taskVideo.RequestData != nil {
		if v, ok := taskVideo.RequestData["model"].(string); ok {
			requestModel = v
		}
	}

	model := s.resolveUpstreamModel(ctx, logVideo, requestModel)

	if taskVideo.RequestData != nil {
		taskVideo.RequestData["model"] = model
	}

	return model
}

func (s *sTaskVideo) resolveUpstreamModel(ctx context.Context, logVideo *entity.LogVideo, requestModel string) string {

	model := requestModel
	if model == "" && logVideo.RequestData != nil {
		if v, ok := logVideo.RequestData["model"].(string); ok {
			model = v
		}
	}
	if model == "" {
		model = logVideo.Model
	}

	if logVideo.RealModel != "" && !gstr.Contains(logVideo.RealModel, "*") {
		model = logVideo.RealModel
	}

	if logVideo.ModelAgentId == "" {
		return model
	}

	agent, err := dao.ModelAgent.FindById(ctx, logVideo.ModelAgentId)
	if err != nil {
		if !errors.Is(err, mongo.ErrNoDocuments) {
			logger.Error(ctx, err)
		}
		return model
	}

	if agent == nil || !agent.IsEnableModelReplace {
		return model
	}

	for i, replaceModel := range agent.ReplaceModels {
		if replaceModel == model {
			if i >= len(agent.TargetModels) {
				break
			}
			logger.Infof(ctx, "sTaskVideo resolveUpstreamModel request.Model: %s replaced %s", model, agent.TargetModels[i])
			model = agent.TargetModels[i]
			break
		}
	}

	return model
}

func resolveVideoUrl(videoUrl string) string {

	if videoUrl == "" {
		return ""
	}

	if config.Cfg.VideoTask.IsEnableStorage {
		return buildStorageUrl(videoUrl)
	}

	return videoUrl
}

func buildStorageUrl(videoUrl string) string {

	if videoUrl == "" {
		return ""
	}

	if config.Cfg.VideoTask.StorageBaseUrl != "" {
		if gstr.HasSuffix(config.Cfg.VideoTask.StorageBaseUrl, "/") {
			videoUrl = gstr.TrimLeftStr(videoUrl, "/")
		} else if !gstr.HasPrefix(videoUrl, "/") {
			videoUrl = "/" + videoUrl
		}
	}

	return config.Cfg.VideoTask.StorageBaseUrl + videoUrl
}

func taskVideoTotalTime(result *entity.TaskVideo) int64 {

	if result.CreatedAt <= 0 {
		return 0
	}

	var end int64
	switch result.Status {
	case "completed", "expired", "deleted":
		end = result.CompletedAt * 1000
	case "failed":
		end = result.UpdatedAt
	default:
		return 0
	}

	if end <= 0 {
		return 0
	}

	return end - result.CreatedAt
}

func pickBestLogVideo(logVideos []*entity.LogVideo) *entity.LogVideo {

	var best *entity.LogVideo

	for _, l := range logVideos {

		if best == nil {
			best = l
			continue
		}

		pl, pb := logVideoStatusPriority(l.Status), logVideoStatusPriority(best.Status)
		if pl < pb || (pl == pb && l.CreatedAt > best.CreatedAt) {
			best = l
		}
	}

	return best
}

func logVideoStatusPriority(status int) int {
	switch status {
	case 1:
		return 0
	case -1:
		return 1
	default:
		return 2
	}
}

func parseVideoCreateTaskId(data []byte) string {

	j := gjson.New(data)
	if id := j.Get("task_id").String(); id != "" {
		return id
	}
	if id := j.Get("request_id").String(); id != "" {
		return id
	}
	if id := j.Get("output.task_id").String(); id != "" {
		return id
	}
	return j.Get("id").String()
}

func pollVideoId(taskVideo *entity.TaskVideo) string {
	if handle := videoPollHandle(taskVideo); handle != "" {
		return handle
	}
	return taskVideo.VideoId
}

// video_id 是网关返回给客户端的查询 ID, 重提后保持不变; job_id 才是当前上游句柄
// "-" 表示已放弃旧句柄, 下一轮必须重新提交, 不能再回退到原 video_id
const videoJobResubmit = "-"

func isVideoJobHandle(id string) bool {
	return id != "" && id != videoJobResubmit
}

func videoPollHandle(taskVideo *entity.TaskVideo) string {
	if taskVideo == nil {
		return ""
	}
	if isVideoJobHandle(taskVideo.JobId) {
		return taskVideo.JobId
	}
	if taskVideo.JobId == videoJobResubmit {
		return ""
	}
	return taskVideo.VideoId
}

func abandonVideoJob(taskVideo *entity.TaskVideo) {
	if taskVideo != nil {
		taskVideo.JobId = videoJobResubmit
	}
}
