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
	ITicket interface {
		// 创建工单
		Create(ctx context.Context, params model.TicketCreateReq) (string, error)
		// 回复工单
		Reply(ctx context.Context, params model.TicketReplyReq) (string, error)
		// 工单详情
		Detail(ctx context.Context, id string) (*model.TicketDetailRes, error)
		// 工单分页列表
		Page(ctx context.Context, params model.TicketPageReq) (*model.TicketPageRes, error)
		// 更新工单状态
		UpdateStatus(ctx context.Context, params model.TicketUpdateStatusReq) error
		// 关闭工单
		Close(ctx context.Context, params model.TicketCloseReq) error
		// 分配工单
		Assign(ctx context.Context, params model.TicketAssignReq) error
		// 删除工单
		Delete(ctx context.Context, id string) error
		// 批量删除工单
		BatchDelete(ctx context.Context, params model.TicketBatchDeleteReq) error
	}
)

var (
	localTicket ITicket
)

func Ticket() ITicket {
	if localTicket == nil {
		panic("implement not found for interface ITicket, forgot register?")
	}
	return localTicket
}

func RegisterTicket(i ITicket) {
	localTicket = i
}
