package router

import (
	"privilege-api-myais/lib"
	"privilege-api-myais/middleware"
	"privilege-api-myais/src/voucher/handler"
	"privilege-api-myais/src/voucher/repository"
	"privilege-api-myais/src/voucher/service"

	"github.com/gofiber/fiber/v2"
)

func Voucher(router fiber.Router) {
	Repo := repository.NewVoucherRepositoryDB(lib.ConnectionDB())
	Service := service.NewVoucherService(Repo)
	Handler := handler.NewVoucherHandler(Service)

	router.Post("/get", middleware.PassMiddleware, Handler.GetVoucher).Name("getVoucher")
	router.Post("/get-active", middleware.PassMiddleware, Handler.GetVoucherActive).Name("getVoucherActive")
	router.Post("/get-today", middleware.PassMiddleware, Handler.GetVoucherToday).Name("getVoucherToday")
	router.Post("/redeem", middleware.PassMiddleware, Handler.RedeemVoucher).Name("redeemVoucher")
}
