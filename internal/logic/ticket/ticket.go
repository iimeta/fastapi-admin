package ticket

import (
	"context"
	"fmt"
	"regexp"
	"strings"

	"github.com/go-redsync/redsync/v4"
	"github.com/go-redsync/redsync/v4/redis/goredis/v9"
	"github.com/gogf/gf/v2/os/gctx"
	"github.com/gogf/gf/v2/os/gtime"
	"github.com/gogf/gf/v2/util/grand"
	"github.com/iimeta/fastapi-admin/v2/internal/config"
	"github.com/iimeta/fastapi-admin/v2/internal/consts"
	"github.com/iimeta/fastapi-admin/v2/internal/dao"
	"github.com/iimeta/fastapi-admin/v2/internal/errors"
	"github.com/iimeta/fastapi-admin/v2/internal/model"
	"github.com/iimeta/fastapi-admin/v2/internal/model/do"
	"github.com/iimeta/fastapi-admin/v2/internal/model/entity"
	"github.com/iimeta/fastapi-admin/v2/internal/service"
	"github.com/iimeta/fastapi-admin/v2/utility/db"
	"github.com/iimeta/fastapi-admin/v2/utility/email"
	"github.com/iimeta/fastapi-admin/v2/utility/logger"
	"github.com/iimeta/fastapi-admin/v2/utility/redis"
	"github.com/iimeta/fastapi-admin/v2/utility/util"
	"go.mongodb.org/mongo-driver/v2/bson"
)

type sTicket struct {
	ticketRedsync *redsync.Redsync
}

func init() {
	service.RegisterTicket(New())
}

func New() service.ITicket {
	return &sTicket{
		ticketRedsync: redsync.New(goredis.NewPool(redis.UniversalClient)),
	}
}

// 创建工单
func (s *sTicket) Create(ctx context.Context, params model.TicketCreateReq) (string, error) {

	if params.Content == "" || params.Content == "\n" {
		return "", errors.New("请输入工单内容")
	}

	userId := service.Session().GetUserId(ctx)
	rid := service.Session().GetRid(ctx)
	role := service.Session().GetRole(ctx)

	userName := ""
	switch role {
	case consts.SESSION_USER:
		user := service.Session().GetUser(ctx)
		if user != nil {
			userName = user.Name
		}
	case consts.SESSION_RESELLER:
		reseller := service.Session().GetReseller(ctx)
		if reseller != nil {
			userName = reseller.Name
		}
	}

	ticket := &do.Ticket{
		TicketNo: "TK" + strings.ToUpper(grand.Letters(8)),
		Title:    params.Title,
		Content:  params.Content,
		Category: params.Category,
		Priority: params.Priority,
		Status:   consts.STATUS_PENDING, // 待处理
		UserId:   userId,
		UserName: userName,
		UserRole: role, // 提交者真实角色
	}

	// 代理商提交工单给管理员处理
	if role == consts.SESSION_RESELLER {
		ticket.AssigneeId = 0
		ticket.AssigneeRole = consts.SESSION_ADMIN
	} else if rid != 0 {
		// 用户有代理商, 分配给代理商
		ticket.AssigneeId = rid
		ticket.AssigneeRole = consts.SESSION_RESELLER
	} else {
		// 用户无代理商, 分配给管理员
		ticket.AssigneeId = 0
		ticket.AssigneeRole = consts.SESSION_ADMIN
	}

	id, err := dao.Ticket.Insert(ctx, ticket)
	if err != nil {
		logger.Error(ctx, err)
		return "", err
	}

	if newTicket, err := dao.Ticket.FindById(ctx, id); err != nil {
		logger.Error(ctx, err)
	} else {
		s.noticeAssignee(ctx, newTicket, "工单创建提醒")
	}

	return id, nil
}

