package handler

import (
	"privilege-api-myais/lib"
	"privilege-api-myais/src/auth/service"

	"github.com/gofiber/fiber/v2"
)

type authHandler struct {
	authService service.AuthService
}

func NewAuthHandler(connect service.AuthService) authHandler {
	return authHandler{connect}
}

type reqParamJWT struct {
	User string `json:"user" validate:"required"`
	// CaseId    string `json:"caseId" validate:"required"`
}

func (h authHandler) GenerateToken(c *fiber.Ctx) error {
	Body := new(reqParamJWT)
	errBodyParser := c.BodyParser(Body)

	responseErr := lib.CheckReqParamValidate(lib.Validate.Struct(Body), errBodyParser, "")
	if responseErr != nil {
		return c.Status(400).JSON(responseErr)
	}

	res, err := h.authService.GenerateToken(Body.User)
	if err != nil {
		return c.Status(500).JSON(lib.ErrorRes("", err))
	}

	return c.Status(200).JSON(res)
}

func (h authHandler) RevokeToken(c *fiber.Ctx) error {
	Body := new(reqParamJWT)
	errBodyParser := c.BodyParser(Body)

	responseErr := lib.CheckReqParamValidate(lib.Validate.Struct(Body), errBodyParser, "")
	if responseErr != nil {
		return c.Status(400).JSON(responseErr)
	}

	res, err := h.authService.RevokeToken(Body.User)
	if err != nil {
		return c.Status(500).JSON(lib.ErrorRes("", err))
	}

	return c.Status(200).JSON(res)
}
