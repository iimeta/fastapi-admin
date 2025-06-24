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
	INotice interface {
		// 新建通知公告
		Create(ctx context.Context, params model.NoticeCreateReq) (string, error)
		// 更新通知公告
		Update(ctx context.Context, params model.NoticeUpdateReq) error
		// 删除通知公告
		Delete(ctx context.Context, id string) error
		// 通知公告详情
		Detail(ctx context.Context, id string) (*model.Notice, error)
		// 通知公告分页列表
		Page(ctx context.Context, params model.NoticePageReq) (*model.NoticePageRes, error)
		// 通知公告列表
		List(ctx context.Context, params model.NoticeListReq) ([]*model.Notice, error)
		// 通知公告批量操作
		BatchOperate(ctx context.Context, params model.NoticeBatchOperateReq) error
		// 额度预警任务
		QuotaWarningTask(ctx context.Context)
	}
)

var (
	localNotice INotice
)

func Notice() INotice {
	if localNotice == nil {
		panic("implement not found for interface INotice, forgot register?")
	}
	return localNotice
}

func RegisterNotice(i INotice) {
	localNotice = i
}
