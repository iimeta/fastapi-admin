package corp

import (
	"context"
	"github.com/gogf/gf/v2/os/gtime"
	"github.com/gogf/gf/v2/text/gstr"
	"github.com/gogf/gf/v2/util/gconv"
	"github.com/iimeta/fastapi-admin/internal/consts"
	"github.com/iimeta/fastapi-admin/internal/dao"
	"github.com/iimeta/fastapi-admin/internal/model"
	"github.com/iimeta/fastapi-admin/internal/model/do"
	"github.com/iimeta/fastapi-admin/internal/service"
	"github.com/iimeta/fastapi-admin/utility/db"
	"github.com/iimeta/fastapi-admin/utility/logger"
	"github.com/iimeta/fastapi-admin/utility/redis"
	"github.com/iimeta/fastapi-admin/utility/util"
	"go.mongodb.org/mongo-driver/bson"
	"regexp"
	"time"
)

type sCorp struct{}

func init() {
	service.RegisterCorp(New())
}

func New() service.ICorp {
	return &sCorp{}
}

// 新建公司
func (s *sCorp) Create(ctx context.Context, params model.CorpCreateReq) (string, error) {

	if params.Sort == 0 {
		if corp, err := dao.Corp.FindOne(ctx, bson.M{}, &dao.FindOptions{SortFields: []string{"-sort"}}); err == nil && corp != nil {
			params.Sort = corp.Sort + 1
		} else {
			params.Sort = 1
		}
	}

	id, err := dao.Corp.Insert(ctx, &do.Corp{
		Name:     gstr.Trim(params.Name),
		Code:     gstr.Trim(params.Code),
		Sort:     params.Sort,
		IsPublic: params.IsPublic,
		Remark:   params.Remark,
		Status:   params.Status,
	})
	if err != nil {
		logger.Error(ctx, err)
		return "", err
	}

	corp, err := dao.Corp.FindById(ctx, id)
	if err != nil {
		logger.Error(ctx, err)
		return "", err
	}

	if _, err = redis.Publish(ctx, consts.CHANGE_CHANNEL_CORP, model.PubMessage{
		Action:  consts.ACTION_CREATE,
		NewData: corp,
	}); err != nil {
		logger.Error(ctx, err)
	}

	return id, nil
}

// 更新公司
func (s *sCorp) Update(ctx context.Context, params model.CorpUpdateReq) error {

	if params.Sort == 0 {
		if corp, err := dao.Corp.FindOne(ctx, bson.M{}, &dao.FindOptions{SortFields: []string{"-sort"}}); err == nil && corp != nil {
			params.Sort = corp.Sort + 1
		} else {
			params.Sort = 1
		}
	}

	oldData, err := dao.Corp.FindById(ctx, params.Id)
	if err != nil {
		logger.Error(ctx, err)
		return err
	}

	corp, err := dao.Corp.FindOneAndUpdateById(ctx, params.Id, &do.Corp{
		Name:     gstr.Trim(params.Name),
		Code:     gstr.Trim(params.Code),
		Sort:     params.Sort,
		IsPublic: params.IsPublic,
		Remark:   params.Remark,
		Status:   params.Status,
	})
	if err != nil {
		logger.Error(ctx, err)
		return err
	}

	if _, err = redis.Publish(ctx, consts.CHANGE_CHANNEL_CORP, model.PubMessage{
		Action:  consts.ACTION_UPDATE,
		OldData: oldData,
		NewData: corp,
	}); err != nil {
		logger.Error(ctx, err)
		return err
	}

	return nil
}

// 更改公司公开状态
func (s *sCorp) ChangePublic(ctx context.Context, params model.CorpChangePublicReq) error {

	if err := dao.Corp.UpdateById(ctx, params.Id, bson.M{
		"is_public": params.IsPublic,
	}); err != nil {
		logger.Error(ctx, err)
		return err
	}

	return nil
}

// 更改公司状态
func (s *sCorp) ChangeStatus(ctx context.Context, params model.CorpChangeStatusReq) error {

	corp, err := dao.Corp.FindOneAndUpdateById(ctx, params.Id, bson.M{
		"status": params.Status,
	})
	if err != nil {
		logger.Error(ctx, err)
		return err
	}

	if _, err = redis.Publish(ctx, consts.CHANGE_CHANNEL_CORP, model.PubMessage{
		Action:  consts.ACTION_STATUS,
		NewData: corp,
	}); err != nil {
		logger.Error(ctx, err)
		return err
	}

	return nil
}

// 删除公司
func (s *sCorp) Delete(ctx context.Context, id string) error {

	corp, err := dao.Corp.FindOneAndDeleteById(ctx, id)
	if err != nil {
		logger.Error(ctx, err)
		return err
	}

	if _, err = redis.Publish(ctx, consts.CHANGE_CHANNEL_CORP, model.PubMessage{
		Action:  consts.ACTION_DELETE,
		OldData: corp,
	}); err != nil {
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
		Sort:      corp.Sort,
		IsPublic:  corp.IsPublic,
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
			"$regex": regexp.QuoteMeta(params.Name),
		}
	}

	if params.Code != "" {
		filter["code"] = bson.M{
			"$regex": regexp.QuoteMeta(params.Code),
		}
	}

	if params.IsPublic != "" {
		filter["is_public"] = gconv.Bool(params.IsPublic)
	}

	if params.Remark != "" {
		filter["remark"] = bson.M{
			"$regex": regexp.QuoteMeta(params.Remark),
		}
	}

	if params.Status != 0 {
		filter["status"] = params.Status
	}

	if len(params.UpdatedAt) > 0 {
		gte := gtime.NewFromStrFormat(params.UpdatedAt[0], time.DateOnly).StartOfDay().TimestampMilli()
		lte := gtime.NewFromStrLayout(params.UpdatedAt[1], time.DateOnly).EndOfDay(true).TimestampMilli()
		filter["updated_at"] = bson.M{
			"$gte": gte,
			"$lte": lte,
		}
	}

	results, err := dao.Corp.FindByPage(ctx, paging, filter, &dao.FindOptions{SortFields: []string{"status", "-updated_at"}})
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
			Sort:      result.Sort,
			IsPublic:  result.IsPublic,
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

	if service.Session().IsUserRole(ctx) {
		filter["is_public"] = true
	}

	results, err := dao.Corp.Find(ctx, filter, &dao.FindOptions{SortFields: []string{"sort", "status", "-updated_at"}})
	if err != nil {
		logger.Error(ctx, err)
		return nil, err
	}

	items := make([]*model.Corp, 0)
	for _, result := range results {
		items = append(items, &model.Corp{
			Id:   result.Id,
			Name: result.Name,
			Code: result.Code,
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
