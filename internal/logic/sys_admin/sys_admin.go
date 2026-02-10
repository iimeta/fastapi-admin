package sys_admin

import (
	"context"
	"fmt"

	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/util/grand"
	"github.com/iimeta/fastapi-admin/v2/internal/consts"
	"github.com/iimeta/fastapi-admin/v2/internal/dao"
	"github.com/iimeta/fastapi-admin/v2/internal/errors"
	"github.com/iimeta/fastapi-admin/v2/internal/model"
	"github.com/iimeta/fastapi-admin/v2/internal/model/do"
	"github.com/iimeta/fastapi-admin/v2/internal/service"
	"github.com/iimeta/fastapi-admin/v2/utility/crypto"
	"github.com/iimeta/fastapi-admin/v2/utility/db"
	"github.com/iimeta/fastapi-admin/v2/utility/logger"
	"github.com/iimeta/fastapi-admin/v2/utility/redis"
	"github.com/iimeta/fastapi-admin/v2/utility/util"
	"go.mongodb.org/mongo-driver/v2/bson"
)

type sSysAdmin struct{}

func init() {
	service.RegisterSysAdmin(New())
}

func New() service.ISysAdmin {
	return &sSysAdmin{}
}

// 管理员更新信息
func (s *sSysAdmin) UpdateInfo(ctx context.Context, params model.UserUpdateInfoReq) error {

	if err := dao.SysAdmin.UpdateById(ctx, service.Session().GetUid(ctx), bson.M{
		"name": params.Name,
	}); err != nil {
		logger.Error(ctx, err)
		return err
	}

	admin := service.Session().GetAdmin(ctx)
	admin.Name = params.Name

	if err := service.Session().UpdateAdminSession(ctx, admin); err != nil {
		logger.Error(ctx, err)
		return err
	}

	return nil
}

// 管理员更改密码
func (s *sSysAdmin) ChangePassword(ctx context.Context, params model.UserChangePasswordReq) (err error) {

	uid := service.Session().GetUserId(ctx)

	defer func() {
		if err != nil {
			val, _ := redis.Incr(ctx, fmt.Sprintf(consts.LOCK_CHANGE_PASSWORD, uid))
			if val == 1 {
				_, _ = redis.Expire(ctx, fmt.Sprintf(consts.LOCK_CHANGE_PASSWORD, uid), 30*60) // 锁定30分钟
			}
		} else {
			_, _ = redis.Del(ctx, fmt.Sprintf(consts.LOCK_CHANGE_PASSWORD, uid))
		}
	}()

	if val, err := redis.GetInt(ctx, fmt.Sprintf(consts.LOCK_CHANGE_PASSWORD, uid)); err == nil && val >= 5 {
		return errors.New("失败次数过多, 请稍后再试")
	}

	admin, err := dao.SysAdmin.FindById(ctx, service.Session().GetUid(ctx))
	if err != nil || admin.Id == "" {
		return errors.New("管理员不存在")
	}

	if !crypto.VerifyPassword(admin.Password, params.OldPassword+admin.Salt) {
		return errors.New("登录密码有误, 请重新输入")
	}

	if err = dao.SysAdmin.ChangePassword(ctx, service.Session().GetUid(ctx), params.NewPassword); err != nil {
		logger.Error(ctx, err)
		return errors.New("修改密码失败")
	}

	return nil
}

// 管理员更改邮箱
func (s *sSysAdmin) ChangeEmail(ctx context.Context, params model.UserChangeEmailReq) error {

	if !service.Common().VerifyCode(ctx, consts.SCENE_CHANGE_EMAIL, params.Email, params.Code) {
		return errors.New("邮件验证码填写错误")
	}

	admin, err := dao.SysAdmin.FindById(ctx, service.Session().GetUid(ctx))
	if err != nil {
		logger.Error(ctx, err)
		return err
	}

	if !crypto.VerifyPassword(admin.Password, params.Password+admin.Salt) {
		return errors.New("登录密码有误, 请重新输入")
	}

	defer func() {
		_ = service.Common().DelCode(ctx, consts.SCENE_CHANGE_EMAIL, params.Email)
	}()

	if admin.Email == params.Email {
		return errors.New("邮箱与原邮箱一致无需修改")
	}

	count, err := dao.SysAdmin.CountDocuments(ctx, bson.M{"email": params.Email})
	if err != nil {
		logger.Error(ctx, err)
		return err
	}

	if count > 0 {
		return errors.New(params.Email + " 邮箱已被其它账号使用")
	}

	if err = dao.SysAdmin.UpdateById(ctx, admin.Id, bson.M{
		"email": params.Email,
	}); err != nil {
		logger.Error(ctx, err)
		return errors.New("邮箱修改失败")
	}

	adminSession := service.Session().GetAdmin(ctx)
	adminSession.Email = params.Email

	if err = service.Session().UpdateAdminSession(ctx, adminSession); err != nil {
		logger.Error(ctx, err)
		return err
	}

	return nil
}

