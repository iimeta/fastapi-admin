package image

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

type sImage struct{}

func init() {
	service.RegisterImage(New())
}

func New() service.IImage {
	return &sImage{}
}

// 绘图日志详情
func (s *sImage) Detail(ctx context.Context, id string) (*model.Image, error) {

	result, err := dao.Image.FindById(ctx, id)
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

	image := &model.Image{
		Id:              result.Id,
		TraceId:         result.TraceId,
		UserId:          result.UserId,
		AppId:           result.AppId,
		GroupId:         result.GroupId,
		GroupName:       result.GroupName,
		Discount:        result.Discount,
		ProviderName:    providerName,
		Model:           result.Model,
		ModelType:       result.ModelType,
		Prompt:          result.Prompt,
		Size:            result.Size,
		N:               len(result.ImageData),
		Quality:         result.Quality,
		ImageData:       result.ImageData,
		ImageQuota:      result.ImageQuota,
		GenerationQuota: result.GenerationQuota,
		InputTokens:     result.InputTokens,
		OutputTokens:    result.OutputTokens,
		TextTokens:      result.TextTokens,
		ImageTokens:     result.ImageTokens,
		TotalTokens:     result.TotalTokens,
		TotalTime:       result.TotalTime,
		ReqTime:         util.FormatDateTime(result.ReqTime),
		ClientIp:        result.ClientIp,
		Retry:           result.Retry,
		Status:          result.Status,
		Host:            result.Host,
		Creator:         util.Desensitize(result.Creator),
	}

	for _, imageData := range result.ImageData {
		//if imageData.URL != "" {
		image.Images = append(image.Images, imageData.URL)
		//} else { // 太大了, 不查
		//	image.Images = append(image.Images, "data:image/png;base64,"+imageData.B64JSON)
		//}
	}

	if image.Status == -1 {

		image.ErrMsg = result.ErrMsg

		// 代理商屏蔽错误
		if service.Session().IsResellerRole(ctx) {
			if config.Cfg.ResellerShieldError.Open && len(config.Cfg.ResellerShieldError.Errors) > 0 {
				for _, shieldError := range config.Cfg.ResellerShieldError.Errors {
					if gstr.Contains(image.ErrMsg, shieldError) {
						image.ErrMsg = "详细错误信息请联系管理员..."
						break
					}
				}
			}
		}

		// 用户屏蔽错误
		if service.Session().IsUserRole(ctx) {
			if config.Cfg.UserShieldError.Open && len(config.Cfg.UserShieldError.Errors) > 0 {
				for _, shieldError := range config.Cfg.UserShieldError.Errors {
					if gstr.Contains(image.ErrMsg, shieldError) {
						image.ErrMsg = "详细错误信息请联系管理员..."
						break
					}
				}
			}
		}

		image.ErrMsg = gstr.Split(image.ErrMsg, " TraceId")[0]
		image.ErrMsg = gstr.Split(image.ErrMsg, " (request id:")[0]
	}

	if service.Session().IsAdminRole(ctx) {
		image.ProviderId = result.ProviderId
		image.ModelId = result.ModelId
		image.ModelName = result.ModelName
		image.Key = util.Desensitize(result.Key)
		image.IsEnablePresetConfig = result.IsEnablePresetConfig
		image.IsEnableModelAgent = result.IsEnableModelAgent
		image.ModelAgentId = result.ModelAgentId
		image.IsEnableForward = result.IsEnableForward
		image.ForwardConfig = result.ForwardConfig
		image.IsSmartMatch = result.IsSmartMatch
		image.IsEnableFallback = result.IsEnableFallback
		image.FallbackConfig = result.FallbackConfig
		image.RealModelId = result.RealModelId
		image.RealModelName = result.RealModelName
		image.RealModel = result.RealModel
		image.RemoteIp = result.RemoteIp
		image.LocalIp = result.LocalIp
		image.InternalTime = result.InternalTime
		image.ErrMsg = result.ErrMsg
		image.IsRetry = result.IsRetry
		image.CreatedAt = util.FormatDateTime(result.CreatedAt)
		image.UpdatedAt = util.FormatDateTime(result.UpdatedAt)

		if result.ModelAgent != nil {

			providerName := result.ModelAgent.ProviderId
			if provider, err := dao.Provider.FindById(ctx, result.ModelAgent.ProviderId); err == nil && provider != nil {
				providerName = provider.Name
			}

			image.ModelAgent = &model.ModelAgent{
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

	return image, nil
}

// 绘图日志分页列表
func (s *sImage) Page(ctx context.Context, params model.ImagePageReq) (*model.ImagePageRes, error) {

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

	results, err := dao.Image.FindByPage(ctx, paging, filter, &dao.FindOptions{SortFields: []string{"-req_time", "status", "-created_at"}})
	if err != nil {
		logger.Error(ctx, err)
		return nil, err
	}

	items := make([]*model.Image, 0)
	for _, result := range results {

		image := &model.Image{
			Id:           result.Id,
			UserId:       result.UserId,
			AppId:        result.AppId,
			Model:        result.Model,
			Prompt:       result.Prompt,
			ImageQuota:   result.ImageQuota,
			InputTokens:  result.InputTokens,
			OutputTokens: result.OutputTokens,
			TextTokens:   result.TextTokens,
			ImageTokens:  result.ImageTokens,
			TotalTokens:  result.TotalTokens,
			TotalTime:    result.TotalTime,
			ReqTime:      util.FormatDateTimeMonth(result.ReqTime),
			Status:       result.Status,
		}

		for _, imageData := range result.ImageData {
			//if imageData.URL != "" {
			image.Images = append(image.Images, imageData.URL)
			//} else { // 太大了, 不查
			//	image.Images = append(image.Images, "data:image/png;base64,"+imageData.B64JSON)
			//}
		}

		if service.Session().IsAdminRole(ctx) {
			image.InternalTime = result.InternalTime
			image.IsSmartMatch = result.IsSmartMatch
		}

		items = append(items, image)
	}

	return &model.ImagePageRes{
		Items: items,
		Paging: &model.Paging{
			Page:     paging.Page,
			PageSize: paging.PageSize,
			Total:    paging.Total,
		},
	}, nil
}

// 绘图日志详情复制字段值
func (s *sImage) CopyField(ctx context.Context, params model.ImageCopyFieldReq) (string, error) {

	result, err := dao.Image.FindById(ctx, params.Id)
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
