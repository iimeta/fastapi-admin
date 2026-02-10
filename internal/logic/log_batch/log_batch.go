package log_batch

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
	"go.mongodb.org/mongo-driver/v2/bson"
)

type sLogBatch struct{}

func init() {
	service.RegisterLogBatch(New())
}

func New() service.ILogBatch {
	return &sLogBatch{}
}

// 批处理日志详情
func (s *sLogBatch) Detail(ctx context.Context, id string) (*model.LogBatch, error) {

	result, err := dao.LogBatch.FindById(ctx, id)
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

	batch := &model.LogBatch{
		Id:           result.Id,
		TraceId:      result.TraceId,
		UserId:       result.UserId,
		AppId:        result.AppId,
		ProviderName: result.ProviderName,
		Model:        result.Model,
		ModelType:    result.ModelType,
		Action:       result.Action,
		BatchId:      result.BatchId,
		RequestData:  result.RequestData,
		ResponseData: result.ResponseData,
		Spend:        common.ConvSpend(result.Spend),
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

	if batch.Status == -1 {

		batch.ErrMsg = result.ErrMsg

		// 代理商屏蔽错误
		if service.Session().IsResellerRole(ctx) {
			if config.Cfg.ResellerShieldError.Open && len(config.Cfg.ResellerShieldError.Errors) > 0 {
				for _, shieldError := range config.Cfg.ResellerShieldError.Errors {
					if gstr.Contains(batch.ErrMsg, shieldError) {
						batch.ErrMsg = "详细错误信息请联系管理员..."
						break
					}
				}
			}
		}

		// 用户屏蔽错误
		if service.Session().IsUserRole(ctx) {
			if config.Cfg.UserShieldError.Open && len(config.Cfg.UserShieldError.Errors) > 0 {
				for _, shieldError := range config.Cfg.UserShieldError.Errors {
					if gstr.Contains(batch.ErrMsg, shieldError) {
						batch.ErrMsg = "详细错误信息请联系管理员..."
						break
					}
				}
			}
		}

		batch.ErrMsg = gstr.Split(batch.ErrMsg, " TraceId")[0]
		batch.ErrMsg = gstr.Split(batch.ErrMsg, " (request id:")[0]
	}

	if service.Session().IsAdminRole(ctx) {

		batch.ProviderId = result.ProviderId
		batch.ModelId = result.ModelId
		batch.ModelName = result.ModelName
		batch.Key = util.Desensitize(result.Key)
		batch.IsEnablePresetConfig = result.IsEnablePresetConfig
		batch.IsEnableModelAgent = result.IsEnableModelAgent
		batch.ModelAgentId = result.ModelAgentId
		batch.IsEnableForward = result.IsEnableForward
		batch.ForwardConfig = result.ForwardConfig
		batch.IsSmartMatch = result.IsSmartMatch
		batch.IsEnableFallback = result.IsEnableFallback
		batch.FallbackConfig = result.FallbackConfig
		batch.RealModelId = result.RealModelId
		batch.RealModelName = result.RealModelName
		batch.RealModel = result.RealModel
		batch.RemoteIp = result.RemoteIp
		batch.LocalIp = result.LocalIp
		batch.InternalTime = result.InternalTime
		batch.ErrMsg = result.ErrMsg
		batch.IsRetry = result.IsRetry
		batch.CreatedAt = util.FormatDateTime(result.CreatedAt)
		batch.UpdatedAt = util.FormatDateTime(result.UpdatedAt)

		if result.ModelAgent != nil {

			providerName := result.ModelAgent.ProviderId
			if provider, err := dao.Provider.FindById(ctx, result.ModelAgent.ProviderId); err == nil && provider != nil {
				providerName = provider.Name
			}

			batch.ModelAgent = &model.ModelAgent{
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

	return batch, nil
}

// 批处理日志分页列表
func (s *sLogBatch) Page(ctx context.Context, params model.LogBatchPageReq) (*model.LogBatchPageRes, error) {

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

	results, err := dao.LogBatch.FindByPage(ctx, paging, filter, &dao.FindOptions{SortFields: []string{"-req_time", "status", "-created_at"}})
	if err != nil {
		logger.Error(ctx, err)
		return nil, err
	}

	items := make([]*model.LogBatch, 0)
	for _, result := range results {

		batch := &model.LogBatch{
			Id:        result.Id,
			UserId:    result.UserId,
			AppId:     result.AppId,
			Model:     result.Model,
			ModelType: result.ModelType,
			Action:    result.Action,
			BatchId:   result.BatchId,
			Spend:     common.ConvSpend(result.Spend),
			TotalTime: result.TotalTime,
			ReqTime:   util.FormatDateTimeMonth(result.ReqTime),
			Status:    result.Status,
		}

		if service.Session().IsAdminRole(ctx) {
			batch.InternalTime = result.InternalTime
			batch.IsSmartMatch = result.IsSmartMatch
		}

		items = append(items, batch)
	}

	return &model.LogBatchPageRes{
		Items: items,
		Paging: &model.Paging{
			Page:     paging.Page,
			PageSize: paging.PageSize,
			Total:    paging.Total,
		},
	}, nil
}

// 批处理日志详情复制字段值
func (s *sLogBatch) CopyField(ctx context.Context, params model.LogBatchCopyFieldReq) (string, error) {

	result, err := dao.LogBatch.FindById(ctx, params.Id)
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
	case "batch_id":
		return result.BatchId, nil
	}

	return "", nil
}
