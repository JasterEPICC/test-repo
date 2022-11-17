package router

import (
	"privilege-api-myais/lib"
	"privilege-api-myais/middleware"
	"privilege-api-myais/src/campaign/handler"
	"privilege-api-myais/src/campaign/repository"
	"privilege-api-myais/src/campaign/service"

	"github.com/gofiber/fiber/v2"
)

func Campaign(router fiber.Router) {
	Repo := repository.NewCampaignRepositoryDB(lib.ConnectionDB())
	Service := service.NewCampaignService(Repo)
	Handler := handler.NewCampaignHandler(Service)

	router.Post("/get-history", middleware.PassMiddleware, Handler.GetPrivilegeRedeemHistory).Name("getPrivilegeRedeemHistory")
	router.Post("/get-info", middleware.PassMiddleware, Handler.GetPrivilegeInfo).Name("getPrivilegeInfo")
	router.Post("/get-list", middleware.PassMiddleware, Handler.GetPrivCompact).Name("getPrivCompact")
	router.Post("/get-nearby", middleware.PassMiddleware, Handler.GetNearByPrivilege).Name("getNearByPrivilege")
	router.Post("/get-privilege-today", middleware.PassMiddleware, Handler.GetPrivToday).Name("getPrivToday")
	router.Post("/get-recommend", middleware.PassMiddleware, Handler.GetCampaignRecommend).Name("getCampaignRecommend")
	router.Post("/get-serenade-exclusive", middleware.PassMiddleware, Handler.GetSerenadeExclusive).Name("getSerenadeExclusive")
}
