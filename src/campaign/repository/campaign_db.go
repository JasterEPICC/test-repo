package repository

import (
	"github.com/jmoiron/sqlx"
)

type campaignRepositoryDB struct {
	db *sqlx.DB
}

func NewCampaignRepositoryDB(connect *sqlx.DB) CampaignRepository {
	return campaignRepositoryDB{connect}
}

func (r campaignRepositoryDB) GetPrivToday(arg map[string]interface{}) (connectDB []PrivToday, err error) {
	query := `
SELECT *
FROM (
	SELECT PTD.EVENT_NAME_EN,
		PTD.EVENT_NAME_TH,
		B.BRAND_NAME_EN,  
		B.BRAND_NAME_TH, 
		FORMAT(PTD.ACTIVATE_DATE, 'dd/MM/yyyy hh:mm:ss') as ACTIVATE_DATE,
		FORMAT(PTD.DEACTIVATE_DATE, 'dd/MM/yyyy hh:mm:ss') as DEACTIVATE_DATE,
		PTD.MSG_EN,
		PTD.MSG_TH, 
		PTD.PRIVILEGE_INFO_ID, 
		PTD.EXT_URL, 
		PTD.IMAGE, 
		PI.POINTS,
		PTD.PRIORITY 
	FROM PRIVUSER.PRIVILEGE_TODAY PTD 				
	INNER JOIN PRIVUSER.INTERNAL_LIST L ON
		(PTD.GROUP_SEGMENT_ID=L.GROUP_SEGMENT_ID AND L.MSISDN = '08')  
	LEFT JOIN PRIVUSER.PRIVILEGE_INFO PI ON (PTD.PRIVILEGE_INFO_ID = PI.PRIVILEGE_INFO_ID) 
	LEFT JOIN PRIVUSER.BRAND B ON (PI.BRAND_ID = B.BRAND_ID) 
	WHERE PTD.ACTIVE=1 AND PTD.CHANNEL='HOMELOYALTY' 
		AND GETDATE() BETWEEN PTD.ACTIVATE_DATE AND PTD.DEACTIVATE_DATE 
	UNION 
	SELECT	PTD.EVENT_NAME_EN,
		PTD.EVENT_NAME_TH,
		B.BRAND_NAME_EN,
		B.BRAND_NAME_TH,
		FORMAT(PTD.ACTIVATE_DATE, 'dd/MM/yyyy hh:mm:ss') as ACTIVATE_DATE,
		FORMAT(PTD.DEACTIVATE_DATE, 'dd/MM/yyyy hh:mm:ss') as DEACTIVATE_DATE,
		PTD.MSG_EN,
		PTD.MSG_TH, 
		PTD.PRIVILEGE_INFO_ID, 
		PTD.EXT_URL,
		PTD.IMAGE, 
		PI.POINTS,
		PTD.PRIORITY  
	FROM PRIVUSER.PRIVILEGE_TODAY PTD 
	LEFT JOIN PRIVUSER.PRIVILEGE_INFO PI ON (PTD.PRIVILEGE_INFO_ID = PI.PRIVILEGE_INFO_ID)  
	LEFT JOIN PRIVUSER.BRAND B ON (PI.BRAND_ID = B.BRAND_ID)  
	WHERE PTD.GROUP_SEGMENT_ID = 0  
		AND PTD.ACTIVE=1 AND PTD.CHANNEL= ?
		AND GETDATE() >= CONVERT(DATETIME,PTD.ACTIVATE_DATE)  AND  GETDATE() <= CONVERT(DATETIME,PTD.DEACTIVATE_DATE)
) R 
ORDER BY R.PRIORITY `

	if err = r.db.Select(&connectDB, query, arg["channel"]); err != nil {
		return nil, err
	}
	return connectDB, nil
}

