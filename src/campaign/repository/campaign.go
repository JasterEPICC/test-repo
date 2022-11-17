package repository

import "database/sql"

type CampaignRepository interface {
	GetPrivToday(map[string]interface{}) ([]PrivToday, error)
	GetPrivilegeInfo(privInfoId string) (*GetPrivilegeInfo, error)
	GetPrivRedeemHistory(map[string]interface{}) ([]PrivRedeemHistory, error)
	GetCampaignRecommend(map[string]interface{}) ([]CampaignRecommend, error)
	GetSerenadeExclusive(map[string]interface{}) ([]SerenadeExclusive, error)
	GetNearByPrivilege(map[string]interface{}) ([]NearByPrivilege, error)
}
type (
	PrivToday struct {
		EventNameEn     sql.NullString `db:"EVENT_NAME_EN"`
		EventNameTh     sql.NullString `db:"EVENT_NAME_TH"`
		BrandNameEn     sql.NullString `db:"BRAND_NAME_EN"`
		BrandNameTh     sql.NullString `db:"BRAND_NAME_TH"`
		ActivateDate    sql.NullString `db:"ACTIVATE_DATE"`
		DeactivateDate  sql.NullString `db:"DEACTIVATE_DATE"`
		MsgEn           sql.NullString `db:"MSG_EN"`
		MsgTh           sql.NullString `db:"MSG_TH"`
		PrivilegeInfoID sql.NullString `db:"PRIVILEGE_INFO_ID"`
		Ext_url         sql.NullString `db:"EXT_URL"`
		Image           sql.NullString `db:"IMAGE"`
		Points          sql.NullString `db:"POINTS"`
		Priority        sql.NullString `db:"PRIORITY"`
	}

	GetPrivilegeInfo struct {
		PrivilegeInfoID       sql.NullInt64  `db:"PRIVILEGE_INFO_ID"`
		UssdNo                sql.NullString `db:"USSD_NO"`
		Points                sql.NullInt64  `db:"POINTS"`
		CategoryID            sql.NullInt64  `db:"CATEGORY_ID"`
		Category              sql.NullString `db:"CATEGORY"`
		CategoryTh            sql.NullString `db:"CATEGORY_TH"`
		CategoryType          sql.NullString `db:"CATEGORY_TYPE"`
		BrandID               sql.NullInt64  `db:"BRAND_ID"`
		BrandNameEn           sql.NullString `db:"BRAND_NAME_EN"`
		BrandNameTh           sql.NullString `db:"BRAND_NAME_TH"`
		BrandLogo             sql.NullString `db:"BRAND_LOGO"`
		PrivilegeInfoImageAll sql.NullString `db:"PRIVILEGE_INFO_IMAGE_ALL"`
		Priority              sql.NullInt64  `db:"PRIORITY"`
		WongnaiIDAll          sql.NullString `db:"WONGNAI_ID_ALL"`
		URL                   sql.NullString `db:"URL"`
		DisplayChannel        sql.NullString `db:"DISPLAY_CHANNEL"`
		RedeemChannel         sql.NullString `db:"REDEEM_CHANNEL"`
		ActivateDate          sql.NullString `db:"ACTIVATE_DATE"`
		DeactivateDate        sql.NullString `db:"DEACTIVATE_DATE"`
		RibbonColor           sql.NullString `db:"RIBBON_COLOR"`
		RibbonMessageEn       sql.NullString `db:"RIBBON_MESSAGE_EN"`
		RibbonMessageTh       sql.NullString `db:"RIBBON_MESSAGE_TH"`
		TotalWinner           sql.NullInt64  `db:"TOTAL_WINNER"`
		QuotaRemain           sql.NullInt64  `db:"quotaRemain"`
		Voucher               sql.NullInt64  `db:"Voucher"`
		VoucherExpire         sql.NullString `db:"voucherExpire"`
		Segment               sql.NullString `db:"Segment"`
		HeadlineEn            sql.NullString `db:"Headline_En"`
		HeadlineTh            sql.NullString `db:"Headline_Th"`
		LocationUsageEn       sql.NullString `db:"Location_Usage_En"`
		LocationUsageTh       sql.NullString `db:"Location_Usage_Th"`
		ConditionEn           sql.NullString `db:"Condition_En"`
		ConditionTh           sql.NullString `db:"Condition_Th"`
		DescEmeraldEn         sql.NullString `db:"Desc_Emerald_En"`
		DescEmeraldTh         sql.NullString `db:"Desc_Emerald_Th"`
		DescGoldEn            sql.NullString `db:"Desc_Gold_En"`
		DescGoldTh            sql.NullString `db:"Desc_Gold_Th"`
		DescPlatinumEn        sql.NullString `db:"Desc_Platinum_En"`
		DescPlatinumTh        sql.NullString `db:"Desc_Platinum_Th"`
		StickerType           sql.NullString `db:"Sticker_Type"`
		CampaignHilightEn     sql.NullString `db:"Campaign_Hilight_En"`
		CampaignHilightTh     sql.NullString `db:"Campaign_Hilight_Th"`
		Feature_En            sql.NullString `db:"Feature_En"`
		Feature_Th            sql.NullString `db:"Feature_Th"`
		DescEn                sql.NullString `db:"Desc_En"`
		DescTh                sql.NullString `db:"Desc_Th"`
		VoucherExpireDate     sql.NullString `db:"Voucher_Expire_Date"`
	}

	PrivRedeemHistory struct {
		BrandNameTh sql.NullString `db:"BRAND_NAME_TH"`
		BrandNameEn sql.NullString `db:"BRAND_NAME_EN"`
		BarcodeType sql.NullString `db:"BARCODE_TYPE"`
		Headline_En sql.NullString `db:"HEADLINE_EN"`
		Headline_Th sql.NullString `db:"HEADLINE_TH"`
		Default_Img sql.NullString `db:"DEFAULT_IMG"`
		Msisdn      sql.NullString `db:"MSISDN"`
		RegTime     sql.NullString `db:"REG_TIME"`
		Msg_Barcode sql.NullString `db:"MSG_BARCODE"`
		MsgReply    sql.NullString `db:"MSG_REPLY"`
		Ext_url     sql.NullString `db:"EXT_URL"`
		CategoryID  sql.NullString `db:"CATEGORY_ID"`
		Category    sql.NullString `db:"CATEGORY"`
		CategoryTh  sql.NullString `db:"CATEGORY_TH"`
	}

	CampaignRecommend struct {
		PrivilegeInfoID sql.NullString `db:"PRIVILEGE_INFO_ID"`
		Points          sql.NullString `db:"POINTS"`
		BrandNameEn     sql.NullString `db:"BRAND_NAME_EN"`
		BrandNameTh     sql.NullString `db:"BRAND_NAME_TH"`
		HeadlineEn      sql.NullString `db:"HEADLINE_EN"`
		HeadlineTh      sql.NullString `db:"HEADLINE_TH"`
		DefaultImg      sql.NullString `db:"DEFAULT_IMG"`
	}

	SerenadeExclusive struct {
		PrivilegeInfoID sql.NullString `db:"PRIVILEGE_INFO_ID"`
		Points          sql.NullString `db:"POINTS"`
		BrandNameEn     sql.NullString `db:"BRAND_NAME_EN"`
		BrandNameTh     sql.NullString `db:"BRAND_NAME_TH"`
		HeadlineEn      sql.NullString `db:"HEADLINE_EN"`
		HeadlineTh      sql.NullString `db:"HEADLINE_TH"`
		DefaultImg      sql.NullString `db:"DEFAULT_IMG"`
	}

	NearByPrivilege struct {
		BrandID             sql.NullString `db:"BRAND_ID"`
		Latitude            sql.NullString `db:"LATITUDE"`
		Longitude           sql.NullString `db:"LONGITUDE"`
		Radius              sql.NullString `db:"RADIUS"`
		UssdNo              sql.NullString `db:"USSD_NO"`
		Points              sql.NullString `db:"POINTS"`
		Caten               sql.NullString `db:"CATEN"`
		Catth               sql.NullString `db:"CATTH"`
		Distance            sql.NullString `db:"DISTANCE"`
		SubCategoryEn       sql.NullString `db:"SUB_CATEGORY_EN"`
		SubCategoryTh       sql.NullString `db:"SUB_CATEGORY_TH"`
		CategoryType        sql.NullString `db:"CATEGORY_TYPE"`
		HeadlineEn          sql.NullString `db:"HEADLINE_EN"`
		HeadlineTh          sql.NullString `db:"HEADLINE_TH"`
		BrandNameEn         sql.NullString `db:"BRAND_NAME_EN"`
		BrandNameTh         sql.NullString `db:"BRAND_NAME_TH"`
		BrandLogo           sql.NullString `db:"BRAND_LOGO"`
		LocationUsageEn     sql.NullString `db:"LOCATION_USAGE_EN"`
		LocationUsageTh     sql.NullString `db:"LOCATION_USAGE_TH"`
		ConditionEn         sql.NullString `db:"CONDITION_EN"`
		ConditionTh         sql.NullString `db:"CONDITION_TH"`
		CampaignHilightEn   sql.NullString `db:"CAMPAIGN_HILIGHT_EN"`
		CampaignHilightTh   sql.NullString `db:"CAMPAIGN_HILIGHT_TH"`
		FeatureEn           sql.NullString `db:"FEATURE_EN"`
		FeatureTh           sql.NullString `db:"FEATURE_TH"`
		PrivImg1            sql.NullString `db:"PRIV_IMG_1"`
		PrivImg2            sql.NullString `db:"PRIV_IMG_2"`
		PrivImg3            sql.NullString `db:"PRIV_IMG_3"`
		PrivImg4            sql.NullString `db:"PRIV_IMG_4"`
		PrivImg5            sql.NullString `db:"PRIV_IMG_5"`
		DefaultImg          sql.NullString `db:"DEFAULT_IMG"`
		URL                 sql.NullString `db:"URL"`
		Priority            sql.NullString `db:"PRIORITY"`
		BranchID            sql.NullString `db:"BRANCH_ID"`
		BranchNameEn        sql.NullString `db:"BRANCH_NAME_EN"`
		BranchNameTh        sql.NullString `db:"BRANCH_NAME_TH"`
		WongnaiID           sql.NullString `db:"WONGNAI_ID"`
		PrivilegeInfoID     sql.NullString `db:"PRIVILEGE_INFO_ID"`
		DisplayChannel      sql.NullString `db:"DISPLAY_CHANNEL"`
		RedeemChannel       sql.NullString `db:"REDEEM_CHANNEL"`
		Segment             sql.NullString `db:"SEGMENT"`
		DescEn              sql.NullString `db:"DESC_EN"`
		DescTh              sql.NullString `db:"DESC_TH"`
		DescEmeraldEn       sql.NullString `db:"DESC_EMERALD_EN"`
		DescEmeraldTh       sql.NullString `db:"DESC_EMERALD_TH"`
		DescGoldEn          sql.NullString `db:"DESC_GOLD_EN"`
		DescGoldTh          sql.NullString `db:"DESC_GOLD_TH"`
		DescPlatinumEn      sql.NullString `db:"DESC_PLATINUM_EN"`
		DescPlatinumTh      sql.NullString `db:"DESC_PLATINUM_TH"`
		SerenadePrivilegeEn sql.NullString `db:"SERENADE_PRIVILEGE_EN"`
		SerenadePrivilegeTh sql.NullString `db:"SERENADE_PRIVILEGE_TH"`
		RibbonMessageEn     sql.NullString `db:"RIBBON_MESSAGE_EN"`
		RibbonMessageTh     sql.NullString `db:"RIBBON_MESSAGE_TH"`
		RibbonColor         sql.NullString `db:"RIBBON_COLOR"`
		Voucher             sql.NullString `db:"VOUCHER"`
		CampaignType        sql.NullString `db:"CAMPAIGN_TYPE"`
	}
)
