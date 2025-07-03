package notice

import (
	"context"
	"github.com/go-redsync/redsync/v4"
	"github.com/go-redsync/redsync/v4/redis/goredis/v9"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gctx"
	"github.com/gogf/gf/v2/os/grpool"
	"github.com/gogf/gf/v2/os/gtime"
	"github.com/iimeta/fastapi-admin/internal/consts"
	"github.com/iimeta/fastapi-admin/internal/dao"
	"github.com/iimeta/fastapi-admin/internal/errors"
	"github.com/iimeta/fastapi-admin/internal/model"
	"github.com/iimeta/fastapi-admin/internal/model/do"
	"github.com/iimeta/fastapi-admin/internal/model/entity"
	"github.com/iimeta/fastapi-admin/internal/service"
	"github.com/iimeta/fastapi-admin/utility/db"
	"github.com/iimeta/fastapi-admin/utility/email"
	"github.com/iimeta/fastapi-admin/utility/logger"
	"github.com/iimeta/fastapi-admin/utility/redis"
	"github.com/iimeta/fastapi-admin/utility/util"
	"go.mongodb.org/mongo-driver/bson"
	"regexp"
	"slices"
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

	if params.Content == "" || params.Content == "<p></p>" {
		return "", errors.New("请输入内容")
	}

	notice := &do.Notice{
		Title:         params.Title,
		Content:       params.Content,
		Category:      params.Category,
		Scope:         params.Scope,
		Users:         params.Users,
		Resellers:     params.Resellers,
		Channels:      params.Channels,
		Priority:      params.Priority,
		ExpiresAt:     util.ConvTimestampMilli(params.ExpiresAt),
		ScheduledTime: util.ConvTimestampMilli(params.ScheduledTime),
		Remark:        params.Remark,
		Status:        params.Status,
		UserId:        service.Session().GetUserId(ctx),
	}

	if notice.Status == 1 {
		notice.PublishTime = gtime.TimestampMilli()
	}

	id, err := dao.Notice.Insert(ctx, notice)
	if err != nil {
		logger.Error(ctx, err)
		return "", err
	}

	if notice.Status == 1 && slices.Contains(notice.Channels, "email") {

		newData, err := dao.Notice.FindById(ctx, id)
		if err != nil {
			logger.Error(ctx, err)
			return "", err
		}

		// 发送通知公告邮件
		if err := grpool.AddWithRecover(gctx.NeverDone(ctx), func(ctx context.Context) {
			if err := s.SendMail(ctx, newData); err != nil {
				logger.Error(ctx, err)
			}
		}, nil); err != nil {
			logger.Error(ctx, err)
		}
	}

	return id, nil
}