// 回复工单
func (s *sTicket) Reply(ctx context.Context, params model.TicketReplyReq) (string, error) {

	if params.Content == "" || params.Content == "\n" {
		return "", errors.New("请输入回复内容")
	}

	ticket, err := dao.Ticket.FindById(ctx, params.TicketId)
	if err != nil {
		logger.Error(ctx, err)
		return "", err
	}

	if ticket.Status == consts.STATUS_CLOSED {
		return "", errors.New("工单已关闭, 无法回复")
	}

	role := service.Session().GetRole(ctx)
	userId := service.Session().GetUserId(ctx)

	// 代理商回复自己提交的工单时, 身份视为用户
	replyRole := role
	if role == consts.SESSION_RESELLER && ticket.UserId == userId {
		replyRole = consts.SESSION_USER
	}

	userName := "admin"

	switch role {
	case consts.SESSION_USER:
		user := service.Session().GetUser(ctx)
		if user != nil {
			userName = user.Name
		} else {
			userName = "user"
		}
	case consts.SESSION_RESELLER:
		reseller := service.Session().GetReseller(ctx)
		if reseller != nil && reseller.Name != reseller.Account {
			userName = reseller.Name
		}
	case consts.SESSION_ADMIN:
		admin := service.Session().GetAdmin(ctx)
		if admin != nil && admin.Name != admin.Account {
			userName = admin.Name
		}
	}

	reply := &do.TicketReply{
		TicketId: params.TicketId,
		Content:  params.Content,
		UserId:   userId,
		UserName: userName,
		Role:     role, // 存储真实角色, 前端根据查看者角色做显示映射
	}

	replyId, err := dao.TicketReply.Insert(ctx, reply)
	if err != nil {
		logger.Error(ctx, err)
		return "", err
	}

	// 更新工单回复数和最后回复时间
	update := bson.M{
		"$inc": bson.M{"reply_count": 1},
		"$set": bson.M{
			"last_reply_at": gtime.TimestampMilli(),
			"updated_at":    gtime.TimestampMilli(),
		},
	}

	// 根据回复者角色自动更新工单状态(使用replyRole判断)
	if replyRole == consts.SESSION_RESELLER || replyRole == consts.SESSION_ADMIN {
		// 工作人员回复 → 已回复(4), 在待处理/处理中/待回复状态时触发
		if ticket.Status == consts.STATUS_PENDING || ticket.Status == consts.STATUS_PROCESSING || ticket.Status == consts.STATUS_AWAITING_REPLY {
			update["$set"].(bson.M)["status"] = consts.STATUS_REPLIED
		}
		update["$set"].(bson.M)["assignee_notice"] = false
	} else {
		// 用户/代理商提交者回复 → 待回复(1), 在待处理/处理中/已回复状态时触发
		if ticket.Status == consts.STATUS_PENDING || ticket.Status == consts.STATUS_PROCESSING || ticket.Status == consts.STATUS_REPLIED {
			update["$set"].(bson.M)["status"] = consts.STATUS_AWAITING_REPLY
		}
		update["$set"].(bson.M)["user_notice"] = false
	}

	if err = dao.Ticket.UpdateOne(ctx, bson.M{"_id": params.TicketId}, update); err != nil {
		logger.Error(ctx, err)
		return "", err
	}

	if newTicket, err := dao.Ticket.FindById(ctx, params.TicketId); err != nil {
		logger.Error(ctx, err)
	} else {
		if replyRole == consts.SESSION_RESELLER || replyRole == consts.SESSION_ADMIN {
			s.noticeUser(ctx, newTicket, "工单回复提醒")
		} else {
			s.noticeAssignee(ctx, newTicket, "工单回复提醒")
		}
	}

	return replyId, nil
}

