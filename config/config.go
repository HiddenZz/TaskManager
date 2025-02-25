package config

import (
	"fmt"
	"net/url"

	"github.com/ilyakaznacheev/cleanenv"
)

type (
	Config struct {
		DataBaseConfiguration
	}

	DataBaseConfiguration struct {
		Host     string `env:"HOST" env-default:"localhost"`
		Port     string `env:"PORT" .env-default:"5432"`
		Name     string `env:"DB" .env-default:"local"`
		User     string `env:"USER" .env-default:"local"`
		Password string `env:"PASSWORD" .env-default:"local"`
	}
)

func Create() *Config {
	var dbCfg DataBaseConfiguration
	if err := cleanenv.ReadConfig(".env", &dbCfg); err != nil {
		panic(fmt.Errorf("fatal error config file: %w", err))
	}
	return &Config{dbCfg}
}

func (db *DataBaseConfiguration) BuildConnectionString(additional map[string]string) string {

	params := url.Values{}
	for k, v := range additional {
		params.Add(k, v)
	}

	connectionString := fmt.Sprintf("postgres://%s:%s@%s:%s/%s", db.User, db.Password, db.Host, db.Port, db.Name)

	if len(params) > 0 {
		connectionString = "?" + params.Encode()
	}

	return connectionString
}
