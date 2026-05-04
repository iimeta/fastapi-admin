package ticket

import (
	"context"

	"github.com/iimeta/fastapi-admin/v2/api/ticket/v1"
	"github.com/iimeta/fastapi-admin/v2/internal/service"
)

func (c *ControllerV1) Reply(ctx context.Context, req *v1.ReplyReq) (res *v1.ReplyRes, err error) {

	_, err = service.Ticket().Reply(ctx, req.TicketReplyReq)

	return
}
