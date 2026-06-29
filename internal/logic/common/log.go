package common

import (
	"context"

	"github.com/gogf/gf/v2/text/gstr"
	"github.com/iimeta/fastapi-admin/v2/internal/dao"
	"go.mongodb.org/mongo-driver/v2/bson"
)

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
