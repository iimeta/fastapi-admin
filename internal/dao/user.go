package dao

import (
	"context"
	"github.com/gogf/gf/v2/util/grand"
	"github.com/iimeta/fastapi-admin/internal/model/do"
	"github.com/iimeta/fastapi-admin/internal/model/entity"
	"github.com/iimeta/fastapi-admin/utility/crypto"
	"github.com/iimeta/fastapi-admin/utility/db"
	"github.com/iimeta/fastapi-admin/utility/logger"
	"go.mongodb.org/mongo-driver/bson"
)

var User = NewUserDao()

type UserDao struct {
	*MongoDB[entity.User]
}

func NewUserDao(database ...string) *UserDao {

	if len(database) == 0 {
		database = append(database, db.DefaultDatabase)
	}

	return &UserDao{
		MongoDB: NewMongoDB[entity.User](database[0], do.USER_COLLECTION),
	}
}

// 判断账号是否存在
func (d *UserDao) IsAccountExist(ctx context.Context, account string) bool {

	total, err := Account.CountDocuments(ctx, bson.M{"account": account})
	if err != nil {
		logger.Error(ctx, err)
		return false
	}

	return total > 0
}

// 根据账号查询用户
func (d *UserDao) FindUserByAccount(ctx context.Context, account string) (*entity.User, error) {

	accountInfo, err := Account.FindOne(ctx, bson.M{"account": account})
	if err != nil {
		return nil, err
	}

	return d.FindById(ctx, accountInfo.Uid)
}

// 根据userId查询用户
func (d *UserDao) FindUserByUserId(ctx context.Context, userId int) (*entity.User, error) {
	return d.FindOne(ctx, bson.M{"user_id": userId})
}

// 根据userIds查询用户列表
func (d *UserDao) FindUserListByUserIds(ctx context.Context, userIds []int) ([]*entity.User, error) {
	return d.Find(ctx, bson.M{"user_id": bson.M{"$in": userIds}})
}

// 判断手机号是否存在
func (d *UserDao) IsPhoneExist(ctx context.Context, phone string) bool {

	if len(phone) == 0 {
		return false
	}

	total, err := d.CountDocuments(ctx, bson.M{"phone": phone})
	if err != nil {
		logger.Error(ctx, err)
		return false
	}

	return total > 0
}

// 判断邮箱是否存在
func (d *UserDao) IsEmailExist(ctx context.Context, email string) bool {

	if len(email) == 0 {
		return false
	}

	total, err := d.CountDocuments(ctx, bson.M{"email": email})
	if err != nil {
		logger.Error(ctx, err)
		return false
	}

	return total > 0
}

func (d *UserDao) CreateAccount(ctx context.Context, account *do.Account) (string, error) {
	return Account.Insert(ctx, account)
}

func (d *UserDao) FindAccount(ctx context.Context, account string) (*entity.Account, error) {

	accountInfo, err := Account.FindOne(ctx, bson.M{"account": account})
	if err != nil {
		return nil, err
	}

	return accountInfo, nil
}

func (d *UserDao) FindAccountByUserId(ctx context.Context, userId int) (*entity.Account, error) {

	accountInfo, err := Account.FindOne(ctx, bson.M{"user_id": userId})
	if err != nil {
		return nil, err
	}

	return accountInfo, nil
}

func (d *UserDao) FindAccountsByUserId(ctx context.Context, userId int) ([]*entity.Account, error) {

	accounts, err := Account.Find(ctx, bson.M{"user_id": userId})
	if err != nil {
		return nil, err
	}

	return accounts, nil
}

func (d *UserDao) ChangeAccountById(ctx context.Context, id, account string) error {
	return Account.UpdateById(ctx, id, bson.M{"account": account})
}

func (d *UserDao) ChangePasswordByUserId(ctx context.Context, userId int, password string) error {

	salt := grand.Letters(8)
	if err := Account.UpdateMany(ctx, bson.M{"user_id": userId}, bson.M{
		"password": crypto.EncryptPassword(password + salt),
		"salt":     salt,
	}); err != nil {
		logger.Error(ctx, err)
		return err
	}

	return nil
}
