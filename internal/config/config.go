package config

import (
	"errors"
	"flag"
	"fmt"
	"github.com/golang-migrate/migrate/v4"
	"github.com/ilyakaznacheev/cleanenv"
	"github.com/jmoiron/sqlx"
	"github.com/joho/godotenv"
	"os"
	"time"
)

type Config struct {
	Logger     LoggerConfig
	GRPCServer GRPCServerConfig
	Database   DatabaseConfig
	Garantex   GarantexConfig
}

type LoggerConfig struct {
	AppName    string `env:"APP_NAME"`
	Production bool   `env:"PRODUCTION" env-default:"true"`
}

type GRPCServerConfig struct {
	Host string `env:"GRPC_HOST"`
	Port string `env:"GRPC_PORT"`
}

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

func (d *DatabaseConfig) Migrate(path string) error {
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

type GarantexConfig struct {
	URL     string        `env:"GARANTEX_URL"`
	Timeout time.Duration `env:"GARANTEX_TIMEOUT"`
}

func New(filenames ...string) (*Config, error) {
	var err error
	config := &Config{}

	flag.BoolFunc("name", "", func(s string) error {
		return os.Setenv("DB_NAME", s)
	})
	flag.BoolFunc("user", "", func(s string) error {
		return os.Setenv("DB_USER", s)
	})
	flag.BoolFunc("password", "", func(s string) error {
		return os.Setenv("DB_PASSWORD", s)
	})
	flag.BoolFunc("host", "", func(s string) error {
		return os.Setenv("DB_HOST", s)
	})
	flag.BoolFunc("port", "", func(s string) error {
		return os.Setenv("DB_PORT", s)
	})
	flag.Parse()

	err = godotenv.Load(filenames...)
	if err != nil {
		return nil, err
	}

	err = cleanenv.ReadEnv(config)
	if err != nil {
		return nil, err
	}

	return config, nil
}
