package config

import (
	"context"
	"fmt"
	"github.com/gogf/gf/v2/container/gvar"
	"github.com/gogf/gf/v2/encoding/gjson"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gcfg"
	"github.com/gogf/gf/v2/os/gctx"
	"github.com/gogf/gf/v2/os/gfsnotify"
	"github.com/iimeta/fastapi-admin/utility/logger"
	"time"
)

var Cfg *Config

func init() {

	file, _ := gcfg.NewAdapterFile()
	path, _ := file.GetFilePath()

	if err := gjson.Unmarshal(gjson.MustEncode(gcfg.Instance().MustData(gctx.New())), &Cfg); err != nil {
		panic(fmt.Sprintf("解析配置文件 %s 错误: %v", path, err))
	}

	// 监听配置文件变化, 热加载
	_, _ = gfsnotify.Add(path, func(event *gfsnotify.Event) {
		ctx := gctx.New()
		if data, err := gcfg.Instance().Data(ctx); err != nil {
			logger.Errorf(ctx, "热加载 获取配置文件 %s 数据错误: %v", path, err)
		} else {
			if err = gjson.Unmarshal(gjson.MustEncode(data), &Cfg); err != nil {
				logger.Errorf(ctx, "热加载 解析配置文件 %s 错误: %v", path, err)
			}
		}
	})
}

// 配置信息
type Config struct {
	AdminServerAddress string `json:"admin_server_address"`
	App                *App   `json:"app"`
	Http               *Http  `json:"http"`
	Email              *Email `json:"email"`
	Debug              bool   `json:"debug"`
}

type App struct {
	Env        string   `json:"env"`
	Debug      bool     `json:"debug"`
	JuheKey    string   `json:"juhe_key"`
	JuheUrl    string   `json:"juhe_url"`
	AdminEmail []string `json:"admin_email"`
}

type Http struct {
	Timeout  time.Duration `json:"timeout"`
	ProxyUrl string        `json:"proxy_url"`
}

// 邮件配置信息
type Email struct {
	Host     string `json:"host"`     // smtp.xxx.com
	Port     int    `json:"port"`     // 端口号
	UserName string `json:"username"` // 登录账号
	Password string `json:"password"` // 登录密码
	FromName string `json:"fromname"` // 发送人名称
}

func Get(ctx context.Context, pattern string, def ...interface{}) (*gvar.Var, error) {

	value, err := g.Cfg().Get(ctx, pattern, def...)
	if err != nil {
		return nil, err
	}

	return value, nil
}

func GetString(ctx context.Context, pattern string, def ...interface{}) string {

	value, err := Get(ctx, pattern, def...)
	if err != nil {
		logger.Error(ctx, err)
	}

	return value.String()
}

func GetInt(ctx context.Context, pattern string, def ...interface{}) int {

	value, err := Get(ctx, pattern, def...)
	if err != nil {
		logger.Error(ctx, err)
	}

	return value.Int()
}

func GetBool(ctx context.Context, pattern string, def ...interface{}) (bool, error) {

	value, err := Get(ctx, pattern, def...)
	if err != nil {
		return false, err
	}

	return value.Bool(), nil
}

func GetMapStrStr(ctx context.Context, pattern string, def ...interface{}) map[string]string {

	value, err := Get(ctx, pattern, def...)
	if err != nil {
		logger.Error(ctx, err)
	}

	return value.MapStrStr()
}