func (r campaignRepositoryDB) GetPrivRedeemHistory(arg map[string]interface{}) (connectDB []PrivRedeemHistory, err error) {
	query := `
SELECT  B.BRAND_NAME_TH, B.BRAND_NAME_EN, B.BARCODE_TYPE, 
		(SELECT TOP 1 HEADLINE_EN FROM PRIVUSER.PRIVILEGE_INFO PI WHERE PI.PROJECT_ID = R.PROJECT_ID AND PI.USSD_NO = M.MSG) HEADLINE_EN, 
		(SELECT TOP 1 HEADLINE_TH FROM PRIVUSER.PRIVILEGE_INFO PI WHERE PI.PROJECT_ID = R.PROJECT_ID AND PI.USSD_NO = M.MSG) HEADLINE_TH, 
		(SELECT TOP 1 DEFAULT_IMG FROM PRIVUSER.PRIVILEGE_INFO PI WHERE PI.PROJECT_ID = R.PROJECT_ID AND PI.USSD_NO = M.MSG) DEFAULT_IMG, 
		R.MSISDN, FORMAT(R.REG_TIME, 'dd/MM/yyyy hh:mm:ss') as REG_TIME,
		M.MSG_BARCODE, M.MSG_REPLY, M.EXT_URL, 
		C.CATEGORY_ID, C.CATEGORY, C.CATEGORY_TH 
FROM PRIVUSER.REGISTER R 
INNER JOIN PRIVUSER.CUSTOMER_PROFILE CUS ON (CUS.MSISDN = R.MSISDN) 
INNER JOIN PRIVUSER.MSG M ON (M.MSG_ID = R.MSG_ID AND M.SHORTCODE_ID = 61) 
INNER JOIN PRIVUSER.BRAND B ON (B.BRAND_ID = M.BRAND_ID) 
INNER JOIN PRIVUSER.PROJECT P ON (P.PROJECT_ID = R.PROJECT_ID AND (P.VOUCHER IS NULL OR P.VOUCHER = 0))
INNER JOIN PRIVUSER.CATEGORY C ON (P.CATEGORY = C.CATEGORY_ID) 
WHERE R.MSISDN = ?
	AND (R.REG_TIME  > DATEADD(DAY, -30, GETDATE())) 
	AND R.REFUND_STATUS IN (0,2,3) 
	AND CUS.REGISTER_DATE < R.REG_TIME
	AND (M.SUBMIT_TIME > DATEADD(DAY, -30, GETDATE())) 
ORDER BY R.REG_TIME DESC `
	if err = r.db.Select(&connectDB, query, arg["msisdn"]); err != nil {
		return nil, err
	}
	return connectDB, nil
}

func (r campaignRepositoryDB) GetCampaignRecommend(arg map[string]interface{}) (connectDB []CampaignRecommend, err error) {
	query := `
SELECT  PI.PRIVILEGE_INFO_ID, 
		PI.POINTS,  
		B.BRAND_NAME_EN, 
		B.BRAND_NAME_TH, 
		PI.HEADLINE_EN,
		PI.HEADLINE_TH,  
		PI.DEFAULT_IMG  
FROM PRIVUSER.PRIVILEGE_INFO PI 
INNER JOIN PRIVUSER.BRAND B ON (PI.BRAND_ID = B.BRAND_ID) 
INNER JOIN (  
		SELECT SOURCE.PIID, ORDER1.TOTAL+100 AS TOTAL, ORDER2.PRIORITY_CARD AS PRIORITY_CARD
		FROM (
			SELECT CUS.PRIVILEGE_INFO PIID  
			FROM PRIVUSER.NBO_CUSTOMER_REDEEM CUS  
			WHERE CUS.CATEGORY_TYPE = ? AND CUS.MSISDN = ?
			UNION   
			SELECT RP.PRIVILEGE_INFO_ID PIID  
			FROM PRIVUSER.RECOMMENDED_PRIVILEGE RP  
			WHERE RP.CATEGORY_TYPE = ?
		) SOURCE 
		LEFT JOIN PRIVUSER.NBO_CUSTOMER_REDEEM ORDER1 ON (SOURCE.PIID = ORDER1.PRIVILEGE_INFO AND ORDER1.MSISDN = ?)  
		LEFT JOIN PRIVUSER.RECOMMENDED_PRIVILEGE ORDER2 ON (SOURCE.PIID = ORDER2.PRIVILEGE_INFO_ID)  	
	) R ON (PI.PRIVILEGE_INFO_ID = R.PIID ) 
WHERE PI.ACTIVE = 0 AND PI.DEFAULT_IMG IS NOT NULL 
	AND GETDATE() BETWEEN PI.ACTIVATE_DATE AND PI.DEACTIVATE_DATE 
ORDER BY 2 DESC , 3 `
	if err = r.db.Select(&connectDB, query, arg["categoryType"], arg["msisdn"], arg["categoryType"], arg["msisdn"]); err != nil {
		return nil, err
	}
	return connectDB, nil
}

