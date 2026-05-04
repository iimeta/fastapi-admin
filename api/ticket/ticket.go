// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package ticket

import (
	"context"

	"github.com/iimeta/fastapi-admin/v2/api/ticket/v1"
)

type ITicketV1 interface {
	Create(ctx context.Context, req *v1.CreateReq) (res *v1.CreateRes, err error)
	Reply(ctx context.Context, req *v1.ReplyReq) (res *v1.ReplyRes, err error)
	Detail(ctx context.Context, req *v1.DetailReq) (res *v1.DetailRes, err error)
	Page(ctx context.Context, req *v1.PageReq) (res *v1.PageRes, err error)
	UpdateStatus(ctx context.Context, req *v1.UpdateStatusReq) (res *v1.UpdateStatusRes, err error)
	Close(ctx context.Context, req *v1.CloseReq) (res *v1.CloseRes, err error)
	Assign(ctx context.Context, req *v1.AssignReq) (res *v1.AssignRes, err error)
	BatchDelete(ctx context.Context, req *v1.BatchDeleteReq) (res *v1.BatchDeleteRes, err error)
}
