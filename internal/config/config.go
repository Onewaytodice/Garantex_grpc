package config

import (
	"flag"
	"fmt"
	"github.com/ilyakaznacheev/cleanenv"
	"github.com/joho/godotenv"
	"os"
	"time"
)

type Config struct {
	Logger     LoggerConfig
	GRPCServer GRPCServerConfig
	HTTPServer HTTPServerConfig
	Database   DatabaseConfig
	Garantex   GarantexConfig
	Trace      TraceConfig
}

type LoggerConfig struct {
	Name       string `env:"APP_NAME"`
	Production bool   `env:"PRODUCTION" env-default:"true"`
}

type GRPCServerConfig struct {
	Host string `env:"GRPC_HOST"`
	Port string `env:"GRPC_PORT"`
}

type HTTPServerConfig struct {
	Host string `env:"HTTP_HOST"`
	Port string `env:"HTTP_PORT"`
}

func CreateAddr(host, port string) string {
	return fmt.Sprintf("%s:%s", host, port)
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
