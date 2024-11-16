package handler

import (
	"github.com/gofiber/fiber/v2"
	"server/internal/entities"
	"server/internal/log"
	"strconv"
)

// CreateCompany
// @Tags company
// @Summary      Create a new company
// @Accept       json
// @Produce      json
// @Param        data body entities.CreateCompanyRequest true "Company creation request data"
// @Success      201 {object} entities.CreateCompanyResponse
// @Failure      400 {object} entities.ErrorResponse
// @Failure      500 {object} entities.ErrorResponse
// @Router       /company/create [post]
func (h *Handler) CreateCompany(c *fiber.Ctx) error {
	var request *entities.CreateCompanyRequest
	if err := c.BodyParser(&request); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}
	h.logger.Debug().Msg("call h.services.CompanyService.CreateCompany")
	res, err := h.services.CompanyService.CreateCompany(c.Context(), request)
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}
	return c.Status(fiber.StatusCreated).JSON(res)
}

// GetCompanyMembers
// @Tags company
// @Summary      Get members of a company by ID
// @Accept       json
// @Produce      json
// @Param        id   path     int true "Company ID"
// @Success      200 {array}  entities.CompanyMemberInfo "List of company members"
// @Failure      400 {object} entities.ErrorResponse
// @Failure      500 {object} entities.ErrorResponse
// @Router       /company/{id}/members [get]
func (h *Handler) GetCompanyMembers(c *fiber.Ctx) error {
	companyIdString := c.Params("id")
	companyId, err := strconv.Atoi(companyIdString)
	if err != nil {
		logEvent := log.CreateLog(h.logger, log.LogsField{Level: "Error", Method: c.Method(),
			Url: c.OriginalURL(), Status: fiber.StatusBadRequest})
		logEvent.Msg(err.Error())
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	request := &entities.GetCompanyMembersInfoRequest{CompanyID: companyId}

	h.logger.Debug().Msg("h.services.CompanyService.GetAllMembers")
	res, err := h.services.CompanyService.GetAllMembers(c.Context(), request)
	if err != nil {
		logEvent := log.CreateLog(h.logger, log.LogsField{Level: "Error", Method: c.Method(),
			Url: c.OriginalURL(), Status: fiber.StatusInternalServerError})
		logEvent.Msg(err.Error())
		return fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}

	logEvent := log.CreateLog(h.logger, log.LogsField{Level: "Info", Method: c.Method(),
		Url: c.OriginalURL(), Status: fiber.StatusOK})
	logEvent.Msg("success")
	return c.Status(fiber.StatusOK).JSON(res)
}
