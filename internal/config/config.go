// Configuration package
// project GophKeeper Yandex Practicum
// Created by Vasiliy Van-Zaam
package config

import (
	"log"

	"github.com/caarlos0/env/v7"
)

type Logger interface {
	Info(args ...any)
	Error(args ...any)
	Warn(args ...any)
	Debug(args ...any)
	Fatal(args ...any)
}

type Config interface {
	Client() ClientConfig
	Logger() Logger
}

type ClientConfig interface {
	FilePath() string
}

type configs struct {
	client ClientConfig
	logger Logger
}

// Logger implements Config.
func (c *configs) Logger() Logger {
	return c.logger
}

// Client implements Config.
func (c *configs) Client() ClientConfig {
	return c.client
}

type clientConfig struct {
	LocalstorePath string `env:"localstore_path" envDefault:"datastore"`
}

// Get localstore file path.
func (c *clientConfig) FilePath() string {
	return c.LocalstorePath
}

func newClientConfig() ClientConfig {
	cfg := clientConfig{}
	if err := env.Parse(&cfg); err != nil {
		log.Printf("%+v\n", err)
	}
	return &cfg
}

// Create new config.
func New(logger Logger) Config {
	return &configs{
		client: newClientConfig(),
		logger: logger,
	}
}
