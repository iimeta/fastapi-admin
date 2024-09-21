package chat

import (
	"context"
	"github.com/iimeta/fastapi-admin/internal/model"
	"github.com/iimeta/fastapi-admin/internal/service"

	"github.com/iimeta/fastapi-admin/api/chat/v1"
)

func (c *ControllerV1) CopyField(ctx context.Context, req *v1.CopyFieldReq) (res *v1.CopyFieldRes, err error) {

	if !service.Auth().Authenticator(ctx, req) {
		return
	}

	value, err := service.Chat().CopyField(ctx, req.ChatCopyFieldReq)
	if err != nil {
		return nil, err
	}

	res = &v1.CopyFieldRes{
		ChatCopyFieldRes: &model.ChatCopyFieldRes{
			Value: value,
		},
	}

	return
}
