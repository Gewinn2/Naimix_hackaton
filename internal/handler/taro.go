package handler

import (
	"github.com/gofiber/fiber/v2"
	"strconv"
)

// GetTaroCard
// @Tags tarot
// @Summary      Get a tarot card by ID
// @Accept       json
// @Produce      json
// @Param        id   path     int true "Tarot card ID"
// @Success      200 {object} entities.TaroCard "Details of the tarot card"
// @Failure      400 {object} entities.ErrorResponse
// @Failure      500 {object} entities.ErrorResponse
// @Router       /tarot/{id} [get]
func (h *Handler) GetTaroCard(c *fiber.Ctx) error {
	IdStr := c.Params("id")
	ID, err := strconv.Atoi(IdStr)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	card, err := h.services.TaroService.GetById(c.Context(), ID)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	return c.Status(fiber.StatusOK).JSON(card)
}
