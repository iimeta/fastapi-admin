package user

import (
	"context"
	"github.com/iimeta/fastapi-admin/internal/service"
)

type sUser struct{}

func init() {
	service.RegisterUser(New())
}

func New() service.IUser {
	return &sUser{}
}

// 用户信息
func (s *sUser) Info(ctx context.Context) error {
	return nil
}
