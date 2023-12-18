package dao

import (
	"github.com/iimeta/fastapi-admin/internal/model/do"
	"github.com/iimeta/fastapi-admin/internal/model/entity"
	"github.com/iimeta/fastapi-admin/utility/db"
)

var Key = NewKeyDao()

type KeyDao struct {
	*MongoDB[entity.Key]
}

func NewKeyDao(database ...string) *KeyDao {

	if len(database) == 0 {
		database = append(database, db.DefaultDatabase)
	}

	return &KeyDao{
		MongoDB: NewMongoDB[entity.Key](database[0], do.KEY_COLLECTION),
	}
}
