package dao

import (
	"github.com/iimeta/fastapi-admin/internal/model/do"
	"github.com/iimeta/fastapi-admin/internal/model/entity"
	"github.com/iimeta/fastapi-admin/utility/db"
)

var Audio = NewAudioDao()

type AudioDao struct {
	*MongoDB[entity.Audio]
}

func NewAudioDao(database ...string) *AudioDao {

	if len(database) == 0 {
		database = append(database, db.DefaultDatabase)
	}

	return &AudioDao{
		MongoDB: NewMongoDB[entity.Audio](database[0], do.AUDIO_COLLECTION),
	}
}
