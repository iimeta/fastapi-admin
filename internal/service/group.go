// ================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// You can delete these comments if you wish manually maintain this interface file.
// ================================================================================

package service

import (
	"context"

	"github.com/iimeta/fastapi-admin/v2/internal/model"
)

type (
	IGroup interface {
		// 新建分组
		Create(ctx context.Context, params model.GroupCreateReq) (id string, err error)
		// 更新分组
		Update(ctx context.Context, params model.GroupUpdateReq) error
		// 更改过期时间
		ChangeExpire(ctx context.Context, params model.GroupChangeExpireReq) error
		// 更改分组公开状态
		ChangePublic(ctx context.Context, params model.GroupChangePublicReq) error
		// 更改分组状态
		ChangeStatus(ctx context.Context, params model.GroupChangeStatusReq) error
		// 删除分组
		Delete(ctx context.Context, id string) error
		// 分组详情
		Detail(ctx context.Context, id string) (*model.Group, error)
		// 分组分页列表
		Page(ctx context.Context, params model.GroupPageReq) (*model.GroupPageRes, error)
		// 分组列表
		List(ctx context.Context, params model.GroupListReq) ([]*model.Group, error)
		// 分组批量操作
		BatchOperate(ctx context.Context, params model.GroupBatchOperateReq) error
		// 公开的分组Ids
		PublicGroups(ctx context.Context) ([]string, error)
		// 根据分组Ids查询分组名称
		GroupNames(ctx context.Context, groups []string) ([]string, error)
		// 根据分组Ids获取模型Ids
		GetModelsByGroups(ctx context.Context, groups ...string) ([]string, error)
		// 根据模型Ids获取分组Ids
		GetGroupsByModels(ctx context.Context, models ...string) ([]string, error)
	}
)

var (
	localGroup IGroup
)

func Group() IGroup {
	if localGroup == nil {
		panic("implement not found for interface IGroup, forgot register?")
	}
	return localGroup
}

func RegisterGroup(i IGroup) {
	localGroup = i
}
