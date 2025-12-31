package notice

import (
	"context"

	"github.com/iimeta/fastapi-admin/v2/api/notice/v1"
	"github.com/iimeta/fastapi-admin/v2/internal/service"
)

func (c *ControllerV1) Create(ctx context.Context, req *v1.CreateReq) (res *v1.CreateRes, err error) {

	_, err = service.Notice().Create(ctx, req.NoticeCreateReq)

	return
}
