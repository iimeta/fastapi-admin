package group

import (
	"context"

	"github.com/iimeta/fastapi-admin/v2/api/group/v1"
	"github.com/iimeta/fastapi-admin/v2/internal/model"
	"github.com/iimeta/fastapi-admin/v2/internal/service"
)

func (c *ControllerV1) TagList(ctx context.Context, req *v1.TagListReq) (res *v1.TagListRes, err error) {

	tags, err := service.Group().TagList(ctx)
	if err != nil {
		return nil, err
	}

	res = &v1.TagListRes{
		GroupTagListRes: &model.GroupTagListRes{
			Tags: tags,
		},
	}

	return
}
