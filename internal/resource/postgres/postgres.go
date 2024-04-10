package postgres

import (
	"database/sql"
	"log/slog"

	_ "github.com/lib/pq"
)

type Res struct {
	DB *sql.DB
}

func MustNew(dbUrl string, log *slog.Logger) *Res {
	log.Info("connecting to postgres")

	db, err := sql.Open("postgres", dbUrl)
	if err != nil {
		panic("failed resource to postgres: " + err.Error())
	}

	err = db.Ping()
	if err != nil {
		panic("failed ping to postgres: " + err.Error())
	}

	log.Info("connected to postgres", slog.String("url", dbUrl))

	return &Res{DB: db}
}
