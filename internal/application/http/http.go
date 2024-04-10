package http

import (
	"log/slog"

	"tp_back/internal/config"
	userHandler "tp_back/internal/http/handler/user"
	mwLogger "tp_back/internal/http/middleware/logger"
	userPRepository "tp_back/internal/repository/postgres/user"
	"tp_back/internal/resource"
	hashService "tp_back/internal/service/hash"
	userService "tp_back/internal/service/user"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

type App struct {
	Router   *chi.Mux
	Resource *resource.Res
}

func New(cfg *config.Cfg, res *resource.Res, log *slog.Logger) *App {
	log.Info("set upping http app")

	router := chi.NewRouter()

	setupMiddleware(router, log)
	setupEndpoints(cfg, router, res, log)

	app := &App{Router: router, Resource: res}

	log.Info("set upped http app", slog.String("env", cfg.Env))

	return app
}

func setupMiddleware(router *chi.Mux, log *slog.Logger) {
	router.Use(middleware.RequestID)
	router.Use(middleware.Logger)
	router.Use(mwLogger.New(log))
	router.Use(middleware.Recoverer)
	router.Use(middleware.URLFormat)
}

func setupEndpoints(cfg *config.Cfg, router *chi.Mux, res *resource.Res, log *slog.Logger) {
	userPR := userPRepository.New(log, res.Psql)

	hashS := hashService.New(log, cfg.Jwt)
	userS := userService.New(log, userPR, hashS)

	userH := userHandler.New(log, userS)
	//projectH := projectHandler.New(log)
	//taskH := taskHandler.New(log)
	//roleH := roleHandler.New(log)
	//permissionH := permissionHandler.New(log)

	router.Route("/api", func(r chi.Router) {
		r.Post("/users/register", userH.Register())
		r.Post("/users/login", userH.Login())

		//r.Get("/projects", projectH.Index())
		//r.Post("/projects", projectH.Store())
		//r.Get("/projects/{uuid}", projectH.Show())
		//r.Put("/projects/{uuid}", projectH.Update())
		//r.Delete("/projects/{uuid}", projectH.Delete())
		//
		//r.Get("/tasks", taskH.Index())
		//r.Post("/tasks", taskH.Store())
		//r.Get("/tasks/{uuid}", taskH.Show())
		//r.Put("/tasks/{uuid}", taskH.Update())
		//r.Delete("/tasks/{uuid}", taskH.Delete())
		//
		//r.Get("/roles", roleH.Index())
		//r.Post("/roles", roleH.Store())
		//
		//r.Get("/permissions", permissionH.Index())
		//r.Post("/permissions", permissionH.Store())
	})
}
