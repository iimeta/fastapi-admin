package sys_config

import (
	"context"
	"github.com/gogf/gf/v2/os/gctx"
	"github.com/gogf/gf/v2/os/grpool"
	"github.com/iimeta/fastapi-admin/internal/config"
	"github.com/iimeta/fastapi-admin/internal/consts"
	"github.com/iimeta/fastapi-admin/internal/dao"
	"github.com/iimeta/fastapi-admin/internal/errors"
	"github.com/iimeta/fastapi-admin/internal/model"
	"github.com/iimeta/fastapi-admin/internal/model/common"
	"github.com/iimeta/fastapi-admin/internal/model/do"
	"github.com/iimeta/fastapi-admin/internal/model/entity"
	"github.com/iimeta/fastapi-admin/internal/service"
	"github.com/iimeta/fastapi-admin/utility/logger"
	"github.com/iimeta/fastapi-admin/utility/redis"
	"github.com/iimeta/fastapi-admin/utility/util"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"time"
)

type sSysConfig struct{}

func init() {

	ctx := gctx.New()
	sSysConfig := New()

	service.RegisterSysConfig(sSysConfig)
	if _, err := sSysConfig.Init(ctx); err != nil {
		panic(err)
	}

	conn, _, err := redis.Subscribe(ctx, consts.CHANGE_CHANNEL_CONFIG)
	if err != nil {
		panic(err)
	}

	if err = grpool.AddWithRecover(ctx, func(ctx context.Context) {
		for {

			msg, err := conn.ReceiveMessage(ctx)
			if err != nil {
				logger.Errorf(ctx, "sSysConfig Subscribe error: %v", err)
				time.Sleep(5 * time.Second)
				if conn, _, err = redis.Subscribe(ctx, consts.CHANGE_CHANNEL_CONFIG); err != nil {
					logger.Errorf(ctx, "sSysConfig Subscribe Reconnect error: %v", err)
				} else {
					logger.Info(ctx, "sSysConfig Subscribe Reconnect success")
				}
				continue
			}

			switch msg.Channel {
			case consts.CHANGE_CHANNEL_CONFIG:
				_, err = service.SysConfig().Init(ctx)
			}

			if err != nil {
				logger.Error(ctx, err)
			}
		}
	}, nil); err != nil {
		panic(err)
	}
}

func New() service.ISysConfig {
	return &sSysConfig{}
}

// 更新配置
func (s *sSysConfig) Update(ctx context.Context, params model.SysConfigUpdateReq) (*entity.SysConfig, error) {

	defer func() {
		if _, err := redis.Publish(ctx, consts.CHANGE_CHANNEL_CONFIG, model.PubMessage{
			Action: consts.ACTION_UPDATE,
		}); err != nil {
			logger.Error(ctx, err)
		}
	}()

	sysConfig := &do.SysConfig{}
	switch params.Action {
	case "core":
		sysConfig = &do.SysConfig{Core: params.Core}
	case "http":
		sysConfig = &do.SysConfig{Http: params.Http}
	case "email":
		sysConfig = &do.SysConfig{Email: params.Email}
	case "statistics":
		sysConfig = &do.SysConfig{Statistics: params.Statistics}
	case "base":
		sysConfig = &do.SysConfig{Base: params.Base}
	case "midjourney":
		sysConfig = &do.SysConfig{Midjourney: params.Midjourney}
	case "log":
		sysConfig = &do.SysConfig{Log: params.Log}
	case "user_shield_error":
		sysConfig = &do.SysConfig{UserShieldError: params.UserShieldError}
	case "auto_disabled_error":
		sysConfig = &do.SysConfig{AutoDisabledError: params.AutoDisabledError}
	case "not_retry_error":
		sysConfig = &do.SysConfig{NotRetryError: params.NotRetryError}
	case "not_shield_error":
		sysConfig = &do.SysConfig{NotShieldError: params.NotShieldError}
	case "debug":
		sysConfig = &do.SysConfig{Debug: params.Debug}
	}

	return dao.SysConfig.FindOneAndUpdate(ctx, bson.M{}, sysConfig)
}

