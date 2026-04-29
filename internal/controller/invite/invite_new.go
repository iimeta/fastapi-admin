package invite

import (
	"github.com/iimeta/fastapi-admin/v2/api/invite"
)

type ControllerV1 struct{}

func NewV1() invite.IInviteV1 {
	return &ControllerV1{}
}