func (r campaignRepositoryDB) GetSerenadeExclusive(arg map[string]interface{}) (connectDB []SerenadeExclusive, err error) {
	query := `
SELECT 	PI.PRIVILEGE_INFO_ID, 
	PI.POINTS,
	B.BRAND_NAME_EN,
	B.BRAND_NAME_TH,
	PI.HEADLINE_EN, 
	PI.HEADLINE_TH,  
	PI.DEFAULT_IMG 
FROM PRIVUSER.PRIVILEGE_INFO PI 
INNER JOIN PRIVUSER.BRAND B ON (PI.BRAND_ID = B.BRAND_ID) 
INNER JOIN (
		SELECT (CASE
			WHEN (SE.PRIVILEGE_INFO_ID != 0) THEN SE.PRIVILEGE_INFO_ID 
			WHEN (SE.PRIVILEGE_INFO_ID = 0) THEN CAM.PRIVILEGE_INFO_ID 
			END) PRIVINFOID 
		FROM PRIVUSER.SERENADE_EXCLUSIVE SE 
		LEFT JOIN 
			(  
				SELECT ROW_NUMBER() OVER ( ORDER BY TOTAL DESC, CUS.CATEGORY_ID ) PRIORITY, CUS.CATEGORY_ID  
				FROM PRIVUSER.CUSTOMER_REDEEM_STAT CUS, PRIVUSER.CATEGORY C 
				WHERE CUS.CATEGORY_ID = C.CATEGORY_ID AND C.CATEGORY_TYPE = ? AND CUS.MSISDN = ? 
			) CUS ON (SE.CARD_TYPE = 'PERSONAL' AND SE.PRIORITY_CATEGORY = CUS.PRIORITY) 
		LEFT JOIN 
			( 
				SELECT C1.* FROM PRIVUSER.CAMPAIGN_REDEEM_STAT C1  
				INNER JOIN (  
				SELECT CATEGORY_ID,MAX(TOTAL) MX FROM PRIVUSER.CAMPAIGN_REDEEM_STAT GROUP BY CATEGORY_ID
				)C2 ON (C1.CATEGORY_ID = C2.CATEGORY_ID AND C1.TOTAL = C2.MX)  
			) CAM ON (CUS.CATEGORY_ID = CAM.CATEGORY_ID ) 
		WHERE SE.CATEGORY_TYPE = ?
	) RESULT ON (PI.PRIVILEGE_INFO_ID = RESULT.PRIVINFOID ) 
WHERE PI.ACTIVE = 0 AND PI.DEFAULT_IMG IS NOT NULL 
AND GETDATE() BETWEEN PI.ACTIVATE_DATE AND PI.DEACTIVATE_DATE 
ORDER BY RESULT.PRIVINFOID  
`
	if err = r.db.Select(&connectDB, query, arg["categoryType"], arg["msisdn"], arg["categoryType"]); err != nil {
		return nil, err
	}
	return connectDB, nil
}

