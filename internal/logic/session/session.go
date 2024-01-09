package session

import (
	"context"
	"github.com/gogf/gf/v2/frame/g"
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

// 保存用户会话信息
func (s *sSession) SaveUser(ctx context.Context, token string, user *model.User) error {

	if user == nil {
		logger.Error(ctx, "user is nil")
		return errors.New("user is nil")
	}

	r := g.RequestFromCtx(ctx)

	r.SetCtxVar("token", token)
	r.SetCtxVar("uid", user.Id)
	r.SetCtxVar("user_id", user.UserId)
	r.SetCtxVar("user", user)

	return nil
}

// 保存管理员会话信息
func (s *sSession) SaveAdmin(ctx context.Context, token string, admin *model.SysAdmin) error {

	if admin == nil {
		logger.Error(ctx, "admin is nil")
		return errors.New("admin is nil")
	}

	r := g.RequestFromCtx(ctx)

	r.SetCtxVar("token", token)
	r.SetCtxVar("uid", admin.Id)
	r.SetCtxVar("admin", admin)

	return nil
}

// 获取会话中Token
func (s *sSession) GetToken(ctx context.Context) string {

	token := ctx.Value("token")
	if token == nil {
		logger.Error(ctx, "token is nil")
		return ""
	}

	return token.(string)
}

// 获取会话中用户主键ID
func (s *sSession) GetUid(ctx context.Context) string {

	uid := ctx.Value("uid")
	if uid == nil {
		logger.Error(ctx, "uid is nil")
		return ""
	}

	return uid.(string)
}

// 获取会话中UserId
func (s *sSession) GetUserId(ctx context.Context) int {

	userId := ctx.Value("user_id")
	if userId == nil {
		logger.Error(ctx, "user_id is nil")
		return 0
	}

	return userId.(int)
}

// 获取会话中用户信息
func (s *sSession) GetUser(ctx context.Context) *model.User {

	value := ctx.Value("user")
	if value != nil {
		return value.(*model.User)
	}

	return nil
}

// 获取会话中用户信息
func (s *sSession) GetAdmin(ctx context.Context) *model.SysAdmin {

	value := ctx.Value("admin")
	if value != nil {
		return value.(*model.SysAdmin)
	}

	return nil
}
