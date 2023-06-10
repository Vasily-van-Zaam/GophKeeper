// Configuration package
// project GophKeeper Yandex Practicum
// Created by Vasiliy Van-Zaam
package config

import (
	"log"

	"github.com/caarlos0/env/v7"
)

type Encryptor interface {
	Encrypt(secret []byte, userData []byte) ([]byte, error)
	Decrypt(secret []byte, data []byte) ([]byte, error)
}

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
	Encryptor() Encryptor
}

type ClientConfig interface {
	FilePath() string
}

type ServerConfig interface {
	SecretKey(version string) string
	RunAddrss() string
	// isRefresh == true  then the refresh  expiration date is returned
	// if true, then the access expiration date is returned
	Expires(isRefresh ...bool) int
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
	client  ClientConfig
	logger  Logger
	server  ServerConfig
	cryptor Encryptor
}

// Encryptor implements Config.
func (c *configs) Encryptor() Encryptor {
	return c.cryptor
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

// newServerConfig("0.0.0", "privite_tserver_token") - for client
// newServerConfig()  - for server from env data.
func newServerConfig(vToken ...string) ServerConfig {
	cfg := serverConfig{}
	if err := env.Parse(&cfg); err != nil {
		log.Printf("%+v\n", err)
	}
	if len(vToken) == 2 {
		cfg.SecretKeys[vToken[0]] = vToken[1]
	}

	return &cfg
}

// Create new config.
// New(logger, crypt) - for server
// New(logger, crypt, "0.0.0", "sprivate_server_token") - for client.
func New(logger Logger, crypt Encryptor, vToken ...string) Config {
	conf := &configs{
		client:  newClientConfig(),
		logger:  logger,
		server:  newServerConfig(vToken...),
		cryptor: crypt,
	}
	return conf
}
