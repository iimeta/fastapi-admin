package dao

import (
	"github.com/iimeta/fastapi-admin/v2/internal/model/entity"
	"github.com/iimeta/fastapi-admin/v2/utility/db"
)

var SiteConfig = NewSiteConfigDao()

type SiteConfigDao struct {
	*MongoDB[entity.SiteConfig]
}

func NewSiteConfigDao(database ...string) *SiteConfigDao {

	if len(database) == 0 {
		database = append(database, db.DefaultDatabase)
	}

	return &SiteConfigDao{
		MongoDB: NewMongoDB[entity.SiteConfig](database[0], SITE_CONFIG),
	}
}
