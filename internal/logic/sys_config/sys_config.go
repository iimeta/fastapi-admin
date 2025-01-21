package sys_config

import (
	"context"
	"github.com/iimeta/fastapi-admin/internal/dao"
	"github.com/iimeta/fastapi-admin/internal/model"
	"github.com/iimeta/fastapi-admin/internal/model/do"
	"github.com/iimeta/fastapi-admin/internal/service"
	"github.com/iimeta/fastapi-admin/utility/logger"
	"github.com/iimeta/fastapi-admin/utility/util"
	"go.mongodb.org/mongo-driver/bson"
)

type sSysConfig struct{}

func init() {
	service.RegisterSysConfig(New())
}

func New() service.ISysConfig {
	return &sSysConfig{}
}

// 更新配置
func (s *sSysConfig) Update(ctx context.Context, params model.SysConfigUpdateReq) error {

	sysConfig := &do.SysConfig{
		Core:       params.Core,
		Http:       params.Http,
		Email:      params.Email,
		Statistics: params.Statistics,
		Api:        params.Api,
		Midjourney: params.Midjourney,
		Gcp:        params.Gcp,
		Log:        params.Log,
		Error:      params.Error,
		Debug:      params.Debug,
	}

	if err := dao.SysConfig.UpdateById(ctx, params.Id, sysConfig); err != nil {
		logger.Error(ctx, err)
		return err
	}

	return nil
}

// 更改配置状态
func (s *sSysConfig) ChangeStatus(ctx context.Context, params model.SysConfigChangeStatusReq) error {

	if err := dao.SysConfig.UpdateById(ctx, params.Id, bson.M{
		"status": params.Status,
	}); err != nil {
		logger.Error(ctx, err)
		return err
	}

	return nil
}

// 配置详情
func (s *sSysConfig) Detail(ctx context.Context) (*model.SysConfig, error) {

	sysConfig, err := dao.SysConfig.FindOne(ctx, bson.M{})
	if err != nil {
		logger.Error(ctx, err)
		return nil, err
	}

	return &model.SysConfig{
		Id:         sysConfig.Id,
		Core:       sysConfig.Core,
		Http:       sysConfig.Http,
		Email:      sysConfig.Email,
		Statistics: sysConfig.Statistics,
		Api:        sysConfig.Api,
		Midjourney: sysConfig.Midjourney,
		Gcp:        sysConfig.Gcp,
		Log:        sysConfig.Log,
		Error:      sysConfig.Error,
		Debug:      sysConfig.Debug,
		Creator:    sysConfig.Creator,
		Updater:    sysConfig.Updater,
		CreatedAt:  util.FormatDateTime(sysConfig.CreatedAt),
		UpdatedAt:  util.FormatDateTime(sysConfig.UpdatedAt),
	}, nil
}
