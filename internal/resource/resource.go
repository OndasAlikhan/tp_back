package resource

import (
	"log/slog"

	"tp_back/internal/config"
	"tp_back/internal/resource/postgres"
)

type Res struct {
	Psql *postgres.Res
}

func MustNew(cfg *config.Cfg, log *slog.Logger) *Res {
	log.Info("connecting to resources")

	res := &Res{Psql: postgres.MustNew(cfg.Postgres.Url, log)}

	log.Info("connected to resources")

	return res
}
