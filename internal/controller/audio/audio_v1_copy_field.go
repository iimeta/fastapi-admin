package audio

import (
	"context"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/text/gstr"
	"github.com/gogf/gf/v2/util/gmeta"
	"github.com/iimeta/fastapi-admin/internal/model"
	"github.com/iimeta/fastapi-admin/internal/service"

	"github.com/iimeta/fastapi-admin/api/audio/v1"
)

func (c *ControllerV1) CopyField(ctx context.Context, req *v1.CopyFieldReq) (res *v1.CopyFieldRes, err error) {

	role := gmeta.Get(req, "role").String()
	if role != "*" && !gstr.Contains(role, service.Session().GetRole(ctx)) {
		g.RequestFromCtx(ctx).Response.WriteJson(g.Map{"code": 401, "message": "Unauthorized"})
		return
	}

	value, err := service.Audio().CopyField(ctx, req.AudioCopyFieldReq)
	if err != nil {
		return nil, err
	}

	res = &v1.CopyFieldRes{
		AudioCopyFieldRes: &model.AudioCopyFieldRes{
			Value: value,
		},
	}

	return
}
