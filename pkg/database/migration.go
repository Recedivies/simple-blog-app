package database

import (
	"blog-app/entities/model"

	"github.com/jinzhu/gorm"
)

func AutoMigrate(db *gorm.DB) {
	db.AutoMigrate(&model.Blog{})
}
