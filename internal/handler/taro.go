package handler

import (
	"github.com/gofiber/fiber/v2"
	log "server/internal/log"
	"strconv"
)

// GetTaroCard
// @Tags tarot
// @Summary      Get a tarot card by category ID
// @Accept       json
// @Produce      json
// @Param        id   path     int true "Tarot card ID"
// @Success      200 {object} entities.TaroCard "Details of the tarot card"
// @Failure      400 {object} entities.ErrorResponse
// @Failure      500 {object} entities.ErrorResponse
// @Router       /tarot/{id} [get]
func (h *Handler) GetTaroCard(c *fiber.Ctx) error {
	CategoryIdStr := c.Params("id")
	CategoryID, err := strconv.Atoi(CategoryIdStr)
	if err != nil {
		logEvent := log.CreateLog(h.logger, log.LogsField{Level: "Error", Method: c.Method(),
			Url: c.OriginalURL(), Status: fiber.StatusBadRequest})
		logEvent.Msg("invalid parameter")
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid ID format"})
	}

	h.logger.Debug().Msg("call h.services.TaroService.GetById")
	card, err := h.services.TaroService.GetCard(c.Context(), CategoryID)
	if err != nil {
		logEvent := log.CreateLog(h.logger, log.LogsField{Level: "Error", Method: c.Method(),
			Url: c.OriginalURL(), Status: fiber.StatusInternalServerError})
		logEvent.Msg("error getting Taro card")
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	logEvent := log.CreateLog(h.logger, log.LogsField{Level: "Info", Method: c.Method(),
		Url: c.OriginalURL(), Status: fiber.StatusOK})
	logEvent.Msg("success")
	return c.Status(fiber.StatusOK).JSON(card)
}
