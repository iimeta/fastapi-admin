package chat

import (
	"context"
	"github.com/gogf/gf/v2/os/gtime"
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

type sChat struct{}

func init() {
	service.RegisterChat(New())
}

func New() service.IChat {
	return &sChat{}
}

// 聊天日志详情
func (s *sChat) Detail(ctx context.Context, id string) (*model.Chat, error) {

	result, err := dao.Chat.FindById(ctx, id)
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

	chat := &model.Chat{
		Id:               result.Id,
		TraceId:          result.TraceId,
		UserId:           result.UserId,
		AppId:            result.AppId,
		Corp:             result.Corp,
		CorpName:         corpName,
		Model:            result.Model,
		Type:             result.Type,
		Stream:           result.Stream,
		Messages:         result.Messages,
		Prompt:           result.Prompt,
		Completion:       result.Completion,
		TextQuota:        result.TextQuota,
		ImageQuotas:      result.ImageQuotas,
		MultimodalQuota:  result.MultimodalQuota,
		PromptTokens:     result.PromptTokens,
		CompletionTokens: result.CompletionTokens,
		TotalTokens:      result.TotalTokens,
		ConnTime:         result.ConnTime,
		Duration:         result.Duration,
		TotalTime:        result.TotalTime,
		ReqTime:          util.FormatDateTime(result.ReqTime),
		ClientIp:         result.ClientIp,
		Retry:            result.Retry,
		Status:           result.Status,
		Creator:          util.Desensitize(result.Creator),
	}

	// todo
	if chat.Status == -1 && service.Session().IsUserRole(ctx) {
		chat.ErrMsg = "详细错误信息请联系管理员..."
	}

	if service.Session().IsAdminRole(ctx) {

		chat.ModelId = result.ModelId
		chat.Name = result.Name
		chat.Key = util.Desensitize(result.Key)
		chat.IsEnablePresetConfig = result.IsEnablePresetConfig
		chat.IsEnableModelAgent = result.IsEnableModelAgent
		chat.ModelAgentId = result.ModelAgentId
		chat.IsEnableForward = result.IsEnableForward
		chat.ForwardConfig = result.ForwardConfig
		chat.IsSmartMatch = result.IsSmartMatch
		chat.IsEnableFallback = result.IsEnableFallback
		chat.FallbackConfig = result.FallbackConfig
		chat.RealModelId = result.RealModelId
		chat.RealModelName = result.RealModelName
		chat.RealModel = result.RealModel
		chat.RemoteIp = result.RemoteIp
		chat.LocalIp = result.LocalIp
		chat.InternalTime = result.InternalTime
		chat.ErrMsg = result.ErrMsg
		chat.IsRetry = result.IsRetry
		chat.CreatedAt = util.FormatDateTime(result.CreatedAt)
		chat.UpdatedAt = util.FormatDateTime(result.UpdatedAt)

		if result.ModelAgent != nil {

			corpName := result.ModelAgent.Corp
			if corp, err := dao.Corp.FindById(ctx, result.ModelAgent.Corp); err == nil && corp != nil {
				corpName = corp.Name
			}

			chat.ModelAgent = &model.ModelAgent{
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
	}

	return chat, nil
}

// 聊天日志分页列表
func (s *sChat) Page(ctx context.Context, params model.ChatPageReq) (*model.ChatPageRes, error) {

	paging := &db.Paging{
		Page:     params.Page,
		PageSize: params.PageSize,
	}

	filter := bson.M{}

	if len(params.ReqTime) > 0 {
		gte := gtime.NewFromStrFormat(params.ReqTime[0], time.DateTime).TimestampMilli()
		lte := gtime.NewFromStrLayout(params.ReqTime[1], time.DateTime).TimestampMilli() + 999
		filter["req_time"] = bson.M{
			"$gte": gte,
			"$lte": lte,
		}
	}

	if params.TraceId != "" {
		filter["trace_id"] = params.TraceId
	}

	if len(params.Models) > 0 {
		filter["model_id"] = bson.M{
			"$in": params.Models,
		}
	}

	if service.Session().IsUserRole(ctx) {
		filter["user_id"] = service.Session().GetUserId(ctx)
		filter["is_smart_match"] = bson.M{"$exists": false}
		filter["is_retry"] = bson.M{"$exists": false}
	} else if params.UserId != 0 {
		filter["user_id"] = params.UserId
	}

	if params.Status != 0 {
		filter["status"] = params.Status
	}

	if params.Status == -100 {
		filter["status"] = bson.M{"$ne": 1}
	}

	if params.AppId != 0 {
		filter["app_id"] = params.AppId
	}

	if params.Key != "" {
		filter["creator"] = params.Key
	}

	if params.TotalTime != 0 {
		filter["total_time"] = bson.M{
			"$gte": params.TotalTime,
		}
	}

	results, err := dao.Chat.FindByPage(ctx, paging, filter, "req_time_-1_model_id_1_user_id_1_status_1_created_at_-1", "-req_time", "status", "-created_at")
	if err != nil {
		logger.Error(ctx, err)
		return nil, err
	}

	items := make([]*model.Chat, 0)
	for _, result := range results {

		chat := &model.Chat{
			Id:               result.Id,
			UserId:           result.UserId,
			AppId:            result.AppId,
			Corp:             result.Corp,
			Model:            result.Model,
			Stream:           result.Stream,
			TextQuota:        result.TextQuota,
			ImageQuotas:      result.ImageQuotas,
			MultimodalQuota:  result.MultimodalQuota,
			PromptTokens:     result.PromptTokens,
			CompletionTokens: result.CompletionTokens,
			TotalTokens:      result.TotalTokens,
			ConnTime:         result.ConnTime,
			Duration:         result.Duration,
			TotalTime:        result.TotalTime,
			ReqTime:          util.FormatDateTimeMonth(result.ReqTime),
			Status:           result.Status,
			Creator:          result.Creator,
		}

		if service.Session().IsAdminRole(ctx) {
			chat.InternalTime = result.InternalTime
			chat.IsSmartMatch = result.IsSmartMatch
		}

		items = append(items, chat)
	}

	return &model.ChatPageRes{
		Items: items,
		Paging: &model.Paging{
			Page:     paging.Page,
			PageSize: paging.PageSize,
			Total:    paging.Total,
		},
	}, nil
}
