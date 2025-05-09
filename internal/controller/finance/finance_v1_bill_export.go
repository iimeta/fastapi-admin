package finance

import (
	"context"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/iimeta/fastapi-admin/internal/service"

	"github.com/iimeta/fastapi-admin/api/finance/v1"
)

func (c *ControllerV1) BillExport(ctx context.Context, req *v1.BillExportReq) (res *v1.BillExportRes, err error) {

	filePath, err := service.Finance().BillExport(ctx, req.FinanceBillExportReq)
	if err != nil {
		return nil, err
	}

	g.RequestFromCtx(ctx).Response.ServeFile(filePath)

	return
}
