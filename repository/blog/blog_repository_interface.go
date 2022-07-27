package blog

import "blog-app/entities/model"

type BlogRepository interface {
	FindAll() (blogs []*model.Blog, err error)
}