func (r campaignRepositoryDB) GetNearByPrivilege(arg map[string]interface{}) (connectDB []NearByPrivilege, err error) {
	query := `
SELECT BRAND_ID, LATITUDE, LONGITUDE, RADIUS, DISTANCE
	USSD_NO, POINTS, CATEN, CATTH, SUB_CATEGORY_EN,
	SUB_CATEGORY_TH, CATEGORY_TYPE, HEADLINE_EN, HEADLINE_TH, BRAND_NAME_EN,
	BRAND_NAME_TH, BRAND_LOGO, LOCATION_USAGE_EN, LOCATION_USAGE_TH, CONDITION_EN,
	CONDITION_TH, CAMPAIGN_HILIGHT_EN, CAMPAIGN_HILIGHT_TH, FEATURE_EN, FEATURE_TH,
	PRIV_IMG_1, PRIV_IMG_2, PRIV_IMG_3, PRIV_IMG_4, PRIV_IMG_5,
	DEFAULT_IMG, URL, PRIORITY, BRANCH_ID, BRANCH_NAME_EN, BRANCH_NAME_TH,
	WONGNAI_ID,PRIVILEGE_INFO_ID,DISPLAY_CHANNEL,REDEEM_CHANNEL,
	SEGMENT, DESC_EN,DESC_TH, DESC_EMERALD_EN, DESC_EMERALD_TH,
	DESC_GOLD_EN, DESC_GOLD_TH, DESC_PLATINUM_EN, DESC_PLATINUM_TH, SERENADE_PRIVILEGE_EN,
	SERENADE_PRIVILEGE_TH, RIBBON_MESSAGE_EN, RIBBON_MESSAGE_TH, RIBBON_COLOR, VOUCHER
	,CAMPAIGN_TYPE
FROM(
	SELECT  P.CATEGORY_TYPE, P.CATEGORY_EN CATEN, P.CATEGORY_TH CATTH, P.BRAND_NAME_EN, P.BRAND_NAME_TH, P.BRAND_LOGO, P.BRANCH_ID,
		P.BRANCH_NAME_EN, P.BRANCH_NAME_TH, P.WONGNAI_ID, P.LATITUDE, P.LONGITUDE, PI.*, PJ.VOUCHER, D.RADIUS,
 		cast( ROUND( d.distance_unit * rad2deg * ( ACOS( ROUND( COS( D.deg2rad * ( d.latpoint ) ), 15, 1) *  ROUND( COS(D.deg2rad * (p.latitude) ),15,1) *  ROUND( COS(D.deg2rad * (d.longpoint - p.longitude) ),15,1) +  ROUND( SIN(D.deg2rad * (d.latpoint)),15,1) *ROUND(SIN(D.deg2rad * (p.latitude)),15,1))),3) as decimal(20,3)) AS distance
	FROM PRIVUSER.PRIVILEGE_INFO_NEARBY P , PRIVUSER.PRIVILEGE_INFO PI
	JOIN (
		SELECT ? AS LATPOINT,
			? AS LONGPOINT,
			? AS RADIUS,
			111.045 AS DISTANCE_UNIT,
			57.2957795 AS RAD2DEG, 0.0174532925 AS DEG2RAD
	) D 
	ON(1=1)
	INNER JOIN PRIVUSER.PROJECT PJ ON PI.PROJECT_ID = PJ.PROJECT_ID
	WHERE P.PRIVILEGE_INFO_ID = PI.PRIVILEGE_INFO_ID
	AND P.LATITUDE BETWEEN D.LATPOINT - (D.RADIUS / D.DISTANCE_UNIT) AND D.LATPOINT + (D.RADIUS / D.DISTANCE_UNIT)
	AND P.LONGITUDE BETWEEN D.LONGPOINT - (D.RADIUS / (D.DISTANCE_UNIT * COS(D.DEG2RAD * (D.LATPOINT)))) AND D.LONGPOINT + (D.RADIUS / (D.DISTANCE_UNIT * COS(D.DEG2RAD * (d.latpoint)))
) `

	haveBrandName := "	AND (LOWER(P.BRAND_NAME_TH) LIKE ? OR lower(p.BRAND_NAME_EN) like ?) "
	haveCategoryType := " AND P.CATEGORY_TYPE = ? "
	Where := "AND (PI.CAMPAIGN_TYPE != 'MERCHANT' OR PI.CAMPAIGN_TYPE IS NULL) ) poi  WHERE distance <= ? ORDER by distance "

	switch arg["Cases"] {
	case 1:
		query = query + haveBrandName + haveCategoryType + Where
		if err := r.db.Select(&connectDB, query, arg["Latitude"], arg["Longitude"], arg["Radius"], arg["BrandName"], arg["BrandName"], arg["CategoryType"], arg["Radius"]); err != nil {
			return nil, err
		}
	case 2:
		query = query + haveCategoryType + Where
		if err := r.db.Select(&connectDB, query, arg["Latitude"], arg["Longitude"], arg["Radius"], arg["CategoryType"], arg["Radius"]); err != nil {
			return nil, err
		}
	case 3:
		query = query + haveBrandName + Where
		if err := r.db.Select(&connectDB, query, arg["Latitude"], arg["Longitude"], arg["Radius"], arg["BrandName"], arg["BrandName"], arg["Radius"]); err != nil {
			return nil, err
		}
	case 4:
		query = query + Where
		if err := r.db.Select(&connectDB, query, arg["Latitude"], arg["Longitude"], arg["Radius"], arg["Radius"]); err != nil {
			return nil, err
		}
	}

	return connectDB, nil
}

