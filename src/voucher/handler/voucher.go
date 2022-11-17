package handler

import (
	"privilege-api-myais/lib"
	"privilege-api-myais/src/voucher/service"
	"strings"

	"github.com/gofiber/fiber/v2"
)

type voucherHandler struct {
	voucherService service.VoucherService
}

func NewVoucherHandler(connect service.VoucherService) voucherHandler {
	return voucherHandler{connect}
}

type reqParamGetVoucher struct {
	TransactionID string `json:"transactionID"  validate:"required"`
	Msisdn        string `json:"msisdn" validate:"required"`
}

func (h voucherHandler) GetVoucher(c *fiber.Ctx) error {
	Body := new(reqParamGetVoucher)

	errBodyParser := c.BodyParser(Body)

	Body.Msisdn = strings.TrimSpace(Body.Msisdn)
	Body.TransactionID = strings.TrimSpace(Body.TransactionID)

	responseErr := lib.CheckReqParamValidate(lib.Validate.Struct(Body), errBodyParser, Body.TransactionID)
	if responseErr != nil {
		return c.Status(400).JSON(responseErr)
	}
	response, err := h.voucherService.GetVoucher(Body.Msisdn)
	if err != nil {
		return c.Status(500).JSON(lib.ErrorRes(Body.TransactionID, err))
	}
	return c.Status(200).JSON(ResponseVoucher(Body.TransactionID, response))
}

type reqParamGetVoucherToday struct {
	TransactionID string `json:"transactionID"  validate:"required"`
	PageNumber    int    `json:"pageNumber" validate:"required"`
	ResultPerPage int    `json:"resultPerPage" validate:"required"`
}

func (h voucherHandler) GetVoucherToday(c *fiber.Ctx) error {
	Body := new(reqParamGetVoucherToday)
	errBodyParser := c.BodyParser(Body)

	Body.TransactionID = strings.TrimSpace(Body.TransactionID)

	responseErr := lib.CheckReqParamValidate(lib.Validate.Struct(Body), errBodyParser, Body.TransactionID)
	if responseErr != nil {
		return c.Status(400).JSON(responseErr)
	}
	response, err := h.voucherService.GetVoucherToday(Body.PageNumber, Body.ResultPerPage)
	if err != nil {
		return c.Status(500).JSON(lib.ErrorRes(Body.TransactionID, err))
	}
	return c.Status(200).JSON(ResponseVoucherToday(Body.TransactionID, response))
}

type reqParamRedeemVoucher struct {
	TransactionID string `json:"transactionID"  validate:"required"`
	Msisdn        string `json:"msisdn"  validate:"required"`
	VoucherType   string `json:"vouchertype"  validate:"required"`
	VoucherId     string `json:"voucherId"  validate:"required"`
}

func (h voucherHandler) RedeemVoucher(c *fiber.Ctx) error {
	Body := new(reqParamRedeemVoucher)
	errBodyParser := c.BodyParser(Body)

	Body.TransactionID = strings.TrimSpace(Body.TransactionID)
	Body.Msisdn = strings.TrimSpace(Body.Msisdn)
	Body.VoucherType = strings.TrimSpace(Body.VoucherType)
	Body.VoucherId = strings.TrimSpace(Body.VoucherId)

	responseErr := lib.CheckReqParamValidate(lib.Validate.Struct(Body), errBodyParser, Body.TransactionID)
	if responseErr != nil {
		return c.Status(400).JSON(responseErr)
	}
	Response, err := h.voucherService.GetRedeemVoucher(Body.Msisdn, Body.VoucherType, Body.VoucherId)
	if err != nil {
		return c.Status(500).JSON(lib.ErrorRes(Body.TransactionID, err))
	}
	return c.Status(200).JSON(ResponseRedeemVoucher(Body.TransactionID, Response))
}

type reqParamVoucherActive struct {
	TransactionID string `json:"transactionID"  validate:"required"`
	Msisdn        string `json:"msisdn"  validate:"required"`
}

func (h voucherHandler) GetVoucherActive(c *fiber.Ctx) error {
	Body := new(reqParamVoucherActive)
	errBodyParser := c.BodyParser(Body)

	Body.TransactionID = strings.TrimSpace(Body.TransactionID)
	Body.Msisdn = strings.TrimSpace(Body.Msisdn)

	responseErr := lib.CheckReqParamValidate(lib.Validate.Struct(Body),errBodyParser,Body.TransactionID)
	if responseErr != nil {
		return c.Status(400).JSON(responseErr)
	}
	Response , err := h.voucherService.GetVoucherActive(Body.Msisdn)
	if err != nil {
		return c.Status(500).JSON(lib.ErrorRes(Body.TransactionID,err))
	}
	return c.Status(200).JSON(ResponseVoucherActive(Body.TransactionID,Response))
}
