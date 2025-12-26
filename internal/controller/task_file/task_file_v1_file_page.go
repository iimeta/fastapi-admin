package task_file

import (
	"context"

	"github.com/iimeta/fastapi-admin/api/task_file/v1"
	"github.com/iimeta/fastapi-admin/internal/service"
)

func (c *ControllerV1) FilePage(ctx context.Context, req *v1.FilePageReq) (res *v1.FilePageRes, err error) {

	pageRes, err := service.TaskFile().Page(ctx, req.TaskFilePageReq)
	if err != nil {
		return nil, err
	}

	res = &v1.FilePageRes{
		TaskFilePageRes: pageRes,
	}

	return
}
