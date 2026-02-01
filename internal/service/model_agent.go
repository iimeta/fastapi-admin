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
	IModelAgent interface {
		// 新建模型代理
		Create(ctx context.Context, params model.ModelAgentCreateReq) (string, error)
		// 更新模型代理
		Update(ctx context.Context, params model.ModelAgentUpdateReq) error
		// 更改模型代理状态
		ChangeStatus(ctx context.Context, params model.ModelAgentChangeStatusReq) error
		// 删除模型代理
		Delete(ctx context.Context, id string) error
		// 模型代理详情
		Detail(ctx context.Context, id string) (*model.ModelAgent, error)
		// 模型代理分页列表
		Page(ctx context.Context, params model.ModelAgentPageReq) (*model.ModelAgentPageRes, error)
		// 模型代理列表
		List(ctx context.Context, params model.ModelAgentListReq) ([]*model.ModelAgent, error)
		// 模型代理批量操作
		BatchOperate(ctx context.Context, params model.ModelAgentBatchOperateReq) error
		// 模型代理名称是否存在
		IsNameExist(ctx context.Context, name string, id ...string) bool
		// 快速填入模型
		QuickFillModel(ctx context.Context, params model.ModelAgentQuickFillModelReq) ([]string, error)
		// 测试模型
		TestModel(ctx context.Context, params model.ModelAgentTestModelReq) (*model.ModelAgentTestModelRes, error)
	}
)

var (
	localModelAgent IModelAgent
)

func ModelAgent() IModelAgent {
	if localModelAgent == nil {
		panic("implement not found for interface IModelAgent, forgot register?")
	}
	return localModelAgent
}

func RegisterModelAgent(i IModelAgent) {
	localModelAgent = i
}
