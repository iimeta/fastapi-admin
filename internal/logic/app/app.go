package model

import (
	"context"
	"github.com/iimeta/fastapi-admin/internal/model"
	"github.com/iimeta/fastapi-admin/internal/service"
)

type sApp struct{}

func init() {
	service.RegisterApp(New())
}

func New() service.IApp {
	return &sApp{}
}

// 新建应用
func (s *sApp) Create(ctx context.Context, params model.AppCreateReq) error {

	return nil
}

// 更新应用
func (s *sApp) Update(ctx context.Context, params model.AppUpdateReq) error {

	return nil
}

// 删除应用
func (s *sApp) Delete(ctx context.Context, id string) error {

	return nil
}

// 应用详情
func (s *sApp) Detail(ctx context.Context, id string) (*model.AppDetailRes, error) {

	return nil, nil
}
