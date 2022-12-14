package repository

import "github.com/jmoiron/sqlx"

type voucherRepositoryDB struct {
	db *sqlx.DB
}

func NewVoucherRepositoryDB(connect *sqlx.DB) VoucherRepository {
	return voucherRepositoryDB{connect}
}

func (r voucherRepositoryDB) GetVoucher(msisdn string) (connectDB []Voucher, err error) {
	query := `
SELECT  RPT.MSISDN, RPT.VOUCHER_TYPE, RPT.VOUCHER_ID, 
	FORMAT(RPT.EXPIRE_DATE, 'dd/MM/yyyy hh:mm:ss') as EXPIRE_DATE,
	RPT.USED_DATE, I.CATEGORY_EN, 
	I.CATEGORY_TH, I.HEADLINE_EN, I.HEADLINE_TH,
	I.DESC_EN, I.DESC_TH,I.DESC_EMERALD_EN, 
	I.DESC_EMERALD_TH, I.DESC_GOLD_EN, I.DESC_GOLD_TH,
	I.DESC_PLATINUM_EN, I.DESC_PLATINUM_TH,I.SERENADE_PRIVILEGE_EN, 
	I.SERENADE_PRIVILEGE_TH, B.BRAND_NAME_EN, B.BRAND_NAME_TH, 
	I.CONDITION_EN, I.CONDITION_TH, FORMAT(I.ACTIVATE_DATE, 'dd/MM/yyyy hh:mm:ss') as ACTIVATE_DATE,
	FORMAT(I.DEACTIVATE_DATE, 'dd/MM/yyyy hh:mm:ss') as DEACTIVATE_DATE, I.PRIVILEGE_INFO_ID,B.BRAND_LOGO, 
	I.PRIVILEGE_INFO_IMAGE_ALL, M.MSG_BARCODE, M.MSG_REPLY,
	SEGMENT, M.EXT_URL
FROM (
	SELECT R.MSISDN, 'PRIVILEGE' VOUCHER_TYPE, R.REG_ID VOUCHER_ID,
		R.REG_TIME REG_TIME, R.EXPIRE_DATE, R.USED_DATE,
		R.PROJECT_ID, R.CHOICE_ID, R.MSG_ID
	FROM PRIVUSER.REGISTER R, PRIVUSER.PROJECT P
	WHERE P.PROJECT_ID = R.PROJECT_ID 
		AND P.VOUCHER = 1 
		AND R.MSISDN = ? 
		AND R.REG_TIME > DATEADD(DAY, -90, GETDATE())
	UNION 
	SELECT PT.MSISDN, 'POINT' VOUCHER_TYPE, PT.POINT_TRAN_ID VOUCHER_ID, 
		PT.TRAN_DATE REG_TIME, PT.EXPIRE_DATE, PT.USED_DATE,
		PT.PROJECT_ID, PT.CHOICE_ID, PT.MSG_ID 	
	FROM PRIVUSER.POINT_TRANSACTION PT, PRIVUSER.PROJECT P
	WHERE P.PROJECT_ID = PT.PROJECT_ID 
		AND P.VOUCHER = 1 
		AND PT.MSISDN = ? 
		AND PT.TRAN_DATE > DATEADD(DAY, -90, GETDATE()) 
) RPT 
INNER JOIN PRIVUSER.PRIVILEGE_INFO I ON RPT.PROJECT_ID = I.PROJECT_ID 
	AND RPT.CHOICE_ID = I.CHOICE_ID
INNER JOIN PRIVUSER.BRAND B ON I.BRAND_ID = B.BRAND_ID 
INNER JOIN PRIVUSER.MSG M ON RPT.MSG_ID = M.MSG_ID 
ORDER BY REG_TIME DESC `

	if err = r.db.Select(&connectDB, query, msisdn, msisdn); err != nil {
		return nil, err
	}
	return connectDB, nil
}

func (r voucherRepositoryDB) GetRedeemVoucher(arg map[string]interface{}) (connectDB []RedeemVoucher, err error) {
	query := `
SELECT R.REFUND_STATUS, M.MSG_REPLY, M.MSG_BARCODE, 
	P.BARCODE_TYPE, R.EXPIRE_DATE, M.MSG 
FROM (
	SELECT MSG_ID, REFUND_STATUS, PROJECT_ID, EXPIRE_DATE 
	FROM PRIVUSER.REGISTER 
	WHERE MSISDN = ?  AND REG_ID = ? 
	UNION 
	SELECT MSG_ID, REFUND_STATUS, PROJECT_ID, EXPIRE_DATE
	FROM PRIVUSER.POINT_TRANSACTION 
	WHERE MSISDN = ? 
		AND POINT_TRAN_ID = ?
) R 
INNER JOIN PRIVUSER.MSG M ON (R.MSG_ID = M.MSG_ID)
INNER JOIN PRIVUSER.PROJECT P ON (R.PROJECT_ID = P.PROJECT_ID) `

	if err = r.db.Select(&connectDB, query, arg["msisdn"], arg["voucherId"], arg["msisdn"], arg["voucherId"]); err != nil {
		return nil, err
	}
	return connectDB, nil
}

