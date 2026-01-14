package log_general

import (
	"context"
	"time"

	"github.com/gogf/gf/v2/os/gtime"
	"github.com/gogf/gf/v2/text/gstr"
	"github.com/iimeta/fastapi-admin/v2/internal/config"
	"github.com/iimeta/fastapi-admin/v2/internal/dao"
	"github.com/iimeta/fastapi-admin/v2/internal/errors"
	"github.com/iimeta/fastapi-admin/v2/internal/logic/common"
	"github.com/iimeta/fastapi-admin/v2/internal/model"
	"github.com/iimeta/fastapi-admin/v2/internal/service"
	"github.com/iimeta/fastapi-admin/v2/utility/db"
	"github.com/iimeta/fastapi-admin/v2/utility/logger"
	"github.com/iimeta/fastapi-admin/v2/utility/util"
	"go.mongodb.org/mongo-driver/bson"
)

type sLogGeneral struct{}

func init() {
	service.RegisterLogGeneral(New())
}

func New() service.ILogGeneral {
	return &sLogGeneral{}
}

// 通用日志详情
func (s *sLogGeneral) Detail(ctx context.Context, id string) (*model.LogGeneral, error) {

	result, err := dao.LogGeneral.FindById(ctx, id)
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

	general := &model.LogGeneral{
		Id:           result.Id,
		TraceId:      result.TraceId,
		UserId:       result.UserId,
		AppId:        result.AppId,
		ProviderName: result.ProviderName,
		Model:        result.Model,
		ModelType:    result.ModelType,
		RequestData:  result.RequestData,
		ResponseData: result.ResponseData,
		Stream:       result.Stream,
		Completion:   result.Completion,
		Spend:        common.ConvSpend(result.Spend),
		ConnTime:     result.ConnTime,
		Duration:     result.Duration,
		TotalTime:    result.TotalTime,
		ReqTime:      util.FormatDateTime(result.ReqTime),
		ClientIp:     result.ClientIp,
		Retry:        result.Retry,
		Status:       result.Status,
		Host:         result.Host,
		Method:       result.Method,
		Path:         result.Path,
		Creator:      util.Desensitize(result.Creator),
	}

	if general.Status == -1 {

		general.ErrMsg = result.ErrMsg

		// 代理商屏蔽错误
		if service.Session().IsResellerRole(ctx) {
			if config.Cfg.ResellerShieldError.Open && len(config.Cfg.ResellerShieldError.Errors) > 0 {
				for _, shieldError := range config.Cfg.ResellerShieldError.Errors {
					if gstr.Contains(general.ErrMsg, shieldError) {
						general.ErrMsg = "详细错误信息请联系管理员..."
						break
					}
				}
			}
		}

		// 用户屏蔽错误
		if service.Session().IsUserRole(ctx) {
			if config.Cfg.UserShieldError.Open && len(config.Cfg.UserShieldError.Errors) > 0 {
				for _, shieldError := range config.Cfg.UserShieldError.Errors {
					if gstr.Contains(general.ErrMsg, shieldError) {
						general.ErrMsg = "详细错误信息请联系管理员..."
						break
					}
				}
			}
		}

		general.ErrMsg = gstr.Split(general.ErrMsg, " TraceId")[0]
		general.ErrMsg = gstr.Split(general.ErrMsg, " (request id:")[0]
	}

	if service.Session().IsAdminRole(ctx) {

		general.ProviderId = result.ProviderId
		general.ModelId = result.ModelId
		general.ModelName = result.ModelName
		general.Key = util.Desensitize(result.Key)
		general.IsEnablePresetConfig = result.IsEnablePresetConfig
		general.IsEnableModelAgent = result.IsEnableModelAgent
		general.ModelAgentId = result.ModelAgentId
		general.IsEnableForward = result.IsEnableForward
		general.ForwardConfig = result.ForwardConfig
		general.IsSmartMatch = result.IsSmartMatch
		general.IsEnableFallback = result.IsEnableFallback
		general.FallbackConfig = result.FallbackConfig
		general.RealModelId = result.RealModelId
		general.RealModelName = result.RealModelName
		general.RealModel = result.RealModel
		general.RemoteIp = result.RemoteIp
		general.LocalIp = result.LocalIp
		general.InternalTime = result.InternalTime
		general.ErrMsg = result.ErrMsg
		general.IsRetry = result.IsRetry
		general.CreatedAt = util.FormatDateTime(result.CreatedAt)
		general.UpdatedAt = util.FormatDateTime(result.UpdatedAt)

		if result.ModelAgent != nil {

			providerName := result.ModelAgent.ProviderId
			if provider, err := dao.Provider.FindById(ctx, result.ModelAgent.ProviderId); err == nil && provider != nil {
				providerName = provider.Name
			}

			general.ModelAgent = &model.ModelAgent{
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

	return general, nil
}

// 通用日志分页列表
func (s *sLogGeneral) Page(ctx context.Context, params model.LogGeneralPageReq) (*model.LogGeneralPageRes, error) {

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
		if service.Session().IsAdminRole(ctx) {
			filter["$or"] = bson.A{
				bson.M{"key": params.Key},
				bson.M{"creator": params.Key},
			}
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

	results, err := dao.LogGeneral.FindByPage(ctx, paging, filter, &dao.FindOptions{SortFields: []string{"-req_time", "status", "-created_at"}})
	if err != nil {
		logger.Error(ctx, err)
		return nil, err
	}

	items := make([]*model.LogGeneral, 0)
	for _, result := range results {

		general := &model.LogGeneral{
			Id:        result.Id,
			UserId:    result.UserId,
			AppId:     result.AppId,
			Model:     result.Model,
			ModelType: result.ModelType,
			Stream:    result.Stream,
			Spend:     common.ConvSpend(result.Spend),
			ConnTime:  result.ConnTime,
			Duration:  result.Duration,
			TotalTime: result.TotalTime,
			ReqTime:   util.FormatDateTimeMonth(result.ReqTime),
			Status:    result.Status,
		}

		if service.Session().IsAdminRole(ctx) {
			general.InternalTime = result.InternalTime
			general.IsSmartMatch = result.IsSmartMatch
		}

		items = append(items, general)
	}

	return &model.LogGeneralPageRes{
		Items: items,
		Paging: &model.Paging{
			Page:     paging.Page,
			PageSize: paging.PageSize,
			Total:    paging.Total,
		},
	}, nil
}

// 通用日志详情复制字段值
func (s *sLogGeneral) CopyField(ctx context.Context, params model.LogGeneralCopyFieldReq) (string, error) {

	result, err := dao.LogGeneral.FindById(ctx, params.Id)
	if err != nil {
		logger.Error(ctx, err)
		return "", err
	}

	if service.Session().IsResellerRole(ctx) && (params.Field == "key" || result.Rid != service.Session().GetRid(ctx)) {
		return "", errors.ERR_UNAUTHORIZED
	}

	if service.Session().IsUserRole(ctx) && (params.Field == "key" || result.UserId != service.Session().GetUserId(ctx)) {
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
