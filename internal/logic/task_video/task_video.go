package task_video

import (
	"context"
	"net/url"
	"regexp"
	"sync"
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

type sTaskVideo struct {
	videoRedsync *redsync.Redsync
	running      sync.Map
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

	taskVideo, err := dao.TaskVideo.FindById(ctx, id)
	if err != nil {
		logger.Error(ctx, err)
		return err
	}

	if taskVideo.Status != "in_progress" && taskVideo.Status != "failed" {
		return errors.New("任务不在进行中或已失败状态, 无法重新生成")
	}

	if taskVideo.TraceId != "" {
		if err = dao.LogVideo.UpdateMany(ctx, bson.M{
			"trace_id": taskVideo.TraceId,
			"$or": bson.A{
				bson.M{"action": bson.M{"$in": []string{"create", "remix"}}},
				bson.M{"status": bson.M{"$in": []int{1, -1}}},
			},
		}, bson.M{
			"status":  1,
			"err_msg": "",
		}); err != nil {
			logger.Error(ctx, err)
			return err
		}
	}

	update := bson.M{
		"status":   "queued",
		"progress": 0,
		"retry":    0,
		"error":    nil,
	}

	// 有原请求体则放弃旧句柄, 下一轮按网关同样方式重新提交; 旧任务缺 RequestData 时保留句柄, 仍按原 video_id 轮询
	if len(taskVideo.RequestData) > 0 {
		update["job_id"] = videoJobResubmit
	}

	if _, err = dao.TaskVideo.FindOneAndUpdate(ctx, bson.M{
		"_id": id,
		"status": bson.M{
			"$in": []string{"in_progress", "failed"},
		},
	}, update); err != nil {

		if errors.Is(err, mongo.ErrNoDocuments) {
			return errors.New("任务不在进行中或已失败状态, 无法重新生成")
		}

		logger.Error(ctx, err)
		return err
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

	taskVideos, err := dao.TaskVideo.Find(ctx, bson.M{"status": bson.M{"$in": []string{"queued", "in_progress"}}}, &dao.FindOptions{SortFields: []string{"created_at"}})
	if err != nil {
		logger.Error(ctx, err)
		return
	}

	availableSlots := -1
	if config.Cfg.VideoTask.ConcurrencyLimit > 0 {

		liveInProgress := 0

		for _, taskVideo := range taskVideos {
			if taskVideo.Status == "in_progress" {
				liveInProgress++
			}
		}

		if availableSlots = config.Cfg.VideoTask.ConcurrencyLimit - liveInProgress; availableSlots < 0 {
			availableSlots = 0
		}
	}

	// 已在跑的协程跳过; queued 受并发限制才提升; 已进行中的每轮只查/提交一次就退出
	var dispatched []*entity.TaskVideo

	for _, taskVideo := range taskVideos {

		if _, live := s.running.Load(taskVideo.Id); live {
			if taskVideo.Status == "in_progress" {
				s.advanceFakeProgress(ctx, now, taskVideo)
			}
			continue
		}

		if taskVideo.Status == "queued" {

			if availableSlots == 0 {
				continue
			}

			if err = dao.TaskVideo.UpdateById(ctx, taskVideo.Id, bson.M{"status": "in_progress", "error": nil}); err != nil {
				logger.Error(ctx, err)
				continue
			}

			taskVideo.Status = "in_progress"
			taskVideo.UpdatedAt = now

			if availableSlots > 0 {
				availableSlots--
			}
		}

		dispatched = append(dispatched, taskVideo)
	}

	for _, taskVideo := range dispatched {

		taskId := taskVideo.Id

		if _, loaded := s.running.LoadOrStore(taskId, struct{}{}); loaded {
			continue
		}

		tv := taskVideo

		if err := grpool.AddWithRecover(gctx.NeverDone(ctx), func(ctx context.Context) {
			defer s.running.Delete(tv.Id)
			ctx, _ = gtrace.WithTraceID(ctx, tv.TraceId)
			s.processVideoTask(ctx, tv)
		}, nil); err != nil {
			s.running.Delete(taskId)
			logger.Error(ctx, err)
		}
	}

	s.cleanExpiredAndFiles(ctx, now)

	if _, err := redis.Set(ctx, consts.TASK_VIDEO_END_TIME_KEY, gtime.TimestampMilli()); err != nil {
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

				for _, fp := range taskVideo.InputFilePaths {
					if fp != "" {
						if err := gfile.RemoveFile(fp); err != nil {
							logger.Error(ctx, err)
						}
					}
				}

				if len(taskVideo.InputFilePaths) > 0 {
					update["input_file_paths"] = nil
				}
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

		for _, fp := range taskVideo.InputFilePaths {
			if fp != "" {
				if err := gfile.RemoveFile(fp); err != nil {
					logger.Error(ctx, err)
				}
			}
		}

		if err := dao.TaskVideo.UpdateById(ctx, taskVideo.Id, bson.M{
			"video_url":        "",
			"file_name":        "",
			"file_path":        "",
			"input_file_paths": nil,
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

func (s *sTaskVideo) requeueTask(ctx context.Context, taskVideo *entity.TaskVideo) {

	if _, err := dao.TaskVideo.FindOneAndUpdate(ctx, videoWorkerFilter(taskVideo), bson.M{
		"status": "queued",
	}); err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return
		}
		logger.Error(ctx, err)
	}
}

func (s *sTaskVideo) failTask(ctx context.Context, taskVideo *entity.TaskVideo, code, message string, responseData map[string]any, logVideoId ...string) {

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

	if _, err := dao.TaskVideo.FindOneAndUpdate(ctx, videoWorkerFilter(taskVideo), bson.M{
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
		return common.ReplaceVideoUrl(buildStorageUrl(videoUrl))
	}

	return common.ReplaceVideoUrl(videoUrl)
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

func isLocalVideoId(id string) bool {
	return gstr.HasPrefix(id, "video_")
}

func videoPollHandle(taskVideo *entity.TaskVideo) string {

	if taskVideo == nil {
		return ""
	}

	if isVideoJobHandle(taskVideo.JobId) {
		if isLocalVideoId(taskVideo.JobId) && taskVideo.JobId == taskVideo.VideoId {
			return ""
		}
		return taskVideo.JobId
	}

	if taskVideo.JobId == videoJobResubmit {
		return ""
	}

	if isLocalVideoId(taskVideo.VideoId) {
		return ""
	}

	return taskVideo.VideoId
}

func (s *sTaskVideo) submitVideoJob(ctx context.Context, taskVideo *entity.TaskVideo, logVideo *entity.LogVideo, provider *entity.Provider, timeout time.Duration) (smodel.VideoJobResponse, string, error) {

	var created smodel.VideoJobResponse

	if len(taskVideo.RequestData) == 0 {
		return created, "missing_request_data", errors.New("missing request data")
	}

	upstreamModel := s.prepareUpstreamModel(ctx, taskVideo, logVideo)

	adapterOpts := &options.AdapterOptions{
		Provider: provider.Code,
		Model:    upstreamModel,
		Key:      logVideo.Key,
		BaseUrl:  logVideo.ModelAgent.BaseUrl,
		Path:     logVideo.ModelAgent.Path,
		Timeout:  timeout,
		ProxyUrl: config.Cfg.Http.ProxyUrl,
	}

	submitErrCode := func(err error) string {
		if ctx.Err() != nil {
			return "timeout"
		}
		return "generation_error"
	}

	// 网关 video.go 落库的是 OpenAI 风格请求体, 重提必须走 VideoCreate 做厂商转换;
	// 百炼/火山等官方接口落库的是原生请求体, 才走 Official 透传.
	if !isOpenAIStyleVideoRequest(taskVideo.RequestData) {

		adapter := sdk.NewAdapterOfficial(ctx, adapterOpts)

		bytes, _, err := adapter.VideoCreateOfficial(ctx, gjson.MustEncode(taskVideo.RequestData))
		if err != nil {
			return created, submitErrCode(err), err
		}

		created.ResponseBytes = bytes
		created.Id = parseVideoCreateTaskId(bytes)

		if created.Id == "" {
			return created, "submit_response_invalid", errors.New("create video task failed: missing task_id")
		}

		return created, "", nil
	}

	adapter := sdk.NewAdapter(ctx, adapterOpts)

	var err error
	if taskVideo.Action == "remix" {

		var req smodel.VideoRemixRequest

		if e := gconv.Scan(taskVideo.RequestData, &req); e != nil {
			return created, "invalid_request_data", e
		}

		if up, ok := taskVideo.RequestData["upstream_video_id"].(string); ok && up != "" {
			req.VideoId = up
		}

		created, err = adapter.VideoRemix(ctx, req)

	} else {

		var req smodel.VideoCreateRequest

		if e := gconv.Scan(taskVideo.RequestData, &req); e != nil {
			return created, "invalid_request_data", e
		}

		req.Model = upstreamModel
		created, err = adapter.VideoCreate(ctx, req)
	}

	if err != nil {
		return created, submitErrCode(err), err
	}

	if created.Id == "" {
		created.Id = parseVideoCreateTaskId(created.ResponseBytes)
	}

	if created.Id == "" {
		return created, "submit_response_invalid", errors.New("create video task failed: missing task_id")
	}

	return created, "", nil
}

func canResubmitVideo(taskVideo *entity.TaskVideo) bool {
	return taskVideo != nil && len(taskVideo.RequestData) > 0 && videoPollHandle(taskVideo) == ""
}

func isOpenAIStyleVideoRequest(data map[string]any) bool {

	if data == nil {
		return false
	}

	if prompt, ok := data["prompt"].(string); ok && prompt != "" {
		return true
	}

	if _, ok := data["input"]; ok {
		return false
	}

	if _, ok := data["content"]; ok {
		return false
	}

	_, hasVideoId := data["video_id"]
	return hasVideoId
}

func videoWorkerFilter(taskVideo *entity.TaskVideo) bson.M {

	filter := bson.M{
		"_id":    taskVideo.Id,
		"status": "in_progress",
	}

	switch {
	case isVideoJobHandle(taskVideo.JobId):
		filter["job_id"] = taskVideo.JobId
	case taskVideo.JobId == videoJobResubmit:
		filter["job_id"] = videoJobResubmit
	default:
		filter["job_id"] = bson.M{"$nin": []string{videoJobResubmit}}
	}

	return filter
}
