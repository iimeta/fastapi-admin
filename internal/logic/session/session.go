package session

import (
	"context"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/iimeta/fastapi-admin/internal/consts"
	"github.com/iimeta/fastapi-admin/internal/errors"
	"github.com/iimeta/fastapi-admin/internal/model"
	"github.com/iimeta/fastapi-admin/internal/service"
	"github.com/iimeta/fastapi-admin/utility/logger"
)

type sSession struct{}

func init() {
	service.RegisterSession(New())
}

func New() service.ISession {
	return &sSession{}
}

// 保存代理商会话信息
func (s *sSession) SaveReseller(ctx context.Context, token string, reseller *model.Reseller) error {

	if reseller == nil {
		logger.Error(ctx, "reseller is nil")
		return errors.New("reseller is nil")
	}

	r := g.RequestFromCtx(ctx)

	r.SetCtxVar(consts.SESSION_TOKEN, token)
	r.SetCtxVar(consts.SESSION_UID, reseller.Id)
	r.SetCtxVar(consts.SESSION_RID, reseller.UserId)
	r.SetCtxVar(consts.SESSION_USER_ID, reseller.UserId)
	r.SetCtxVar(consts.SESSION_RESELLER, reseller)
	r.SetCtxVar(consts.SESSION_ROLE, consts.SESSION_RESELLER)
	r.SetCtxVar(consts.SESSION_CREATOR, reseller.Id)

	return nil
}

// 保存用户会话信息
func (s *sSession) SaveUser(ctx context.Context, token string, user *model.User) error {

	if user == nil {
		logger.Error(ctx, "user is nil")
		return errors.New("user is nil")
	}

	r := g.RequestFromCtx(ctx)

	r.SetCtxVar(consts.SESSION_TOKEN, token)
	r.SetCtxVar(consts.SESSION_UID, user.Id)
	r.SetCtxVar(consts.SESSION_RID, user.Rid)
	r.SetCtxVar(consts.SESSION_USER_ID, user.UserId)
	r.SetCtxVar(consts.SESSION_USER, user)
	r.SetCtxVar(consts.SESSION_ROLE, consts.SESSION_USER)
	r.SetCtxVar(consts.SESSION_CREATOR, user.Id)

	return nil
}

// 保存管理员会话信息
func (s *sSession) SaveAdmin(ctx context.Context, token string, admin *model.SysAdmin) error {

	if admin == nil {
		logger.Error(ctx, "admin is nil")
		return errors.New("admin is nil")
	}

	r := g.RequestFromCtx(ctx)

	r.SetCtxVar(consts.SESSION_TOKEN, token)
	r.SetCtxVar(consts.SESSION_UID, admin.Id)
	r.SetCtxVar(consts.SESSION_USER_ID, admin.UserId)
	r.SetCtxVar(consts.SESSION_ADMIN, admin)
	r.SetCtxVar(consts.SESSION_ROLE, consts.SESSION_ADMIN)
	r.SetCtxVar(consts.SESSION_CREATOR, admin.Id)

	return nil
}

// 获取会话中Token
func (s *sSession) GetToken(ctx context.Context) string {

	token := ctx.Value(consts.SESSION_TOKEN)
	if token == nil {
		logger.Error(ctx, "token is nil")
		return ""
	}

	return token.(string)
}

// 获取会话中用户主键ID
func (s *sSession) GetUid(ctx context.Context) string {

	uid := ctx.Value(consts.SESSION_UID)
	if uid == nil {
		return ""
	}

	return uid.(string)
}

// 获取会话中代理商ID
func (s *sSession) GetRid(ctx context.Context) int {

	rid := ctx.Value(consts.SESSION_RID)
	if rid == nil {
		return 0
	}

	return rid.(int)
}

// 获取会话中UserId
func (s *sSession) GetUserId(ctx context.Context) int {

	userId := ctx.Value(consts.SESSION_USER_ID)
	if userId == nil {
		logger.Info(ctx, "user_id is nil")
		return 0
	}

	return userId.(int)
}

// 获取会话中角色
func (s *sSession) GetRole(ctx context.Context) string {

	role := ctx.Value(consts.SESSION_ROLE)
	if role == nil {
		return "nil"
	}

	return role.(string)
}

// 获取会话中创建人
func (s *sSession) GetCreator(ctx context.Context) string {

	creator := ctx.Value(consts.SESSION_CREATOR)
	if creator == nil {
		logger.Error(ctx, "creator is nil")
		return ""
	}

	return creator.(string)
}

// 获取会话中代理商信息
func (s *sSession) GetReseller(ctx context.Context) *model.Reseller {

	reseller := ctx.Value(consts.SESSION_RESELLER)
	if reseller == nil {
		logger.Error(ctx, "reseller is nil")
		return nil
	}

	return reseller.(*model.Reseller)
}

// 获取会话中用户信息
func (s *sSession) GetUser(ctx context.Context) *model.User {

	user := ctx.Value(consts.SESSION_USER)
	if user == nil {
		logger.Error(ctx, "user is nil")
		return nil
	}

	return user.(*model.User)
}

// 获取会话中管理员信息
func (s *sSession) GetAdmin(ctx context.Context) *model.SysAdmin {

	admin := ctx.Value(consts.SESSION_ADMIN)
	if admin == nil {
		logger.Error(ctx, "admin is nil")
		return nil
	}

	return admin.(*model.SysAdmin)
}

// 判断获取会话中角色是否为代理商
func (s *sSession) IsResellerRole(ctx context.Context) bool {
	return s.GetRole(ctx) == consts.SESSION_RESELLER
}

// 判断获取会话中角色是否为用户
func (s *sSession) IsUserRole(ctx context.Context) bool {
	return s.GetRole(ctx) == consts.SESSION_USER
}

// 判断获取会话中角色是否为管理员
func (s *sSession) IsAdminRole(ctx context.Context) bool {
	return s.GetRole(ctx) == consts.SESSION_ADMIN
}

// 更新代理商会话信息
func (s *sSession) UpdateResellerSession(ctx context.Context, reseller *model.Reseller) error {
	return service.Auth().UpdateResellerByToken(ctx, s.GetToken(ctx), reseller)
}

// 更新用户会话信息
func (s *sSession) UpdateUserSession(ctx context.Context, user *model.User) error {
	return service.Auth().UpdateUserByToken(ctx, s.GetToken(ctx), user)
}

// 更新管理员会话信息
func (s *sSession) UpdateAdminSession(ctx context.Context, admin *model.SysAdmin) error {
	return service.Auth().UpdateAdminByToken(ctx, s.GetToken(ctx), admin)
}
