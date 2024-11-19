package config

import (
	"log"
	"os"
	"time"
)

type Config struct {
	Env string
	DB
	HTTPServer
}

type HTTPServer struct {
	ApiAddres      string
	ApiTimeout     time.Duration
	ApiIdleTimeout time.Duration
}

type DB struct {
	PostgresHost   string
	PostgresDBName string
	PostgresPort   string
	PostgresUser   string
	PostgresPass   string
}

func MustLoad() *Config {
	var cfg Config
	var err error
	cfg.Env = os.Getenv("ENVIRONMENT")
	cfg.PostgresHost = os.Getenv("POSTGRES_HOST")
	cfg.PostgresPort = os.Getenv("POSTGRES_PORT")
	cfg.PostgresUser = os.Getenv("POSTGRES_USER")
	cfg.PostgresPass = os.Getenv("POSTGRES_PASSWORD")
	cfg.PostgresDBName = os.Getenv("POSTGRES_DB")
	cfg.ApiAddres = os.Getenv("API_HOST")
	cfg.ApiTimeout, err = time.ParseDuration(os.Getenv("API_TIMEOUT"))
	if err != nil {
		log.Fatal("Cannot parse API_TIMEOUT env var")
	}
	cfg.ApiIdleTimeout, err = time.ParseDuration(os.Getenv("API_IDLE_TIMEOUT"))
	if err != nil {
		log.Fatal("Cannot parse API_TIMEOUT env var")
	}
	if cfg.Env == "" {
		cfg.Env = "local"
	}
	if cfg.PostgresHost == "" {
		cfg.PostgresHost = "postgres"
	}
	if cfg.PostgresPort == "" {
		cfg.PostgresPort = "5432"
	}
	if cfg.PostgresUser == "" {
		cfg.PostgresUser = "postgres"
	}
	if cfg.PostgresPass == "" {
		cfg.PostgresPass = "postgres"
	}
	if cfg.PostgresDBName == "" {
		cfg.PostgresDBName = "postgres"
	}
	if cfg.ApiAddres == "" {
		cfg.ApiAddres = "0.0.0.0:8080"
	}
	if cfg.ApiTimeout == 0 {
		cfg.ApiTimeout = time.Second * 5
	}
	if cfg.ApiIdleTimeout == 0 {
		cfg.ApiIdleTimeout = time.Second * 60
	}

	return &cfg
}
