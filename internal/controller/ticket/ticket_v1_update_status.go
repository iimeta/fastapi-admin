package ticket

import (
	"context"

	"github.com/iimeta/fastapi-admin/v2/api/ticket/v1"
	"github.com/iimeta/fastapi-admin/v2/internal/service"
)

func (c *ControllerV1) UpdateStatus(ctx context.Context, req *v1.UpdateStatusReq) (res *v1.UpdateStatusRes, err error) {

	err = service.Ticket().UpdateStatus(ctx, req.TicketUpdateStatusReq)

	return
}
