package open

import (
	"context"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/iimeta/fastapi-admin/v2/api/open/v1"
	"github.com/iimeta/fastapi-admin/v2/internal/service"
)

func (c *ControllerV1) UserAgreement(ctx context.Context, req *v1.UserAgreementReq) (res *v1.UserAgreementRes, err error) {

	if req.Domain == "" {
		req.Domain = g.RequestFromCtx(ctx).GetHost()
	}

	content, err := service.Open().UserAgreement(ctx, req.SysConfigReq)
	if err != nil {
		return nil, err
	}

	res = &v1.UserAgreementRes{
		Content: content,
	}

	return
}
