package notice

import (
	"context"
	"github.com/go-redsync/redsync/v4"
	"github.com/go-redsync/redsync/v4/redis/goredis/v9"
	"github.com/gogf/gf/v2/os/gtime"
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

type sNotice struct {
	noticeRedsync *redsync.Redsync
}

func init() {
	service.RegisterNotice(New())
}

func New() service.INotice {
	return &sNotice{
		noticeRedsync: redsync.New(goredis.NewPool(redis.UniversalClient)),
	}
}

// 新建通知公告
func (s *sNotice) Create(ctx context.Context, params model.NoticeCreateReq) (string, error) {

	id, err := dao.Notice.Insert(ctx, &do.Notice{
		Title:         params.Title,
		Content:       params.Content,
		Category:      params.Category,
		Scope:         params.Scope,
		Users:         params.Users,
		Resellers:     params.Resellers,
		Methods:       params.Methods,
		Priority:      params.Priority,
		ExpiresAt:     util.ConvTimestampMilli(params.ExpiresAt),
		ScheduledTime: util.ConvTimestampMilli(params.ScheduledTime),
		Remark:        params.Remark,
		Status:        params.Status,
	})
	if err != nil {
		logger.Error(ctx, err)
		return "", err
	}

	return id, nil
}

// 更新通知公告
func (s *sNotice) Update(ctx context.Context, params model.NoticeUpdateReq) error {

	if err := dao.Notice.UpdateById(ctx, params.Id, &do.Notice{
		Title:         params.Title,
		Content:       params.Content,
		Category:      params.Category,
		Scope:         params.Scope,
		Users:         params.Users,
		Resellers:     params.Resellers,
		Methods:       params.Methods,
		Priority:      params.Priority,
		ExpiresAt:     util.ConvTimestampMilli(params.ExpiresAt),
		ScheduledTime: util.ConvTimestampMilli(params.ScheduledTime),
		Remark:        params.Remark,
		Status:        params.Status,
	}); err != nil {
		logger.Error(ctx, err)
		return err
	}

	return nil
}

// 删除通知公告
func (s *sNotice) Delete(ctx context.Context, id string) error {

	if _, err := dao.Notice.DeleteById(ctx, id); err != nil {
		logger.Error(ctx, err)
		return err
	}

	return nil
}

// 通知公告详情
func (s *sNotice) Detail(ctx context.Context, id string) (*model.Notice, error) {

	notice, err := dao.Notice.FindById(ctx, id)
	if err != nil {
		logger.Error(ctx, err)
		return nil, err
	}

	return &model.Notice{
		Id:            notice.Id,
		Title:         notice.Title,
		Content:       notice.Content,
		Category:      notice.Category,
		Scope:         notice.Scope,
		Users:         notice.Users,
		Resellers:     notice.Resellers,
		Methods:       notice.Methods,
		Priority:      notice.Priority,
		ExpiresAt:     util.FormatDateTime(notice.ExpiresAt),
		ScheduledTime: util.FormatDateTime(notice.ScheduledTime),
		Remark:        notice.Remark,
		Status:        notice.Status,
		Reads:         notice.Reads,
		Rid:           notice.Rid,
		Creator:       notice.Creator,
		Updater:       notice.Updater,
		CreatedAt:     util.FormatDateTime(notice.CreatedAt),
		UpdatedAt:     util.FormatDateTime(notice.UpdatedAt),
	}, nil
}

// 通知公告分页列表
func (s *sNotice) Page(ctx context.Context, params model.NoticePageReq) (*model.NoticePageRes, error) {

	paging := &db.Paging{
		Page:     params.Page,
		PageSize: params.PageSize,
	}

	filter := bson.M{}

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

	results, err := dao.Notice.FindByPage(ctx, paging, filter, &dao.FindOptions{SortFields: []string{"-updated_at"}})
	if err != nil {
		logger.Error(ctx, err)
		return nil, err
	}

	items := make([]*model.Notice, 0)
	for _, result := range results {
		items = append(items, &model.Notice{
			Id:            result.Id,
			Title:         result.Title,
			Content:       result.Content,
			Category:      result.Category,
			Scope:         result.Scope,
			Users:         result.Users,
			Resellers:     result.Resellers,
			Methods:       result.Methods,
			Priority:      result.Priority,
			ExpiresAt:     util.FormatDateTime(result.ExpiresAt),
			ScheduledTime: util.FormatDateTime(result.ScheduledTime),
			Remark:        result.Remark,
			Status:        result.Status,
			Reads:         result.Reads,
			Rid:           result.Rid,
			CreatedAt:     util.FormatDateTimeMonth(result.CreatedAt),
			UpdatedAt:     util.FormatDateTimeMonth(result.UpdatedAt),
		})
	}

	return &model.NoticePageRes{
		Items: items,
		Paging: &model.Paging{
			Page:     paging.Page,
			PageSize: paging.PageSize,
			Total:    paging.Total,
		},
	}, nil
}

// 通知公告列表
func (s *sNotice) List(ctx context.Context, params model.NoticeListReq) ([]*model.Notice, error) {

	filter := bson.M{}

	results, err := dao.Notice.Find(ctx, filter, &dao.FindOptions{SortFields: []string{"-updated_at"}})
	if err != nil {
		logger.Error(ctx, err)
		return nil, err
	}

	items := make([]*model.Notice, 0)
	for _, result := range results {
		items = append(items, &model.Notice{
			Id: result.Id,
		})
	}

	return items, nil
}

// 通知公告批量操作
func (s *sNotice) BatchOperate(ctx context.Context, params model.NoticeBatchOperateReq) error {

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
