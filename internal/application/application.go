package application

import (
	"log/slog"

	"tp_back/internal/application/http"
	"tp_back/internal/config"
	"tp_back/internal/resource"
)

type App struct {
	HttpApp *http.App
}

func New(cfg *config.Cfg, res *resource.Res, log *slog.Logger) *App {
	log.Info(" applications")

	app := &App{HttpApp: http.New(cfg, res, log)}

	log.Info("started applications")

	return app
}