// 更改配置状态
func (s *sSysConfig) ChangeStatus(ctx context.Context, params model.SysConfigChangeStatusReq) error {

	defer func() {
		if _, err := redis.Publish(ctx, consts.CHANGE_CHANNEL_CONFIG, model.PubMessage{
			Action: consts.ACTION_STATUS,
		}); err != nil {
			logger.Error(ctx, err)
		}
	}()

	if err := dao.SysConfig.UpdateOne(ctx, bson.M{}, bson.M{
		params.Action + ".open": params.Open,
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
		Id:                sysConfig.Id,
		Core:              sysConfig.Core,
		Http:              sysConfig.Http,
		Email:             sysConfig.Email,
		Statistics:        sysConfig.Statistics,
		Base:              sysConfig.Base,
		Midjourney:        sysConfig.Midjourney,
		Log:               sysConfig.Log,
		UserShieldError:   sysConfig.UserShieldError,
		AutoDisabledError: sysConfig.AutoDisabledError,
		NotRetryError:     sysConfig.NotRetryError,
		NotShieldError:    sysConfig.NotShieldError,
		Debug:             sysConfig.Debug,
		Creator:           sysConfig.Creator,
		Updater:           sysConfig.Updater,
		CreatedAt:         util.FormatDateTime(sysConfig.CreatedAt),
		UpdatedAt:         util.FormatDateTime(sysConfig.UpdatedAt),
	}, nil
}

// 重置配置
func (s *sSysConfig) Reset(ctx context.Context, params model.SysConfigResetReq) (*entity.SysConfig, error) {

	sysConfigUpdateReq := model.SysConfigUpdateReq{
		Action: params.Action,
	}

	switch params.Action {
	case "core":
		sysConfigUpdateReq.Core = s.Default().Core
	case "http":
		sysConfigUpdateReq.Http = s.Default().Http
	case "email":
		sysConfigUpdateReq.Email = s.Default().Email
	case "statistics":
		sysConfigUpdateReq.Statistics = s.Default().Statistics
	case "base":
		sysConfigUpdateReq.Base = s.Default().Base
	case "midjourney":
		sysConfigUpdateReq.Midjourney = s.Default().Midjourney
	case "log":
		sysConfigUpdateReq.Log = s.Default().Log
	case "user_shield_error":
		sysConfigUpdateReq.UserShieldError = s.Default().UserShieldError
	case "auto_disabled_error":
		sysConfigUpdateReq.AutoDisabledError = s.Default().AutoDisabledError
	case "not_retry_error":
		sysConfigUpdateReq.NotRetryError = s.Default().NotRetryError
	case "not_shield_error":
		sysConfigUpdateReq.NotShieldError = s.Default().NotShieldError
	}

	return s.Update(ctx, sysConfigUpdateReq)
}

// 初始化配置
func (s *sSysConfig) Init(ctx context.Context) (sysConfig *entity.SysConfig, err error) {

	defer func() {
		if err == nil && sysConfig != nil {
			config.Reload(ctx, sysConfig)
		}
	}()

	sysConfig, err = dao.SysConfig.FindOne(ctx, bson.M{})
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			id, err := dao.SysConfig.Insert(ctx, s.Default())
			if err != nil {
				logger.Error(ctx, err)
				return nil, err
			}
			return dao.SysConfig.FindById(ctx, id)
		} else {
			logger.Error(ctx, err)
			return nil, err
		}
	}

	if sysConfig.Core == nil {
		return s.Reset(ctx, model.SysConfigResetReq{Action: "core"})
	}

	if sysConfig.Http == nil {
		return s.Reset(ctx, model.SysConfigResetReq{Action: "http"})
	}

	if sysConfig.Email == nil {
		return s.Reset(ctx, model.SysConfigResetReq{Action: "email"})
	}

	if sysConfig.Statistics == nil {
		return s.Reset(ctx, model.SysConfigResetReq{Action: "statistics"})
	}

	if sysConfig.Base == nil {
		return s.Reset(ctx, model.SysConfigResetReq{Action: "base"})
	}

	if sysConfig.Midjourney == nil {
		return s.Reset(ctx, model.SysConfigResetReq{Action: "midjourney"})
	}

	if sysConfig.Log == nil {
		return s.Reset(ctx, model.SysConfigResetReq{Action: "log"})
	}

	if sysConfig.UserShieldError == nil {
		return s.Reset(ctx, model.SysConfigResetReq{Action: "user_shield_error"})
	}

	if sysConfig.AutoDisabledError == nil {
		return s.Reset(ctx, model.SysConfigResetReq{Action: "auto_disabled_error"})
	}

	if sysConfig.NotRetryError == nil {
		return s.Reset(ctx, model.SysConfigResetReq{Action: "not_retry_error"})
	}

	if sysConfig.NotShieldError == nil {
		return s.Reset(ctx, model.SysConfigResetReq{Action: "not_shield_error"})
	}

	return sysConfig, nil
}

