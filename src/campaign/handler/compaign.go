package handler

import (
	"privilege-api-myais/lib"
	"privilege-api-myais/src/campaign/service"

	"strings"

	"github.com/gofiber/fiber/v2"
)

type campaignHandler struct {
	campaignService service.CampaignService
}

func NewCampaignHandler(connect service.CampaignService) campaignHandler {
	return campaignHandler{connect}
}

type reqParamPrivRedeemHistory struct {
	TransactionID string `json:"transactionID" validate:"required"`
	Msisdn        string `json:"msisdn" validate:"required"`
	PageNumber    int    `json:"pageNumber" validate:"required"`
	ResultPerPage int    `json:"resultPerPage" validate:"required"`
}

func (h campaignHandler) GetPrivilegeRedeemHistory(c *fiber.Ctx) error {
	Body := new(reqParamPrivRedeemHistory)
	errBodyParser := c.BodyParser(Body)

	Body.TransactionID = strings.TrimSpace(Body.TransactionID)
	Body.Msisdn = strings.TrimSpace(Body.Msisdn)

	responseErr := lib.CheckReqParamValidate(lib.Validate.Struct(Body), errBodyParser, Body.TransactionID)
	if responseErr != nil {
		return c.Status(400).JSON(responseErr)
	}
	response, err := h.campaignService.GetPrivRedeemHistory(Body.Msisdn, Body.PageNumber, Body.ResultPerPage)
	if err != nil {
		return c.Status(500).JSON(lib.ErrorRes(Body.TransactionID, err))
	}
	return c.Status(200).JSON(ResponsePrivRedeemHistory(Body.TransactionID, response))
}

type reqGetInfo struct {
	TransactionID string `json:"transactionID" validate:"required"`
	PrivInfoId    string `json:"privInfoId" validate:"required"`
}

type reqGetNearByPrivilege struct {
	TransactionID string  `json:"transactionID" validate:"required"`
	LocationName  string  `json:"locationName"`
	BrandName     string  `json:"brandName"`
	LatCurrent    string  `json:"latCurrent" validate:"required"`
	LonCurrent    string  `json:"lonCurrent" validate:"required"`
	CategoryType  string  `json:"categoryType"`
	Radius        float64 `json:"radius"`
	PageNumber    int     `json:"pageNumber"`
	ResultPerPage int     `json:"resultPerPage"`
}

func (h campaignHandler) GetNearByPrivilege(c *fiber.Ctx) error {
	Body := new(reqGetNearByPrivilege)
	errBodyParser := c.BodyParser(Body)

	Body.TransactionID = strings.TrimSpace(Body.TransactionID)
	Body.LonCurrent = strings.TrimSpace(Body.LonCurrent)
	Body.LocationName = strings.TrimSpace(Body.LocationName)
	Body.BrandName = strings.TrimSpace(Body.BrandName)
	Body.LatCurrent = strings.TrimSpace(Body.LatCurrent)
	Body.LonCurrent = strings.TrimSpace(Body.LonCurrent)

	responseErr := lib.CheckReqParamValidate(lib.Validate.Struct(Body), errBodyParser, Body.TransactionID)
	if responseErr != nil {
		return c.Status(400).JSON(responseErr)
	}
	response, err := h.campaignService.GetNearByPrivilege(Body.BrandName, Body.CategoryType, Body.LatCurrent, Body.LonCurrent, Body.Radius, Body.PageNumber, Body.ResultPerPage)
	if err != nil {
		return c.Status(500).JSON(lib.ErrorRes(Body.TransactionID, err))
	}

	return c.Status(200).JSON(ResponseNearByPrivilege(Body.TransactionID, Body.PageNumber, Body.ResultPerPage, response))
}

type reqGetPrivToday struct {
	TransactionID string `json:"transactionID"  validate:"required"`
	Msisdn        string `json:"msisdn" validate:"required"`
	Channel       string `json:"channel" validate:"required"`
}

func (h campaignHandler) GetPrivToday(c *fiber.Ctx) error {
	Body := new(reqGetPrivToday)
	errBodyParser := c.BodyParser(Body)

	Body.Channel = strings.TrimSpace(Body.Channel)
	Body.Msisdn = strings.TrimSpace(Body.Msisdn)
	Body.TransactionID = strings.TrimSpace(Body.TransactionID)

	responseErr := lib.CheckReqParamValidate(lib.Validate.Struct(Body), errBodyParser, Body.TransactionID)
	if responseErr != nil {
		return c.Status(400).JSON(responseErr)
	}
	response, err := h.campaignService.GetPrivToday(Body.Msisdn, Body.Channel)
	if err != nil {
		return c.Status(500).JSON(lib.ErrorRes(Body.TransactionID, err))
	}

	return c.Status(200).JSON(ResponsePrivToday(Body.TransactionID, response))
}

