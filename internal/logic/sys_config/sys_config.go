package sys_config

import (
	"context"
	"time"

	"github.com/gogf/gf/v2/os/gcron"
	"github.com/gogf/gf/v2/os/gctx"
	"github.com/gogf/gf/v2/os/grpool"
	"github.com/gogf/gf/v2/os/gtimer"
	"github.com/iimeta/fastapi-admin/v2/internal/config"
	"github.com/iimeta/fastapi-admin/v2/internal/consts"
	"github.com/iimeta/fastapi-admin/v2/internal/dao"
	"github.com/iimeta/fastapi-admin/v2/internal/errors"
	"github.com/iimeta/fastapi-admin/v2/internal/model"
	"github.com/iimeta/fastapi-admin/v2/internal/model/common"
	"github.com/iimeta/fastapi-admin/v2/internal/model/do"
	"github.com/iimeta/fastapi-admin/v2/internal/model/entity"
	"github.com/iimeta/fastapi-admin/v2/internal/service"
	"github.com/iimeta/fastapi-admin/v2/internal/task"
	"github.com/iimeta/fastapi-admin/v2/utility/logger"
	"github.com/iimeta/fastapi-admin/v2/utility/redis"
	"github.com/iimeta/fastapi-admin/v2/utility/util"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type sSysConfig struct{}

func init() {

	ctx := gctx.New()
	sSysConfig := New()

	service.RegisterSysConfig(sSysConfig)
	if _, err := sSysConfig.Init(ctx); err != nil {
		panic(err)
	}

	_, _ = gcron.AddSingleton(gctx.New(), "0 0/30 * * * ?", func(ctx context.Context) {
		_, _ = service.SysConfig().Init(ctx)
	})

	_ = gtimer.AddSingleton(gctx.New(), 30*time.Minute, func(ctx context.Context) {
		_, _ = service.SysConfig().Init(ctx)
	})

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
			case config.Cfg.Core.ChannelPrefix + consts.CHANGE_CHANNEL_CONFIG:
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
	//case "midjourney":
	//	sysConfig = &do.SysConfig{Midjourney: params.Midjourney}
	case "log":
		sysConfig = &do.SysConfig{Log: params.Log}
	case "user_login_register":
		sysConfig = &do.SysConfig{UserLoginRegister: params.UserLoginRegister}
	case "user_shield_error":
		sysConfig = &do.SysConfig{UserShieldError: params.UserShieldError}
	case "reseller_login_register":
		sysConfig = &do.SysConfig{ResellerLoginRegister: params.ResellerLoginRegister}
	case "reseller_shield_error":
		sysConfig = &do.SysConfig{ResellerShieldError: params.ResellerShieldError}
	case "admin_login":
		sysConfig = &do.SysConfig{AdminLogin: params.AdminLogin}
	case "auto_disabled_error":
		sysConfig = &do.SysConfig{AutoDisabledError: params.AutoDisabledError}
	case "auto_enable_error":
		sysConfig = &do.SysConfig{AutoEnableError: params.AutoEnableError}
	case "not_retry_error":
		sysConfig = &do.SysConfig{NotRetryError: params.NotRetryError}
	case "not_shield_error":
		sysConfig = &do.SysConfig{NotShieldError: params.NotShieldError}
	case "notice":
		sysConfig = &do.SysConfig{Notice: params.Notice}
	case "quota":
		if params.Quota != nil {
			params.Quota.Threshold *= consts.QUOTA_DEFAULT_UNIT
		}
		sysConfig = &do.SysConfig{Quota: params.Quota}
	case "quota_task":
		sysConfig = &do.SysConfig{QuotaTask: params.QuotaTask}
	case "video_task":
		sysConfig = &do.SysConfig{VideoTask: params.VideoTask}
	case "file_task":
		sysConfig = &do.SysConfig{FileTask: params.FileTask}
	case "batch_task":
		sysConfig = &do.SysConfig{BatchTask: params.BatchTask}
	case "service_unavailable":
		sysConfig = &do.SysConfig{ServiceUnavailable: params.ServiceUnavailable}
	case "general_api":
		sysConfig = &do.SysConfig{GeneralApi: params.GeneralApi}
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

	if sysConfig.Quota != nil {
		sysConfig.Quota.Threshold /= consts.QUOTA_DEFAULT_UNIT
	}

	return &model.SysConfig{
		Id:         sysConfig.Id,
		Core:       sysConfig.Core,
		Http:       sysConfig.Http,
		Email:      sysConfig.Email,
		Statistics: sysConfig.Statistics,
		Base:       sysConfig.Base,
		//Midjourney:            sysConfig.Midjourney,
		Log:                   sysConfig.Log,
		UserLoginRegister:     sysConfig.UserLoginRegister,
		UserShieldError:       sysConfig.UserShieldError,
		ResellerLoginRegister: sysConfig.ResellerLoginRegister,
		ResellerShieldError:   sysConfig.ResellerShieldError,
		AdminLogin:            sysConfig.AdminLogin,
		AutoDisabledError:     sysConfig.AutoDisabledError,
		AutoEnableError:       sysConfig.AutoEnableError,
		NotRetryError:         sysConfig.NotRetryError,
		NotShieldError:        sysConfig.NotShieldError,
		Notice:                sysConfig.Notice,
		Quota:                 sysConfig.Quota,
		QuotaTask:             sysConfig.QuotaTask,
		VideoTask:             sysConfig.VideoTask,
		FileTask:              sysConfig.FileTask,
		BatchTask:             sysConfig.BatchTask,
		ServiceUnavailable:    sysConfig.ServiceUnavailable,
		GeneralApi:            sysConfig.GeneralApi,
		Debug:                 sysConfig.Debug,
		Creator:               sysConfig.Creator,
		Updater:               sysConfig.Updater,
		CreatedAt:             util.FormatDateTime(sysConfig.CreatedAt),
		UpdatedAt:             util.FormatDateTime(sysConfig.UpdatedAt),
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
	//case "midjourney":
	//	sysConfigUpdateReq.Midjourney = s.Default().Midjourney
	case "log":
		sysConfigUpdateReq.Log = s.Default().Log
	case "user_login_register":
		sysConfigUpdateReq.UserLoginRegister = s.Default().UserLoginRegister
	case "user_shield_error":
		sysConfigUpdateReq.UserShieldError = s.Default().UserShieldError
	case "reseller_login_register":
		sysConfigUpdateReq.ResellerLoginRegister = s.Default().ResellerLoginRegister
	case "reseller_shield_error":
		sysConfigUpdateReq.ResellerShieldError = s.Default().ResellerShieldError
	case "admin_login":
		sysConfigUpdateReq.AdminLogin = s.Default().AdminLogin
	case "auto_disabled_error":
		sysConfigUpdateReq.AutoDisabledError = s.Default().AutoDisabledError
	case "auto_enable_error":
		sysConfigUpdateReq.AutoEnableError = s.Default().AutoEnableError
	case "not_retry_error":
		sysConfigUpdateReq.NotRetryError = s.Default().NotRetryError
	case "not_shield_error":
		sysConfigUpdateReq.NotShieldError = s.Default().NotShieldError
	case "reset_api_error":

		keys, err := redis.Keys(ctx, "api:error:*")
		if err != nil {
			logger.Error(ctx, err)
			return nil, err
		}

		if len(keys) > 0 {
			if _, err = redis.Del(ctx, keys...); err != nil {
				logger.Error(ctx, err)
				return nil, err
			}
		}

		return nil, nil
	case "notice":
		sysConfigUpdateReq.Notice = s.Default().Notice
	case "quota":
		sysConfigUpdateReq.Quota = s.Default().Quota
	case "quota_task":
		sysConfigUpdateReq.QuotaTask = s.Default().QuotaTask
	case "video_task":
		sysConfigUpdateReq.VideoTask = s.Default().VideoTask
	case "file_task":
		sysConfigUpdateReq.FileTask = s.Default().FileTask
	case "batch_task":
		sysConfigUpdateReq.BatchTask = s.Default().BatchTask
	case "service_unavailable":
		sysConfigUpdateReq.ServiceUnavailable = s.Default().ServiceUnavailable
	case "general_api":
		sysConfigUpdateReq.GeneralApi = s.Default().GeneralApi
	}

	return s.Update(ctx, sysConfigUpdateReq)
}

// 刷新配置
func (s *sSysConfig) Refresh(ctx context.Context, params model.SysConfigRefreshReq) error {

	switch params.Action {
	case "refresh_api_cache":
		if _, err := redis.Publish(ctx, consts.REFRESH_CHANNEL_API, model.PubMessage{
			Action: consts.ACTION_CACHE,
		}); err != nil {
			logger.Error(ctx, err)
			return err
		}
	}

	return nil
}

// 系统配置
func (s *sSysConfig) Config(ctx context.Context) (*model.SysConfig, error) {

	sysConfig, err := dao.SysConfig.FindOne(ctx, bson.M{})
	if err != nil {
		logger.Error(ctx, err)
		return nil, err
	}

	return &model.SysConfig{
		UserLoginRegister:     sysConfig.UserLoginRegister,
		ResellerLoginRegister: sysConfig.ResellerLoginRegister,
		AdminLogin:            sysConfig.AdminLogin,
	}, nil
}

// 初始化配置
func (s *sSysConfig) Init(ctx context.Context) (sysConfig *entity.SysConfig, err error) {

	defer func() {
		if err == nil && sysConfig != nil {
			config.Reload(ctx, sysConfig)
			task.Init(ctx)
		}
	}()

	if sysConfig, err = dao.SysConfig.FindOne(ctx, bson.M{}); err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			id, err := dao.SysConfig.Insert(ctx, s.Default())
			if err != nil {
				logger.Error(ctx, err)
				return nil, err
			}
			return dao.SysConfig.FindById(ctx, id)
		}
		logger.Error(ctx, err)
		return nil, err
	}

	if sysConfig.Core == nil {
		if sysConfig, err = s.Reset(ctx, model.SysConfigResetReq{Action: "core"}); err != nil {
			logger.Error(ctx, err)
			return nil, err
		}
	}

	if sysConfig.Notice == nil {
		if sysConfig, err = s.Reset(ctx, model.SysConfigResetReq{Action: "notice"}); err != nil {
			logger.Error(ctx, err)
			return nil, err
		}
	}

	if sysConfig.Quota == nil {
		if sysConfig, err = s.Reset(ctx, model.SysConfigResetReq{Action: "quota"}); err != nil {
			logger.Error(ctx, err)
			return nil, err
		}
	}

	if sysConfig.QuotaTask == nil {
		if sysConfig, err = s.Reset(ctx, model.SysConfigResetReq{Action: "quota_task"}); err != nil {
			logger.Error(ctx, err)
			return nil, err
		}
	}

	if sysConfig.ResellerLoginRegister == nil {
		if sysConfig, err = s.Reset(ctx, model.SysConfigResetReq{Action: "reseller_login_register"}); err != nil {
			logger.Error(ctx, err)
			return nil, err
		}
	}

	if sysConfig.ResellerShieldError == nil {
		if sysConfig, err = s.Reset(ctx, model.SysConfigResetReq{Action: "reseller_shield_error"}); err != nil {
			logger.Error(ctx, err)
			return nil, err
		}
	}

	if sysConfig.ServiceUnavailable == nil {
		if sysConfig, err = s.Reset(ctx, model.SysConfigResetReq{Action: "service_unavailable"}); err != nil {
			logger.Error(ctx, err)
			return nil, err
		}
	}

	if sysConfig.VideoTask == nil {
		if sysConfig, err = s.Reset(ctx, model.SysConfigResetReq{Action: "video_task"}); err != nil {
			logger.Error(ctx, err)
			return nil, err
		}
	}

	if sysConfig.FileTask == nil {
		if sysConfig, err = s.Reset(ctx, model.SysConfigResetReq{Action: "file_task"}); err != nil {
			logger.Error(ctx, err)
			return nil, err
		}
	}

	if sysConfig.BatchTask == nil {
		if sysConfig, err = s.Reset(ctx, model.SysConfigResetReq{Action: "batch_task"}); err != nil {
			logger.Error(ctx, err)
			return nil, err
		}
	}

	if sysConfig.GeneralApi == nil {
		if sysConfig, err = s.Reset(ctx, model.SysConfigResetReq{Action: "general_api"}); err != nil {
			logger.Error(ctx, err)
			return nil, err
		}
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
			Timeout: 300,
		},
		Email: &common.Email{
			Host:     "smtp.xxx.com",
			Port:     465,
			UserName: "xxx@xxx.com",
			Password: "xxx",
			FromName: "智元 Fast API",
			Interval: 3000,
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
			ShortTimeout:            300,
			LongTimeout:             600,
		},
		//Midjourney: &common.Midjourney{
		//	CdnUrl:          "https://cdn.xxx.com",
		//	ApiBaseUrl:      "https://xxx/mj",
		//	ApiSecret:       "xxx",
		//	ApiSecretHeader: "mj-api-secret",
		//	CdnOriginalUrl:  "https://cdn.discordapp.com",
		//},
		Log: &common.Log{
			Open: true,
			TextRecords: []string{
				"prompt",
				"completion",
				"messages",
				"image",
				"audio",
			},
			TextReserve:    90,
			ImageReserve:   90,
			AudioReserve:   90,
			VideoReserve:   90,
			FileReserve:    90,
			BatchReserve:   90,
			GeneralReserve: 90,
			Status:         []int{1, 2, 3, -1},
			Cron:           "0 0 2 * * ?",
		},
		UserLoginRegister: &common.UserLoginRegister{
			AccountLogin:  true,
			EmailLogin:    true,
			EmailRegister: true,
			EmailRetrieve: true,
			SessionExpire: 3600 * 6,
		},
		UserShieldError: &common.UserShieldError{
			Open: true,
			Errors: []string{
				"http",
				"tcp",
				"No available",
				"quota",
				"All key error.",
				"All model agent error.",
				"All model agent key error.",
			},
		},
		ResellerLoginRegister: &common.ResellerLoginRegister{
			AccountLogin:  true,
			EmailLogin:    true,
			EmailRegister: false,
			EmailRetrieve: true,
			SessionExpire: 3600 * 6,
		},
		ResellerShieldError: &common.ResellerShieldError{
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
		AdminLogin: &common.AdminLogin{
			AccountLogin:  true,
			EmailLogin:    true,
			EmailRetrieve: true,
			SessionExpire: 3600 * 6,
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
				"on requests per min (RPM): Limit",
				"on tokens per min (TPM): Limit",
			},
		},
		AutoEnableError: &common.AutoEnableError{
			Open: true,
			EnableErrors: []common.EnableError{
				{
					Cron:       "0 * * * * ?",
					EnableTime: 20,
					Error:      "on requests per min (RPM): Limit",
				},
				{
					Cron:       "0 0 0/2 * * ?",
					EnableTime: 60 * 60 * 2,
					Error:      "on tokens per min (TPM): Limit",
				},
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
		Notice: &common.Notice{
			Open:        false,
			Cron:        "0 * * * * ?",
			LockMinutes: 10,
		},
		Quota: &common.Quota{
			Warning:           true,
			Threshold:         100 * consts.QUOTA_DEFAULT_UNIT,
			ExpiredWarning:    true,
			ExpiredThreshold:  3,
			ExhaustedNotice:   true,
			ExpiredNotice:     true,
			ExpiredClear:      false,
			ExpiredClearDefer: 5,
		},
		QuotaTask: &common.QuotaTask{
			Open:        true,
			Cron:        "0 * * * * ?",
			LockMinutes: 10,
		},
		VideoTask: &common.VideoTask{
			Open:            true,
			Cron:            "0/20 * * * * ?",
			LockMinutes:     30,
			IsEnableStorage: true,
		},
		FileTask: &common.FileTask{
			Open:            true,
			Cron:            "0/20 * * * * ?",
			LockMinutes:     30,
			IsEnableStorage: true,
		},
		BatchTask: &common.BatchTask{
			Open:        true,
			Cron:        "0/20 * * * * ?",
			LockMinutes: 30,
		},
		ServiceUnavailable: &common.ServiceUnavailable{
			Open: false,
			IpWhitelist: []string{
				"127.0.0.1",
				"::1",
				"172.17.0.1",
			},
		},
		GeneralApi: &common.GeneralApi{
			Open:        false,
			IpWhitelist: []string{},
		},
		Debug: &common.Debug{
			Open: false,
		},
	}
}
