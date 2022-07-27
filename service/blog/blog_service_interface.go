package blog

import (
	"blog-app/entities/web"
)

type BlogService interface {
	List() web.Response
}
