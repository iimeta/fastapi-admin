package model

import (
	"context"
	"fmt"
	"github.com/gogf/gf/v2/encoding/gjson"
	"github.com/iimeta/fastapi-admin/internal/model"
	"github.com/iimeta/fastapi-admin/internal/service"
)

type sModel struct{}

func init() {
	service.RegisterModel(New())
}

func New() service.IModel {
	return &sModel{}
}

// 新建模型
func (s *sModel) Create(ctx context.Context, params model.ModelCreateReq) error {

	fmt.Println(gjson.MustEncodeString(params))

	return nil
}

// 更新模型
func (s *sModel) Update(ctx context.Context, params model.ModelUpdateReq) error {

	return nil
}

// 删除模型
func (s *sModel) Delete(ctx context.Context, id string) error {

	return nil
}

// 模型详情
func (s *sModel) Detail(ctx context.Context, id string) (*model.ModelDetailRes, error) {

	return nil, nil
}
