package routes

import (
	blogController "blog-app/controllers"
	blogRepo "blog-app/repository/blog"
	blogService "blog-app/service/blog"

	"github.com/gofiber/fiber/v2"
	"github.com/jinzhu/gorm"
)

func RegisterBlogRoutes(app *fiber.App, db *gorm.DB) {
	blogRepo := blogRepo.NewBlogRepository(db)
	blogService := blogService.NewBlogService(&blogRepo)
	blogController := blogController.NewBlogController(&blogService)

	blogGroup := app.Group("/api/blogs")

	// connect blogs routes
	blogGroup.Get("", blogController.List)
}
