// ================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// You can delete these comments if you wish manually maintain this interface file.
// ================================================================================

package service

import (
	"context"

	"github.com/iimeta/fastapi-admin/v2/internal/model"
	"github.com/iimeta/fastapi-admin/v2/internal/model/do"
)

type (
	INoticeTemplate interface {
		// 初始化通知模板
		Init(ctx context.Context)
		// 默认通知模板
		Default() []*do.NoticeTemplate
		// 新建通知模板
		Create(ctx context.Context, params model.NoticeTemplateCreateReq) (string, error)
		// 更新通知模板
		Update(ctx context.Context, params model.NoticeTemplateUpdateReq) error
		// 更改通知模板公开状态
		ChangePublic(ctx context.Context, params model.NoticeTemplateChangePublicReq) error
		// 更改通知模板状态
		ChangeStatus(ctx context.Context, params model.NoticeTemplateChangeStatusReq) error
		// 删除通知模板
		Delete(ctx context.Context, id string) error
		// 通知模板详情
		Detail(ctx context.Context, id string) (*model.NoticeTemplate, error)
		// 通知模板分页列表
		Page(ctx context.Context, params model.NoticeTemplatePageReq) (*model.NoticeTemplatePageRes, error)
		// 通知模板列表
		List(ctx context.Context, params model.NoticeTemplateListReq) ([]*model.NoticeTemplate, error)
		// 根据使用场景获取通知模板
		GetNoticeTemplateByScene(ctx context.Context, scene string, channels []string) (*model.NoticeTemplate, error)
		// 通知模板批量操作
		BatchOperate(ctx context.Context, params model.NoticeTemplateBatchOperateReq) error
	}
)

var (
	localNoticeTemplate INoticeTemplate
)

func NoticeTemplate() INoticeTemplate {
	if localNoticeTemplate == nil {
		panic("implement not found for interface INoticeTemplate, forgot register?")
	}
	return localNoticeTemplate
}

func RegisterNoticeTemplate(i INoticeTemplate) {
	localNoticeTemplate = i
}