// 默认配置
func (s *sSysConfig) Default() *do.SysConfig {
	return &do.SysConfig{
		Core: &common.Core{
			SecretKeyPrefix: "sk-FastAPI",
			ErrorPrefix:     "fastapi",
		},
		Http: &common.Http{
			Timeout: 60,
		},
		Email: &common.Email{
			Host:     "smtp.xxx.com",
			Port:     465,
			UserName: "xxx@xxx.com",
			Password: "xxx",
			FromName: "智元 Fast API",
		},
		Statistics: &common.Statistics{
			Open:        true,
			Cron:        "0 0/5 * * * ?",
			Limit:       1000,
			LockMinutes: 30,
		},
		Base: &common.Base{
			ErrRetry:                3,
			ModelKeyErrDisable:      100000,
			ModelAgentErrDisable:    100000,
			ModelAgentKeyErrDisable: 100000,
		},
		Midjourney: &common.Midjourney{
			CdnUrl:          "https://cdn.xxx.com",
			ApiBaseUrl:      "https://xxx/mj",
			ApiSecret:       "xxx",
			ApiSecretHeader: "mj-api-secret",
			CdnOriginalUrl:  "https://cdn.discordapp.com",
		},
		Log: &common.Log{
			Open: true,
			Records: []string{
				"prompt",
				"completion",
				"messages",
				"image",
			},
		},
		UserShieldError: &common.UserShieldError{
			Open: true,
			Errors: []string{
				"TraceId",
				"http",
				"tcp",
				"No available",
				"quota",
				"All key error.",
				"All model agent error.",
				"All model agent key error.",
			},
		},
		AutoDisabledError: &common.AutoDisabledError{
			Open: true,
			Errors: []string{
				"Incorrect API key provided or has been disabled.",
				"You exceeded your current quota.",
				"The OpenAI account associated with this API key has been deactivated.",
				"PERMISSION_DENIED",
				"BILLING_DISABLED",
				"ACCESS_TOKEN_EXPIRED",
				"is not allowed to use Publisher Model",
				"Resource has been exhausted",
				"IAM_PERMISSION_DENIED",
				"SERVICE_DISABLED",
				"ACCOUNT_STATE_INVALID",
			},
		},
		NotRetryError: &common.NotRetryError{
			Open: true,
			Errors: []string{
				"Please reduce the length of the messages.",
			},
		},
		NotShieldError: &common.NotShieldError{
			Open: true,
			Errors: []string{
				"Please reduce the length of the messages.",
			},
		},
		Debug: &common.Debug{
			Open: false,
		},
	}
}