// 工单详情
func (s *sTicket) Detail(ctx context.Context, id string) (*model.TicketDetailRes, error) {

	ticket, err := dao.Ticket.FindById(ctx, id)
	if err != nil {
		logger.Error(ctx, err)
		return nil, err
	}

	// 获取回复列表
	replies, err := dao.TicketReply.Find(ctx, bson.M{"ticket_id": id}, &dao.FindOptions{SortFields: []string{"created_at"}})
	if err != nil {
		logger.Error(ctx, err)
		return nil, err
	}

	replyItems := make([]*model.TicketReply, 0)
	for _, reply := range replies {
		replyItems = append(replyItems, &model.TicketReply{
			Id:        reply.Id,
			TicketId:  reply.TicketId,
			Content:   reply.Content,
			UserId:    reply.UserId,
			UserName:  reply.UserName,
			Role:      reply.Role,
			CreatedAt: util.FormatDateTime(reply.CreatedAt),
		})
	}

	return &model.TicketDetailRes{
		Ticket: &model.Ticket{
			Id:           ticket.Id,
			TicketNo:     ticket.TicketNo,
			Title:        ticket.Title,
			Content:      ticket.Content,
			Category:     ticket.Category,
			Priority:     ticket.Priority,
			Status:       ticket.Status,
			UserId:       ticket.UserId,
			UserName:     ticket.UserName,
			UserRole:     ticket.UserRole,
			AssigneeId:   ticket.AssigneeId,
			AssigneeRole: ticket.AssigneeRole,
			ReplyCount:   ticket.ReplyCount,
			LastReplyAt:  util.FormatDateTime(ticket.LastReplyAt),
			Rid:          ticket.Rid,
			Creator:      ticket.Creator,
			Updater:      ticket.Updater,
			CreatedAt:    util.FormatDateTime(ticket.CreatedAt),
			UpdatedAt:    util.FormatDateTime(ticket.UpdatedAt),
		},
		Replies: replyItems,
	}, nil
}

// 工单分页列表
func (s *sTicket) Page(ctx context.Context, params model.TicketPageReq) (*model.TicketPageRes, error) {

	paging := &db.Paging{
		Page:     params.Page,
		PageSize: params.PageSize,
	}

	filter := bson.M{}

	if params.TicketNo != "" {
		filter["ticket_no"] = bson.M{
			"$regex":   regexp.QuoteMeta(params.TicketNo),
			"$options": "i",
		}
	}

	if params.Title != "" {
		filter["title"] = bson.M{
			"$regex": regexp.QuoteMeta(params.Title),
		}
	}

	if params.Category != "" {
		filter["category"] = params.Category
	}

	if params.Priority != 0 {
		filter["priority"] = params.Priority
	}

	if params.Status != 0 {
		filter["status"] = params.Status
	}

	if params.UserName != "" {
		filter["$or"] = bson.A{
			bson.M{"user_name": bson.M{
				"$regex": regexp.QuoteMeta(params.UserName),
			}},
			bson.M{"user_id": bson.M{
				"$regex": regexp.QuoteMeta(params.UserName),
			}},
		}
	}

	if params.CreatedAtStart != "" && params.CreatedAtEnd != "" {
		if start := gtime.NewFromStr(params.CreatedAtStart); start != nil {
			if end := gtime.NewFromStr(params.CreatedAtEnd); end != nil {
				filter["created_at"] = bson.M{
					"$gte": start.StartOfDay().TimestampMilli(),
					"$lte": end.EndOfDay(true).TimestampMilli(),
				}
			}
		}
	}

	if params.Scope == "my" {
		userId := service.Session().GetUserId(ctx)
		filter["user_id"] = userId
	} else {
		role := service.Session().GetRole(ctx)
		if role == consts.SESSION_RESELLER {
			userId := service.Session().GetUserId(ctx)
			filter["assignee_id"] = userId
			filter["assignee_role"] = consts.SESSION_RESELLER
		}
	}

	results, err := dao.Ticket.FindByPage(ctx, paging, filter, &dao.FindOptions{SortFields: []string{"-priority", "status", "-updated_at"}})
	if err != nil {
		logger.Error(ctx, err)
		return nil, err
	}

	items := make([]*model.Ticket, 0)
	for _, result := range results {
		items = append(items, &model.Ticket{
			Id:           result.Id,
			TicketNo:     result.TicketNo,
			Title:        result.Title,
			Category:     result.Category,
			Priority:     result.Priority,
			Status:       result.Status,
			UserId:       result.UserId,
			UserName:     result.UserName,
			AssigneeId:   result.AssigneeId,
			AssigneeRole: result.AssigneeRole,
			ReplyCount:   result.ReplyCount,
			LastReplyAt:  util.FormatDateTime(result.LastReplyAt),
			Rid:          result.Rid,
			CreatedAt:    util.FormatDateTime(result.CreatedAt),
			UpdatedAt:    util.FormatDateTime(result.UpdatedAt),
		})
	}

	return &model.TicketPageRes{
		Items: items,
		Paging: &model.Paging{
			Page:     paging.Page,
			PageSize: paging.PageSize,
			Total:    paging.Total,
		},
	}, nil
}

