package middleware

import (
	"privilege-api-myais/lib"
	"privilege-api-myais/src/auth/repository"
	"privilege-api-myais/src/auth/service"

	"github.com/gofiber/fiber/v2"
)

type reqParamToken struct {
	Signature string `json:"signature" validate:"required"`
}

func PassMiddleware(c *fiber.Ctx) error {
	// fmt.Println(c.Route().Name)
	return c.Next()
}

func AuthenRequired(c *fiber.Ctx) error {
	listCommandId, errGet := getCommandId(c)
	if errGet != nil {
		return c.Status(400).JSON(errGet)
	}

	validate := false

	for _, detail := range listCommandId {
		if c.Route().Name == detail.Command {
			validate = true
			break
		}
	}

	if validate {
		return c.Next()
	} else {
		return c.Status(400).JSON(lib.ErrorRes("", lib.NewError(40301, "User unauthorized.")))
	}
}

func getCommandId(c *fiber.Ctx) ([]service.CommandResponse, *lib.TemplateError) {
	Repo := repository.NewAuthRepositoryDB(lib.ConnectionDB())
	Service := service.NewAuthService(Repo)

	Body := new(reqParamToken)
	errBodyParser := c.BodyParser(Body)

	responseErr := lib.CheckReqParamValidate(lib.Validate.Struct(Body), errBodyParser, "")
	if responseErr != nil {
		return nil, responseErr
	}

	listCommandId, _ := Service.GetCommand(Body.Signature)
	return listCommandId, nil
}
