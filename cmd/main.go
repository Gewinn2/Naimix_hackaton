package main

import (
	"fmt"
	"log"
	_ "server/docs"
	"server/internal/config"
	"server/internal/handler"
	logger "server/internal/log"
	"server/internal/repository/postgres"
	"server/internal/service/domain"

	"github.com/joho/godotenv"
)

// @title Harmony Compass API
// @version 1.0
// @BasePath /
func main() {
	if err := godotenv.Load("./.env"); err != nil {
		log.Fatal("could not load .env configuration: " + err.Error())
	}
	conf := config.NewConfig()
	log := logger.InitLogger(conf)
	dbConn, err := postgres.NewDatabase(conf)
	if err != nil {
		log.Fatal().Msg(fmt.Sprintf("could not initialize database connection: %s", err))
	}

	repository := postgres.NewRepository(dbConn.GetDB())
	services := domain.NewService(repository, conf)
	handlers := handler.NewHandler(services, log, conf)

	app := handlers.Router()
	app.Listen(":8080")
}
