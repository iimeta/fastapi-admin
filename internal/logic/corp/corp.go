package model

import (
	"context"
	"github.com/gogf/gf/v2/text/gstr"
	"github.com/gogf/gf/v2/util/gconv"
	"github.com/iimeta/fastapi-admin/internal/consts"
	"github.com/iimeta/fastapi-admin/internal/dao"
	"github.com/iimeta/fastapi-admin/internal/errors"
	"github.com/iimeta/fastapi-admin/internal/model"
	"github.com/iimeta/fastapi-admin/internal/model/do"
	"github.com/iimeta/fastapi-admin/internal/service"
	"github.com/iimeta/fastapi-admin/utility/db"
	"github.com/iimeta/fastapi-admin/utility/logger"
	"github.com/iimeta/fastapi-admin/utility/util"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type sCorp struct{}

func init() {
	service.RegisterCorp(New())
}

func New() service.ICorp {
	return &sCorp{}
}

// 新建公司
func (s *sCorp) Create(ctx context.Context, params model.CorpCreateReq) error {

	if s.IsNameExist(ctx, params.Name) {
		return errors.Newf("名称 \"%s\" 已存在", params.Name)
	}

	if s.IsCodeExist(ctx, params.Code) {
		return errors.Newf("代码 \"%s\" 已存在", params.Code)
	}

	if _, err := dao.Corp.Insert(ctx, &do.Corp{
		Name:   gstr.Trim(params.Name),
		Code:   gstr.Trim(params.Code),
		Remark: params.Remark,
		Status: params.Status,
	}); err != nil {
		logger.Error(ctx, err)
		return err
	}

	return nil
}

// 更新公司
func (s *sCorp) Update(ctx context.Context, params model.CorpUpdateReq) error {

	if s.IsNameExist(ctx, params.Name, params.Id) {
		return errors.Newf("名称 \"%s\" 已存在", params.Name)
	}

	if s.IsCodeExist(ctx, params.Code, params.Id) {
		return errors.Newf("代码 \"%s\" 已存在", params.Code)
	}

	if err := dao.Corp.UpdateById(ctx, params.Id, &do.Corp{
		Name:   gstr.Trim(params.Name),
		Code:   params.Code,
		Remark: params.Remark,
		Status: params.Status,
	}); err != nil {
		logger.Error(ctx, err)
		return err
	}

	return nil
}

// 更改公司状态
func (s *sCorp) ChangeStatus(ctx context.Context, params model.CorpChangeStatusReq) error {

	if err := dao.Corp.UpdateById(ctx, params.Id, bson.M{
		"status": params.Status,
	}); err != nil {
		logger.Error(ctx, err)
		return err
	}

	return nil
}

// 删除公司
func (s *sCorp) Delete(ctx context.Context, id string) error {

	if _, err := dao.Corp.DeleteById(ctx, id); err != nil {
		logger.Error(ctx, err)
		return err
	}

	return nil
}

// 公司详情
func (s *sCorp) Detail(ctx context.Context, id string) (*model.Corp, error) {

	corp, err := dao.Corp.FindById(ctx, id)
	if err != nil {
		logger.Error(ctx, err)
		return nil, err
	}

	return &model.Corp{
		Id:        corp.Id,
		Name:      corp.Name,
		Code:      corp.Code,
		Remark:    corp.Remark,
		Status:    corp.Status,
		Creator:   corp.Creator,
		Updater:   corp.Updater,
		CreatedAt: util.FormatDateTime(corp.CreatedAt),
		UpdatedAt: util.FormatDateTime(corp.UpdatedAt),
	}, nil
}

// 公司分页列表
func (s *sCorp) Page(ctx context.Context, params model.CorpPageReq) (*model.CorpPageRes, error) {

	paging := &db.Paging{
		Page:     params.Page,
		PageSize: params.PageSize,
	}

	filter := bson.M{}

	if params.Name != "" {
		filter["name"] = bson.M{
			"$regex": params.Name,
		}
	}

	if params.Code != "" {
		filter["code"] = bson.M{
			"$regex": params.Code,
		}
	}

	if params.Status != 0 {
		filter["status"] = params.Status
	}

	results, err := dao.Corp.FindByPage(ctx, paging, filter, "status", "-updated_at")
	if err != nil {
		logger.Error(ctx, err)
		return nil, err
	}

	items := make([]*model.Corp, 0)
	for _, result := range results {
		items = append(items, &model.Corp{
			Id:        result.Id,
			Name:      result.Name,
			Code:      result.Code,
			Remark:    result.Remark,
			Status:    result.Status,
			CreatedAt: util.FormatDateTimeMonth(result.CreatedAt),
			UpdatedAt: util.FormatDateTimeMonth(result.UpdatedAt),
		})
	}

	return &model.CorpPageRes{
		Items: items,
		Paging: &model.Paging{
			Page:     paging.Page,
			PageSize: paging.PageSize,
			Total:    paging.Total,
		},
	}, nil
}

// 公司列表
func (s *sCorp) List(ctx context.Context, params model.CorpListReq) ([]*model.Corp, error) {

	filter := bson.M{
		"status": 1,
	}

	results, err := dao.Corp.Find(ctx, filter, "-updated_at")
	if err != nil {
		logger.Error(ctx, err)
		return nil, err
	}

	items := make([]*model.Corp, 0)
	for _, result := range results {
		items = append(items, &model.Corp{
			Id:     result.Id,
			Name:   result.Name,
			Code:   result.Code,
			Status: result.Status,
		})
	}

	return items, nil
}

// 公司批量操作
func (s *sCorp) BatchOperate(ctx context.Context, params model.CorpBatchOperateReq) error {

	switch params.Action {
	case consts.ACTION_STATUS:
		for _, id := range params.Ids {
			if err := s.ChangeStatus(ctx, model.CorpChangeStatusReq{
				Id:     id,
				Status: gconv.Int(params.Value),
			}); err != nil {
				logger.Error(ctx, err)
				return err
			}
		}
	case consts.ACTION_DELETE:
		for _, id := range params.Ids {
			if err := s.Delete(ctx, id); err != nil {
				logger.Error(ctx, err)
				return err
			}
		}
	}

	return nil
}

// 公司名称是否存在
func (s *sCorp) IsNameExist(ctx context.Context, name string, id ...string) bool {

	corp, err := dao.Corp.FindOne(ctx, bson.M{"name": gstr.Trim(name)})
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return false
		}
		logger.Error(ctx, err)
		return true
	}

	if corp != nil {
		if len(id) > 0 && corp.Id == id[0] {
			return false
		}
		return true
	}

	return false
}

// 公司代码是否存在
func (s *sCorp) IsCodeExist(ctx context.Context, code string, id ...string) bool {

	corp, err := dao.Corp.FindOne(ctx, bson.M{"code": gstr.Trim(code)})
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return false
		}
		logger.Error(ctx, err)
		return true
	}

	if corp != nil {
		if len(id) > 0 && corp.Id == id[0] {
			return false
		}
		return true
	}

	return false
}
