package controllers

import (
	"blog-app/entities/model"
	"fmt"

	blogService "blog-app/service/blog"

	"github.com/gofiber/fiber/v2"
)

type BlogController struct {
	BlogService blogService.BlogService
}

func NewBlogController(blogService *blogService.BlogService) BlogController {
	return BlogController{BlogService: *blogService}
}

func (controller *BlogController) List(c *fiber.Ctx) error {
	responses := controller.BlogService.List()
	return c.Status(responses.StatusCode).JSON(responses)
}

var blogs = []*model.Blog{
	{
		Id:     1,
		Title:  "Test Blog 1",
		Author: "Test Author 1",
		Url:    "Test Url 1",
		Likes:  1,
	},
	{
		Id:     2,
		Title:  "Test Blog 2",
		Author: "Test Author 2",
		Url:    "Test Url 2",
		Likes:  2,
	},
}

func GetBlogs(c *fiber.Ctx) error {
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"success": "true",
		"data": fiber.Map{
			"blogs": blogs,
		},
	})
}

func CreateBlog(c *fiber.Ctx) error {
	type Request struct {
		Title  string `json:"title"`
		Author string `json:"author"`
		Url    string `json:"url"`
	}

	var body Request

	err := c.BodyParser(&body)

	if err != nil {
		fmt.Println(err.Error())
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"success": false,
			"message": "Cannot parse JSON",
		})
	}

	blog := &model.Blog{
		Id:     len(blogs) + 1,
		Title:  body.Title,
		Author: body.Author,
		Url:    body.Url,
		Likes:  0,
	}

	// append in blogs
	blogs = append(blogs, blog)

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"success": true,
		"data": fiber.Map{
			"blog": blog,
		},
	})
}
