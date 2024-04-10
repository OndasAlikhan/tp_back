package main

import (
	"context"
	"log/slog"
	"net/http"
	"os"
	"os/exec"
	"os/signal"
	"syscall"
	"time"

	"tp_back/internal/application"
	"tp_back/internal/config"
	"tp_back/internal/lib/sl"
	"tp_back/internal/resource"
)

func main() {
	cfg := config.MustLoad()

	log := sl.Setup(cfg.Env)

	log.Debug("debug message enabled")

	mustRunMigration(cfg, log)

	res := resource.MustNew(cfg, log)

	defer res.Psql.DB.Close()

	app := application.New(cfg, res, log)

	done := make(chan os.Signal, 1)
	signal.Notify(done, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)

	log.Info("starting http server", slog.String("address", cfg.Http.Address))

	httpSrv := &http.Server{
		Addr:         cfg.Address,
		Handler:      app.HttpApp.Router,
		ReadTimeout:  cfg.Http.Timeout,
		WriteTimeout: cfg.Http.Timeout,
		IdleTimeout:  cfg.Http.IdleTimeout,
	}

	go func() {
		if err := httpSrv.ListenAndServe(); err != nil {
			log.Error("failed to start http server", sl.Error(err))
		}
	}()

	log.Info("started http server")

	<-done
	log.Info("stopping http server")

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if err := httpSrv.Shutdown(ctx); err != nil {
		log.Error("failed to stop server", sl.Error(err))

		return
	}

	log.Info("stopped http server")
}

func mustRunMigration(cfg *config.Cfg, log *slog.Logger) {
	log.Info("running migration", slog.String("path", cfg.Postgres.MigrationsPath), slog.String("url", cfg.Postgres.Url))

	cmd := exec.Command("go", "run", "./cmd/migrator/main.go", "--path", cfg.Postgres.MigrationsPath, "--url", cfg.Postgres.Url)
	output, err := cmd.CombinedOutput()

	log.Info("migration output", slog.String("output", string(output)))
	if err != nil {
		panic("migration failed: " + err.Error())
	}

	log.Info("ran migration")
}