// 更新工单状态
func (s *sTicket) UpdateStatus(ctx context.Context, params model.TicketUpdateStatusReq) error {

	ticket, err := dao.Ticket.FindById(ctx, params.Id)
	if err != nil {
		logger.Error(ctx, err)
		return err
	}

	role := service.Session().GetRole(ctx)

	// 验证状态流转合法性
	if !s.isValidTransition(ticket.Status, params.Status, role) {
		return errors.New("无效的状态变更")
	}

	if err = dao.Ticket.UpdateById(ctx, params.Id, bson.M{
		"status":          params.Status,
		"user_notice":     false,
		"assignee_notice": false,
		"updated_at":      gtime.TimestampMilli(),
	}); err != nil {
		logger.Error(ctx, err)
		return err
	}

	if newTicket, err := dao.Ticket.FindById(ctx, params.Id); err != nil {
		logger.Error(ctx, err)
	} else {
		if role == consts.SESSION_RESELLER || role == consts.SESSION_ADMIN {
			s.noticeUser(ctx, newTicket, "工单状态变更提醒")
		} else {
			s.noticeAssignee(ctx, newTicket, "工单状态变更提醒")
		}
	}

	return nil
}

// 关闭工单
func (s *sTicket) Close(ctx context.Context, params model.TicketCloseReq) error {

	ticket, err := dao.Ticket.FindById(ctx, params.Id)
	if err != nil {
		logger.Error(ctx, err)
		return err
	}

	if ticket.Status == consts.STATUS_CLOSED {
		return errors.New("工单已关闭")
	}

	if err = dao.Ticket.UpdateById(ctx, params.Id, bson.M{
		"status":     consts.STATUS_CLOSED,
		"updated_at": gtime.TimestampMilli(),
	}); err != nil {
		logger.Error(ctx, err)
		return err
	}

	return nil
}

// 分配工单
func (s *sTicket) Assign(ctx context.Context, params model.TicketAssignReq) error {

	if _, err := dao.Ticket.FindById(ctx, params.Id); err != nil {
		logger.Error(ctx, err)
		return err
	}

	if err := dao.Ticket.UpdateById(ctx, params.Id, bson.M{
		"assignee_id":   params.AssigneeId,
		"assignee_role": consts.SESSION_ADMIN,
		"updated_at":    gtime.TimestampMilli(),
	}); err != nil {
		logger.Error(ctx, err)
		return err
	}

	return nil
}

