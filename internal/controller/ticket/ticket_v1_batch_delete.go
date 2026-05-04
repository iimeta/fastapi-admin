package ticket

import (
	"context"

	"github.com/iimeta/fastapi-admin/v2/api/ticket/v1"
	"github.com/iimeta/fastapi-admin/v2/internal/service"
)

func (c *ControllerV1) BatchDelete(ctx context.Context, req *v1.BatchDeleteReq) (res *v1.BatchDeleteRes, err error) {

	err = service.Ticket().BatchDelete(ctx, req.TicketBatchDeleteReq)

	return
}