// 更新通知公告
func (s *sNotice) Update(ctx context.Context, params model.NoticeUpdateReq) error {

	if params.Content == "" || params.Content == "<p></p>" {
		return errors.New("请输入内容")
	}

	notice := &do.Notice{
		Title:         params.Title,
		Content:       params.Content,
		Category:      params.Category,
		Scope:         params.Scope,
		Users:         params.Users,
		Resellers:     params.Resellers,
		Channels:      params.Channels,
		Priority:      params.Priority,
		ExpiresAt:     util.ConvTimestampMilli(params.ExpiresAt),
		ScheduledTime: util.ConvTimestampMilli(params.ScheduledTime),
		Remark:        params.Remark,
		Status:        params.Status,
	}

	newData, err := dao.Notice.FindOneAndUpdateById(ctx, params.Id, notice)
	if err != nil {
		logger.Error(ctx, err)
		return err
	}

	if newData.Status == 1 && newData.PublishTime == 0 && slices.Contains(newData.Channels, "email") {

		if err = dao.Notice.UpdateById(ctx, params.Id, bson.M{"publish_time": gtime.TimestampMilli()}); err != nil {
			logger.Error(ctx, err)
			return err
		}

		// 发送通知公告邮件
		if err := grpool.AddWithRecover(gctx.NeverDone(ctx), func(ctx context.Context) {
			if err := s.SendMail(ctx, newData); err != nil {
				logger.Error(ctx, err)
			}
		}, nil); err != nil {
			logger.Error(ctx, err)
		}
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
		Channels:      notice.Channels,
		Priority:      notice.Priority,
		ExpiresAt:     util.FormatDateTime(notice.ExpiresAt),
		ScheduledTime: util.FormatDateTime(notice.ScheduledTime),
		Remark:        notice.Remark,
		Status:        notice.Status,
		Reads:         notice.Reads,
		UserId:        notice.UserId,
		PublishTime:   util.FormatDateTime(notice.PublishTime),
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
			Category:      result.Category,
			Scope:         result.Scope,
			Users:         result.Users,
			Resellers:     result.Resellers,
			Channels:      result.Channels,
			Priority:      result.Priority,
			ExpiresAt:     util.FormatDateTime(result.ExpiresAt),
			ScheduledTime: util.FormatDateTime(result.ScheduledTime),
			Remark:        result.Remark,
			Status:        result.Status,
			Reads:         result.Reads,
			UserId:        result.UserId,
			PublishTime:   util.FormatDateTime(result.PublishTime),
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

// 发送通知公告邮件
func (s *sNotice) SendMail(ctx context.Context, notice *entity.Notice) (err error) {

	var (
		users     []*entity.User
		resellers []*entity.Reseller
	)

	if notice.Scope == 1 || notice.Scope == 2 {
		if users, err = dao.User.Find(ctx, bson.M{"status": 1}); err != nil {
			logger.Error(ctx, err)
			return err
		}
	}

	if notice.Scope == 4 || notice.Scope == 6 {
		if users, err = dao.User.Find(ctx, bson.M{"user_id": bson.M{"$in": notice.Users}, "status": 1}); err != nil {
			logger.Error(ctx, err)
			return err
		}
	}

	if notice.Scope == 1 || notice.Scope == 3 {
		if resellers, err = dao.Reseller.Find(ctx, bson.M{"status": 1}); err != nil {
			logger.Error(ctx, err)
			return err
		}
	}

	if notice.Scope == 5 || notice.Scope == 6 {
		if resellers, err = dao.Reseller.Find(ctx, bson.M{"user_id": bson.M{"$in": notice.Resellers}, "status": 1}); err != nil {
			logger.Error(ctx, err)
			return err
		}
	}

	for _, user := range users {

		if err := g.Validator().Data(user.Email).Rules("email").Run(ctx); err != nil {
			logger.Infof(ctx, "sNotice SendMail user: %d, error: %v", user.UserId, err)
			continue
		}

		account, err := dao.Account.FindOne(ctx, bson.M{"user_id": user.UserId, "status": 1}, &dao.FindOptions{SortFields: []string{"-updated_at"}})
		if err != nil {
			logger.Error(ctx, err)
			continue
		}

		if account == nil {
			logger.Infof(ctx, "sNotice SendMail user: %d, 因无可用账号, 不发送通知公告邮件", user.UserId)
			continue
		}

		dialer := email.NewDefaultDialer()

		if user.Rid > 0 {

			isConfigEmail := false

			if account.LoginDomain != "" {
				siteConfig := service.SiteConfig().GetSiteConfigByDomain(ctx, account.LoginDomain)
				if siteConfig != nil && siteConfig.Rid == user.Rid && siteConfig.Host != "" {
					dialer = email.NewDialer(siteConfig.Host, siteConfig.Port, siteConfig.UserName, siteConfig.Password)
					isConfigEmail = true
				}
			}

			if !isConfigEmail {
				siteConfigs := service.SiteConfig().GetSiteConfigsByRid(ctx, user.Rid)
				for _, siteConfig := range siteConfigs {
					if siteConfig != nil && siteConfig.Host != "" {
						dialer = email.NewDialer(siteConfig.Host, siteConfig.Port, siteConfig.UserName, siteConfig.Password)
						isConfigEmail = true
						break
					}
				}
			}

			if !isConfigEmail {
				logger.Infof(ctx, "sNotice SendMail 因代理商: %d, 所有站点未配置邮箱, 不发送通知公告邮件", user.Rid)
				continue
			}

		} else if account.LoginDomain != "" {
			siteConfig := service.SiteConfig().GetSiteConfigByDomain(ctx, account.LoginDomain)
			if siteConfig != nil && siteConfig.Host != "" {
				dialer = email.NewDialer(siteConfig.Host, siteConfig.Port, siteConfig.UserName, siteConfig.Password)
			} else {
				logger.Infof(ctx, "sNotice SendMail 因站点 %s 未配置邮箱, 默认使用系统配置邮箱", account.LoginDomain)
			}
		}

		if err = email.SendMail(email.NewMessage([]string{user.Email}, notice.Title, notice.Content), dialer); err != nil {
			logger.Errorf(ctx, "sNotice SendMail user: %d, email: %s, SendMail %s error: %v", user.UserId, user.Email, notice.Title, err)
			continue
		}

		logger.Infof(ctx, "sNotice SendMail user: %d, email: %s, SendMail %s success", user.UserId, user.Email, notice.Title)
	}

	for _, reseller := range resellers {

		if err := g.Validator().Data(reseller.Email).Rules("email").Run(ctx); err != nil {
			logger.Infof(ctx, "sNotice SendMail reseller: %d, error: %v", reseller.UserId, err)
			continue
		}

		account, err := dao.ResellerAccount.FindOne(ctx, bson.M{"user_id": reseller.UserId, "status": 1}, &dao.FindOptions{SortFields: []string{"-updated_at"}})
		if err != nil {
			logger.Error(ctx, err)
			continue
		}

		if account == nil {
			logger.Infof(ctx, "sNotice SendMail reseller: %d, 因无可用账号, 不发送通知公告邮件", reseller.UserId)
			continue
		}

		dialer := email.NewDefaultDialer()

		if account.LoginDomain != "" {
			siteConfig := service.SiteConfig().GetSiteConfigByDomain(ctx, account.LoginDomain)
			if siteConfig != nil && siteConfig.Host != "" {
				dialer = email.NewDialer(siteConfig.Host, siteConfig.Port, siteConfig.UserName, siteConfig.Password)
			} else {
				logger.Infof(ctx, "sNotice SendMail 因站点 %s 未配置邮箱, 默认使用系统配置邮箱", account.LoginDomain)
			}
		}

		if err = email.SendMail(email.NewMessage([]string{reseller.Email}, notice.Title, notice.Content), dialer); err != nil {
			logger.Errorf(ctx, "sNotice SendMail reseller: %d, email: %s, SendMail %s error: %v", reseller.UserId, reseller.Email, notice.Title, err)
			continue
		}

		logger.Infof(ctx, "sNotice SendMail reseller: %d, email: %s, SendMail %s success", reseller.UserId, reseller.Email, notice.Title)
	}

	return nil
}

// 通知公告批量操作
func (s *sNotice) BatchOperate(ctx context.Context, params model.NoticeBatchOperateReq) error {

	switch params.Action {
	case consts.ACTION_NOTICE:
		if notices, err := dao.Notice.Find(ctx, bson.M{"_id": bson.M{"$in": params.Ids}, "status": 1}); err != nil {
			logger.Error(ctx, err)
			return err
		} else {
			for _, notice := range notices {
				if err := grpool.AddWithRecover(gctx.NeverDone(ctx), func(ctx context.Context) {
					if err := s.SendMail(ctx, notice); err != nil {
						logger.Error(ctx, err)
					}
				}, nil); err != nil {
					logger.Error(ctx, err)
				}
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
