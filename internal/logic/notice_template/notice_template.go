package notice_template

import (
	"context"
	"github.com/go-redsync/redsync/v4"
	"github.com/go-redsync/redsync/v4/redis/goredis/v9"
	"github.com/gogf/gf/v2/os/gtime"
	"github.com/iimeta/fastapi-admin/internal/consts"
	"github.com/iimeta/fastapi-admin/internal/dao"
	"github.com/iimeta/fastapi-admin/internal/errors"
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

type sNoticeTemplate struct {
	noticeRedsync *redsync.Redsync
}

func init() {
	service.RegisterNoticeTemplate(New())
}

func New() service.INoticeTemplate {
	return &sNoticeTemplate{
		noticeRedsync: redsync.New(goredis.NewPool(redis.UniversalClient)),
	}
}

// 新建通知公告模板
func (s *sNoticeTemplate) Create(ctx context.Context, params model.NoticeTemplateCreateReq) (string, error) {

	if params.Content == "" || params.Content == "<p></p>" {
		return "", errors.New("请输入内容")
	}

	notice := &do.NoticeTemplate{
		Name:     params.Name,
		Action:   params.Action,
		Content:  params.Content,
		Category: params.Category,
		IsPublic: params.IsPublic,
		Remark:   params.Remark,
		Status:   params.Status,
		UserId:   service.Session().GetUserId(ctx),
	}

	id, err := dao.NoticeTemplate.Insert(ctx, notice)
	if err != nil {
		logger.Error(ctx, err)
		return "", err
	}

	return id, nil
}

// 更新通知公告模板
func (s *sNoticeTemplate) Update(ctx context.Context, params model.NoticeTemplateUpdateReq) error {

	if params.Content == "" || params.Content == "<p></p>" {
		return errors.New("请输入内容")
	}

	notice := &do.NoticeTemplate{
		Name:     params.Name,
		Action:   params.Action,
		Content:  params.Content,
		Category: params.Category,
		IsPublic: params.IsPublic,
		Remark:   params.Remark,
		Status:   params.Status,
	}

	if err := dao.NoticeTemplate.UpdateById(ctx, params.Id, notice); err != nil {
		logger.Error(ctx, err)
		return err
	}

	return nil
}

// 删除通知公告模板
func (s *sNoticeTemplate) Delete(ctx context.Context, id string) error {

	if _, err := dao.NoticeTemplate.DeleteById(ctx, id); err != nil {
		logger.Error(ctx, err)
		return err
	}

	return nil
}

// 通知公告模板详情
func (s *sNoticeTemplate) Detail(ctx context.Context, id string) (*model.NoticeTemplate, error) {

	notice, err := dao.NoticeTemplate.FindById(ctx, id)
	if err != nil {
		logger.Error(ctx, err)
		return nil, err
	}

	return &model.NoticeTemplate{
		Id:        notice.Id,
		Name:      notice.Name,
		Action:    notice.Action,
		Content:   notice.Content,
		Category:  notice.Category,
		IsPublic:  notice.IsPublic,
		Remark:    notice.Remark,
		Status:    notice.Status,
		UserId:    notice.UserId,
		Rid:       notice.Rid,
		Creator:   notice.Creator,
		Updater:   notice.Updater,
		CreatedAt: util.FormatDateTime(notice.CreatedAt),
		UpdatedAt: util.FormatDateTime(notice.UpdatedAt),
	}, nil
}

// 通知公告模板分页列表
func (s *sNoticeTemplate) Page(ctx context.Context, params model.NoticeTemplatePageReq) (*model.NoticeTemplatePageRes, error) {

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

	if params.Content != "" {
		filter["content"] = bson.M{
			"$regex": regexp.QuoteMeta(params.Content),
		}
	}

	if params.Category != 0 {
		filter["category"] = params.Category
	}

	if params.Remark != "" {
		filter["remark"] = bson.M{
			"$regex": regexp.QuoteMeta(params.Remark),
		}
	}

	if params.Status != 0 {
		filter["status"] = params.Status
	}

	if len(params.PublishTime) > 0 {
		gte := gtime.NewFromStrFormat(params.PublishTime[0], time.DateOnly).StartOfDay().TimestampMilli()
		lte := gtime.NewFromStrLayout(params.PublishTime[1], time.DateOnly).EndOfDay(true).TimestampMilli()
		filter["publish_time"] = bson.M{
			"$gte": gte,
			"$lte": lte,
		}
	}

	results, err := dao.NoticeTemplate.FindByPage(ctx, paging, filter, &dao.FindOptions{SortFields: []string{"-updated_at"}})
	if err != nil {
		logger.Error(ctx, err)
		return nil, err
	}

	items := make([]*model.NoticeTemplate, 0)
	for _, result := range results {
		items = append(items, &model.NoticeTemplate{
			Id:        result.Id,
			Name:      result.Name,
			Action:    result.Action,
			Category:  result.Category,
			IsPublic:  result.IsPublic,
			Remark:    result.Remark,
			Status:    result.Status,
			UserId:    result.UserId,
			Rid:       result.Rid,
			CreatedAt: util.FormatDateTimeMonth(result.CreatedAt),
			UpdatedAt: util.FormatDateTimeMonth(result.UpdatedAt),
		})
	}

	return &model.NoticeTemplatePageRes{
		Items: items,
		Paging: &model.Paging{
			Page:     paging.Page,
			PageSize: paging.PageSize,
			Total:    paging.Total,
		},
	}, nil
}

// 通知公告模板列表
func (s *sNoticeTemplate) List(ctx context.Context, params model.NoticeTemplateListReq) ([]*model.NoticeTemplate, error) {

	filter := bson.M{}

	results, err := dao.NoticeTemplate.Find(ctx, filter, &dao.FindOptions{SortFields: []string{"-updated_at"}})
	if err != nil {
		logger.Error(ctx, err)
		return nil, err
	}

	items := make([]*model.NoticeTemplate, 0)
	for _, result := range results {
		items = append(items, &model.NoticeTemplate{
			Id: result.Id,
		})
	}

	return items, nil
}

// 通知公告模板批量操作
func (s *sNoticeTemplate) BatchOperate(ctx context.Context, params model.NoticeTemplateBatchOperateReq) error {

	switch params.Action {
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
