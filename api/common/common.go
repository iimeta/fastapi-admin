// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package common

import (
	"context"

	"github.com/iimeta/fastapi-admin/v2/api/common/v1"
)

type ICommonV1 interface {
	SendSms(ctx context.Context, req *v1.SendSmsReq) (res *v1.SendSmsRes, err error)
	SendEmail(ctx context.Context, req *v1.SendEmailReq) (res *v1.SendEmailRes, err error)
}
