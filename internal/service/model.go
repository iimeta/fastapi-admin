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
	IModel interface {
		// 新建模型
		Create(ctx context.Context, params model.ModelCreateReq) error
		// 更新模型
		Update(ctx context.Context, params model.ModelUpdateReq) error
		// 更改模型状态
		ChangeStatus(ctx context.Context, params model.ModelChangeStatusReq) error
		// 删除模型
		Delete(ctx context.Context, id string) error
		// 模型详情
		Detail(ctx context.Context, id string) (*model.Model, error)
		// 模型分页列表
		Page(ctx context.Context, params model.ModelPageReq) (*model.ModelPageRes, error)
		// 模型列表
		List(ctx context.Context, params model.ModelListReq) ([]*model.Model, error)
		// 模型批量操作
		BatchOperate(ctx context.Context, params model.ModelBatchOperateReq) (err error)
		// 公开的模型Ids
		PublicModels(ctx context.Context) ([]string, error)
		// 根据模型Ids查询模型名称
		ModelNames(ctx context.Context, models []string) ([]string, error)
		// 模型名称是否存在
		IsNameExist(ctx context.Context, name string, id ...string) bool
	}
)

var (
	localModel IModel
)

func Model() IModel {
	if localModel == nil {
		panic("implement not found for interface IModel, forgot register?")
	}
	return localModel
}

func RegisterModel(i IModel) {
	localModel = i
}
