package repository

import (
	"database/sql"
)

type VoucherRepository interface {
	GetVoucher(msisdn string) ([]Voucher, error)
	GetVoucherToday(map[string]interface{}) ([]VoucherToday, error)
	GetRedeemVoucher(map[string]interface{}) ([]RedeemVoucher, error)
	GetVoucherActive(map[string]interface{}) ([]VoucherActive, error)
}

type (
	Voucher struct {
		Msisdn                sql.NullString `db:"MSISDN"`
		VoucherType           sql.NullString `db:"VOUCHER_TYPE"`
		VoucherID             sql.NullString `db:"VOUCHER_ID"`
		Used_date             sql.NullString `db:"USED_DATE"`
		ExpireDate            sql.NullString `db:"EXPIRE_DATE"`
		CategoryEn            sql.NullString `db:"CATEGORY_EN"`
		CategoryTh            sql.NullString `db:"CATEGORY_TH"`
		HeadlineEn            sql.NullString `db:"HEADLINE_EN"`
		HeadlineTh            sql.NullString `db:"HEADLINE_TH"`
		DescEn                sql.NullString `db:"DESC_EN"`
		DescTh                sql.NullString `db:"DESC_TH"`
		DescEmeraldEn         sql.NullString `db:"DESC_EMERALD_EN"`
		DescEmeraldTh         sql.NullString `db:"DESC_EMERALD_TH"`
		DescGoldEn            sql.NullString `db:"DESC_GOLD_EN"`
		DescGoldTh            sql.NullString `db:"DESC_GOLD_TH"`
		DescPlatinumEn        sql.NullString `db:"DESC_PLATINUM_EN"`
		DescPlatinumTh        sql.NullString `db:"DESC_PLATINUM_TH"`
		SerenadePrivilegeEn   sql.NullString `db:"SERENADE_PRIVILEGE_EN"`
		SerenadePrivilegeTh   sql.NullString `db:"SERENADE_PRIVILEGE_TH"`
		Condition_En          sql.NullString `db:"CONDITION_EN"`
		Condition_th          sql.NullString `db:"CONDITION_TH"`
		BrandNameEn           sql.NullString `db:"BRAND_NAME_EN"`
		BrandNameTh           sql.NullString `db:"BRAND_NAME_TH"`
		ActivateDate          sql.NullString `db:"ACTIVATE_DATE"`
		DeactivateDate        sql.NullString `db:"DEACTIVATE_DATE"`
		PrivilegeInfoID       sql.NullString `db:"PRIVILEGE_INFO_ID"`
		BrandLogo             sql.NullString `db:"BRAND_LOGO"`
		PrivilegeInfoImageAll sql.NullString `db:"PRIVILEGE_INFO_IMAGE_ALL"`
		MsgBarcode            sql.NullString `db:"MSG_BARCODE"`
		MsgReply              sql.NullString `db:"MSG_REPLY"`
		Segment               sql.NullString `db:"SEGMENT"`
		Ext_url               sql.NullString `db:"EXT_URL"`
	}
	RedeemVoucher struct {
		RefundStatus sql.NullString `db:"REFUND_STATUS"`
		MsgReply     sql.NullString `db:"MSG_REPLY"`
		Msg_Barcode  sql.NullString `db:"MSG_BARCODE"`
		BarcodeType  sql.NullString `db:"BARCODE_TYPE"`
		ExpireDate   sql.NullString `db:"EXPIRE_DATE"`
		Msg          sql.NullString `db:"MSG"`
	}
	VoucherToday struct {
		PrivilegeInfoID sql.NullInt64  `db:"PRIVILEGE_INFO_ID"`
		Points          sql.NullInt64  `db:"POINTS"`
		BrandNameEn     sql.NullString `db:"BRAND_NAME_EN"`
		BrandNameTh     sql.NullString `db:"BRAND_NAME_TH"`
		HeadlineEn      sql.NullString `db:"HEADLINE_EN"`
		HeadlineTh      sql.NullString `db:"HEADLINE_TH"`
		DefaultImg      sql.NullString `db:"DEFAULT_IMG"`
		DeactivateDate  sql.NullString `db:"DEACTIVATE_DATE"`
		Quota           sql.NullInt64  `db:"QUOTA"`
	}
	VoucherActive struct {
		RegID sql.NullString `db:"REG_ID"`
	}
)
