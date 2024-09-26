package midjourney

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

	if service.Session().IsUserRole(ctx) && result.UserId != service.Session().GetUserId(ctx) {
		return nil, errors.ERR_UNAUTHORIZED
	}

	corpName := result.Corp
	if corp, err := dao.Corp.FindById(ctx, result.Corp); err == nil && corp != nil {
		corpName = corp.Name
	}

	midjourney := &model.Midjourney{
		Id:               result.Id,
		TraceId:          result.TraceId,
		UserId:           result.UserId,
		AppId:            result.AppId,
		Corp:             result.Corp,
		CorpName:         corpName,
		Model:            result.Model,
		Type:             result.Type,
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
		Creator:          util.Desensitize(result.Creator)}

	// todo
	if midjourney.Status == -1 && service.Session().IsUserRole(ctx) {
		midjourney.ErrMsg = "详细错误信息请联系管理员..."
	}

	if service.Session().IsAdminRole(ctx) {

		midjourney.ModelId = result.ModelId
		midjourney.Name = result.Name
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

			corpName := result.ModelAgent.Corp
			if corp, err := dao.Corp.FindById(ctx, result.ModelAgent.Corp); err == nil && corp != nil {
				corpName = corp.Name
			}

			midjourney.ModelAgent = &model.ModelAgent{
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

	results, err := dao.Midjourney.FindByPage(ctx, paging, filter, "", "-req_time", "status", "-created_at")
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
			Corp:             result.Corp,
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
