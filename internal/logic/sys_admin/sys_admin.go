package sys_admin

import (
	"context"
	"github.com/gogf/gf/v2/util/grand"
	"github.com/iimeta/fastapi-admin/internal/dao"
	"github.com/iimeta/fastapi-admin/internal/model"
	"github.com/iimeta/fastapi-admin/internal/model/do"
	"github.com/iimeta/fastapi-admin/internal/service"
	"github.com/iimeta/fastapi-admin/utility/crypto"
	"github.com/iimeta/fastapi-admin/utility/db"
	"github.com/iimeta/fastapi-admin/utility/logger"
	"go.mongodb.org/mongo-driver/bson"
)

type sSysAdmin struct{}

func init() {
	service.RegisterSysAdmin(New())
}

func New() service.ISysAdmin {
	return &sSysAdmin{}
}

// 新建管理员
func (s *sSysAdmin) Create(ctx context.Context, params model.SysAdminCreateReq) error {

	salt := grand.Letters(8)

	if _, err := dao.SysAdmin.Insert(ctx, &do.SysAdmin{
		Name:     params.Name,
		Avatar:   params.Avatar,
		Gender:   params.Gender,
		Phone:    params.Phone,
		Email:    params.Email,
		Account:  params.Account,
		Password: crypto.EncryptPassword(params.Password + salt),
		Salt:     salt,
		Remark:   params.Remark,
		Status:   params.Status,
	}); err != nil {
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
		Gender: params.Gender,
		Phone:  params.Phone,
		Email:  params.Email,
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
		Gender:    admin.Gender,
		Phone:     admin.Phone,
		Email:     admin.Email,
		Account:   admin.Account,
		Password:  admin.Password,
		Salt:      admin.Salt,
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

	results, err := dao.SysAdmin.FindByPage(ctx, paging, filter, "-updated_at")
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
			Gender:    result.Gender,
			Phone:     result.Phone,
			Email:     result.Email,
			Account:   result.Account,
			Password:  result.Password,
			Salt:      result.Salt,
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
