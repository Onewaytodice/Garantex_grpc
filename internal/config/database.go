package config

import (
	"errors"
	"fmt"
	"github.com/golang-migrate/migrate/v4"
	"github.com/jmoiron/sqlx"
)

type DatabaseConfig struct {
	Name     string `env:"DB_NAME"`
	User     string `env:"DB_USER"`
	Password string `env:"DB_PASSWORD"`
	Host     string `env:"DB_HOST"`
	Port     string `env:"DB_PORT"`
}

func (d *DatabaseConfig) createDataSourceName() string {
	return fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable",
		d.User, d.Password, d.Host, d.Port, d.Name)
}

func (d *DatabaseConfig) UpMigrations(path string) error {
	m, err := migrate.New(path, d.createDataSourceName())
	if err != nil {
		return err
	}

	if err = m.Up(); err != nil && !errors.Is(err, migrate.ErrNoChange) {
		return err
	}
	return nil
}

func (d *DatabaseConfig) Connect() (*sqlx.DB, error) {
	db, err := sqlx.Connect("postgres", d.createDataSourceName())
	if err != nil {
		return nil, err
	}
	return db, nil
}
