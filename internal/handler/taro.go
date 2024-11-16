package handler

import (
	"github.com/gofiber/fiber/v2"
	"log"
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
	CategoryIdStr := c.Params("id")
	CategoryID, err := strconv.Atoi(CategoryIdStr)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid ID format"})
	}

	log.Println("Calling service to get Taro card")
	card, err := h.services.TaroService.GetById(c.Context(), CategoryID)
	if err != nil {
		log.Printf("Error getting Taro card: %v", err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	return c.Status(fiber.StatusOK).JSON(card)

}
