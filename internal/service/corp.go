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
	ICorp interface {
		// 新建公司
		Create(ctx context.Context, params model.CorpCreateReq) error
		// 更新公司
		Update(ctx context.Context, params model.CorpUpdateReq) error
		// 更改公司状态
		ChangeStatus(ctx context.Context, params model.CorpChangeStatusReq) error
		// 删除公司
		Delete(ctx context.Context, id string) error
		// 公司详情
		Detail(ctx context.Context, id string) (*model.Corp, error)
		// 公司分页列表
		Page(ctx context.Context, params model.CorpPageReq) (*model.CorpPageRes, error)
		// 公司列表
		List(ctx context.Context, params model.CorpListReq) ([]*model.Corp, error)
		// 公司批量操作
		BatchOperate(ctx context.Context, params model.CorpBatchOperateReq) error
		// 公司名称是否存在
		IsNameExist(ctx context.Context, name string, id ...string) bool
		// 公司代码是否存在
		IsCodeExist(ctx context.Context, code string, id ...string) bool
	}
)

var (
	localCorp ICorp
)

func Corp() ICorp {
	if localCorp == nil {
		panic("implement not found for interface ICorp, forgot register?")
	}
	return localCorp
}

func RegisterCorp(i ICorp) {
	localCorp = i
}
