package entity

import (
	"github.com/gogf/gf/v2/util/gmeta"
	"github.com/iimeta/fastapi-admin/internal/model/common"
)

type SysConfig struct {
	gmeta.Meta `role:"admin" bson:"-"`
	Id         string             `bson:"_id,omitempty"`        // ID
	Core       *common.Core       `bson:"core,omitempty"`       // 核心
	Http       *common.Http       `bson:"http,omitempty"`       // HTTP
	Email      *common.Email      `bson:"email,omitempty"`      // 邮箱
	Statistics *common.Statistics `bson:"statistics,omitempty"` // 统计
	Api        *common.Api        `bson:"api,omitempty"`        // API
	Midjourney *common.Midjourney `bson:"midjourney,omitempty"` // Midjourney
	Gcp        *common.Gcp        `bson:"gcp,omitempty"`        // GCP
	Log        *common.Log        `bson:"log,omitempty"`        // 日志
	Error      *common.Error      `bson:"error,omitempty"`      // 错误配置
	Debug      *common.Debug      `bson:"debug,omitempty"`      // 调试
	Creator    string             `bson:"creator,omitempty"`    // 创建人
	Updater    string             `bson:"updater,omitempty"`    // 更新人
	CreatedAt  int64              `bson:"created_at,omitempty"` // 创建时间
	UpdatedAt  int64              `bson:"updated_at,omitempty"` // 更新时间
}
