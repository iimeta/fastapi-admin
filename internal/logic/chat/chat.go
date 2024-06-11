package key

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

// 聊天详情
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
		Id:            result.Id,
		TraceId:       result.TraceId,
		UserId:        result.UserId,
		AppId:         result.AppId,
		Corp:          result.Corp,
		CorpName:      corpName,
		Model:         result.Model,
		Type:          result.Type,
		Stream:        result.Stream,
		Prompt:        result.Prompt,
		Completion:    result.Completion,
		BillingMethod: result.BillingMethod,
		TotalTokens:   result.TotalTokens,
		TotalPrice:    util.QuotaConv(result.TotalTokens),
		FixedQuota:    result.FixedQuota,
		ConnTime:      result.ConnTime,
		Duration:      result.Duration,
		TotalTime:     result.TotalTime,
		ReqTime:       util.FormatDateTime(result.ReqTime),
		ClientIp:      result.ClientIp,
		Status:        result.Status,
		Creator:       result.Creator,
	}

	if result.BillingMethod == 1 {
		chat.PromptRatio = result.PromptRatio
		chat.CompletionRatio = result.CompletionRatio
		chat.PromptTokens = result.PromptTokens
		chat.CompletionTokens = result.CompletionTokens
	}

	for _, message := range result.Messages {
		chat.Messages = append(chat.Messages, model.Message{
			Role:    message.Role,
			Content: message.Content,
		})
	}

	// todo
	if chat.Status == -1 && service.Session().IsUserRole(ctx) {
		chat.ErrMsg = "详细错误信息请联系管理员..."
	}

	if service.Session().IsAdminRole(ctx) {

		chat.ModelId = result.ModelId
		chat.Name = result.Name
		chat.Key = result.Key
		chat.PromptRatio = result.PromptRatio
		chat.CompletionRatio = result.CompletionRatio
		chat.PromptTokens = result.PromptTokens
		chat.CompletionTokens = result.CompletionTokens
		chat.IsEnableModelAgent = result.IsEnableModelAgent
		chat.ModelAgentId = result.ModelAgentId
		chat.IsEnableForward = result.IsEnableForward
		chat.IsSmartMatch = result.IsSmartMatch
		chat.IsEnableFallback = result.IsEnableFallback
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

		if result.ForwardConfig != nil {
			chat.ForwardConfig = &model.ForwardConfig{
				ForwardRule:   result.ForwardConfig.ForwardRule,
				MatchRule:     result.ForwardConfig.MatchRule,
				TargetModel:   result.ForwardConfig.TargetModel,
				DecisionModel: result.ForwardConfig.DecisionModel,
				Keywords:      result.ForwardConfig.Keywords,
				TargetModels:  result.ForwardConfig.TargetModels,
				ContentLength: result.ForwardConfig.ContentLength,
			}
		}

		if result.FallbackConfig != nil {
			chat.FallbackConfig = &model.FallbackConfig{
				FallbackModel:     result.FallbackConfig.FallbackModel,
				FallbackModelName: result.FallbackConfig.FallbackModelName,
			}
		}

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

		if result.Retry != nil {
			chat.Retry = &model.Retry{
				IsRetry:    result.Retry.IsRetry,
				RetryCount: result.Retry.RetryCount,
				ErrMsg:     result.Retry.ErrMsg,
			}
		}
	}

	return chat, nil
}

// 聊天分页列表
func (s *sChat) Page(ctx context.Context, params model.ChatPageReq) (*model.ChatPageRes, error) {

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
		filter["is_smart_match"] = bson.M{"$exists": false}
		filter["is_retry"] = bson.M{"$exists": false}
	} else if params.UserId != 0 {
		filter["user_id"] = params.UserId
	}

	if params.AppId != 0 {
		filter["app_id"] = params.AppId
	}

	if params.Key != "" {
		filter["creator"] = params.Key
	}

	if len(params.Models) > 0 {
		filter["model_id"] = bson.M{
			"$in": params.Models,
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

	if len(params.ReqTime) > 0 {
		gte := gtime.NewFromStrFormat(params.ReqTime[0], time.DateTime).TimestampMilli()
		lte := gtime.NewFromStrLayout(params.ReqTime[1], time.DateTime).TimestampMilli() + 999
		filter["req_time"] = bson.M{
			"$gte": gte,
			"$lte": lte,
		}
	}

	results, err := dao.Chat.FindByPage(ctx, paging, filter, "-req_time", "status", "-created_at")
	if err != nil {
		logger.Error(ctx, err)
		return nil, err
	}

	items := make([]*model.Chat, 0)
	for _, result := range results {

		chat := &model.Chat{
			Id:            result.Id,
			UserId:        result.UserId,
			AppId:         result.AppId,
			Corp:          result.Corp,
			Model:         result.Model,
			Stream:        result.Stream,
			BillingMethod: result.BillingMethod,
			TotalTokens:   result.TotalTokens,
			TotalPrice:    util.QuotaConv(result.TotalTokens),
			ConnTime:      result.ConnTime,
			Duration:      result.Duration,
			TotalTime:     result.TotalTime,
			ReqTime:       util.FormatDateTimeMonth(result.ReqTime),
			Status:        result.Status,
			Creator:       result.Creator,
		}

		if result.BillingMethod == 1 {
			chat.PromptRatio = result.PromptRatio
			chat.CompletionRatio = result.CompletionRatio
			chat.PromptTokens = result.PromptTokens
			chat.CompletionTokens = result.CompletionTokens
		}

		if service.Session().IsAdminRole(ctx) {
			chat.PromptRatio = result.PromptRatio
			chat.CompletionRatio = result.CompletionRatio
			chat.PromptTokens = result.PromptTokens
			chat.CompletionTokens = result.CompletionTokens
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
