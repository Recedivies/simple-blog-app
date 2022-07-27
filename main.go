package main

import (
	"blog-app/config"
	"blog-app/pkg/database"
	"blog-app/pkg/middleware"
	"blog-app/pkg/utils"
	"blog-app/routes"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

func init() {
	// Setup Config (Database)
	config.InitConfig()
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix

	log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr})

	database.InitDatabase(config.Config)
}

func main() {
	// Setup Fiber
	app := fiber.New()

	// Setup Middleware
	middleware.RegisterMiddleware(app)

	// Setup Routes
	routes.RegisterBlogRoutes(app, database.DB)

	// Run Server
	utils.StartServer(app)
}
