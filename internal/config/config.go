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
	Server() ServerConfig
}

type ClientConfig interface {
	FilePath() string
}

type ServerConfig interface {
	SecretKey(version string) string
	RunAddrss() string
	Expires(name ...bool) int
}
type serverConfig struct {
	SecretKeys        map[string]string `env:"server_secret_keys" envDefault:"0.0.0:secret_key_version_0.0.0,0.0.1:secret_key_version_0.0.1"`
	RunAddress        string            `env:"server_run_address" envDefault:":3200"`
	AccessExpiresMin  int               `env:"access_expires_min" envDefault:"10"`
	RefreshExpiresMin int               `env:"refresh_expires_min" envDefault:"43200"`
}

// Expires implements ServerConfig.
func (c *serverConfig) Expires(isRefresh ...bool) int {
	if len(isRefresh) > 0 && isRefresh[0] {
		return c.RefreshExpiresMin
	}
	return c.AccessExpiresMin
}

// RunAddrss implements ServerConfig.
func (c *serverConfig) RunAddrss() string {
	return c.RunAddress
}

// GetSecretKey implements ServerConfig.
func (c *serverConfig) SecretKey(version string) string {
	return c.SecretKeys[version]
}

type configs struct {
	client ClientConfig
	logger Logger
	server ServerConfig
}

// Server implements Config.
func (c *configs) Server() ServerConfig {
	return c.server
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
	LocalstorePath string `env:"client_localstore_path" envDefault:"datastore"`
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
func newServerConfig() ServerConfig {
	cfg := serverConfig{}
	if err := env.Parse(&cfg); err != nil {
		log.Printf("%+v\n", err)
	}
	return &cfg
}

// Create new config.
func New(logger Logger) Config {
	conf := &configs{
		client: newClientConfig(),
		logger: logger,
		server: newServerConfig(),
	}
	return conf
}