// 验证状态流转合法性
// 状态: 1-待回复, 2-待处理, 3-处理中, 4-已回复, 5-已解决, 6-已关闭
func (s *sTicket) isValidTransition(currentStatus, targetStatus int, role string) bool {

	switch currentStatus {
	case consts.STATUS_AWAITING_REPLY: // 待回复 → 处理中/已解决/已关闭
		if targetStatus == consts.STATUS_PROCESSING && (role == consts.SESSION_RESELLER || role == consts.SESSION_ADMIN) {
			return true
		}
		if targetStatus == consts.STATUS_RESOLVED && (role == consts.SESSION_RESELLER || role == consts.SESSION_ADMIN) {
			return true
		}
		if targetStatus == consts.STATUS_CLOSED {
			return true
		}
	case consts.STATUS_PENDING: // 待处理 → 处理中/已关闭
		if targetStatus == consts.STATUS_PROCESSING && (role == consts.SESSION_RESELLER || role == consts.SESSION_ADMIN) {
			return true
		}
		if targetStatus == consts.STATUS_CLOSED {
			return true
		}
	case consts.STATUS_PROCESSING: // 处理中 → 已解决/已关闭
		if targetStatus == consts.STATUS_RESOLVED && (role == consts.SESSION_RESELLER || role == consts.SESSION_ADMIN) {
			return true
		}
		if targetStatus == consts.STATUS_CLOSED {
			return true
		}
	case consts.STATUS_REPLIED: // 已回复 → 已解决/已关闭
		if targetStatus == consts.STATUS_RESOLVED && (role == consts.SESSION_RESELLER || role == consts.SESSION_ADMIN) {
			return true
		}
		if targetStatus == consts.STATUS_CLOSED {
			return true
		}
	case consts.STATUS_RESOLVED: // 已解决 → 待处理(重新打开)/已关闭
		if targetStatus == consts.STATUS_PENDING {
			return true // 重新打开
		}
		if targetStatus == consts.STATUS_CLOSED {
			return true
		}
	case consts.STATUS_CLOSED: // 已关闭 → 待处理(重新打开)
		if targetStatus == consts.STATUS_PENDING {
			return true // 重新打开
		}
	}

	return false
}

// 删除工单
func (s *sTicket) Delete(ctx context.Context, id string) error {

	// 删除工单
	if _, err := dao.Ticket.DeleteById(ctx, id); err != nil {
		logger.Error(ctx, err)
		return err
	}

	// 删除关联回复
	if _, err := dao.TicketReply.DeleteMany(ctx, bson.M{"ticket_id": id}); err != nil {
		logger.Error(ctx, err)
	}

	return nil
}

// 批量删除工单
func (s *sTicket) BatchDelete(ctx context.Context, params model.TicketBatchDeleteReq) error {

	role := service.Session().GetRole(ctx)
	if role != consts.SESSION_ADMIN {
		return errors.New("无权限操作")
	}

	for _, id := range params.Ids {
		if err := s.Delete(ctx, id); err != nil {
			logger.Error(ctx, err)
			return err
		}
	}

	return nil
}

func (s *sTicket) noticeUser(ctx context.Context, ticket *entity.Ticket, action string) {

	if ticket == nil || ticket.UserNotice || ticket.Status == consts.STATUS_CLOSED || config.Cfg.Ticket == nil || !config.Cfg.Ticket.Notice {
		return
	}

	to, name, dialer, err := s.userMailTarget(ctx, ticket)
	if err != nil {
		logger.Infof(ctx, "sTicket noticeUser ticket: %s, error: %v", ticket.Id, err)
		return
	}

	if err = s.sendNoticeMail(ctx, to, name, ticket, action, dialer); err != nil {
		logger.Errorf(ctx, "sTicket noticeUser ticket: %s, email: %s, error: %v", ticket.Id, to, err)
		return
	}

	if err = dao.Ticket.UpdateById(gctx.New(), ticket.Id, bson.M{"user_notice": true}); err != nil {
		logger.Error(ctx, err)
	}
}

