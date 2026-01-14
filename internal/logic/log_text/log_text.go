package log_text

import (
	"context"
	"fmt"
	"time"

	"github.com/gogf/gf/v2/os/gctx"
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
	"github.com/iimeta/fastapi-admin/v2/internal/service"
	"github.com/iimeta/fastapi-admin/v2/utility/db"
	"github.com/iimeta/fastapi-admin/v2/utility/logger"
	"github.com/iimeta/fastapi-admin/v2/utility/util"
	"go.mongodb.org/mongo-driver/bson"
)

type sLogText struct{}

func init() {
	service.RegisterLogText(New())
}

func New() service.ILogText {
	return &sLogText{}
}

// 文本日志详情
func (s *sLogText) Detail(ctx context.Context, id string) (*model.LogText, error) {

	result, err := dao.LogText.FindById(ctx, id)
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

	text := &model.LogText{
		Id:           result.Id,
		TraceId:      result.TraceId,
		UserId:       result.UserId,
		AppId:        result.AppId,
		ProviderName: result.ProviderName,
		Model:        result.Model,
		ModelType:    result.ModelType,
		Stream:       result.Stream,
		Messages:     result.Messages,
		Prompt:       result.Prompt,
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

	if text.Status == -1 {

		text.ErrMsg = result.ErrMsg

		// 代理商屏蔽错误
		if service.Session().IsResellerRole(ctx) {
			if config.Cfg.ResellerShieldError.Open && len(config.Cfg.ResellerShieldError.Errors) > 0 {
				for _, shieldError := range config.Cfg.ResellerShieldError.Errors {
					if gstr.Contains(text.ErrMsg, shieldError) {
						text.ErrMsg = "详细错误信息请联系管理员..."
						break
					}
				}
			}
		}

		// 用户屏蔽错误
		if service.Session().IsUserRole(ctx) {
			if config.Cfg.UserShieldError.Open && len(config.Cfg.UserShieldError.Errors) > 0 {
				for _, shieldError := range config.Cfg.UserShieldError.Errors {
					if gstr.Contains(text.ErrMsg, shieldError) {
						text.ErrMsg = "详细错误信息请联系管理员..."
						break
					}
				}
			}
		}

		text.ErrMsg = gstr.Split(text.ErrMsg, " TraceId")[0]
		text.ErrMsg = gstr.Split(text.ErrMsg, " (request id:")[0]
	}

	if service.Session().IsAdminRole(ctx) {
		text.ProviderId = result.ProviderId
		text.ModelId = result.ModelId
		text.ModelName = result.ModelName
		text.Key = util.Desensitize(result.Key)
		text.IsEnablePresetConfig = result.IsEnablePresetConfig
		text.IsEnableModelAgent = result.IsEnableModelAgent
		text.ModelAgentId = result.ModelAgentId
		text.IsEnableForward = result.IsEnableForward
		text.ForwardConfig = result.ForwardConfig
		text.IsSmartMatch = result.IsSmartMatch
		text.IsEnableFallback = result.IsEnableFallback
		text.FallbackConfig = result.FallbackConfig
		text.RealModelId = result.RealModelId
		text.RealModelName = result.RealModelName
		text.RealModel = result.RealModel
		text.RemoteIp = result.RemoteIp
		text.LocalIp = result.LocalIp
		text.InternalTime = result.InternalTime
		text.ErrMsg = result.ErrMsg
		text.IsRetry = result.IsRetry
		text.CreatedAt = util.FormatDateTime(result.CreatedAt)
		text.UpdatedAt = util.FormatDateTime(result.UpdatedAt)

		if result.ModelAgent != nil {

			providerName := result.ModelAgent.ProviderId
			if provider, err := dao.Provider.FindById(ctx, result.ModelAgent.ProviderId); err == nil && provider != nil {
				providerName = provider.Name
			}

			text.ModelAgent = &model.ModelAgent{
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

	return text, nil
}

// 文本日志分页列表
func (s *sLogText) Page(ctx context.Context, params model.LogTextPageReq) (*model.LogTextPageRes, error) {

	paging := &db.Paging{
		Page:     params.Page,
		PageSize: params.PageSize,
	}

	filter := bson.M{}
	index := ""

	if len(params.ReqTime) > 0 && params.TraceId == "" {
		gte := gtime.NewFromStrFormat(params.ReqTime[0], time.DateTime).TimestampMilli()
		lte := gtime.NewFromStrLayout(params.ReqTime[1], time.DateTime).TimestampMilli() + 999
		filter["req_time"] = bson.M{
			"$gte": gte,
			"$lte": lte,
		}
	}

	if params.TraceId != "" {
		filter["trace_id"] = gstr.Trim(params.TraceId)
	}

	if len(params.Models) > 0 {
		filter["model_id"] = bson.M{
			"$in": params.Models,
		}
		index = "req_time_-1_model_id_1_user_id_1_status_1_created_at_-1"
	}

	if len(params.ModelAgents) > 0 && service.Session().IsAdminRole(ctx) {
		filter["model_agent_id"] = bson.M{
			"$in": params.ModelAgents,
		}
	}

	if service.Session().IsResellerRole(ctx) {
		filter["rid"] = service.Session().GetRid(ctx)
		filter["is_smart_match"] = bson.M{"$exists": false}
		filter["is_retry"] = bson.M{"$exists": false}
	}

	if service.Session().IsUserRole(ctx) {
		filter["user_id"] = service.Session().GetUserId(ctx)
		filter["is_smart_match"] = bson.M{"$exists": false}
		filter["is_retry"] = bson.M{"$exists": false}
	} else if params.UserId != 0 {
		filter["user_id"] = params.UserId
		index = "req_time_-1_model_id_1_user_id_1_status_1_created_at_-1"
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
		if service.Session().IsAdminRole(ctx) {
			filter["$or"] = bson.A{
				bson.M{"key": params.Key},
				bson.M{"creator": params.Key},
			}
		} else {
			filter["creator"] = params.Key
		}
	}

	if params.TotalTime != 0 {
		filter["total_time"] = bson.M{
			"$gte": params.TotalTime,
		}
	}

	findOptions := &dao.FindOptions{
		SortFields:    []string{"-req_time", "status", "-created_at"},
		Index:         index,
		IncludeFields: []string{"_id", "user_id", "app_id", "model", "model_type", "stream", "spend", "conn_time", "duration", "total_time", "req_time", "status", "internal_time", "is_smart_match"},
	}

	results, err := dao.LogText.FindByPage(ctx, paging, filter, findOptions)
	if err != nil {
		logger.Error(ctx, err)
		return nil, err
	}

	items := make([]*model.LogText, 0)
	for _, result := range results {

		text := &model.LogText{
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
			text.InternalTime = result.InternalTime
			text.IsSmartMatch = result.IsSmartMatch
		}

		items = append(items, text)
	}

	return &model.LogTextPageRes{
		Items: items,
		Paging: &model.Paging{
			Page:     paging.Page,
			PageSize: paging.PageSize,
			Total:    paging.Total,
		},
	}, nil
}

// 文本日志导出
func (s *sLogText) Export(ctx context.Context, params model.LogTextExportReq) (string, error) {

	filter := bson.M{
		"status": bson.M{"$in": []int{1, -1, 2}},
	}

	if len(params.Ids) > 0 {
		filter["_id"] = bson.M{"$in": params.Ids}
	} else {
		filter["req_time"] = bson.M{
			"$gte": gtime.NewFromStrFormat(params.ReqTime[0], time.DateTime).TimestampMilli(),
			"$lte": gtime.NewFromStrLayout(params.ReqTime[1], time.DateTime).TimestampMilli() + 999,
		}
	}

	if service.Session().IsResellerRole(ctx) {
		filter["rid"] = service.Session().GetRid(ctx)
	}

	if service.Session().IsUserRole(ctx) {
		filter["user_id"] = service.Session().GetUserId(ctx)
	} else {
		if params.UserId != 0 {
			filter["user_id"] = params.UserId
		}
	}

	findOptions := &dao.FindOptions{
		SortFields:    []string{"-req_time", "status", "-created_at"},
		IncludeFields: []string{"user_id", "app_id", "model", "spend", "req_time", "key", "creator"},
	}

	results, err := dao.LogText.Find(ctx, filter, findOptions)
	if err != nil {
		logger.Error(ctx, err)
		return "", err
	}

	colFieldMap := make(map[string]string)
	colFieldMap["请求时间"] = "ReqTime"
	colFieldMap["模型"] = "Model"
	colFieldMap["输入Token数"] = "InputTokens"
	colFieldMap["输出Token数"] = "OutputTokens"
	colFieldMap["花费"] = "TotalTokens"

	var titleCols []string

	if service.Session().IsResellerRole(ctx) {
		titleCols = append(titleCols, "请求时间", "用户ID", "模型", "输入Token数", "输出Token数", "花费")
		colFieldMap["用户ID"] = "UserId"
	}

	if service.Session().IsUserRole(ctx) {
		titleCols = append(titleCols, "请求时间", "应用ID", "密钥", "模型", "输入Token数", "输出Token数", "花费")
		colFieldMap["应用ID"] = "AppId"
		colFieldMap["密钥"] = "Creator"
	}

	if service.Session().IsAdminRole(ctx) {
		titleCols = append(titleCols, "请求时间", "用户ID", "模型", "输入Token数", "输出Token数", "花费", "密钥")
		colFieldMap["用户ID"] = "UserId"
		colFieldMap["密钥"] = "Key"
	}

	filePath := fmt.Sprintf("./resource/export/text_%d.xlsx", gtime.TimestampMilli())

	values := make([]any, 0)
	for _, result := range results {

		value := &model.LogTextExport{
			ReqTime:     util.FormatDateTime(result.ReqTime),
			UserId:      result.UserId,
			AppId:       result.AppId,
			Model:       result.Model,
			TotalTokens: gconv.String(common.ConvQuotaUnitReverse(int(result.Spend.TotalSpendTokens))),
			Key:         result.Key,
			Creator:     result.Creator,
		}

		if result.Spend.Text != nil {
			value.InputTokens = result.Spend.Text.InputTokens
			value.OutputTokens = result.Spend.Text.OutputTokens
		} else if result.Spend.TieredText != nil {
			value.InputTokens = result.Spend.TieredText.InputTokens
			value.OutputTokens = result.Spend.TieredText.OutputTokens
		}

		values = append(values, value)
	}

	if err = util.ExcelExport("文本日志", titleCols, colFieldMap, values, filePath); err != nil {
		return "", err
	}

	return filePath, nil
}

// 文本日志批量操作
func (s *sLogText) BatchOperate(ctx context.Context, params model.LogTextBatchOperateReq) error {

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

			if _, err := dao.LogText.DeleteMany(ctx, filter); err != nil {
				logger.Error(ctx, err)
			}

		case consts.ACTION_DELETE:
			if _, err := dao.LogText.DeleteMany(ctx, bson.M{"_id": bson.M{"$in": params.Ids}}); err != nil {
				logger.Error(ctx, err)
			}
		}

	}, nil); err != nil {
		logger.Error(ctx, err)
		return err
	}

	return nil
}

// 文本日志详情复制字段值
func (s *sLogText) CopyField(ctx context.Context, params model.LogTextCopyFieldReq) (string, error) {

	result, err := dao.LogText.FindById(ctx, params.Id)
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
