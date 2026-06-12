package task_image

import (
	"context"

	"github.com/iimeta/fastapi-admin/v2/api/task_image/v1"
	"github.com/iimeta/fastapi-admin/v2/internal/service"
)

func (c *ControllerV1) Regenerate(ctx context.Context, req *v1.RegenerateReq) (res *v1.RegenerateRes, err error) {

	err = service.TaskImage().Regenerate(ctx, req.Id)

	return
}
