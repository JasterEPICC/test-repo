package handler

import (
	"privilege-api-myais/src/voucher/service"
	"strings"
	"time"
)

type templateVoucher struct {
	TransactionID string                    `json:"transactionID"`
	Status        string                    `json:"status"`
	Description   string                    `json:"description"`
	TotalVoucher  int                       `json:"totalVoucher"`
	VoucherArr    []service.VoucherResponse `json:"voucherArr"`
}

func ResponseVoucher(id string, arg []service.VoucherResponse) *templateVoucher {
	response := new(templateVoucher)
	response.TransactionID = id
	if len(arg) != 0 {
		response.Status = "20000"
		response.Description = "SUCCESS"
	} else {
		response.Status = "20002"
		response.Description = "Voucher not found."
	}
	response.VoucherArr = arg
	response.TotalVoucher = len(arg)

	return response
}

type templateRedeemVoucher struct {
	TransactionID string `json:"transactionID"`
	HttpStatus    string `json:"httpStatus"`
	Status        string `json:"status"`
	Description   string `json:"description"`
	Msg           string `json:"msg"`
	MsgBarcode    string `json:"msgBarcode"`
	BarcodeType   string `json:"barcodeType"`
}

func ResponseRedeemVoucher(id string, arg []service.RedeemVoucherResponse) *templateRedeemVoucher {
	response := new(templateRedeemVoucher)
	response.TransactionID = id

	if len(arg) != 0 {
		timeSplit := strings.Split(arg[0].ExpireDate, " ")
		if arg[0].RefundStatus == "3" {
			response.Status = "20003"
			response.HttpStatus = "200"
			response.Description = "Voucher used."
			response.Msg = ""
			response.MsgBarcode = ""
			response.BarcodeType = ""
		} else if u, _ := time.Parse("2006-01-02", timeSplit[0]); time.Now().After(u) {
			response.Status = "20003"
			response.HttpStatus = "200"
			response.Description = "Voucher expired."
			response.Msg = ""
			response.MsgBarcode = ""
			response.BarcodeType = ""
		} else {
			response.Status = "20000"
			response.HttpStatus = "200"
			response.Description = "SUCCESS"
			response.Msg = arg[0].MsgReply
			response.MsgBarcode = arg[0].ExpireDate
			response.BarcodeType = arg[0].BarcodeType
		}
	} else {
		response.Status = "20002"
		response.HttpStatus = "200"
		response.Description = "Voucher invalid."
		response.Msg = ""
		response.MsgBarcode = ""
		response.BarcodeType = ""

	}
	return response
}

type templateVoucherToday struct {
	TransactionID     string                         `json:"transactionID"`
	Status            string                         `json:"status"`
	Description       string                         `json:"description"`
	TotalVoucherToday int                            `json:"totalVoucherToday"`
	VoucherTodayArr   []service.VoucherTodayResponse `json:"voucherTodayArr"`
}

func ResponseVoucherToday(id string, arg []service.VoucherTodayResponse) *templateVoucherToday {
	response := new(templateVoucherToday)
	response.TransactionID = id
	if len(arg) != 0 {
		response.Status = "20000"
		response.Description = "SUCCESS"
	} else {
		response.Status = "20001"
		response.Description = "Data not found."
	}
	response.VoucherTodayArr = arg
	response.TotalVoucherToday = len(arg)

	return response
}


type templateVoucherActive struct { 
	TransactionID     string                         `json:"transactionID"`
	HttpStatus         string                     	 `json:"httpstatus"`
	Status            string                         `json:"status"`
	Description       string                         `json:"description"`
	Count 			  int                            `json:"count"`
}
func ResponseVoucherActive(id string , arg[]service.VoucherActiveResponse) *templateVoucherActive{
	response := new(templateVoucherActive)
	response.TransactionID = id
	if len(arg) != 0 {
		response.HttpStatus = "200"
		response.Status  = "20000"
		response.Description = "SUCCESS"
	}else{
		response.HttpStatus = "200"
		response.Status  = "20001"
		response.Description = "Data not found."
	}
	response.Count = len(arg)
	return response
}
