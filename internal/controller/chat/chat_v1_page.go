package chat

import (
	"context"
	"github.com/iimeta/fastapi-admin/internal/service"

	"github.com/iimeta/fastapi-admin/api/chat/v1"
)

func (c *ControllerV1) Page(ctx context.Context, req *v1.PageReq) (res *v1.PageRes, err error) {

	if !service.Auth().Authenticator(ctx, req) {
		return
	}

	chatPageRes, err := service.Chat().Page(ctx, req.ChatPageReq)
	if err != nil {
		return nil, err
	}

	res = &v1.PageRes{
		ChatPageRes: chatPageRes,
	}

	return
}
