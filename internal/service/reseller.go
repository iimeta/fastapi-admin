// ================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// You can delete these comments if you wish manually maintain this interface file.
// ================================================================================

package service

import (
	"context"

	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/iimeta/fastapi-admin/internal/model"
)

type (
	IReseller interface {
		// 代理商更新信息
		UpdateInfo(ctx context.Context, params model.UserUpdateInfoReq) error
		// 代理商更改密码
		ChangePassword(ctx context.Context, params model.UserChangePasswordReq) (err error)
		// 代理商更改邮箱
		ChangeEmail(ctx context.Context, params model.UserChangeEmailReq) error
		// 代理商更改头像
		ChangeAvatar(ctx context.Context, file *ghttp.UploadFile) error
		// 根据userId获取代理商信息
		GetResellerByUserId(ctx context.Context, userId int) (*model.Reseller, error)
	}
)

var (
	localReseller IReseller
)

func Reseller() IReseller {
	if localReseller == nil {
		panic("implement not found for interface IReseller, forgot register?")
	}
	return localReseller
}

func RegisterReseller(i IReseller) {
	localReseller = i
}
