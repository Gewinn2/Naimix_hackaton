package handler

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/swagger"
	"github.com/rs/zerolog"
	_ "server/docs"
	"server/internal/config"
	"server/internal/log"
	"server/internal/service"
	"server/pkg"
)

type Handler struct {
	services *service.Service
	logger   *zerolog.Logger
	conf     *config.Config
}

func NewHandler(services *service.Service, logger *zerolog.Logger, conf *config.Config) *Handler {
	return &Handler{services: services, logger: logger, conf: conf}
}

func (h *Handler) Router() *fiber.App {
	f := fiber.New(fiber.Config{
		//Prefork:       true,
		//CaseSensitive: true,
		//StrictRouting: true,
	})

	f.Use(cors.New(cors.Config{
		AllowOrigins:     "*",
		AllowCredentials: true,
		AllowHeaders:     "Origin, Content-Type, Accept, Authorization",
		AllowMethods:     "GET, HEAD, PUT, PATCH, POST, DELETE",
	}))

	f.Options("/*", func(c *fiber.Ctx) error {
		c.Set("Access-Control-Allow-Origin", "http://localhost:3000,http://localhost:8080,http://lab.easy4.team") // Или ваш список доменов
		c.Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		c.Set("Access-Control-Allow-Headers", "Origin, Content-Type, Accept, Authorization")
		c.Set("Access-Control-Allow-Credentials", "true")
		return c.SendStatus(fiber.StatusNoContent)
	})

	f.Use(log.RequestLogger(h.logger))
	f.Get("/swagger/*", swagger.HandlerDefault)

	f.Post("/signup", h.SignUp)
	f.Post("/login", h.Login)
	f.Get("/users", h.GetAllUsers)
	f.Post("/company/create", h.CreateCompany)
	f.Get("/company/:id/members", h.GetCompanyMembers)
	f.Get("/tarot/:id", h.GetTaroCard)

	f.Post("/cosmogram", h.GetCosmogram)
	authGroup := f.Group("/auth")
	authGroup.Use(func(c *fiber.Ctx) error {
		return pkg.WithJWTAuth(c, h.conf.Application.SigningKey)
	})
	// authGroup.Post("/cosmogram", h.GetCosmogram)

	return f
}
