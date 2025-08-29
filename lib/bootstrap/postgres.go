package bootstrap

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
	"github.com/vrischmann/envconfig"
)

type psqlConfig struct {
	Host         string `envconfig:"POSTGRES_HOST"`
	Username     string `envconfig:"POSTGRES_USER"`
	Password     string `envconfig:"POSTGRES_PASSWORD"`
	DatabaseName string `envconfig:"default=postgres"`
}

func initPostgres() (*sql.DB, error) {
	cfg := psqlConfig{}

	if err := envconfig.Init(&cfg); err != nil {
		return nil, err
	}

	db, err := sql.Open("postgres", fmt.Sprintf("host=%s user=%s password=%s dbname=%s sslmode=disable",
		cfg.Host, cfg.Username, cfg.Password, cfg.DatabaseName))

	if err != nil {
		return nil, err
	}

	return db, nil
}
