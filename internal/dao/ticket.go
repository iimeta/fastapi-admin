package dao

import (
	"github.com/iimeta/fastapi-admin/v2/internal/model/entity"
	"github.com/iimeta/fastapi-admin/v2/utility/db"
)

var Ticket = NewTicketDao()

type TicketDao struct {
	*MongoDB[entity.Ticket]
}

func NewTicketDao(database ...string) *TicketDao {

	if len(database) == 0 {
		database = append(database, db.DefaultDatabase)
	}

	return &TicketDao{
		MongoDB: NewMongoDB[entity.Ticket](database[0], TICKET),
	}
}
