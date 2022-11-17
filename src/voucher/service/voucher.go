package service

type VoucherService interface {
	GetVoucher(msisdn string) ([]VoucherResponse, error)
	GetVoucherToday(pageNumber, resultPerPage int) ([]VoucherTodayResponse, error)
	GetRedeemVoucher(msisdn, voucherType, voucherId string) ([]RedeemVoucherResponse, error)
	GetVoucherActive(msisdn string)([]VoucherActiveResponse,error)
}

type (
	VoucherResponse struct {
		Msisdn                string `json:"msisdn"`
		VoucherType           string `json:"voucher_type"`
		VoucherID             string `json:"voucher_id"`
		ExpireDate            string `json:"expire_date"`
		CategoryEn            string `json:"category_en"`
		CategoryTh            string `json:"category_th"`
		HeadlineEn            string `json:"headline_en"`
		HeadlineTh            string `json:"headline_th"`
		DescEn                string `json:"desc_en"`
		DescTh                string `json:"desc_th"`
		DescEmeraldEn         string `json:"desc_emerald_en"`
		DescEmeraldTh         string `json:"desc_emerald_th"`
		DescGoldEn            string `json:"desc_gold_en"`
		DescGoldTh            string `json:"desc_gold_th"`
		DescPlatinumEn        string `json:"desc_platinum_en"`
		DescPlatinumTh        string `json:"desc_platinum_th"`
		Condition_En          string `json:"condition_en"`
		Condition_th          string `json:"condition_th"`
		BrandNameEn           string `json:"brand_name_en"`
		BrandNameTh           string `json:"brand_name_th"`
		ActivateDate          string `json:"activate_date"`
		DeactivateDate        string `json:"deactivate_date"`
		PrivilegeInfoID       string `json:"privilege_info_id"`
		BrandLogo             string `json:"brand_logo"`
		PrivilegeInfoImageAll string `json:"privilege_info_image_all"`
	}
	RedeemVoucherResponse struct {
		RefundStatus string `json:"refundStatus"`
		MsgReply     string `json:"msgReply"`
		Msg_Barcode  string `json:"msgBarcode"`
		BarcodeType  string `json:"barcodeType"`
		ExpireDate   string `json:"expireDate"`
		Msg          string `json:"msg"`
	}
	VoucherTodayResponse struct {
		PrivilegeInfoID int    `json:"Privilege_Info_Id"`
		Points          int    `json:"Points"`
		BrandNameEn     string `json:"Brand_Name_En"`
		BrandNameTh     string `json:"Brand_Name_Th"`
		HeadlineEn      string `json:"Headline_En"`
		HeadlineTh      string `json:"Headline_Th"`
		DefaultImg      string `json:"Default_Img"`
		DeactivateDate  string `json:"Deactivate_Date"`
		Quota           int    `json:"quota"`
		MsgReply        string `json:"msg_reply"`
		Msg_Barcode     string `json:"msg_barcode"`
		BarcodeType     string `json:"barcode_type"`
	}
	VoucherActiveResponse struct {
		RegID string `json:"reg_id"`
	}
)
