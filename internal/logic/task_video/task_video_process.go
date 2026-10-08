package task_video

import (
	"context"
	"encoding/json"
	"time"

	"github.com/gogf/gf/v2/os/gctx"
	"github.com/gogf/gf/v2/os/gfile"
	"github.com/gogf/gf/v2/os/grpool"
	"github.com/gogf/gf/v2/os/gtime"
	"github.com/gogf/gf/v2/text/gstr"
	"github.com/gogf/gf/v2/util/gconv"
	"github.com/iimeta/fastapi-admin/v2/internal/config"
	"github.com/iimeta/fastapi-admin/v2/internal/dao"
	"github.com/iimeta/fastapi-admin/v2/internal/errors"
	"github.com/iimeta/fastapi-admin/v2/internal/logic/common"
	"github.com/iimeta/fastapi-admin/v2/internal/model/entity"
	"github.com/iimeta/fastapi-admin/v2/internal/service"
	"github.com/iimeta/fastapi-admin/v2/utility/logger"
	"github.com/iimeta/fastapi-admin/v2/utility/util"
	sdk "github.com/iimeta/fastapi-sdk/v2"
	sconsts "github.com/iimeta/fastapi-sdk/v2/consts"
	smodel "github.com/iimeta/fastapi-sdk/v2/model"
	"github.com/iimeta/fastapi-sdk/v2/options"
	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

// 每轮 Task 只走一步: 无上游句柄则提交, 有句柄则查询一次, 然后退出等下一轮
func (s *sTaskVideo) processVideoTask(ctx context.Context, taskVideo *entity.TaskVideo) {

	logVideos, err := dao.LogVideo.Find(ctx, bson.M{
		"trace_id": taskVideo.TraceId,
		"action":   bson.M{"$in": []string{"create", "remix"}},
		"status":   bson.M{"$in": []int{1, -1}},
	}, &dao.FindOptions{SortFields: []string{"-created_at"}})
	if err != nil {
		logger.Error(ctx, err)
		s.failTask(ctx, taskVideo, "log_not_found", err.Error(), nil)
		return
	}

	if len(logVideos) == 0 {
		logVideos, err = dao.LogVideo.Find(ctx, bson.M{
			"trace_id": taskVideo.TraceId,
			"status":   bson.M{"$in": []int{1, -1}},
		}, &dao.FindOptions{SortFields: []string{"-created_at"}})
		if err != nil {
			logger.Error(ctx, err)
			s.failTask(ctx, taskVideo, "log_not_found", err.Error(), nil)
			return
		}
	}

	logVideo := pickBestLogVideo(logVideos)

	if logVideo == nil {
		s.requeueTask(ctx, taskVideo)
		return
	}

	if logVideo.Status == -1 && !canResubmitVideo(taskVideo) {
		logger.Infof(ctx, "sTaskVideo processVideoTask task: %s all log_video failed, mark task failed directly", taskVideo.Id)
		s.failTask(ctx, taskVideo, "log_failed", logVideo.ErrMsg, nil, logVideo.Id)
		return
	}

	if logVideo.Status == -1 {
		logger.Infof(ctx, "sTaskVideo processVideoTask task: %s log_video failed but request_data present, resubmit instead of log_failed", taskVideo.Id)
	}

	if logVideo.ModelAgent == nil {
		logger.Errorf(ctx, "sTaskVideo processVideoTask task: %s log_video: %s has no model_agent", taskVideo.Id, logVideo.Id)
		s.failTask(ctx, taskVideo, "model_agent_not_found", "log_video has no model_agent", nil, logVideo.Id)
		return
	}

	provider, err := dao.Provider.FindById(ctx, logVideo.ModelAgent.ProviderId)
	if err != nil {
		logger.Error(ctx, err)
		s.failTask(ctx, taskVideo, "provider_not_found", err.Error(), nil, logVideo.Id)
		return
	}

	if isAttemptTimeout(taskVideo) {
		logger.Errorf(ctx, "sTaskVideo processVideoTask task: %s attempt timeout, job_id: %s", taskVideo.Id, taskVideo.JobId)
		s.retryOrFail(ctx, taskVideo, logVideo, "timeout", errors.New("generation timeout"), nil, true)
		return
	}

	httpTimeout := videoHTTPTimeout()

	if videoPollHandle(taskVideo) == "" {
		s.submitOnce(ctx, taskVideo, logVideo, provider, httpTimeout)
		return
	}

	s.pollOnce(ctx, taskVideo, logVideo, provider, httpTimeout)
}

func (s *sTaskVideo) submitOnce(ctx context.Context, taskVideo *entity.TaskVideo, logVideo *entity.LogVideo, provider *entity.Provider, timeout time.Duration) {

	if taskVideo.Retry > 0 {

		agentChanged, e := s.switchRetryAgent(ctx, taskVideo, logVideo, []string{logVideo.ModelAgentId}, []string{logVideo.Key}, taskVideo.Retry)
		if e != nil {
			logger.Errorf(ctx, "sTaskVideo submitOnce task: %s switch agent failed: %v, skip", taskVideo.Id, e)
			return
		}

		if agentChanged && logVideo.ModelAgent != nil {

			p, err := dao.Provider.FindById(ctx, logVideo.ModelAgent.ProviderId)
			if err != nil {
				logger.Error(ctx, err)
				s.failTask(ctx, taskVideo, "provider_not_found", err.Error(), nil, logVideo.Id)
				return
			}

			provider = p
		}
	}

	taskCtx, cancel := context.WithTimeout(ctx, timeout)
	created, errCode, err := s.submitVideoJob(taskCtx, taskVideo, logVideo, provider, timeout)

	cancel()

	if err != nil {
		s.retryOrFail(ctx, taskVideo, logVideo, errCode, err, util.ConvToMap(created.ResponseBytes), true)
		return
	}

	jobId := created.Id
	if jobId == "" || (isLocalVideoId(jobId) && jobId == taskVideo.VideoId) {
		s.retryOrFail(ctx, taskVideo, logVideo, "submit_response_invalid", errors.New("create video task failed: missing or local task_id"), util.ConvToMap(created.ResponseBytes), true)
		return
	}

	if !s.persistJobId(ctx, taskVideo, jobId, false) {
		return
	}

	if created.Status == "failed" || created.Status == "expired" || created.Status == "deleted" || created.Status == "cancelled" {

		message := created.Status

		if created.Error != nil && created.Error.Message != "" {
			message = created.Error.Message
		}

		s.failTask(ctx, taskVideo, "async_"+created.Status, message, util.ConvToMap(created.ResponseBytes), logVideo.Id)

		return
	}

	if created.Status == "completed" {
		s.completeOnce(ctx, taskVideo, logVideo, provider, created, timeout)
	}
}

func (s *sTaskVideo) pollOnce(ctx context.Context, taskVideo *entity.TaskVideo, logVideo *entity.LogVideo, provider *entity.Provider, timeout time.Duration) {

	jobId := videoPollHandle(taskVideo)
	if jobId == "" {
		return
	}

	if !isVideoJobHandle(taskVideo.JobId) {
		if !s.persistJobId(ctx, taskVideo, jobId, true) {
			return
		}
	}

	taskCtx, cancel := context.WithTimeout(ctx, timeout)
	job, err := s.retrieveVideoJob(taskCtx, logVideo, provider, jobId, timeout)

	cancel()

	if err != nil {
		logger.Errorf(ctx, "sTaskVideo pollOnce task: %s retrieve job %s error: %v, wait next tick", taskVideo.Id, jobId, err)
		s.advanceFakeProgress(ctx, gtime.TimestampMilli(), taskVideo)
		return
	}

	switch job.Status {
	case "completed":
		s.completeOnce(ctx, taskVideo, logVideo, provider, job, timeout)
	case "failed", "expired", "deleted", "cancelled":

		message := job.Status
		if job.Error != nil && job.Error.Message != "" {
			message = job.Error.Message
		}

		s.failTask(ctx, taskVideo, "async_"+job.Status, message, util.ConvToMap(job.ResponseBytes), logVideo.Id)

	case "queued", "in_progress", "running":
		s.updateRunningProgress(ctx, taskVideo, job.Progress)
		s.advanceFakeProgress(ctx, gtime.TimestampMilli(), taskVideo)
	default:
		logger.Errorf(ctx, "sTaskVideo pollOnce task: %s job %s unknown status: %s, wait next tick", taskVideo.Id, jobId, job.Status)
		s.advanceFakeProgress(ctx, gtime.TimestampMilli(), taskVideo)
	}
}

func (s *sTaskVideo) completeOnce(ctx context.Context, taskVideo *entity.TaskVideo, logVideo *entity.LogVideo, provider *entity.Provider, retrieve smodel.VideoJobResponse, timeout time.Duration) {

	videoUrl, fileName, filePath, errCode, err := s.storeCompletedMedia(ctx, taskVideo, logVideo, provider, retrieve, timeout)
	if err != nil {
		s.retryOrFail(ctx, taskVideo, logVideo, errCode, err, util.ConvToMap(retrieve.ResponseBytes), false)
		return
	}

	s.finishCompleted(ctx, taskVideo, logVideo, provider, retrieve, videoUrl, fileName, filePath)
}

func (s *sTaskVideo) storeCompletedMedia(ctx context.Context, taskVideo *entity.TaskVideo, logVideo *entity.LogVideo, provider *entity.Provider, retrieve smodel.VideoJobResponse, timeout time.Duration) (videoUrl, fileName, filePath, errCode string, err error) {

	if config.Cfg.VideoTask.IsEnableStorage {

		var content []byte

		content, err = s.downloadVideo(ctx, logVideo, provider, taskVideo, timeout)
		if err != nil {
			logger.Error(ctx, err)
			return "", "", "", "storage_failed", err
		}

		if len(content) == 0 {
			return "", "", "", "storage_failed", errors.New("downloaded video content is empty")
		}

		storageDir := config.Cfg.VideoTask.StorageDir

		if storageDir == "" {
			storageDir = "./resource/public/video/"
		} else if !gstr.HasSuffix(storageDir, "/") {
			storageDir = storageDir + "/"
		}

		fileName = taskVideo.VideoId + "_video.mp4"

		if err = gfile.PutBytes(storageDir+fileName, content); err != nil {
			logger.Error(ctx, err)
			return "", "", "", "storage_failed", err
		}

		filePath = storageDir + fileName

		if gstr.HasPrefix(storageDir, "./resource/public/") {
			videoUrl = "/public/" + gstr.TrimLeftStr(storageDir, "./resource/public/") + fileName
		} else if config.Cfg.VideoTask.StorageBaseUrl == "" {
			videoUrl = "/open/video/" + fileName
		} else {
			videoUrl = fileName
		}

		return videoUrl, fileName, filePath, "", nil
	}

	if retrieve.VideoUrl != "" {
		return retrieve.VideoUrl, "", "", "", nil
	}

	return "", "", "", "no_video", errors.New("no video in completed job")
}

func (s *sTaskVideo) finishCompleted(ctx context.Context, taskVideo *entity.TaskVideo, logVideo *entity.LogVideo, provider *entity.Provider, retrieve smodel.VideoJobResponse, videoUrl, fileName, filePath string) {

	completedAt := gtime.TimestampMilli() / 1000
	var expiresAt int64

	if config.Cfg.VideoTask.IsEnableStorage {
		if videoUrl != "" && config.Cfg.VideoTask.StorageExpiresAt > 0 {
			expiresAt = gtime.NewFromTimeStamp(completedAt).Add(config.Cfg.VideoTask.StorageExpiresAt * time.Minute).Unix()
		}
	} else if retrieve.ExpiresAt != nil {
		expiresAt = *retrieve.ExpiresAt
	} else if config.Cfg.VideoTask.StorageExpiresAt > 0 {
		expiresAt = gtime.NewFromTimeStamp(completedAt).Add(config.Cfg.VideoTask.StorageExpiresAt * time.Minute).Unix()
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

	if _, err := dao.TaskVideo.FindOneAndUpdate(ctx, videoWorkerFilter(taskVideo), update); err != nil {

		if errors.Is(err, mongo.ErrNoDocuments) {
			logger.Infof(ctx, "sTaskVideo finishCompleted task: %s already handled by another worker, skip", taskVideo.Id)
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

	if retrieve.Usage != nil {
		common.Billing(ctx, *retrieve.Usage, &logVideo.Spend)
	}

	if provider.Code == sconsts.PROVIDER_BAILIAN && logVideo.Spend.VideoGeneration != nil && logVideo.Spend.VideoGeneration.Seconds == 0 {
		if seconds := gconv.Int(retrieve.Seconds); seconds > 0 {
			common.RecalcVideoSecondsSpend(&logVideo.Spend, seconds)
			if err := dao.TaskVideo.UpdateById(ctx, taskVideo.Id, bson.M{"seconds": seconds}); err != nil {
				logger.Error(ctx, err)
			}
		}
	}

	if err := common.RecordSpend(ctx, logVideo.UserId, logVideo.AppId, logVideo.Creator, logVideo.Rid, logVideo.Key, logVideo.Spend); err != nil {
		logger.Error(ctx, err)
		return
	}

	if err := dao.LogVideo.UpdateById(ctx, logVideo.Id, bson.M{
		"spend":          logVideo.Spend,
		"model_agent_id": logVideo.ModelAgentId,
		"model_agent":    logVideo.ModelAgent,
		"key":            logVideo.Key,
	}); err != nil {
		logger.Error(ctx, err)
	}
}

func (s *sTaskVideo) persistJobId(ctx context.Context, taskVideo *entity.TaskVideo, jobId string, preserveUpdatedAt bool) bool {

	filter := videoWorkerFilter(taskVideo)
	update := bson.M{"job_id": jobId}
	if preserveUpdatedAt && taskVideo.UpdatedAt > 0 {
		update["updated_at"] = taskVideo.UpdatedAt
	}

	if _, err := dao.TaskVideo.FindOneAndUpdate(ctx, filter, update); err != nil {

		if errors.Is(err, mongo.ErrNoDocuments) {
			logger.Infof(ctx, "sTaskVideo persistJobId task: %s already handled, skip", taskVideo.Id)
			return false
		}

		logger.Error(ctx, err)
		return false
	}

	taskVideo.JobId = jobId
	if !preserveUpdatedAt {
		taskVideo.UpdatedAt = gtime.TimestampMilli()
	}

	return true
}

func (s *sTaskVideo) retryOrFail(ctx context.Context, taskVideo *entity.TaskVideo, logVideo *entity.LogVideo, errCode string, err error, responseData map[string]any, abandon bool) {

	message := ""
	if err != nil {
		message = err.Error()
	}

	if err != nil {
		isRetry, isDisabled := common.IsNeedRetry(err)

		if isDisabled && logVideo != nil && logVideo.Key != "" {
			if e := grpool.AddWithRecover(gctx.NeverDone(ctx), func(ctx context.Context) {
				service.Key().AutoDisabled(ctx, logVideo.Key, err.Error())
			}, nil); e != nil {
				logger.Error(ctx, e)
			}
		}

		if errCode != "timeout" && !isRetry {
			logger.Errorf(ctx, "sTaskVideo retryOrFail task: %s failed: %s, error: %v, no need to retry by config", taskVideo.Id, errCode, err)

			logId := ""
			if logVideo != nil {
				logId = logVideo.Id
			}

			s.failTask(ctx, taskVideo, errCode, message, responseData, logId)
			return
		}
	}

	retryCount := config.Cfg.VideoTask.RetryCount
	if retryCount < 0 {
		retryCount = 0
	}

	logId := ""
	if logVideo != nil {
		logId = logVideo.Id
	}

	if taskVideo.Retry >= retryCount {
		logger.Error(ctx, err)
		s.failTask(ctx, taskVideo, errCode, message, responseData, logId)
		return
	}

	if latest, e := dao.TaskVideo.FindById(ctx, taskVideo.Id); e != nil {
		logger.Error(ctx, e)
	} else if latest.Status != "in_progress" {
		logger.Infof(ctx, "sTaskVideo retryOrFail task: %s status is %s, skip", taskVideo.Id, latest.Status)
		return
	} else if isVideoJobHandle(taskVideo.JobId) && latest.JobId != taskVideo.JobId {
		logger.Infof(ctx, "sTaskVideo retryOrFail task: %s job handle changed (%s -> %s), skip", taskVideo.Id, taskVideo.JobId, latest.JobId)
		return
	}

	filter := videoWorkerFilter(taskVideo)

	update := bson.M{"retry": taskVideo.Retry + 1}
	if abandon {
		update["job_id"] = videoJobResubmit
	}

	logger.Errorf(ctx, "sTaskVideo retryOrFail task: %s failed: %s, retry: %d/%d, abandon: %v", taskVideo.Id, errCode, taskVideo.Retry+1, retryCount, abandon)

	if _, e := dao.TaskVideo.FindOneAndUpdate(ctx, filter, update); e != nil {
		if !errors.Is(e, mongo.ErrNoDocuments) {
			logger.Error(ctx, e)
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

func (s *sTaskVideo) updateRunningProgress(ctx context.Context, taskVideo *entity.TaskVideo, progress int) {

	if progress <= 0 {
		return
	}

	if progress <= 1 {
		progress = progress * 100
	}

	if progress >= 100 {
		progress = 99
	}

	if progress > taskVideo.Progress {
		taskVideo.Progress = progress
	}

	filter := videoWorkerFilter(taskVideo)
	filter["progress"] = bson.M{"$not": bson.M{"$gte": progress}}

	update := bson.M{"progress": progress}
	if taskVideo.UpdatedAt > 0 {
		update["updated_at"] = taskVideo.UpdatedAt
	}

	if _, err := dao.TaskVideo.FindOneAndUpdate(ctx, filter, update); err != nil && !errors.Is(err, mongo.ErrNoDocuments) {
		logger.Error(ctx, err)
	}
}

func videoHTTPTimeout() time.Duration {

	if config.Cfg.Http.Timeout > 0 {
		return config.Cfg.Http.Timeout * time.Second
	}

	if config.Cfg.Base.LongTimeout > 0 {
		return config.Cfg.Base.LongTimeout * time.Second
	}

	return 60 * time.Second
}

func videoAttemptTimeout() time.Duration {

	if config.Cfg.VideoTask.Timeout > 0 {
		return config.Cfg.VideoTask.Timeout * time.Second
	}

	if config.Cfg.Base.LongTimeout > 0 {
		return config.Cfg.Base.LongTimeout * time.Second
	}

	return 0
}

func isAttemptTimeout(taskVideo *entity.TaskVideo) bool {

	timeout := videoAttemptTimeout()

	if timeout <= 0 || taskVideo == nil || taskVideo.UpdatedAt <= 0 {
		return false
	}

	return gtime.TimestampMilli()-taskVideo.UpdatedAt >= timeout.Milliseconds()
}
