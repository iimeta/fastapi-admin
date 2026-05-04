package dao

import (
	"github.com/iimeta/fastapi-admin/v2/internal/model/entity"
	"github.com/iimeta/fastapi-admin/v2/utility/db"
)

var TicketReply = NewTicketReplyDao()

type TicketReplyDao struct {
	*MongoDB[entity.TicketReply]
}

func NewTicketReplyDao(database ...string) *TicketReplyDao {

	if len(database) == 0 {
		database = append(database, db.DefaultDatabase)
	}

	return &TicketReplyDao{
		MongoDB: NewMongoDB[entity.TicketReply](database[0], TICKET_REPLY),
	}
}
