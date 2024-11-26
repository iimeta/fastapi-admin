package audio

import (
	"context"
	"github.com/gogf/gf/v2/os/gtime"
	"github.com/gogf/gf/v2/text/gstr"
	"github.com/iimeta/fastapi-admin/internal/config"
	"github.com/iimeta/fastapi-admin/internal/dao"
	"github.com/iimeta/fastapi-admin/internal/errors"
	"github.com/iimeta/fastapi-admin/internal/model"
	"github.com/iimeta/fastapi-admin/internal/service"
	"github.com/iimeta/fastapi-admin/utility/db"
	"github.com/iimeta/fastapi-admin/utility/logger"
	"github.com/iimeta/fastapi-admin/utility/util"
	"go.mongodb.org/mongo-driver/bson"
	"time"
)

type sAudio struct{}

func init() {
	service.RegisterAudio(New())
}

func New() service.IAudio {
	return &sAudio{}
}

// 音频日志详情
func (s *sAudio) Detail(ctx context.Context, id string) (*model.Audio, error) {

	result, err := dao.Audio.FindById(ctx, id)
	if err != nil {
		logger.Error(ctx, err)
		return nil, err
	}

	if service.Session().IsUserRole(ctx) && result.UserId != service.Session().GetUserId(ctx) {
		return nil, errors.ERR_UNAUTHORIZED
	}

	corpName := result.Corp
	if corp, err := dao.Corp.FindById(ctx, result.Corp); err == nil && corp != nil {
		corpName = corp.Name
	}

	audio := &model.Audio{
		Id:          result.Id,
		TraceId:     result.TraceId,
		UserId:      result.UserId,
		AppId:       result.AppId,
		Corp:        result.Corp,
		CorpName:    corpName,
		Model:       result.Model,
		Type:        result.Type,
		Input:       result.Input,
		Text:        result.Text,
		Characters:  result.Characters,
		Minute:      result.Minute,
		AudioQuota:  result.AudioQuota,
		TotalTokens: result.TotalTokens,
		TotalTime:   result.TotalTime,
		ReqTime:     util.FormatDateTime(result.ReqTime),
		ClientIp:    result.ClientIp,
		Retry:       result.Retry,
		Status:      result.Status,
		Host:        result.Host,
		Creator:     util.Desensitize(result.Creator),
	}

	if audio.Status == -1 && service.Session().IsUserRole(ctx) {
		audio.ErrMsg = "详细错误信息请联系管理员..."
		if len(config.Cfg.Error.ShieldUser) > 0 {
			audio.ErrMsg = result.ErrMsg
			for _, shieldError := range config.Cfg.Error.ShieldUser {
				if gstr.Contains(result.ErrMsg, shieldError) {
					audio.ErrMsg = "详细错误信息请联系管理员..."
					break
				}
			}
		}
	}

	if service.Session().IsAdminRole(ctx) {

		audio.ModelId = result.ModelId
		audio.Name = result.Name
		audio.Key = util.Desensitize(result.Key)
		audio.IsEnablePresetConfig = result.IsEnablePresetConfig
		audio.IsEnableModelAgent = result.IsEnableModelAgent
		audio.ModelAgentId = result.ModelAgentId
		audio.IsEnableForward = result.IsEnableForward
		audio.ForwardConfig = result.ForwardConfig
		audio.IsSmartMatch = result.IsSmartMatch
		audio.IsEnableFallback = result.IsEnableFallback
		audio.FallbackConfig = result.FallbackConfig
		audio.RealModelId = result.RealModelId
		audio.RealModelName = result.RealModelName
		audio.RealModel = result.RealModel
		audio.RemoteIp = result.RemoteIp
		audio.LocalIp = result.LocalIp
		audio.InternalTime = result.InternalTime
		audio.ErrMsg = result.ErrMsg
		audio.IsRetry = result.IsRetry
		audio.CreatedAt = util.FormatDateTime(result.CreatedAt)
		audio.UpdatedAt = util.FormatDateTime(result.UpdatedAt)

		if result.ModelAgent != nil {

			corpName := result.ModelAgent.Corp
			if corp, err := dao.Corp.FindById(ctx, result.ModelAgent.Corp); err == nil && corp != nil {
				corpName = corp.Name
			}

			audio.ModelAgent = &model.ModelAgent{
				Corp:     result.ModelAgent.Corp,
				CorpName: corpName,
				Name:     result.ModelAgent.Name,
				BaseUrl:  result.ModelAgent.BaseUrl,
				Path:     result.ModelAgent.Path,
				Weight:   result.ModelAgent.Weight,
				Remark:   result.ModelAgent.Remark,
				Status:   result.ModelAgent.Status,
			}
		}

		if audio.Type == 6 {
			audio.Input = result.FilePath
		}
	}

	return audio, nil
}

