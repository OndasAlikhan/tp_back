package main

import (
	"errors"
	"flag"
	"fmt"

	"tp_back/internal/config"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

func main() {
	var path, url string
	flag.StringVar(&path, "path", "", "Migrations path")
	flag.StringVar(&url, "url", "", "Postgres url")
	var down bool
	flag.BoolVar(&down, "down", false, "Down or up")
	flag.Parse()

	if path == "" || url == "" {
		cfg := config.MustLoad()
		if path == "" {
			path = cfg.Postgres.MigrationsPath
		}
		if url == "" {
			url = cfg.Postgres.Url
		}
	}

	m, err := migrate.New("file://"+path, url)
	if err != nil {
		panic("failed when connecting to postgres: " + err.Error())
	}

	if !down {
		if err := m.Up(); err != nil {
			if errors.Is(err, migrate.ErrNoChange) {
				fmt.Println("no migrations to apply")

				return
			}

			panic("failed when applying migrations: " + err.Error())
		}

		fmt.Println("migrations applied")
	} else {
		if err := m.Down(); err != nil {
			if errors.Is(err, migrate.ErrNoChange) {
				fmt.Println("no migrations to revert")

				return
			}

			panic("failed when reverting migrations: " + err.Error())
		}

		fmt.Println("migrations reverted")
	}
}
