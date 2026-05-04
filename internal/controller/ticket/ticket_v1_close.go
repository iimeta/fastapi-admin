package ticket

import (
	"context"

	"github.com/iimeta/fastapi-admin/v2/api/ticket/v1"
	"github.com/iimeta/fastapi-admin/v2/internal/service"
)

func (c *ControllerV1) Close(ctx context.Context, req *v1.CloseReq) (res *v1.CloseRes, err error) {

	err = service.Ticket().Close(ctx, req.TicketCloseReq)

	return
}
