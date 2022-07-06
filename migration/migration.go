package migration

import (
	"github.com/itsapep/golang-api-with-gin/model"
	"gorm.io/gorm"
)

func DbMigration(db *gorm.DB) {
	db.AutoMigrate(
		&model.Product{},
	)
}
