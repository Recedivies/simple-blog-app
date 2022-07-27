package blog

import (
	"blog-app/entities/model"

	"github.com/jinzhu/gorm"
)

type blogRepository struct {
	db *gorm.DB
}

func (b blogRepository) FindAll() (blogs []*model.Blog, err error) {
	err = b.db.Find(blogs).Error
	return blogs, err
}

func NewBlogRepository(db *gorm.DB) BlogRepository {
	return &blogRepository{db: db}
}
