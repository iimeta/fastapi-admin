// ================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// You can delete these comments if you wish manually maintain this interface file.
// ================================================================================

package service

import (
	"context"

	"github.com/iimeta/fastapi-admin/internal/model"
	"github.com/iimeta/fastapi-admin/internal/model/entity"
	"github.com/iimeta/fastapi-admin/utility/email"
)

type (
	INotice interface {
		// 新建消息通知
		Create(ctx context.Context, params model.NoticeCreateReq) (string, error)
		// 更新消息通知
		Update(ctx context.Context, params model.NoticeUpdateReq) error
		// 删除消息通知
		Delete(ctx context.Context, id string) error
		// 消息通知详情
		Detail(ctx context.Context, id string) (*model.Notice, error)
		// 消息通知分页列表
		Page(ctx context.Context, params model.NoticePageReq) (*model.NoticePageRes, error)
		// 消息通知列表
		List(ctx context.Context, params model.NoticeListReq) ([]*model.Notice, error)
		// 发送消息通知
		Send(ctx context.Context, notice *entity.Notice) (err error)
		// 发送消息通知邮件
		SendMail(ctx context.Context, dialer *email.Dialer, to string, title string, content string) error
		// 消息通知批量操作
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