func (s *sTicket) noticeAssignee(ctx context.Context, ticket *entity.Ticket, action string) {

	if ticket == nil || ticket.AssigneeNotice || ticket.Status == consts.STATUS_CLOSED || config.Cfg.Ticket == nil || !config.Cfg.Ticket.Notice {
		return
	}

	to, name, dialer, err := s.assigneeMailTarget(ctx, ticket)
	if err != nil {
		logger.Infof(ctx, "sTicket noticeAssignee ticket: %s, error: %v", ticket.Id, err)
		return
	}

	if err = s.sendNoticeMail(ctx, to, name, ticket, action, dialer); err != nil {
		logger.Errorf(ctx, "sTicket noticeAssignee ticket: %s, email: %s, error: %v", ticket.Id, to, err)
		return
	}

	if err = dao.Ticket.UpdateById(gctx.New(), ticket.Id, bson.M{"assignee_notice": true}); err != nil {
		logger.Error(ctx, err)
	}
}

func (s *sTicket) userMailTarget(ctx context.Context, ticket *entity.Ticket) (string, string, *email.Dialer, error) {

	switch ticket.UserRole {
	case consts.SESSION_RESELLER:
		reseller, err := dao.Reseller.FindOne(ctx, bson.M{"user_id": ticket.UserId})
		if err != nil {
			return "", "", nil, err
		}
		dialer, err := s.dialerForReseller(ctx, reseller)
		if err != nil {
			return "", "", nil, err
		}
		return reseller.Email, reseller.Name, dialer, email.Verify(reseller.Email)
	default:
		user, err := dao.User.FindOne(ctx, bson.M{"user_id": ticket.UserId})
		if err != nil {
			return "", "", nil, err
		}
		dialer, err := s.dialerForUser(ctx, user)
		if err != nil {
			return "", "", nil, err
		}
		return user.Email, user.Name, dialer, email.Verify(user.Email)
	}
}

func (s *sTicket) assigneeMailTarget(ctx context.Context, ticket *entity.Ticket) (string, string, *email.Dialer, error) {

	switch ticket.AssigneeRole {
	case consts.SESSION_RESELLER:
		reseller, err := dao.Reseller.FindOne(gctx.New(), bson.M{"user_id": ticket.AssigneeId})
		if err != nil {
			return "", "", nil, err
		}
		dialer, err := s.dialerForReseller(ctx, reseller)
		if err != nil {
			return "", "", nil, err
		}
		return reseller.Email, reseller.Name, dialer, email.Verify(reseller.Email)
	case consts.SESSION_ADMIN:
		filter := bson.M{"status": 1}
		if ticket.AssigneeId > 0 {
			filter["user_id"] = ticket.AssigneeId
		} else {
			filter["is_super_admin"] = true
		}
		admin, err := dao.SysAdmin.FindOne(gctx.New(), filter, &dao.FindOptions{SortFields: []string{"-is_super_admin", "-updated_at"}})
		if err != nil {
			return "", "", nil, err
		}
		dialer := s.dialerByDomain(ctx, admin.LoginDomain)
		return admin.Email, admin.Name, dialer, email.Verify(admin.Email)
	default:
		return "", "", nil, errors.New("工单处理人角色无效")
	}
}

func (s *sTicket) dialerForUser(ctx context.Context, user *entity.User) (*email.Dialer, error) {

	dialer := email.NewDefaultDialer()
	account, err := dao.Account.FindOne(ctx, bson.M{"user_id": user.UserId, "status": 1}, &dao.FindOptions{SortFields: []string{"-updated_at"}})
	if err != nil {
		logger.Error(ctx, err)
	}

	if user.Rid > 0 {
		if account != nil {
			if siteConfig := service.SiteConfig().GetSiteConfigByDomain(ctx, account.LoginDomain); siteConfig != nil && siteConfig.Rid == user.Rid && siteConfig.Host != "" {
				return email.NewDialer(siteConfig.Host, siteConfig.Port, siteConfig.UserName, siteConfig.Password, siteConfig.FromName), nil
			}
		}
		for _, siteConfig := range service.SiteConfig().GetSiteConfigsByRid(ctx, user.Rid) {
			if siteConfig != nil && siteConfig.Host != "" {
				return email.NewDialer(siteConfig.Host, siteConfig.Port, siteConfig.UserName, siteConfig.Password, siteConfig.FromName), nil
			}
		}
		return nil, errors.Newf("代理商 %d 未配置发信邮箱", user.Rid)
	}

	if account != nil {
		dialer = s.dialerByDomain(ctx, account.LoginDomain)
	}

	return dialer, nil
}

