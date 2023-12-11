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

// 保存会话信息
func (s *sSession) Save(ctx context.Context, user *model.User) error {

	if user == nil {
		logger.Error(ctx, "user is nil")
		return errors.New("user is nil")
	}

	r := g.RequestFromCtx(ctx)

	r.SetCtxVar("uid", user.Id)
	r.SetCtxVar("user_id", user.UserId)
	r.SetCtxVar("user", user)

	return nil
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

	user, err := service.User().GetUserById(ctx, s.GetUserId(ctx))
	if err != nil {
		logger.Error(ctx, err)
		return nil
	}

	g.RequestFromCtx(ctx).SetCtxVar("user", user)

	return user
}
