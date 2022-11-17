package service

type CampaignService interface {
	GetPrivToday(msisdn, channel string) ([]PrivTodayResponse, error)
	GetPrivilegeInfo(privInfoId string) (*GetCampaignResponse, error)
	GetPrivRedeemHistory(msisdn string, pageNumber, resultPerPage int) (*PrivRedeemHistoryResponse, error)
	GetCampaignRecommend(msisdn, categoryType string) ([]CampaignRecommendResponse, error)
	GetSerenadeExclusive(msisdn, categoryType string) ([]SerenadeExclusiveResponse, error)
	GetNearByPrivilege(name, cattype, latitude, longitude string, radius float64, pageNumber, resultPerPage int) ([]NearByPrivilegeResponse, error)
}

type (
	PrivTodayResponse struct {
		EventNameEn     string `json:"event_name_en"`
		EventNameTh     string `json:"event_name_th"`
		BrandNameEn     string `json:"brand_name_en"`
		BrandNameTh     string `json:"brand_name_th"`
		ActivateDate    string `json:"activate_date"`
		DeactivateDate  string `json:"deactivate_date"`
		MsgEn           string `json:"msg_en"`
		MsgTh           string `json:"msg_th"`
		PrivilegeInfoID string `json:"privilege_info_id"`
		Ext_url         string `json:"ext_url"`
		Image           string `json:"image"`
		Points          string `json:"points"`
	}

	GetCampaignResponse struct {
		TransactionID  string `json:"transactionID"`
		Status         string `json:"status"`
		Description    string `json:"description"`
		PrivInfoID     int64  `json:"privInfoId"`
		UssdNo         string `json:"ussdNo"`
		Points         int64  `json:"points"`
		CategoryID     int64  `json:"categoryId"`
		CategoryNameEn string `json:"categoryNameEn"`
		CategoryNameTh string `json:"categoryNameTh"`
		CategoryType   string `json:"categoryType"`
		BrandID        int64  `json:"brandId"`
		BrandNameEn    string `json:"brandNameEn"`
		BrandNameTh    string `json:"brandNameTh"`
		BrandLogo      string `json:"brandLogo"`
		ImageAll       string `json:"imageAll"`
		Priority       int64  `json:"priority"`
		WongnaiID      string `json:"wongnaiId"`
		URL            string `json:"url"`
		DisplayChannel string `json:"displayChannel"`
		RedeemChannel  string `json:"redeemChannel"`
		ActivateDate   string `json:"activateDate"`
		DeactivateDate string `json:"deactivateDate"`
		RibbonColor    string `json:"ribbonColor"`
		RibbonMsgEn    string `json:"ribbonMsgEn"`
		RibbonMsgTh    string `json:"ribbonMsgTh"`
		Quota          int64  `json:"quota"`
		QuotaRemain    int64  `json:"quotaRemain"`
		StickerType    string `json:"stickerType"`
		Voucher        int64  `json:"voucher"`
		VoucherExpire  string `json:"voucherExpire"`
		Segment        string `json:"segment"`
		HeadLineEn     string `json:"headLineEn"`
		HeadLineTh     string `json:"headLineTh"`
		LocationEN     string `json:"locationEN"`
		LocationTH     string `json:"locationTH"`
		ConditionEN    string `json:"conditionEN"`
		ConditionTH    string `json:"conditionTH"`
		HilightEn      string `json:"hilightEn"`
		HilightTh      string `json:"hilightTh"`
		FeatureEn      string `json:"featureEn"`
		FeatureTh      string `json:"featureTh"`
		DescEn         string `json:"descEn"`
		DescTh         string `json:"descTh"`
		DescEmeraldEn  string `json:"descEmeraldEn"`
		DescEmeraldTh  string `json:"descEmeraldTh"`
		DescGoldEn     string `json:"descGoldEn"`
		DescGoldTh     string `json:"descGoldTh"`
		DescPlatinumEn string `json:"descPlatinumEn"`
		DescPlatinumTh string `json:"descPlatinumTh"`
	}

	PrivRedeemHistoryResponse struct {
		TransactionID             string                              `json:"transactionID"`
		HttpStatus                string                              `json:"httpStatus"`
		Status                    string                              `json:"status"`
		Description               string                              `json:"description"`
		PageNumber                int                                 `json:"pageNumber"`
		TotalPage                 int                                 `json:"totalPage"`
		ResultPerPage             int                                 `json:"resultPerPage"`
		ResultAvailable           int                                 `json:"resultAvailable"`
		TotalResultReturned       int                                 `json:"totalResultReturned"`
		SystemTime                string                              `json:"systemTime"`
		PrivilegeRedeemHistoryArr []PrivilegeRedeemHistoryArrResponse `json:"privilegeRedeemHistoryArr"`
	}

	PrivilegeRedeemHistoryArrResponse struct {
		BrandNameTh string `json:"brandNameTH"`
		BrandNameEn string `json:"brandNameEN"`
		BarcodeType string `json:"barcodeType"`
		DescEn      string `json:"descEN"`
		DescTh      string `json:"descTH"`
		DefaultImg  string `json:"defaultImg"`
		Msisdn      string `json:"msisdn"`
		TranDate    string `json:"tranDate"`
		MsgBarcode  string `json:"msgBarcode"`
		MsgReply    string `json:"msgReply"`
	}

	CampaignRecommendResponse struct {
		PrivilegeInfoID string `json:"PrivilegeInfoID"`
		Points          string `json:"Points"`
		BrandNameEn     string `json:"BrandNameEn"`
		BrandNameTh     string `json:"BrandNameTh"`
		HeadlineEn      string `json:"HeadlineEn"`
		HeadlineTh      string `json:"HeadlineTh"`
		DefaultImg      string `json:"DefaultImg"`
	}
	SerenadeExclusiveResponse struct {
		PrivilegeInfoID string `json:"PrivilegeInfoID"`
		Points          string `json:"Points"`
		BrandNameEn     string `json:"BrandNameEn"`
		BrandNameTh     string `json:"BrandNameTh"`
		HeadlineEn      string `json:"HeadlineEn"`
		HeadlineTh      string `json:"HeadlineTh"`
		DefaultImg      string `json:"DefaultImg"`
	}

	NearByPrivilegeResponse struct {
		BrandID         string                          `json:"brandId"`
		Latitude        string                          `json:"latitude"`
		Longitude       string                          `json:"longitude"`
		Distance        string                          `json:"distance"`
		Radius          string                          `json:"radius"`
		UssdNo          string                          `json:"ussdNo"`
		Points          string                          `json:"points"`
		CategoryEN      string                          `json:"categoryEN"`
		CategoryTH      string                          `json:"categoryTH"`
		SubCategoryEN   string                          `json:"subCategoryEN"`
		SubCategoryTH   string                          `json:"subCategoryTH"`
		CategoryType    string                          `json:"categoryType"`
		HeadLineEN      string                          `json:"headLineEN"`
		HeadLineTH      string                          `json:"headLineTH"`
		BrandNameEN     string                          `json:"brandNameEN"`
		BrandNameTH     string                          `json:"brandNameTH"`
		BrandLogo       string                          `json:"brandLogo"`
		LocationEN      string                          `json:"locationEN"`
		LocationTH      string                          `json:"locationTH"`
		ConditionEN     string                          `json:"conditionEN"`
		ConditionTH     string                          `json:"conditionTH"`
		HilightEN       string                          `json:"hilightEN"`
		HilightTH       string                          `json:"hilightTH"`
		FeatureEN       string                          `json:"featureEN"`
		FeatureTH       string                          `json:"featureTH"`
		PrivImg1        string                          `json:"privImg1"`
		PrivImg2        string                          `json:"privImg2"`
		PrivImg3        string                          `json:"privImg3"`
		PrivImg4        string                          `json:"privImg4"`
		PrivImg5        string                          `json:"privImg5"`
		PrivImg6        string                          `json:"privImg6"`
		PrivImg7        string                          `json:"privImg7"`
		PrivImg8        string                          `json:"privImg8"`
		PrivImg9        string                          `json:"privImg9"`
		PrivImg10       string                          `json:"privImg10"`
		DefaultImg      string                          `json:"defaultImg"`
		URL             string                          `json:"url"`
		Priority        string                          `json:"priority"`
		BranchID        string                          `json:"branchId"`
		BranchNameEN    string                          `json:"branchNameEN"`
		BranchNameTH    string                          `json:"branchNameTH"`
		WongnaiID       string                          `json:"wongnaiId"`
		PrivilegeInfoID string                          `json:"privilegeInfoId"`
		DisplayChannel  string                          `json:"displayChannel"`
		RedeemChannel   string                          `json:"redeemChannel"`
		CustomerDescEN  string                          `json:"customerDescEN"`
		CustomerDescTH  string                          `json:"customerDescTH"`
		ActivateDate    string                          `json:"activateDate"`
		DeactivateDate  string                          `json:"deactivateDate"`
		RibbonMsgEN     string                          `json:"ribbonMsgEN"`
		RibbonMsgTH     string                          `json:"ribbonMsgTH"`
		RibbonMsgColor  string                          `json:"ribbonMsgColor"`
		Voucher         string                          `json:"voucher"`
		TotalDesc       string                          `json:"totalDesc"`
		DescriptionArr  []NearByPrivilegeDesArrResponse `json:"descriptionArr"`
		Segment         string                          `json:"segment"`
	}

	NearByPrivilegeDesArrResponse struct {
		Segment string `json:"segment"`
		DescEN  string `json:"descEN"`
		DescTH  string `json:"descTH"`
	}
)
