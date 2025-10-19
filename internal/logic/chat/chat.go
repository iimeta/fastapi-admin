package chat

import (
	"context"
	"fmt"
	"time"

	"github.com/gogf/gf/v2/os/gctx"
	"github.com/gogf/gf/v2/os/grpool"
	"github.com/gogf/gf/v2/os/gtime"
	"github.com/gogf/gf/v2/text/gstr"
	"github.com/gogf/gf/v2/util/gconv"
	"github.com/iimeta/fastapi-admin/internal/config"
	"github.com/iimeta/fastapi-admin/internal/consts"
	"github.com/iimeta/fastapi-admin/internal/dao"
	"github.com/iimeta/fastapi-admin/internal/errors"
	"github.com/iimeta/fastapi-admin/internal/logic/common"
	"github.com/iimeta/fastapi-admin/internal/model"
	"github.com/iimeta/fastapi-admin/internal/service"
	"github.com/iimeta/fastapi-admin/utility/db"
	"github.com/iimeta/fastapi-admin/utility/logger"
	"github.com/iimeta/fastapi-admin/utility/util"
	"go.mongodb.org/mongo-driver/bson"
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

	if service.Session().IsResellerRole(ctx) && result.Rid != service.Session().GetRid(ctx) {
		return nil, errors.ERR_UNAUTHORIZED
	}

	if service.Session().IsUserRole(ctx) && result.UserId != service.Session().GetUserId(ctx) {
		return nil, errors.ERR_UNAUTHORIZED
	}

	chat := &model.Chat{
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
		Spend:        common.ConvSpendPricingToPrice(result.Spend),
		ConnTime:     result.ConnTime,
		Duration:     result.Duration,
		TotalTime:    result.TotalTime,
		ReqTime:      util.FormatDateTime(result.ReqTime),
		ClientIp:     result.ClientIp,
		Retry:        result.Retry,
		Status:       result.Status,
		Host:         result.Host,
		Creator:      util.Desensitize(result.Creator),
	}

	if chat.Status == -1 {

		chat.ErrMsg = result.ErrMsg

		// 代理商屏蔽错误
		if service.Session().IsResellerRole(ctx) {
			if config.Cfg.ResellerShieldError.Open && len(config.Cfg.ResellerShieldError.Errors) > 0 {
				for _, shieldError := range config.Cfg.ResellerShieldError.Errors {
					if gstr.Contains(chat.ErrMsg, shieldError) {
						chat.ErrMsg = "详细错误信息请联系管理员..."
						break
					}
				}
			}
		}

		// 用户屏蔽错误
		if service.Session().IsUserRole(ctx) {
			if config.Cfg.UserShieldError.Open && len(config.Cfg.UserShieldError.Errors) > 0 {
				for _, shieldError := range config.Cfg.UserShieldError.Errors {
					if gstr.Contains(chat.ErrMsg, shieldError) {
						chat.ErrMsg = "详细错误信息请联系管理员..."
						break
					}
				}
			}
		}

		chat.ErrMsg = gstr.Split(chat.ErrMsg, " TraceId")[0]
		chat.ErrMsg = gstr.Split(chat.ErrMsg, " (request id:")[0]
	}

	if service.Session().IsAdminRole(ctx) {
		chat.ProviderId = result.ProviderId
		chat.ModelId = result.ModelId
		chat.ModelName = result.ModelName
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

			providerName := result.ModelAgent.ProviderId
			if provider, err := dao.Provider.FindById(ctx, result.ModelAgent.ProviderId); err == nil && provider != nil {
				providerName = provider.Name
			}

			chat.ModelAgent = &model.ModelAgent{
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

	return chat, nil
}

// 聊天日志分页列表
func (s *sChat) Page(ctx context.Context, params model.ChatPageReq) (*model.ChatPageRes, error) {

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
			filter["key"] = params.Key
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

	results, err := dao.Chat.FindByPage(ctx, paging, filter, findOptions)
	if err != nil {
		logger.Error(ctx, err)
		return nil, err
	}

	items := make([]*model.Chat, 0)
	for _, result := range results {

		chat := &model.Chat{
			Id:        result.Id,
			UserId:    result.UserId,
			AppId:     result.AppId,
			Model:     result.Model,
			ModelType: result.ModelType,
			Stream:    result.Stream,
			Spend:     common.ConvSpendPricingToPrice(result.Spend),
			ConnTime:  result.ConnTime,
			Duration:  result.Duration,
			TotalTime: result.TotalTime,
			ReqTime:   util.FormatDateTimeMonth(result.ReqTime),
			Status:    result.Status,
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

// 聊天导出
func (s *sChat) Export(ctx context.Context, params model.ChatExportReq) (string, error) {

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

	results, err := dao.Chat.Find(ctx, filter, findOptions)
	if err != nil {
		logger.Error(ctx, err)
		return "", err
	}

	colFieldMap := make(map[string]string)
	colFieldMap["请求时间"] = "ReqTime"
	colFieldMap["模型"] = "Model"
	colFieldMap["输入Token数"] = "InputTokens"
	colFieldMap["输出Token数"] = "OutputTokens"
	colFieldMap["花费($)"] = "TotalTokens"

	var titleCols []string

	if service.Session().IsResellerRole(ctx) {
		titleCols = append(titleCols, "请求时间", "用户ID", "模型", "输入Token数", "输出Token数", "花费($)")
		colFieldMap["用户ID"] = "UserId"
	}

	if service.Session().IsUserRole(ctx) {
		titleCols = append(titleCols, "请求时间", "应用ID", "密钥", "模型", "输入Token数", "输出Token数", "花费($)")
		colFieldMap["应用ID"] = "AppId"
		colFieldMap["密钥"] = "Creator"
	}

	if service.Session().IsAdminRole(ctx) {
		titleCols = append(titleCols, "请求时间", "用户ID", "模型", "输入Token数", "输出Token数", "花费($)", "密钥")
		colFieldMap["用户ID"] = "UserId"
		colFieldMap["密钥"] = "Key"
	}

	filePath := fmt.Sprintf("./resource/export/chat_%d.xlsx", gtime.TimestampMilli())

	values := make([]interface{}, 0)
	for _, result := range results {

		value := &model.ChatExport{
			ReqTime:     util.FormatDateTime(result.ReqTime),
			UserId:      result.UserId,
			AppId:       result.AppId,
			Model:       result.Model,
			TotalTokens: gconv.String(util.ConvQuota(result.Spend.TotalSpendTokens)),
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

	if err = util.ExcelExport("聊天日志", titleCols, colFieldMap, values, filePath); err != nil {
		return "", err
	}

	return filePath, nil
}

// 聊天批量操作
func (s *sChat) BatchOperate(ctx context.Context, params model.ChatBatchOperateReq) error {

	if err := grpool.AddWithRecover(gctx.NeverDone(ctx), func(ctx context.Context) {

		switch params.Action {
		case consts.ACTION_TIME:

			reqTime := params.Value.([]interface{})
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

			if _, err := dao.Chat.DeleteMany(ctx, filter); err != nil {
				logger.Error(ctx, err)
			}

		case consts.ACTION_DELETE:
			if _, err := dao.Chat.DeleteMany(ctx, bson.M{"_id": bson.M{"$in": params.Ids}}); err != nil {
				logger.Error(ctx, err)
			}
		}

	}, nil); err != nil {
		logger.Error(ctx, err)
		return err
	}

	return nil
}

// 聊天日志详情复制字段值
func (s *sChat) CopyField(ctx context.Context, params model.ChatCopyFieldReq) (string, error) {

	result, err := dao.Chat.FindById(ctx, params.Id)
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
