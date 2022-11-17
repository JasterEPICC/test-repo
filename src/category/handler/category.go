package handler

import (
	"privilege-api-myais/lib"
	"privilege-api-myais/src/category/service"
	"strings"

	"github.com/gofiber/fiber/v2"
)

type categoryHandler struct {
	categoryService service.CategoryService
}

func NewCategoryHandler(connect service.CategoryService) categoryHandler {
	return categoryHandler{connect}
}

type reqParamGetCategory struct {
	TransactionID string `json:"transactionID" validate:"required"`
}

func (h categoryHandler) GetCategory(c *fiber.Ctx) error {
	Body := new(reqParamGetCategory)
	errBodyParser := c.BodyParser(Body)

	Body.TransactionID = strings.TrimSpace(Body.TransactionID)

	responseErr := lib.CheckReqParamValidate(lib.Validate.Struct(Body), errBodyParser, Body.TransactionID)
	if responseErr != nil {
		return c.Status(400).JSON(responseErr)
	}

	Response, err := h.categoryService.GetCategory()
	if err != nil {
		return c.Status(500).JSON(lib.ErrorRes(Body.TransactionID, err))
	}
	return c.Status(200).JSON(ResponseCategory(Body.TransactionID, Response))
}
