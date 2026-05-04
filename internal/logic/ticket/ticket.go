package ticket

import (
	"context"
	"regexp"
	"strings"

	"github.com/gogf/gf/v2/os/gtime"
	"github.com/gogf/gf/v2/util/grand"
	"github.com/iimeta/fastapi-admin/v2/internal/consts"
	"github.com/iimeta/fastapi-admin/v2/internal/dao"
	"github.com/iimeta/fastapi-admin/v2/internal/errors"
	"github.com/iimeta/fastapi-admin/v2/internal/model"
	"github.com/iimeta/fastapi-admin/v2/internal/model/do"
	"github.com/iimeta/fastapi-admin/v2/internal/service"
	"github.com/iimeta/fastapi-admin/v2/utility/db"
	"github.com/iimeta/fastapi-admin/v2/utility/logger"
	"github.com/iimeta/fastapi-admin/v2/utility/util"
	"go.mongodb.org/mongo-driver/v2/bson"
)

type sTicket struct{}

func init() {
	service.RegisterTicket(New())
}

func New() service.ITicket {
	return &sTicket{}
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
	} else {
		// 用户/代理商提交者回复 → 待回复(1), 在待处理/处理中/已回复状态时触发
		if ticket.Status == consts.STATUS_PENDING || ticket.Status == consts.STATUS_PROCESSING || ticket.Status == consts.STATUS_REPLIED {
			update["$set"].(bson.M)["status"] = consts.STATUS_AWAITING_REPLY
		}
	}

	if err = dao.Ticket.UpdateOne(ctx, bson.M{"_id": params.TicketId}, update); err != nil {
		logger.Error(ctx, err)
		return "", err
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
		"status":     params.Status,
		"updated_at": gtime.TimestampMilli(),
	}); err != nil {
		logger.Error(ctx, err)
		return err
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
