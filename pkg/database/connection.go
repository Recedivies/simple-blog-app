package database

import (
	"blog-app/config"
	"fmt"

	"github.com/jinzhu/gorm"
	_ "github.com/lib/pq" // postgres database driver
	"github.com/rs/zerolog/log"
)

var DB *gorm.DB

func InitDatabase(config config.AppConfig) *gorm.DB {
	DBURL := fmt.Sprintf("host=%s port=%s user=%s dbname=%s sslmode=disable password=%s", config.DB.Host, config.DB.Port, config.DB.User, config.DB.Name, config.DB.Password)
	DB, err := gorm.Open("postgres", DBURL)
	if err != nil {
		error := fmt.Sprintf("Cannot connect to %s database\n", config.DB.Driver)
		log.Error().Err(err).Msg(error)
		return nil
	}
	log.Info().Msg("connected to database")

	return DB
}
