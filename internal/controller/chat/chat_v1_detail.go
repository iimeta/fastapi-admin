package chat

import (
	"context"
	"github.com/iimeta/fastapi-admin/internal/model"
	"github.com/iimeta/fastapi-admin/internal/service"

	"github.com/iimeta/fastapi-admin/api/chat/v1"
)

func (c *ControllerV1) Detail(ctx context.Context, req *v1.DetailReq) (res *v1.DetailRes, err error) {

	chat, err := service.Chat().Detail(ctx, req.Id)
	if err != nil {
		return nil, err
	}

	res = &v1.DetailRes{
		ChatDetailRes: &model.ChatDetailRes{
			Chat: chat,
		},
	}

	return
}