// 音频日志分页列表
func (s *sAudio) Page(ctx context.Context, params model.AudioPageReq) (*model.AudioPageRes, error) {

	paging := &db.Paging{
		Page:     params.Page,
		PageSize: params.PageSize,
	}

	filter := bson.M{}

	if params.TraceId != "" {
		filter["trace_id"] = params.TraceId
	}

	if service.Session().IsUserRole(ctx) {
		filter["user_id"] = service.Session().GetUserId(ctx)
		filter["is_retry"] = bson.M{"$exists": false}
	} else if params.UserId != 0 {
		filter["user_id"] = params.UserId
	}

	if params.AppId != 0 {
		filter["app_id"] = params.AppId
	}

	if params.Key != "" {
		if service.Session().IsAdminRole(ctx) {
			filter["key"] = params.Key
		} else {
			filter["creator"] = params.Key
		}
	}

	if len(params.Models) > 0 {
		filter["model_id"] = bson.M{
			"$in": params.Models,
		}
	}

	if len(params.ModelAgents) > 0 && service.Session().IsAdminRole(ctx) {
		filter["model_agent_id"] = bson.M{
			"$in": params.ModelAgents,
		}
	}

	if params.TotalTime != 0 {
		filter["total_time"] = bson.M{
			"$gte": params.TotalTime,
		}
	}

	if params.Status != 0 {
		filter["status"] = params.Status
	}

	if params.Status == -100 {
		filter["status"] = bson.M{"$ne": 1}
	}

	if len(params.ReqTime) > 0 && params.TraceId == "" {
		gte := gtime.NewFromStrFormat(params.ReqTime[0], time.DateTime).TimestampMilli()
		lte := gtime.NewFromStrLayout(params.ReqTime[1], time.DateTime).TimestampMilli() + 999
		filter["req_time"] = bson.M{
			"$gte": gte,
			"$lte": lte,
		}
	}

	results, err := dao.Audio.FindByPage(ctx, paging, filter, "", "-req_time", "status", "-created_at")
	if err != nil {
		logger.Error(ctx, err)
		return nil, err
	}

	items := make([]*model.Audio, 0)
	for _, result := range results {

		audio := &model.Audio{
			Id:          result.Id,
			UserId:      result.UserId,
			AppId:       result.AppId,
			Corp:        result.Corp,
			Model:       result.Model,
			Characters:  result.Characters,
			Minute:      result.Minute,
			TotalTokens: result.TotalTokens,
			TotalTime:   result.TotalTime,
			ReqTime:     util.FormatDateTimeMonth(result.ReqTime),
			Status:      result.Status,
		}

		if service.Session().IsAdminRole(ctx) {
			audio.InternalTime = result.InternalTime
			audio.IsSmartMatch = result.IsSmartMatch
		}

		items = append(items, audio)
	}

	return &model.AudioPageRes{
		Items: items,
		Paging: &model.Paging{
			Page:     paging.Page,
			PageSize: paging.PageSize,
			Total:    paging.Total,
		},
	}, nil
}

// 音频日志详情复制字段值
func (s *sAudio) CopyField(ctx context.Context, params model.AudioCopyFieldReq) (string, error) {

	result, err := dao.Audio.FindById(ctx, params.Id)
	if err != nil {
		logger.Error(ctx, err)
		return "", err
	}

	if service.Session().IsUserRole(ctx) && result.UserId != service.Session().GetUserId(ctx) {
		return "", errors.ERR_UNAUTHORIZED
	}

	switch params.Field {
	case "key":
		return result.Key, nil
	case "creator":
		return result.Creator, nil
	}

	return "", nil
}
