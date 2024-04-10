package config

import (
	"log"
	"os"
	"time"

	"github.com/ilyakaznacheev/cleanenv"
)

type Cfg struct {
	Env      string `yaml:"env" env-required:"true"`
	Http     `yaml:"http"`
	Postgres `yaml:"postgres"`
	Jwt      `yaml:"jwt"`
}

type Http struct {
	Address     string        `yaml:"address" env-required:"true"`
	Timeout     time.Duration `yaml:"timeout" env-default:"4s"`
	IdleTimeout time.Duration `yaml:"idle_timeout" env-default:"60s"`
}

type Postgres struct {
	Url            string `yaml:"url" env-required:"true"`
	MigrationsPath string `yaml:"migrations_path" env-required:"true"`
}

type Jwt struct {
	SecretKey string `yaml:"secret_key" env-required:"true"`
}

func MustLoad() *Cfg {
	log.Println("starting to get CONFIG_PATH variable")
	path := mustFetchPathEnv()

	return MustLoadPath(path)
}

func MustLoadPath(path string) *Cfg {
	if _, err := os.Stat(path); os.IsNotExist(err) {
		panic("config file does not exist: " + path)
	}

	var config Cfg
	if err := cleanenv.ReadConfig(path, &config); err != nil {
		panic("failed to read config: " + err.Error())
	}

	return &config
}

func mustFetchPathEnv() string {
	path := os.Getenv("CONFIG_PATH")
	if path == "" {
		panic("CONFIG_PATH variable is not set")
	}

	return path
}
