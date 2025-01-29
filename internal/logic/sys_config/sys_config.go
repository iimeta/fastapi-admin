package sys_config

import (
	"context"
	"github.com/gogf/gf/v2/os/gctx"
	"github.com/iimeta/fastapi-admin/internal/dao"
	"github.com/iimeta/fastapi-admin/internal/errors"
	"github.com/iimeta/fastapi-admin/internal/model"
	"github.com/iimeta/fastapi-admin/internal/model/common"
	"github.com/iimeta/fastapi-admin/internal/model/do"
	"github.com/iimeta/fastapi-admin/internal/model/entity"
	"github.com/iimeta/fastapi-admin/internal/service"
	"github.com/iimeta/fastapi-admin/utility/logger"
	"github.com/iimeta/fastapi-admin/utility/util"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type sSysConfig struct{}

func init() {
	sSysConfig := New()
	service.RegisterSysConfig(sSysConfig)
	_, err := sSysConfig.Default(gctx.New())
	if err != nil {
		panic(err)
	}
}

func New() service.ISysConfig {
	return &sSysConfig{}
}

// 更新配置
func (s *sSysConfig) Update(ctx context.Context, params model.SysConfigUpdateReq) (*entity.SysConfig, error) {

	sysConfig := &do.SysConfig{}
	switch params.Action {
	case "core":
		sysConfig = &do.SysConfig{
			Core: params.Core,
		}
	case "http":
		sysConfig = &do.SysConfig{
			Http: params.Http,
		}
	case "email":
		sysConfig = &do.SysConfig{
			Email: params.Email,
		}
	case "statistics":
		sysConfig = &do.SysConfig{
			Statistics: params.Statistics,
		}
	case "api":
		sysConfig = &do.SysConfig{
			Api: params.Api,
		}
	case "midjourney":
		sysConfig = &do.SysConfig{
			Midjourney: params.Midjourney,
		}
	case "log":
		sysConfig = &do.SysConfig{
			Log: params.Log,
		}
	case "error":
		sysConfig = &do.SysConfig{
			Error: params.Error,
		}
	case "debug":
		sysConfig = &do.SysConfig{
			Debug: params.Debug,
		}
	}

	return dao.SysConfig.FindOneAndUpdate(ctx, bson.M{}, sysConfig)
}

// 更改配置状态
func (s *sSysConfig) ChangeStatus(ctx context.Context, params model.SysConfigChangeStatusReq) error {

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
		Id:         sysConfig.Id,
		Core:       sysConfig.Core,
		Http:       sysConfig.Http,
		Email:      sysConfig.Email,
		Statistics: sysConfig.Statistics,
		Api:        sysConfig.Api,
		Midjourney: sysConfig.Midjourney,
		Log:        sysConfig.Log,
		Error:      sysConfig.Error,
		Debug:      sysConfig.Debug,
		Creator:    sysConfig.Creator,
		Updater:    sysConfig.Updater,
		CreatedAt:  util.FormatDateTime(sysConfig.CreatedAt),
		UpdatedAt:  util.FormatDateTime(sysConfig.UpdatedAt),
	}, nil
}

// 重置配置
func (s *sSysConfig) Reset(ctx context.Context, params model.SysConfigResetReq) (*entity.SysConfig, error) {

	sysConfigUpdateReq := model.SysConfigUpdateReq{
		Action: params.Action,
	}
	switch params.Action {
	case "core":
		sysConfigUpdateReq.Core = &common.Core{
			SecretKeyPrefix: "sk-FastAPI",
			ErrorPrefix:     "fastapi",
		}
	case "http":
		sysConfigUpdateReq.Http = &common.Http{
			Timeout: 60,
		}
	case "email":
		sysConfigUpdateReq.Email = &common.Email{
			Host:     "smtp.xxx.com",
			Port:     465,
			UserName: "xxx@xxx.com",
			Password: "xxx",
			FromName: "智元 Fast API",
		}
	case "statistics":
		sysConfigUpdateReq.Statistics = &common.Statistics{
			Open:        true,
			Cron:        "0 0/5 * * * ?",
			Limit:       1000,
			LockMinutes: 30,
		}
	case "api":
		sysConfigUpdateReq.Api = &common.Api{
			Retry:                   3,
			ModelKeyErrDisable:      100000,
			ModelAgentErrDisable:    100000,
			ModelAgentKeyErrDisable: 100000,
		}
	case "midjourney":
		sysConfigUpdateReq.Midjourney = &common.Midjourney{
			CdnUrl:          "https://cdn.xxx.com",
			ApiBaseUrl:      "https://xxx/mj",
			ApiSecret:       "xxx",
			ApiSecretHeader: "mj-api-secret",
			CdnOriginalUrl:  "https://cdn.discordapp.com",
		}
	case "log":
		sysConfigUpdateReq.Log = &common.Log{
			Open: true,
			Records: []string{
				"prompt",
				"completion",
				"messages",
				"image",
			},
		}
	case "error":
		sysConfigUpdateReq.Error = &common.Error{
			Open: true,
			ShieldUser: []string{
				"TraceId",
				"http",
				"tcp",
				"No available",
				"quota",
				"All key error.",
				"All model agent error.",
				"All model agent key error.",
			},
			AutoDisabled: []string{
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
			NotRetry: []string{
				"Please reduce the length of the messages.",
			},
			NotShield: []string{
				"Please reduce the length of the messages.",
			},
		}
	case "all":
		if _, err := dao.SysConfig.DeleteOne(ctx, bson.M{}); err != nil {
			logger.Error(ctx, err)
			return nil, err
		}
		return s.Default(ctx)
	}

	return s.Update(ctx, sysConfigUpdateReq)
}

// 默认配置
func (s *sSysConfig) Default(ctx context.Context) (*entity.SysConfig, error) {

	sysConfig, err := dao.SysConfig.FindOne(ctx, bson.M{})
	if err != nil && errors.Is(err, mongo.ErrNoDocuments) {

		id, err := dao.SysConfig.Insert(ctx, &do.SysConfig{
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
			Api: &common.Api{
				Retry:                   3,
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
			Error: &common.Error{
				Open: true,
				ShieldUser: []string{
					"TraceId",
					"http",
					"tcp",
					"No available",
					"quota",
					"All key error.",
					"All model agent error.",
					"All model agent key error.",
				},
				AutoDisabled: []string{
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
				NotRetry: []string{
					"Please reduce the length of the messages.",
				},
				NotShield: []string{
					"Please reduce the length of the messages.",
				},
			},
		})
		if err != nil {
			logger.Error(ctx, err)
			return nil, err
		}

		return dao.SysConfig.FindById(ctx, id)
	}

	return sysConfig, err
}
