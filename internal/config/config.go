package config

import "os"

type Application struct {
	ProductionType  string
	SigningKey      string
	TokenExpiration string
}

type Db struct {
	Host     string
	Port     string
	User     string
	Password string
	Database string
}

type Config struct {
	Application Application
	Db          Db
}

func NewConfig() *Config {
	return &Config{
		Application: Application{
			ProductionType:  os.Getenv("PRODUCTION_TYPE"),
			SigningKey:      os.Getenv("SIGNING_KEY"),
			TokenExpiration: os.Getenv("TOKEN_EXPIRATION"),
		},
		Db: Db{
			Host:     os.Getenv("POSTGRESQL_HOST"),
			Port:     os.Getenv("POSTGRESQL_PORT"),
			User:     os.Getenv("POSTGRESQL_USER"),
			Password: os.Getenv("POSTGRESQL_PASSWORD"),
			Database: os.Getenv("POSTGRESQL_DB"),
		},
	}
}
