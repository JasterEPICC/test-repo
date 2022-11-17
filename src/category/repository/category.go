package repository

import "database/sql"

type CategoryRepository interface {
	GetCategory() ([]Category, error)
}

type (
	Category struct {
		CategoryID            sql.NullInt64  `db:"CATEGORY_ID"`
		Category              sql.NullString `db:"CATEGORY"`
		CategoryTh            sql.NullString `db:"CATEGORY_TH"`
		Type                  sql.NullInt64  `db:"TYPE"`
		CategoryType          sql.NullString `db:"CATEGORY_TYPE"`
		DisplayChannel        sql.NullString `db:"DISPLAY_CHANNEL"`
		CategoryEnDesc        sql.NullString `db:"CATEGORY_EN_DESC"`
		CategoryThDesc        sql.NullString `db:"CATEGORY_TH_DESC"`
		CategoryIcon          sql.NullString `db:"CATEGORY_ICON"`
		CategoryImage         sql.NullString `db:"CATEGORY_IMAGE"`
		CategoryURL           sql.NullString `db:"CATEGORY_URL"`
		Priority              sql.NullInt64  `db:"PRIORITY"`
		MobileIconActive      sql.NullString `db:"MOBILE_ICON_ACTIVE"`
		MobileIconInactive    sql.NullString `db:"MOBILE_ICON_INACTIVE"`
		SerenadeIconActive    sql.NullString `db:"SERENADE_ICON_ACTIVE"`
		SerenadeIconInactive  sql.NullString `db:"SERENADE_ICON_INACTIVE"`
		PrivilegeIconActive   sql.NullString `db:"PRIVILEGE_ICON_ACTIVE"`
		PrivilegeIconInactive sql.NullString `db:"PRIVILEGE_ICON_INACTIVE"`
		CategoryMsgTh         sql.NullString `db:"CATEGORY_MSG_TH"`
		CategoryMsgEn         sql.NullString `db:"CATEGORY_MSG_EN"`
	}
)
