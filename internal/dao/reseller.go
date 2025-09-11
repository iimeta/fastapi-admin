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

var Reseller = NewResellerDao()

type ResellerDao struct {
	*MongoDB[entity.Reseller]
}

func NewResellerDao(database ...string) *ResellerDao {

	if len(database) == 0 {
		database = append(database, db.DefaultDatabase)
	}

	return &ResellerDao{
		MongoDB: NewMongoDB[entity.Reseller](database[0], RESELLER),
	}
}

var ResellerAccount = NewResellerAccountDao()

type ResellerAccountDao struct {
	*MongoDB[entity.ResellerAccount]
}

func NewResellerAccountDao(database ...string) *ResellerAccountDao {

	if len(database) == 0 {
		database = append(database, db.DefaultDatabase)
	}

	return &ResellerAccountDao{
		MongoDB: NewMongoDB[entity.ResellerAccount](database[0], RESELLER_ACCOUNT),
	}
}

// 判断账号是否存在
func (d *ResellerDao) IsAccountExist(ctx context.Context, account string) bool {

	total, err := ResellerAccount.CountDocuments(ctx, bson.M{"account": account})
	if err != nil {
		logger.Error(ctx, err)
		return false
	}

	return total > 0
}

// 根据账号查询代理商
func (d *ResellerDao) FindResellerByAccount(ctx context.Context, account string) (*entity.Reseller, error) {

	accountInfo, err := ResellerAccount.FindOne(ctx, bson.M{"account": account})
	if err != nil {
		return nil, err
	}

	return d.FindById(ctx, accountInfo.Uid)
}

// 根据userId查询代理商
func (d *ResellerDao) FindResellerByUserId(ctx context.Context, userId int) (*entity.Reseller, error) {
	return d.FindOne(ctx, bson.M{"user_id": userId})
}

// 根据userIds查询代理商列表
func (d *ResellerDao) FindResellerListByUserIds(ctx context.Context, userIds []int) ([]*entity.Reseller, error) {
	return d.Find(ctx, bson.M{"user_id": bson.M{"$in": userIds}})
}

// 判断手机号是否存在
func (d *ResellerDao) IsPhoneExist(ctx context.Context, phone string) bool {

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
func (d *ResellerDao) IsEmailExist(ctx context.Context, email string) bool {

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

func (d *ResellerDao) CreateAccount(ctx context.Context, account *do.ResellerAccount) (string, error) {
	return ResellerAccount.Insert(ctx, account)
}

func (d *ResellerDao) FindAccount(ctx context.Context, account string) (*entity.ResellerAccount, error) {

	accountInfo, err := ResellerAccount.FindOne(ctx, bson.M{"account": account})
	if err != nil {
		return nil, err
	}

	return accountInfo, nil
}

func (d *ResellerDao) FindAccountByUserId(ctx context.Context, userId int) (*entity.ResellerAccount, error) {

	accountInfo, err := ResellerAccount.FindOne(ctx, bson.M{"user_id": userId})
	if err != nil {
		return nil, err
	}

	return accountInfo, nil
}

func (d *ResellerDao) FindAccountsByUserId(ctx context.Context, userId int) ([]*entity.ResellerAccount, error) {

	accounts, err := ResellerAccount.Find(ctx, bson.M{"user_id": userId})
	if err != nil {
		return nil, err
	}

	return accounts, nil
}

func (d *ResellerDao) ChangeAccountById(ctx context.Context, id, account string) error {
	return ResellerAccount.UpdateById(ctx, id, bson.M{"account": account})
}

func (d *ResellerDao) ChangePasswordByUserId(ctx context.Context, userId int, password string) error {

	salt := grand.Letters(8)
	if err := ResellerAccount.UpdateMany(ctx, bson.M{"user_id": userId}, bson.M{
		"password": crypto.EncryptPassword(password + salt),
		"salt":     salt,
	}); err != nil {
		logger.Error(ctx, err)
		return err
	}

	return nil
}