func (s *sTicket) dialerForReseller(ctx context.Context, reseller *entity.Reseller) (*email.Dialer, error) {

	account, err := dao.ResellerAccount.FindOne(ctx, bson.M{"user_id": reseller.UserId, "status": 1}, &dao.FindOptions{SortFields: []string{"-updated_at"}})
	if err != nil {
		logger.Error(ctx, err)
	}

	if account != nil {
		return s.dialerByDomain(ctx, account.LoginDomain), nil
	}

	return email.NewDefaultDialer(), nil
}

func (s *sTicket) dialerByDomain(ctx context.Context, domain string) *email.Dialer {

	if domain != "" {
		if siteConfig := service.SiteConfig().GetSiteConfigByDomain(ctx, domain); siteConfig != nil && siteConfig.Host != "" {
			return email.NewDialer(siteConfig.Host, siteConfig.Port, siteConfig.UserName, siteConfig.Password, siteConfig.FromName)
		}
	}

	if siteConfig, err := dao.SiteConfig.FindOne(gctx.New(), bson.M{"user_id": 1, "status": 1}, &dao.FindOptions{SortFields: []string{"-updated_at"}}); err == nil && siteConfig != nil && siteConfig.Host != "" {
		return email.NewDialer(siteConfig.Host, siteConfig.Port, siteConfig.UserName, siteConfig.Password, siteConfig.FromName)
	}

	return email.NewDefaultDialer()
}

func (s *sTicket) sendNoticeMail(ctx context.Context, to, name string, ticket *entity.Ticket, action string, dialer *email.Dialer) error {

	if to == "" {
		return errors.New("收件邮箱为空")
	}

	title := fmt.Sprintf("工单%s：%s", action, ticket.Title)
	content := fmt.Sprintf(`<div style="max-width:600px;margin:0 auto;background:#fff;border-radius:8px;overflow:hidden;border:1px solid #e5e7eb;font-family:Arial,'Microsoft YaHei',sans-serif;line-height:1.6;color:#333"><div style="background:#2563eb;padding:24px;text-align:center;color:#fff"><span style="font-size:22px;font-weight:700">%s</span></div><div style="padding:28px"><p style="margin:0 0 14px">尊敬的 <strong>%s</strong>：</p><p style="margin:0 0 18px">工单 <strong>%s</strong> 有新的动态，请及时查看处理。</p><div style="background:#f8fafc;border-radius:8px;padding:16px;border:1px solid #e2e8f0"><p style="margin:0 0 8px"><strong>工单编号：</strong>%s</p><p style="margin:0 0 8px"><strong>标题：</strong>%s</p><p style="margin:0"><strong>当前状态：</strong>%s</p></div></div><div style="padding:18px;text-align:center;font-size:12px;color:#9ca3af;background:#f9fafb"><p style="margin:0">此为系统邮件，请勿直接回复</p></div></div>`,
		action, name, ticket.TicketNo, ticket.TicketNo, ticket.Title, s.statusName(ticket.Status))

	return email.SendMailTask(ctx, email.NewMessage([]string{to}, title, content), dialer)
}

func (s *sTicket) statusName(status int) string {
	switch status {
	case consts.STATUS_AWAITING_REPLY:
		return "待回复"
	case consts.STATUS_PENDING:
		return "待处理"
	case consts.STATUS_PROCESSING:
		return "处理中"
	case consts.STATUS_REPLIED:
		return "已回复"
	case consts.STATUS_RESOLVED:
		return "已解决"
	case consts.STATUS_CLOSED:
		return "已关闭"
	default:
		return "未知"
	}
}
