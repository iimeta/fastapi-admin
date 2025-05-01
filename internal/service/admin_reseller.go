// ================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// You can delete these comments if you wish manually maintain this interface file.
// ================================================================================

package service

import (
	"context"

	"github.com/iimeta/fastapi-admin/internal/model"
)

type (
	IAdminReseller interface {
		// 新建代理商
		Create(ctx context.Context, params model.ResellerCreateReq) (err error)
		// 更新代理商
		Update(ctx context.Context, params model.ResellerUpdateReq) error
		// 更改代理商额度过期时间
		ChangeQuotaExpire(ctx context.Context, params model.ResellerChangeQuotaExpireReq) error
		// 更改代理商状态
		ChangeStatus(ctx context.Context, params model.ResellerChangeStatusReq) error
		// 删除代理商
		Delete(ctx context.Context, id string) error
		// 代理商详情
		Detail(ctx context.Context, id string) (*model.Reseller, error)
		// 代理商分页列表
		Page(ctx context.Context, params model.ResellerPageReq) (*model.ResellerPageRes, error)
		// 代理商列表
		List(ctx context.Context, params model.ResellerListReq) ([]*model.Reseller, error)
		// 代理商充值
		Recharge(ctx context.Context, params model.ResellerRechargeReq) error
		// 代理商权限
		Permissions(ctx context.Context, params model.ResellerPermissionsReq) error
		// 代理商批量操作
		BatchOperate(ctx context.Context, params model.ResellerBatchOperateReq) error
	}
)

var (
	localAdminReseller IAdminReseller
)

func AdminReseller() IAdminReseller {
	if localAdminReseller == nil {
		panic("implement not found for interface IAdminReseller, forgot register?")
	}
	return localAdminReseller
}

func RegisterAdminReseller(i IAdminReseller) {
	localAdminReseller = i
}
