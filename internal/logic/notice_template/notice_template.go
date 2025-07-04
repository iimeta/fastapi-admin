package notice_template

import (
	"context"
	"github.com/go-redsync/redsync/v4"
	"github.com/go-redsync/redsync/v4/redis/goredis/v9"
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

// 新建通知模板
func (s *sNoticeTemplate) Create(ctx context.Context, params model.NoticeTemplateCreateReq) (string, error) {

	if params.Content == "" || params.Content == "<p></p>" {
		return "", errors.New("请输入内容")
	}

	notice := &do.NoticeTemplate{
		Name:      params.Name,
		Scenes:    params.Scenes,
		Title:     params.Title,
		Content:   params.Content,
		Channels:  params.Channels,
		IsPublic:  params.IsPublic,
		Remark:    params.Remark,
		Status:    params.Status,
		Variables: util.GetTemplateVariables(params.Title, params.Content),
		UserId:    service.Session().GetUserId(ctx),
	}

	id, err := dao.NoticeTemplate.Insert(ctx, notice)
	if err != nil {
		logger.Error(ctx, err)
		return "", err
	}

	return id, nil
}

// 更新通知模板
func (s *sNoticeTemplate) Update(ctx context.Context, params model.NoticeTemplateUpdateReq) error {

	if params.Content == "" || params.Content == "<p></p>" {
		return errors.New("请输入内容")
	}

	notice := &do.NoticeTemplate{
		Name:      params.Name,
		Scenes:    params.Scenes,
		Title:     params.Title,
		Content:   params.Content,
		Channels:  params.Channels,
		IsPublic:  params.IsPublic,
		Remark:    params.Remark,
		Status:    params.Status,
		Variables: util.GetTemplateVariables(params.Title, params.Content),
	}

	if err := dao.NoticeTemplate.UpdateById(ctx, params.Id, notice); err != nil {
		logger.Error(ctx, err)
		return err
	}

	return nil
}

// 删除通知模板
func (s *sNoticeTemplate) Delete(ctx context.Context, id string) error {

	if _, err := dao.NoticeTemplate.DeleteById(ctx, id); err != nil {
		logger.Error(ctx, err)
		return err
	}

	return nil
}

// 通知模板详情
func (s *sNoticeTemplate) Detail(ctx context.Context, id string) (*model.NoticeTemplate, error) {

	notice, err := dao.NoticeTemplate.FindById(ctx, id)
	if err != nil {
		logger.Error(ctx, err)
		return nil, err
	}

	return &model.NoticeTemplate{
		Id:        notice.Id,
		Name:      notice.Name,
		Scenes:    notice.Scenes,
		Title:     notice.Title,
		Content:   notice.Content,
		Channels:  notice.Channels,
		Variables: notice.Variables,
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

// 通知模板分页列表
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

	if len(params.Scenes) > 0 {
		filter["scenes"] = bson.M{
			"$in": params.Scenes,
		}
	}

	if params.Title != "" {
		filter["title"] = bson.M{
			"$regex": regexp.QuoteMeta(params.Title),
		}
	}

	if params.Content != "" {
		filter["content"] = bson.M{
			"$regex": regexp.QuoteMeta(params.Content),
		}
	}

	if len(params.Channels) > 0 {
		filter["channels"] = bson.M{
			"$in": params.Channels,
		}
	}

	if params.Remark != "" {
		filter["remark"] = bson.M{
			"$regex": regexp.QuoteMeta(params.Remark),
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
			Scenes:    result.Scenes,
			Title:     result.Title,
			Channels:  result.Channels,
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

// 通知模板列表
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

// 根据使用场景获取通知模板
func (s *sNoticeTemplate) GetNoticeTemplateByScene(ctx context.Context, scene string) (*model.NoticeTemplate, error) {

	notice, err := dao.NoticeTemplate.FindOne(ctx, bson.M{"scenes": bson.M{"$in": []string{scene}}, "status": 1}, &dao.FindOptions{SortFields: []string{"-updated_at"}})
	if err != nil {
		logger.Error(ctx, err)
		return nil, err
	}

	return &model.NoticeTemplate{
		Id:        notice.Id,
		Name:      notice.Name,
		Scenes:    notice.Scenes,
		Title:     notice.Title,
		Content:   notice.Content,
		Channels:  notice.Channels,
		Variables: notice.Variables,
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

// 通知模板批量操作
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
