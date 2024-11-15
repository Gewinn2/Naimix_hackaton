package config

import "os"

type Application struct {
	ProductionType string
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
			ProductionType: os.Getenv("PRODUCTION_TYPE"),
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
