package common

import (
	"context"
	"encoding/json"
	"reflect"
	"slices"

	"github.com/gogf/gf/v2/text/gstr"
	"github.com/iimeta/fastapi-admin/v2/internal/config"
	"github.com/iimeta/fastapi-admin/v2/internal/dao"
	"github.com/iimeta/fastapi-admin/v2/internal/service"
	"go.mongodb.org/mongo-driver/v2/bson"
)

// 内部错误(与 API 端 errors.Error 中强制屏蔽的错误一致), 不享受"不屏蔽错误"白名单, 仍按屏蔽配置处理
var internalErrors = []string{
	"No available key.",
	"No available model agent.",
	"Model agent has been disabled.",
	"No available model agent key.",
	"All key error.",
	"All model agent error.",
	"All model agent key error.",
	"Model has been disabled.",
}

// 判断记录是否命中"不屏蔽错误"白名单(需开启"不屏蔽错误"开关), 与 API 端共用同一份白名单配置:
// userId 为记录所属用户ID, creator 为产生该记录的用户请求密钥(sk-)
func isNotShieldWhitelist(userId int, creator string, errMsg string) bool {

	if config.Cfg.NotShieldError == nil || !config.Cfg.NotShieldError.Open {
		return false
	}

	for _, internalError := range internalErrors {
		if gstr.Contains(errMsg, internalError) {
			return false
		}
	}

	// 按用户ID匹配
	if userId != 0 && slices.Contains(config.Cfg.NotShieldError.WhitelistUserIds, userId) {
		return true
	}

	// 按用户请求密钥(sk-)匹配
	if creator != "" && slices.Contains(config.Cfg.NotShieldError.WhitelistKeys, creator) {
		return true
	}

	return false
}

// 按角色转换日志错误信息: 管理员返回原文; 代理商/用户按配置屏蔽, 并去掉 TraceId / request id;
// userId、creator 为记录所属的用户ID和用户请求密钥, 用户查看命中"不屏蔽错误"白名单的记录时不按"用户屏蔽错误"屏蔽
func ConvErrMsg(ctx context.Context, errMsg string, status int, userId int, creator string) string {

	if service.Session().IsAdminRole(ctx) {
		return errMsg
	}

	if status != -1 || errMsg == "" {
		return ""
	}

	if service.Session().IsResellerRole(ctx) {
		if config.Cfg.ResellerShieldError != nil && config.Cfg.ResellerShieldError.Open && len(config.Cfg.ResellerShieldError.Errors) > 0 {
			for _, shieldError := range config.Cfg.ResellerShieldError.Errors {
				if gstr.Contains(errMsg, shieldError) {
					errMsg = "详细错误信息请联系管理员..."
					break
				}
			}
		}
	}

	if service.Session().IsUserRole(ctx) && !isNotShieldWhitelist(userId, creator, errMsg) {
		if config.Cfg.UserShieldError != nil && config.Cfg.UserShieldError.Open && len(config.Cfg.UserShieldError.Errors) > 0 {
			for _, shieldError := range config.Cfg.UserShieldError.Errors {
				if gstr.Contains(errMsg, shieldError) {
					errMsg = "详细错误信息请联系管理员..."
					break
				}
			}
		}
	}

	errMsg = gstr.Split(errMsg, " TraceId")[0]
	errMsg = gstr.Split(errMsg, " (request id:")[0]

	return errMsg
}

// 批量查询应用名称, 返回 appId -> 应用名称
func GetAppNames(ctx context.Context, appIds []int) map[int]string {

	appNames := make(map[int]string)

	if len(appIds) == 0 {
		return appNames
	}

	apps, err := dao.App.Find(ctx, bson.M{"app_id": bson.M{"$in": appIds}}, &dao.FindOptions{
		IncludeFields: []string{"app_id", "name"},
	})
	if err != nil {
		return appNames
	}

	for _, app := range apps {
		appNames[app.AppId] = app.Name
	}

	return appNames
}

