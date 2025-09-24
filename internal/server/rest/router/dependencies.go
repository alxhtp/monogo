package router

import (
	"github.com/alxhtp/monogo/config"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

type Dependencies struct {
	App *fiber.App
	DB  *gorm.DB
	Cfg *config.Config
}

func NewDependencies(app *fiber.App, db *gorm.DB, cfg *config.Config) *Dependencies {
	return &Dependencies{
		App: app,
		DB:  db,
		Cfg: cfg,
	}
}