func (r voucherRepositoryDB) GetVoucherToday(arg map[string]interface{}) (connectDB []VoucherToday, err error) {
	query := ` 
SELECT PRIVILEGE_INFO_ID, POINTS, BRAND_NAME_EN,  BRAND_NAME_TH, HEADLINE_EN, HEADLINE_TH, DEFAULT_IMG, DEACTIVATE_DATE, QUOTA
FROM (
	SELECT ROW_NUMBER() OVER (ORDER BY VT.PRIORITY, PI.ACTIVATE_DATE DESC) LINE_NUMBER, PI.PRIVILEGE_INFO_ID,PI.POINTS,
	PI.HEADLINE_EN,PI.HEADLINE_TH,PI.DEFAULT_IMG, FORMAT(PI.DEACTIVATE_DATE, 'dd/MM/yyyy hh:mm:ss') as DEACTIVATE_DATE, B.BRAND_NAME_EN,B.BRAND_NAME_TH,
	(CASE
		WHEN (P.INVALID_PERIOD = 0 AND P.DECREASE_WINNER_TYPE = 1) THEN P.WINNER_COUNTER
		WHEN (P.INVALID_PERIOD = 0 AND P.DECREASE_WINNER_TYPE = 3) THEN PC.WINNER_COUNTER
		WHEN (P.INVALID_PERIOD != 0 AND P.DECREASE_WINNER_TYPE = 1) THEN S1.WINNER_COUNTER
		WHEN (P.INVALID_PERIOD != 0 AND P.DECREASE_WINNER_TYPE = 3) THEN S2.WINNER_COUNTER
		ELSE 99 END
	) QUOTA
	FROM PRIVUSER.PRIVILEGE_INFO PI
	INNER JOIN PRIVUSER.VOUCHER_TODAY_PRIORITY VT 
		ON (PI.PRIVILEGE_INFO_ID = VT.PRIVILEGE_INFO_ID AND VT.CHANNEL = 'HOMELOYALTY')
	INNER JOIN PRIVUSER.PROJECT P 
		ON (PI.PROJECT_ID = P.PROJECT_ID AND P.VOUCHER = 1)
	INNER JOIN PRIVUSER.BRAND B 
		ON (PI.BRAND_ID = B.BRAND_ID)
	LEFT JOIN PRIVUSER.PROJECT_CHOICE PC 
		ON (PI.PROJECT_ID = PC.PROJECT_ID AND PI.CHOICE_ID = PC.CHOICE_ID)
	LEFT JOIN PRIVUSER.RECURRING_SLOT S1
		ON (PI.PROJECT_ID = S1.PROJECT_ID AND GETDATE() BETWEEN S1.ACTIVATE_DATE AND S1.DEACTIVATE_DATE)
	LEFT JOIN PRIVUSER.RECURRING_SLOT S2 
		ON (PI.PROJECT_ID = S2.PROJECT_ID AND PI.CHOICE_ID = S2.CHOICE_ID AND GETDATE() BETWEEN S2.ACTIVATE_DATE AND S2.DEACTIVATE_DATE)
	WHERE PI.ACTIVE = 0 
		AND PI.DEFAULT_IMG IS NOT NULL
		AND GETDATE() BETWEEN PI.ACTIVATE_DATE AND PI.DEACTIVATE_DATE
	) R
WHERE LINE_NUMBER BETWEEN ? AND ? ORDER BY LINE_NUMBER `

	if err = r.db.Select(&connectDB, query, arg["pageNumber"], arg["resultPerPage"]); err != nil {
		return nil, err
	}
	return connectDB, nil
}

func (r voucherRepositoryDB) GetVoucherActive(arg map[string]interface{}) (connectDB []VoucherActive, err error) {
	query := ` 
SELECT * 
FROM ( 
	SELECT R.REG_ID 
	FROM PRIVUSER.REGISTER R, PRIVUSER.PROJECT P 
	WHERE P.PROJECT_ID = R.PROJECT_ID 
		AND P.VOUCHER = 1 
		AND R.MSISDN = ?
		AND R.REG_TIME > DATEADD(DAY, -90, GETDATE()) 
		AND R.USED_DATE IS NULL 
		AND R.EXPIRE_DATE > GETDATE() 
	UNION 
	SELECT PT.POINT_TRAN_ID 
	FROM PRIVUSER.POINT_TRANSACTION PT, PRIVUSER.PROJECT P 
	WHERE P.PROJECT_ID = PT.PROJECT_ID
		AND P.VOUCHER = 1 
		AND PT.MSISDN = ?
		AND PT.TRAN_DATE > DATEADD(DAY, -90, GETDATE()) 
		AND PT.USED_DATE IS NULL 
		AND PT.EXPIRE_DATE > GETDATE()
) active
	`
	if err = r.db.Select(&connectDB, query, arg["msisdn"], arg["msisdn"]); err != nil {
		return nil, err
	}
	return connectDB, nil
}
