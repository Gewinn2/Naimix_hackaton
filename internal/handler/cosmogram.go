package handler

import (
	"github.com/gofiber/fiber/v2"
	"server/internal/entities"
	"server/internal/log"
)

// GetCosmogram
// @Tags cosmogram
// @Summary      Get cosmogram
// @Accept       json
// @Produce      json
// @Param data body entities.CosmogramRequestBody true "Get cosmogram data"
// @Success 200 {object} entities.CosmogramResponseBody
// @Failure 400 {object} entities.ErrorResponse
// @Failure 500 {object} entities.ErrorResponse
// @Router       /cosmogram [post]
func (h *Handler) GetCosmogram(c *fiber.Ctx) error {
	var cosmogramRequestBody entities.CosmogramRequestBody
	if err := c.BodyParser(&cosmogramRequestBody); err != nil {
		logEvent := log.CreateLog(h.logger, log.LogsField{Level: "Error", Method: c.Method(),
			Url: c.OriginalURL(), Status: fiber.StatusBadRequest})
		logEvent.Msg(err.Error())
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	h.logger.Debug().Msg("call h.services.CosmogramService.Get")
	responses, err := h.services.CosmogramService.Get(cosmogramRequestBody)
	if err != nil {
		logEvent := log.CreateLog(h.logger, log.LogsField{Level: "Error", Method: c.Method(),
			Url: c.OriginalURL(), Status: fiber.StatusInternalServerError})
		logEvent.Msg(err.Error())
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	logEvent := log.CreateLog(h.logger, log.LogsField{Level: "Info", Method: c.Method(),
		Url: c.OriginalURL(), Status: fiber.StatusOK})
	logEvent.Msg("success")
	return c.Status(fiber.StatusOK).JSON(responses)
}
