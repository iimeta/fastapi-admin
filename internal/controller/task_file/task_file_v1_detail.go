package task_file

import (
	"context"

	"github.com/iimeta/fastapi-admin/v2/api/task_file/v1"
	"github.com/iimeta/fastapi-admin/v2/internal/model"
	"github.com/iimeta/fastapi-admin/v2/internal/service"
)

func (c *ControllerV1) Detail(ctx context.Context, req *v1.DetailReq) (res *v1.DetailRes, err error) {

	taskFile, err := service.TaskFile().Detail(ctx, req.Id)
	if err != nil {
		return nil, err
	}

	res = &v1.DetailRes{
		TaskFileDetailRes: &model.TaskFileDetailRes{
			TaskFile: taskFile,
		},
	}

	return
}