// 批量查询密钥名称, 返回 密钥 -> 密钥名称; 名称为空时取密钥的后5位
func GetKeyNames(ctx context.Context, keys []string) map[string]string {

	keyNames := make(map[string]string)

	if len(keys) == 0 {
		return keyNames
	}

	appKeys, err := dao.AppKey.Find(ctx, bson.M{"key": bson.M{"$in": keys}}, &dao.FindOptions{
		IncludeFields: []string{"key", "name"},
	})
	if err != nil {
		return keyNames
	}

	for _, appKey := range appKeys {
		if appKey.Name != "" {
			keyNames[appKey.Key] = appKey.Name
		} else {
			keyNames[appKey.Key] = lastN(appKey.Key, 5)
		}
	}

	// 兜底: 未查到记录的密钥, 直接取后5位
	for _, key := range keys {
		if _, ok := keyNames[key]; !ok {
			keyNames[key] = lastN(key, 5)
		}
	}

	return keyNames
}

func lastN(s string, n int) string {
	if gstr.LenRune(s) <= n {
		return s
	}
	return gstr.SubStrRune(s, gstr.LenRune(s)-n, n)
}

func ConvTaskErrMsg(ctx context.Context, err any, userId int, creator string) string {
	return ConvErrMsg(ctx, extractErrMsg(err), -1, userId, creator)
}

func ShieldTaskError(ctx context.Context, err any, userId int, creator string) {

	if err == nil {
		return
	}

	switch v := err.(type) {
	case map[string]any:
		shieldMapMessages(ctx, v, userId, creator)
	case bson.M:
		shieldMapMessages(ctx, v, userId, creator)
	default:
		shieldStructMessage(ctx, err, userId, creator)
	}
}

func shieldStructMessage(ctx context.Context, err any, userId int, creator string) {

	rv := reflect.ValueOf(err)
	if rv.Kind() == reflect.Ptr {
		if rv.IsNil() {
			return
		}
		rv = rv.Elem()
	}
	if rv.Kind() != reflect.Struct || !rv.IsValid() {
		return
	}

	f := rv.FieldByName("Message")
	if f.IsValid() && f.CanSet() && f.Kind() == reflect.String {
		f.SetString(ConvErrMsg(ctx, f.String(), -1, userId, creator))
	}
}

func shieldMapMessages(ctx context.Context, m map[string]any, userId int, creator string) {

	if m == nil {
		return
	}

	if msg, ok := m["message"].(string); ok {
		m["message"] = ConvErrMsg(ctx, msg, -1, userId, creator)
	}

	if data, ok := m["data"].([]any); ok {
		for _, item := range data {
			if im, ok := item.(map[string]any); ok {
				if msg, ok := im["message"].(string); ok {
					im["message"] = ConvErrMsg(ctx, msg, -1, userId, creator)
				}
			}
		}
	}
}

func extractErrMsg(err any) string {

	if err == nil {
		return ""
	}

	switch v := err.(type) {
	case string:
		return v
	case map[string]any:
		return mapErrMsg(v)
	case bson.M:
		return mapErrMsg(v)
	}

	b, e := json.Marshal(err)
	if e != nil {
		return ""
	}

	var m map[string]any
	if json.Unmarshal(b, &m) == nil {
		if msg := mapErrMsg(m); msg != "" {
			return msg
		}
	}

	s := string(b)
	if s == "null" || s == "{}" || s == "[]" || s == `""` {
		return ""
	}

	return s
}

func mapErrMsg(m map[string]any) string {

	if m == nil {
		return ""
	}

	if msg, ok := m["message"].(string); ok && msg != "" {
		return msg
	}

	if code, ok := m["code"].(string); ok && code != "" {
		return code
	}

	if data, ok := m["data"].([]any); ok && len(data) > 0 {
		msgs := make([]string, 0, len(data))
		for _, item := range data {
			if im, ok := item.(map[string]any); ok {
				if msg, ok := im["message"].(string); ok && msg != "" {
					msgs = append(msgs, msg)
				}
			}
		}
		if len(msgs) > 0 {
			return gstr.Join(msgs, "; ")
		}
	}

	return ""
}
