package ticket

import (
	"context"

	"github.com/iimeta/fastapi-admin/v2/api/ticket/v1"
	"github.com/iimeta/fastapi-admin/v2/internal/service"
)

func (c *ControllerV1) Assign(ctx context.Context, req *v1.AssignReq) (res *v1.AssignRes, err error) {

	err = service.Ticket().Assign(ctx, req.TicketAssignReq)

	return
}
