package dao

import (
	"github.com/iimeta/fastapi-admin/internal/model/entity"
	"github.com/iimeta/fastapi-admin/utility/db"
)

var Text = NewTextDao()

type TextDao struct {
	*MongoDB[entity.Text]
}

func NewTextDao(database ...string) *TextDao {

	if len(database) == 0 {
		database = append(database, db.DefaultDatabase)
	}

	return &TextDao{
		MongoDB: NewMongoDB[entity.Text](database[0], TEXT),
	}
}
