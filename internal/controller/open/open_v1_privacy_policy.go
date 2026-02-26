package open

import (
	"context"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/iimeta/fastapi-admin/v2/api/open/v1"
	"github.com/iimeta/fastapi-admin/v2/internal/service"
)

func (c *ControllerV1) PrivacyPolicy(ctx context.Context, req *v1.PrivacyPolicyReq) (res *v1.PrivacyPolicyRes, err error) {

	if req.Domain == "" {
		req.Domain = g.RequestFromCtx(ctx).GetHost()
	}

	content, err := service.Open().PrivacyPolicy(ctx, req.SysConfigReq)
	if err != nil {
		return nil, err
	}

	res = &v1.PrivacyPolicyRes{
		Content: content,
	}

	return
}