type reqParamCampaignRecommend struct {
	TransactionID string `json:"transactionID"  validate:"required"`
	Msisdn        string `json:"msisdn" validate:"required"`
	CategoryType  string `json:"categoryType" validate:"required"`
}

func (h campaignHandler) GetCampaignRecommend(c *fiber.Ctx) error {
	Body := new(reqParamCampaignRecommend)
	errBodyParser := c.BodyParser(Body)

	Body.TransactionID = strings.TrimSpace(Body.TransactionID)
	Body.Msisdn = strings.TrimSpace(Body.Msisdn)
	Body.CategoryType = strings.TrimSpace(Body.CategoryType)

	responseErr := lib.CheckReqParamValidate(lib.Validate.Struct(Body), errBodyParser, Body.TransactionID)
	if responseErr != nil {
		return c.Status(400).JSON(responseErr)
	}
	response, err := h.campaignService.GetCampaignRecommend(Body.Msisdn, Body.CategoryType)
	if err != nil {
		return c.Status(500).JSON(lib.ErrorRes(Body.TransactionID, err))
	}
	return c.Status(200).JSON(ResponseCampaignRecommend(Body.TransactionID, response))
}

type reqParamSerenadeExclusive struct {
	TransactionID string `json:"transactionID"  validate:"required"`
	Msisdn        string `json:"msisdn" validate:"required"`
	CategoryType  string `json:"categoryType" validate:"required"`
}

func (h campaignHandler) GetSerenadeExclusive(c *fiber.Ctx) error {
	Body := new(reqParamSerenadeExclusive)
	errBodyParser := c.BodyParser(Body)

	Body.TransactionID = strings.TrimSpace(Body.TransactionID)
	Body.Msisdn = strings.TrimSpace(Body.Msisdn)
	Body.CategoryType = strings.TrimSpace(Body.CategoryType)

	responseErr := lib.CheckReqParamValidate(lib.Validate.Struct(Body), errBodyParser, Body.TransactionID)
	if responseErr != nil {
		return c.Status(400).JSON(responseErr)
	}
	response, err := h.campaignService.GetSerenadeExclusive(Body.Msisdn, Body.CategoryType)
	if err != nil {
		return c.Status(500).JSON(lib.ErrorRes(Body.TransactionID, err))
	}
	return c.Status(200).JSON(ResponseSerenadeExclusive(Body.TransactionID, response))
}

func (h campaignHandler) GetPrivilegeInfo(c *fiber.Ctx) error {

	Body := new(reqGetInfo)
	errBodyParser := c.BodyParser(Body)

	Body.TransactionID = strings.TrimSpace(Body.TransactionID)
	// Body.UserName = strings.TrimSpace(Body.UserName)
	// Body.PassWord = strings.TrimSpace(Body.PassWord)
	// Body.IpAddress = strings.TrimSpace(Body.IpAddress)
	// Body.SortBy = strings.TrimSpace(Body.SortBy)
	// Body.SortType = strings.TrimSpace(Body.SortType)
	// Body.PrivDesc = strings.TrimSpace(Body.PrivDesc)
	// Body.PrivCategory = strings.TrimSpace(Body.PrivCategory)
	// Body.PrivCode = strings.TrimSpace(Body.PrivCode)
	// Body.LocationName = strings.TrimSpace(Body.LocationName)
	// Body.CategoryType = strings.TrimSpace(Body.CategoryType)
	// Body.PrivInfoId = strings.TrimSpace(Body.PrivInfoId)
	// Body.PriorityChannel = strings.TrimSpace(Body.PriorityChannel)
	// Body.DisplayChannel = strings.TrimSpace(Body.DisplayChannel)
	// Body.Url = strings.TrimSpace(Body.Url)

	responseErr := lib.CheckReqParamValidate(lib.Validate.Struct(Body), errBodyParser, Body.TransactionID)
	if responseErr != nil {
		return c.Status(400).JSON(responseErr)
	}

	response, err := h.campaignService.GetPrivilegeInfo(Body.PrivInfoId)
	if err != nil {
		return c.Status(500).JSON(lib.ErrorRes(Body.TransactionID, err))
	}
	return c.Status(200).JSON(ResponseCampaign(Body.TransactionID, response))
}

func (h campaignHandler) GetPrivCompact(c *fiber.Ctx) error {
	return nil
}
