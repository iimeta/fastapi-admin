package notice_template

import (
	"context"

	"github.com/iimeta/fastapi-admin/v2/api/notice_template/v1"
	"github.com/iimeta/fastapi-admin/v2/internal/service"
)

func (c *ControllerV1) BatchOperate(ctx context.Context, req *v1.BatchOperateReq) (res *v1.BatchOperateRes, err error) {

	err = service.NoticeTemplate().BatchOperate(ctx, req.NoticeTemplateBatchOperateReq)

	return
}