func (r campaignRepositoryDB) GetPrivilegeInfo(privInfoId string) (*GetPrivilegeInfo, error) {
	connectDB := GetPrivilegeInfo{}
	query := ` 
SELECT I.PRIVILEGE_INFO_ID,  I.USSD_NO, I.POINTS, 
	I.CATEGORY_ID, C.CATEGORY, C.CATEGORY_TH, 
	C.CATEGORY_TYPE, I.BRAND_ID, B.BRAND_NAME_EN, 
	B.BRAND_NAME_TH, B.BRAND_LOGO, I.PRIVILEGE_INFO_IMAGE_ALL, 
	I.PRIORITY, I.WONGNAI_ID_ALL, I.DISPLAY_CHANNEL,
	I.REDEEM_CHANNEL, I.ACTIVATE_DATE, I.DEACTIVATE_DATE,
	I.RIBBON_COLOR, I.RIBBON_MESSAGE_EN, I.RIBBON_MESSAGE_TH,
	P.TOTAL_WINNER,
	(CASE 
		WHEN (p.Invalid_Period = 0 AND p.Decrease_Winner_Type = 1) THEN p.Winner_Counter 
		WHEN (p.Invalid_Period = 0 AND p.Decrease_Winner_Type = 3) THEN pc.Winner_Counter 
		WHEN (p.Invalid_Period != 0 AND p.Decrease_Winner_Type = 1) THEN s1.Winner_Counter 
		WHEN (p.Invalid_Period != 0 AND p.Decrease_Winner_Type = 3) THEN s2.Winner_Counter 
		ELSE 99 
	END) 
	quotaRemain, i.Sticker_Type, p.Voucher, 
	P.Voucher_Expire_Date, i.Segment, i.Headline_En, 
	i.Headline_Th, i.Location_Usage_En, i.Location_Usage_Th, 
	i.Condition_En, i.Condition_Th, i.Campaign_Hilight_En, 
	i.Campaign_Hilight_Th, i.Feature_En, i.Feature_Th, 
	i.Desc_En, i.Desc_Th, i.Desc_Emerald_En, 
	i.Desc_Emerald_Th, i.Desc_Gold_En, i.Desc_Gold_Th, 
	i.Desc_Platinum_En, i.Desc_Platinum_Th 
FROM PRIVUSER.PRIVILEGE_INFO AS I
INNER JOIN PRIVUSER.CATEGORY AS C ON I.CATEGORY_ID = C.CATEGORY_ID
INNER JOIN PRIVUSER.BRAND AS B ON I.BRAND_ID = B.BRAND_ID
INNER JOIN PRIVUSER.PROJECT AS P ON I.PROJECT_ID = P.PROJECT_ID
LEFT JOIN PRIVUSER.PROJECT_CHOICE AS PC ON I.PROJECT_ID = PC.PROJECT_ID AND I.CHOICE_ID = PC.CHOICE_ID
LEFT JOIN PRIVUSER.RECURRING_SLOT S1 ON I.PROJECT_ID = S1.PROJECT_ID AND 
	GETDATE() BETWEEN S1.ACTIVATE_DATE AND S1.DEACTIVATE_DATE
LEFT JOIN PRIVUSER.RECURRING_SLOT S2 ON I.PROJECT_ID = S2.PROJECT_ID AND I.CHOICE_ID = S2.CHOICE_ID AND GETDATE() BETWEEN S2.ACTIVATE_DATE AND S2.DEACTIVATE_DATE
WHERE I.ACTIVE = 0
	AND I.PRIVILEGE_INFO_IMAGE_ALL IS NOT NULL
	AND GETDATE() BETWEEN I.ACTIVATE_DATE AND I.DEACTIVATE_DATE
	AND I.PRIVILEGE_INFO_ID = ? `

	// SELECT q.*
	// from (
	//      select ROW_NUMBER() over (order by PRIV_IMG_1) as r ,n.*
	//      FROM (
	//          select (
	//              SELECT string_agg(pim.FILE_NAME,', ') as elements
	//              FROM (
	//                  select * , row_number() over (order by privilege_info_image_id desc) as img_order
	//                  from privuser.PRIVILEGE_INFO_IMAGE px
	//             ) pim
	//                   where pim.privilege_info_id = m.privilege_info_id
	//               ) as PRIV_IMG_1,
	//             (SELECT string_agg(pim.priority,',') elements
	//               FROM (
	//                   select * ,row_number() over (order by privilege_info_image_id desc) as img_order
	//                   from privuser.PRIVILEGE_INFO_IMAGE px
	//               ) pim
	//               where pim.privilege_info_id = m.privilege_info_id
	//           ) as PRIV_IMG_2,
	//           m.PRIV_IMG_3,m.PRIV_IMG_4,m.PRIV_IMG_5,m.PRIVILEGE_INFO_ID,m.USSD_NO,m.DESC_EN,m.DESC_TH,
	//           m.PRIORITY,m.URL,m.BIZ_ID,m.MODIFIED_BY,m.MODIFIED_DATE,m.SHORTCODE_ID,m.SERVICE_ID,m.PROJECT_ID,
	//           m.CHOICE_ID,m.START_DATE,m.ACTIVATE_DATE,m.DEACTIVATE_DATE,m.STOP_DATE,m.CATEGORY_TH,m.SUB_CATEGORY_TH,m.BIZ_NAME,m.CONDITION_TH,
	//           m.DEFAULT_IMG,m.ACTIVE,m.CONDITION_EN,m.SERVICE_TYPE,m.CATEGORY_ID,m.CREATED_DATE,m.CAMPAIGN_ID, m.CAMPAIGN_STATUS ,m.REDEEM_CHANNEL,m.DISPLAY_CHANNEL
	//           from privuser.PRIVILEGE_INFO m, privuser.PRIVILEGE_INFO_IMAGE mm
	//           where m.privilege_info_id = mm.privilege_info_id
	//             and m.activate_date <= getdate()
	//             and m.deactivate_date >= getdate()
	//             and m.active = 0
	//         group by m.PRIVILEGE_INFO_ID,m.USSD_NO,m.DESC_EN,m.DESC_TH,
	//             m.PRIORITY,m.PRIV_IMG_3,m.PRIV_IMG_4,m.PRIV_IMG_5,m.URL,m.BIZ_ID,m.MODIFIED_BY,m.MODIFIED_DATE,m.SHORTCODE_ID,m.SERVICE_ID,m.PROJECT_ID,
	//             m.CHOICE_ID,m.START_DATE,m.ACTIVATE_DATE,m.DEACTIVATE_DATE,m.STOP_DATE,m.CATEGORY_TH,m.SUB_CATEGORY_TH,m.BIZ_NAME,m.CONDITION_TH,
	//             m.DEFAULT_IMG,m.ACTIVE,m.CONDITION_EN,m.SERVICE_TYPE,m.CATEGORY_ID,m.CREATED_DATE,m.CAMPAIGN_ID,m.CAMPAIGN_STATUS,m.REDEEM_CHANNEL,m.DISPLAY_CHANNEL
	//     ) n
	// ) q
	// where r Between 10 and 200

	if err := r.db.Get(&connectDB, query, privInfoId); err != nil {
		return nil, err
	}

	return &connectDB, nil
}
