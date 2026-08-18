package log_batch

import (
	"context"
	"time"

	"github.com/gogf/gf/v2/os/gctx"
	"github.com/gogf/gf/v2/os/grpool"
	"github.com/gogf/gf/v2/os/gtime"
	"github.com/gogf/gf/v2/text/gstr"
	"github.com/gogf/gf/v2/util/gconv"
	"github.com/iimeta/fastapi-admin/v2/internal/consts"
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

	batch.ErrMsg = common.ConvErrMsg(ctx, result.ErrMsg, result.Status)

	if service.Session().IsAdminRole(ctx) {

		batch.ProviderId = result.ProviderId
		batch.ProviderCode = result.ProviderCode
		batch.ModelId = result.ModelId
		batch.ModelName = result.ModelName
		batch.Key = util.Desensitize(result.Key)
		batch.IsEnablePresetConfig = result.IsEnablePresetConfig
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

	if len(params.Actions) > 0 {
		filter["action"] = bson.M{
			"$in": params.Actions,
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
			Id:           result.Id,
			UserId:       result.UserId,
			AppId:        result.AppId,
			ProviderName: result.ProviderName,
			Model:        result.Model,
			ModelType:    result.ModelType,
			Action:       result.Action,
			BatchId:      result.BatchId,
			Spend:        common.ConvSpend(result.Spend),
			TotalTime:    result.TotalTime,
			ReqTime:      util.FormatDateTimeMonth(result.ReqTime),
			Status:       result.Status,
			ErrMsg:       common.ConvErrMsg(ctx, result.ErrMsg, result.Status),
		}

		if service.Session().IsAdminRole(ctx) {
			batch.ProviderCode = result.ProviderCode
			batch.InternalTime = result.InternalTime
			batch.IsSmartMatch = result.IsSmartMatch
		}

		items = append(items, batch)
	}

	if service.Session().IsUserRole(ctx) {

		appIds := make([]int, 0)
		keys := make([]string, 0)
		for _, result := range results {
			appIds = append(appIds, result.AppId)
			keys = append(keys, result.Creator)
		}

		appNames := common.GetAppNames(ctx, appIds)
		keyNames := common.GetKeyNames(ctx, keys)

		for i, result := range results {
			items[i].AppName = appNames[result.AppId]
			items[i].KeyName = keyNames[result.Creator]
		}
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

// 批处理日志批量操作
func (s *sLogBatch) BatchOperate(ctx context.Context, params model.LogBatchBatchOperateReq) error {

	if err := grpool.AddWithRecover(gctx.NeverDone(ctx), func(ctx context.Context) {

		switch params.Action {
		case consts.ACTION_TIME:

			reqTime := params.Value.([]any)
			filter := bson.M{
				"req_time": bson.M{
					"$gte": gtime.NewFromStrFormat(gconv.String(reqTime[0]), time.DateTime).TimestampMilli(),
					"$lte": gtime.NewFromStrLayout(gconv.String(reqTime[1]), time.DateTime).TimestampMilli() + 999,
				},
			}

			if params.UserId != 0 {
				filter["user_id"] = params.UserId
			}

			if len(params.Status) != 4 {
				filter["status"] = bson.M{"$in": params.Status}
			}

			if _, err := dao.LogBatch.DeleteMany(ctx, filter); err != nil {
				logger.Error(ctx, err)
			}

		case consts.ACTION_DELETE:
			if _, err := dao.LogBatch.DeleteMany(ctx, bson.M{"_id": bson.M{"$in": params.Ids}}); err != nil {
				logger.Error(ctx, err)
			}
		}

	}, nil); err != nil {
		logger.Error(ctx, err)
		return err
	}

	return nil
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
