package utils

import (
	"blog-app/config"
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/rs/zerolog/log"
)

func fiberConnURL() string {
	PORT := config.Config.App.Port
	if PORT == "" {
		PORT = "8080"
	}

	return fmt.Sprintf("0.0.0.0:%s", PORT)
}

func StartServer(a *fiber.App) {
	fiberConnURL := fiberConnURL()

	if err := a.Listen(fiberConnURL); err != nil {
		log.Error().Err(err).Msg("Server is not running")
	}
}
