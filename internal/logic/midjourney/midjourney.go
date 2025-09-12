package midjourney

import (
	"context"
	"time"

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
)

type sMidjourney struct{}

func init() {
	service.RegisterMidjourney(New())
}

func New() service.IMidjourney {
	return &sMidjourney{}
}

// Midjourney详情
func (s *sMidjourney) Detail(ctx context.Context, id string) (*model.Midjourney, error) {

	result, err := dao.Midjourney.FindById(ctx, id)
	if err != nil {
		logger.Error(ctx, err)
		return nil, err
	}

	if service.Session().IsResellerRole(ctx) && result.Rid != service.Session().GetRid(ctx) {
		return nil, errors.ERR_UNAUTHORIZED
	}

	if service.Session().IsUserRole(ctx) && result.UserId != service.Session().GetUserId(ctx) {
		return nil, errors.ERR_UNAUTHORIZED
	}

	providerName := result.ProviderId
	if provider, err := dao.Provider.FindById(ctx, result.ProviderId); err == nil && provider != nil {
		providerName = provider.Name
	}

	midjourney := &model.Midjourney{
		Id:               result.Id,
		TraceId:          result.TraceId,
		UserId:           result.UserId,
		AppId:            result.AppId,
		GroupId:          result.GroupId,
		GroupName:        result.GroupName,
		Discount:         result.Discount,
		ProviderName:     providerName,
		Model:            result.Model,
		ModelType:        result.ModelType,
		Prompt:           result.Prompt,
		MidjourneyQuotas: result.MidjourneyQuotas,
		TotalTokens:      result.TotalTokens,
		ConnTime:         result.ConnTime,
		Duration:         result.Duration,
		TotalTime:        result.TotalTime,
		ReqTime:          util.FormatDateTime(result.ReqTime),
		ClientIp:         result.ClientIp,
		Retry:            result.Retry,
		Status:           result.Status,
		Host:             result.Host,
		Creator:          util.Desensitize(result.Creator),
	}

	if midjourney.Status == -1 && service.Session().IsResellerRole(ctx) {
		midjourney.ErrMsg = "详细错误信息请联系管理员..."
		if config.Cfg.ResellerShieldError.Open && len(config.Cfg.ResellerShieldError.Errors) > 0 {
			midjourney.ErrMsg = result.ErrMsg
			for _, shieldError := range config.Cfg.ResellerShieldError.Errors {
				if gstr.Contains(result.ErrMsg, shieldError) {
					midjourney.ErrMsg = "详细错误信息请联系管理员..."
					break
				}
			}
		}
	}

	if midjourney.Status == -1 && service.Session().IsUserRole(ctx) {

		midjourney.ErrMsg = result.ErrMsg

		// 用户屏蔽错误
		if config.Cfg.UserShieldError.Open && len(config.Cfg.UserShieldError.Errors) > 0 {
			for _, shieldError := range config.Cfg.UserShieldError.Errors {
				if gstr.Contains(midjourney.ErrMsg, shieldError) {
					midjourney.ErrMsg = "详细错误信息请联系管理员..."
					break
				}
			}
		}

		midjourney.ErrMsg = gstr.Split(midjourney.ErrMsg, " TraceId")[0]
		midjourney.ErrMsg = gstr.Split(midjourney.ErrMsg, " (request id:")[0]
	}

	if service.Session().IsAdminRole(ctx) {

		midjourney.ProviderId = result.ProviderId
		midjourney.ModelId = result.ModelId
		midjourney.ModelName = result.ModelName
		midjourney.Key = result.Key
		midjourney.IsEnablePresetConfig = result.IsEnablePresetConfig
		midjourney.IsEnableModelAgent = result.IsEnableModelAgent
		midjourney.ModelAgentId = result.ModelAgentId
		midjourney.IsEnableForward = result.IsEnableForward
		midjourney.ForwardConfig = result.ForwardConfig
		midjourney.IsSmartMatch = result.IsSmartMatch
		midjourney.IsEnableFallback = result.IsEnableFallback
		midjourney.FallbackConfig = result.FallbackConfig
		midjourney.RealModelId = result.RealModelId
		midjourney.RealModelName = result.RealModelName
		midjourney.RealModel = result.RealModel
		midjourney.RemoteIp = result.RemoteIp
		midjourney.LocalIp = result.LocalIp
		midjourney.InternalTime = result.InternalTime
		midjourney.ErrMsg = result.ErrMsg
		midjourney.IsRetry = result.IsRetry
		midjourney.CreatedAt = util.FormatDateTime(result.CreatedAt)
		midjourney.UpdatedAt = util.FormatDateTime(result.UpdatedAt)

		if result.ModelAgent != nil {

			providerName := result.ModelAgent.ProviderId
			if provider, err := dao.Provider.FindById(ctx, result.ModelAgent.ProviderId); err == nil && provider != nil {
				providerName = provider.Name
			}

			midjourney.ModelAgent = &model.ModelAgent{
				ProviderId:   result.ModelAgent.ProviderId,
				ProviderName: providerName,
				Name:         result.ModelAgent.Name,
				BaseUrl:      result.ModelAgent.BaseUrl,
				Path:         result.ModelAgent.Path,
				Weight:       result.ModelAgent.Weight,
				Remark:       result.ModelAgent.Remark,
			}
		}
	}

	return midjourney, nil
}

// Midjourney分页列表
func (s *sMidjourney) Page(ctx context.Context, params model.MidjourneyPageReq) (*model.MidjourneyPageRes, error) {

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
		filter["is_retry"] = bson.M{"$exists": false}
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

	if len(params.ReqTime) > 0 && params.TraceId == "" {
		gte := gtime.NewFromStrFormat(params.ReqTime[0], time.DateTime).TimestampMilli()
		lte := gtime.NewFromStrLayout(params.ReqTime[1], time.DateTime).TimestampMilli() + 999
		filter["req_time"] = bson.M{
			"$gte": gte,
			"$lte": lte,
		}
	}

	results, err := dao.Midjourney.FindByPage(ctx, paging, filter, &dao.FindOptions{SortFields: []string{"-req_time", "status", "-created_at"}})
	if err != nil {
		logger.Error(ctx, err)
		return nil, err
	}

	items := make([]*model.Midjourney, 0)
	for _, result := range results {

		midjourney := &model.Midjourney{
			Id:               result.Id,
			UserId:           result.UserId,
			AppId:            result.AppId,
			Model:            result.Model,
			MidjourneyQuotas: result.MidjourneyQuotas,
			TotalTokens:      result.TotalTokens,
			ConnTime:         result.ConnTime,
			Duration:         result.Duration,
			TotalTime:        result.TotalTime,
			ReqTime:          util.FormatDateTimeMonth(result.ReqTime),
			Status:           result.Status,
		}

		if service.Session().IsAdminRole(ctx) {
			midjourney.InternalTime = result.InternalTime
			midjourney.IsSmartMatch = result.IsSmartMatch
		}

		items = append(items, midjourney)
	}

	return &model.MidjourneyPageRes{
		Items: items,
		Paging: &model.Paging{
			Page:     paging.Page,
			PageSize: paging.PageSize,
			Total:    paging.Total,
		},
	}, nil
}
