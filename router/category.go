package router

import (
	"privilege-api-myais/lib"
	"privilege-api-myais/middleware"
	"privilege-api-myais/src/category/handler"
	"privilege-api-myais/src/category/repository"
	"privilege-api-myais/src/category/service"

	"github.com/gofiber/fiber/v2"
)

func Category(router fiber.Router) {
	Repo := repository.NewCategoryRepositoryDB(lib.ConnectionDB())
	Service := service.NewCategoryService(Repo)
	Handler := handler.NewCategoryHandler(Service)

	router.Post("/get", middleware.PassMiddleware, Handler.GetCategory).Name("getCategory")
}
