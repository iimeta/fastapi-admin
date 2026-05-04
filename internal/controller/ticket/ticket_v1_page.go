package ticket

import (
	"context"

	"github.com/iimeta/fastapi-admin/v2/api/ticket/v1"
	"github.com/iimeta/fastapi-admin/v2/internal/service"
)

func (c *ControllerV1) Page(ctx context.Context, req *v1.PageReq) (res *v1.PageRes, err error) {

	ticketPageRes, err := service.Ticket().Page(ctx, req.TicketPageReq)
	if err != nil {
		return nil, err
	}

	res = &v1.PageRes{
		TicketPageRes: ticketPageRes,
	}

	return
}