// 管理员更改头像
func (s *sSysAdmin) ChangeAvatar(ctx context.Context, file *ghttp.UploadFile) error {

	if file.Size > 1024*1024*8 {
		return errors.New("头像文件过大, 请更换或压缩后再上传")
	}

	root := "./resource"
	path := "/public/avatar/"

	filename, err := file.Save(root+path, true)
	if err != nil {
		logger.Error(ctx, err)
		return err
	}

	if err = dao.SysAdmin.UpdateById(ctx, service.Session().GetUid(ctx), bson.M{
		"avatar": path + filename,
	}); err != nil {
		logger.Error(ctx, err)
		return err
	}

	admin := service.Session().GetAdmin(ctx)
	admin.Avatar = path + filename

	if err = service.Session().UpdateAdminSession(ctx, admin); err != nil {
		logger.Error(ctx, err)
		return err
	}

	return nil
}

// 新建管理员
func (s *sSysAdmin) Create(ctx context.Context, params model.SysAdminCreateReq) error {

	count, err := dao.SysAdmin.EstimatedDocumentCount(ctx)
	if err != nil {
		logger.Error(ctx, err)
		return err
	}

	id := util.GenerateId()

	creator := service.Session().GetUid(ctx)
	if creator == "" {
		creator = id
	}

	salt := grand.Letters(8)

	sysAdmin := &do.SysAdmin{
		Id:       id,
		UserId:   int(count + 1),
		Name:     params.Name,
		Avatar:   params.Avatar,
		Email:    params.Email,
		Phone:    params.Phone,
		Account:  params.Account,
		Password: crypto.EncryptPassword(params.Password + salt),
		Salt:     salt,
		Remark:   params.Remark,
		Status:   params.Status,
		Creator:  creator,
	}

	if count == 0 {
		sysAdmin.IsSuperAdmin = true
		sysAdmin.IsSysAdmin = true
	}

	if _, err = dao.SysAdmin.Insert(ctx, sysAdmin); err != nil {
		logger.Error(ctx, err)
		return err
	}

	return nil
}

// 更新管理员
func (s *sSysAdmin) Update(ctx context.Context, params model.SysAdminUpdateReq) error {

	if err := dao.SysAdmin.UpdateById(ctx, params.Id, &do.SysAdmin{
		Name:   params.Name,
		Avatar: params.Avatar,
		Email:  params.Email,
		Phone:  params.Phone,
		Remark: params.Remark,
		Status: params.Status,
	}); err != nil {
		logger.Error(ctx, err)
		return err
	}

	return nil
}

// 删除管理员
func (s *sSysAdmin) Delete(ctx context.Context, id string) error {

	if _, err := dao.SysAdmin.DeleteById(ctx, id); err != nil {
		logger.Error(ctx, err)
		return err
	}

	return nil
}

// 管理员详情
func (s *sSysAdmin) Detail(ctx context.Context, id string) (*model.SysAdmin, error) {

	admin, err := dao.SysAdmin.FindById(ctx, id)
	if err != nil {
		logger.Error(ctx, err)
		return nil, err
	}

	return &model.SysAdmin{
		Id:        admin.Id,
		Name:      admin.Name,
		Avatar:    admin.Avatar,
		Email:     admin.Email,
		Phone:     admin.Phone,
		Account:   admin.Account,
		LoginIP:   admin.LoginIP,
		LoginTime: admin.LoginTime,
		Remark:    admin.Remark,
		Status:    admin.Status,
		Creator:   admin.Creator,
		Updater:   admin.Updater,
		CreatedAt: admin.CreatedAt,
		UpdatedAt: admin.UpdatedAt,
	}, nil
}

// 管理员分页列表
func (s *sSysAdmin) Page(ctx context.Context, params model.SysAdminPageReq) (*model.SysAdminPageRes, error) {

	paging := &db.Paging{
		Page:     params.Page,
		PageSize: params.PageSize,
	}

	filter := bson.M{}

	results, err := dao.SysAdmin.FindByPage(ctx, paging, filter, &dao.FindOptions{SortFields: []string{"-updated_at"}})
	if err != nil {
		logger.Error(ctx, err)
		return nil, err
	}

	items := make([]*model.SysAdmin, 0)
	for _, result := range results {
		items = append(items, &model.SysAdmin{
			Id:        result.Id,
			Name:      result.Name,
			Avatar:    result.Avatar,
			Email:     result.Email,
			Phone:     result.Phone,
			Account:   result.Account,
			LoginIP:   result.LoginIP,
			LoginTime: result.LoginTime,
			Remark:    result.Remark,
			Status:    result.Status,
			Creator:   result.Creator,
			Updater:   result.Updater,
			CreatedAt: result.CreatedAt,
			UpdatedAt: result.UpdatedAt,
		})
	}

	return &model.SysAdminPageRes{
		Items: items,
		Paging: &model.Paging{
			Page:     paging.Page,
			PageSize: paging.PageSize,
			Total:    paging.Total,
		},
	}, nil
}
