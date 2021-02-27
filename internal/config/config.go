package config

import (
	"strings"
	"sync"

	"github.com/jinzhu/configor"
)

type Config struct {
	DB        DB
	Service   Service
	LogLevel  string `default:"INFO" env:"LOG_LEVEL"`
	LogFormat string `default:"text" env:"LOG_FORMAT"`
}

type DB struct {
	Client    string `default:"mysql" env:"DB_CLIENT"`
	Host      string `default:"0.0.0.0" env:"DB_HOST"`
	User      string `default:"root" env:"DB_USER"`
	Password  string `required:"true" env:"DB_PASSWORD"`
	Port      uint   `default:"3306" env:"DB_PORT"`
	Database  string `default:"goboiler" env:"DB_DATABASE"`
	Migration struct {
		Autoload bool   `env:"DB_RUN_MIGRATION"`
		Path     string `default:"./database/migration" env:"DB_MIGRATION_PATH"`
	}
	MaxIdleConnections int    `default:"2" env:"DB_MAX_IDLE_CONN"`
	MaxOpenConnections int    `default:"0" env:"DB_MAX_OPEN_CONN"`
	MaxConnLifeTime    int    `default:"90" env:"DB_MAX_CONN_LIFETIME"`
	Debug              bool   `default:"false" env:"DB_DEBUG"`
	TxConnKey          string `default:"txConn" env:"TX_CONN_KEY"`
}

type Service struct {
	Name   string `default:"go-boiler" env:"SERVICE_NAME"`
	Scheme string `default:"http" env:"SERVICE_SCHEME"`
	Host   string `default:"0.0.0.0" env:"SERVICE_HOST"`
	Port   string `default:"9000" env:"SERVICE_PORT"`
}

var config *Config
var configLock = &sync.Mutex{}

func Instance() Config {
	if config == nil {
		err := Load()
		if err != nil {
			panic(err)
		}
	}
	return *config
}

func Load() error {
	tmpConfig := Config{}
	err := configor.Load(&tmpConfig)
	if err != nil {
		return err
	}

	configLock.Lock()
	defer configLock.Unlock()
	config = &tmpConfig

	return nil
}

func (Config) String() string {
	sb := strings.Builder{}
	return sb.String()
}
