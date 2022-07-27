package blog

import (
	"blog-app/entities/web"
	blogRepo "blog-app/repository/blog"

	"github.com/gofiber/fiber/v2"
)

type blogService struct {
	blogRepository blogRepo.BlogRepository
}

func (b blogService) List() web.Response {
	blogs, err := b.blogRepository.FindAll()

	if err != nil {
		return web.Response{
			StatusCode: fiber.StatusBadRequest,
			Message:    "Bad Request",
			Data:       nil,
			Error:      "Bad Request",
		}
	}

	return web.Response{
		StatusCode: fiber.StatusOK,
		Message:    "Get All Blogs successful",
		Data:       blogs,
	}
}

func NewBlogService(blogRepository *blogRepo.BlogRepository) BlogService {
	return &blogService{blogRepository: *blogRepository}
}
