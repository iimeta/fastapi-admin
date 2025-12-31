package finance

import (
	"context"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/iimeta/fastapi-admin/v2/api/finance/v1"
	"github.com/iimeta/fastapi-admin/v2/internal/service"
)

func (c *ControllerV1) BillExport(ctx context.Context, req *v1.BillExportReq) (res *v1.BillExportRes, err error) {

	filePath, err := service.Finance().BillExport(ctx, req.FinanceBillExportReq)
	if err != nil {
		return nil, err
	}

	g.RequestFromCtx(ctx).Response.ServeFile(filePath)

	return
}
