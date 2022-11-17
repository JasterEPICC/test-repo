package service

import (
	"database/sql"
	"fmt"
	"privilege-api-myais/lib"
	"privilege-api-myais/src/campaign/repository"
	"strconv"
	"strings"
	"time"
)

type campaignService struct {
	campaignRepository repository.CampaignRepository
}

func NewCampaignService(connect repository.CampaignRepository) CampaignService {
	return campaignService{connect}
}

func (s campaignService) GetPrivToday(msisdn, channel string) (res []PrivTodayResponse, err error) {
	req := map[string]interface{}{
		"msisdn":  msisdn,
		"channel": channel,
	}
	PrivToday, errPrivToday := s.campaignRepository.GetPrivToday(req)
	if errPrivToday != nil {
		lib.LogInfoSQL(fmt.Sprintf("errPrivToday: %v", errPrivToday))
		return nil, lib.NewError(50001, "Have error occur while query data.")
	}
	for _, detail := range PrivToday {
		res = append(res, PrivTodayResponse{
			EventNameEn:     detail.BrandNameEn.String,
			EventNameTh:     detail.EventNameTh.String,
			BrandNameEn:     detail.BrandNameEn.String,
			BrandNameTh:     detail.BrandNameTh.String,
			ActivateDate:    detail.ActivateDate.String,
			DeactivateDate:  detail.DeactivateDate.String,
			MsgEn:           detail.MsgEn.String,
			MsgTh:           detail.MsgTh.String,
			PrivilegeInfoID: detail.PrivilegeInfoID.String,
			Ext_url:         detail.Ext_url.String,
			Image:           detail.Image.String,
			Points:          detail.Points.String,
		})
	}
	return res, nil
}

func (s campaignService) GetPrivilegeInfo(privInfoId string) (*GetCampaignResponse, error) {

	connectDB, err := s.campaignRepository.GetPrivilegeInfo(privInfoId)
	if err != nil {
		fmt.Println("Error HERE ???", err, "||", connectDB)
		lib.LogInfoSQL(fmt.Sprintf("GetPrivilageInfo: %v", err))

		if err == sql.ErrNoRows {
			return nil, lib.NewError(20001, "data not found")
		} else {
			return nil, lib.NewError(50001, "data not found")
		}
	}

	// req := map[string]interface{}{
	// 	"page" : 1 ,

	// }

	res := GetCampaignResponse{}
	res.PrivInfoID = connectDB.PrivilegeInfoID.Int64
	res.UssdNo = connectDB.UssdNo.String
	res.Points = connectDB.Points.Int64
	res.CategoryID = connectDB.CategoryID.Int64
	res.CategoryNameEn = connectDB.Category.String
	res.CategoryNameTh = connectDB.CategoryTh.String
	res.CategoryType = connectDB.CategoryType.String
	res.BrandID = connectDB.BrandID.Int64
	res.BrandNameEn = connectDB.BrandNameEn.String
	res.BrandNameTh = connectDB.BrandNameTh.String
	res.BrandLogo = connectDB.BrandLogo.String
	res.ImageAll = connectDB.PrivilegeInfoImageAll.String
	res.Priority = connectDB.Priority.Int64
	res.WongnaiID = connectDB.WongnaiIDAll.String
	res.URL = connectDB.URL.String
	res.DisplayChannel = connectDB.DisplayChannel.String
	res.RedeemChannel = connectDB.RedeemChannel.String
	res.ActivateDate = connectDB.ActivateDate.String
	res.DeactivateDate = connectDB.DeactivateDate.String
	res.RibbonColor = connectDB.RibbonColor.String
	res.RibbonMsgEn = connectDB.RibbonMessageEn.String
	res.RibbonMsgTh = connectDB.RibbonMessageTh.String
	res.Quota = connectDB.TotalWinner.Int64
	res.QuotaRemain = connectDB.QuotaRemain.Int64
	res.StickerType = connectDB.StickerType.String
	res.Voucher = connectDB.Voucher.Int64
	res.VoucherExpire = connectDB.VoucherExpire.String
	res.Segment = connectDB.Segment.String
	res.HeadLineEn = connectDB.HeadlineEn.String
	res.HeadLineTh = connectDB.HeadlineTh.String
	res.LocationEN = connectDB.LocationUsageEn.String
	res.LocationTH = connectDB.LocationUsageTh.String
	res.ConditionEN = connectDB.ConditionEn.String
	res.ConditionTH = connectDB.ConditionTh.String
	res.HilightEn = connectDB.CampaignHilightEn.String
	res.HilightTh = connectDB.CampaignHilightTh.String
	res.FeatureEn = connectDB.Feature_En.String
	res.FeatureTh = connectDB.Feature_Th.String
	res.DescEn = connectDB.DescEn.String
	res.DescTh = connectDB.DescTh.String
	res.DescEmeraldEn = connectDB.DescEmeraldEn.String
	res.DescEmeraldTh = connectDB.DescEmeraldTh.String
	res.DescGoldEn = connectDB.DescGoldEn.String
	res.DescGoldTh = connectDB.DescGoldTh.String
	res.DescPlatinumEn = connectDB.DescPlatinumEn.String
	res.DescPlatinumTh = connectDB.DescPlatinumTh.String

	return &res, nil
}

func (s campaignService) GetPrivRedeemHistory(msisdn string, pageNumber, resultPerPage int) (*PrivRedeemHistoryResponse, error) {
	if (resultPerPage < 1 || resultPerPage > 50) || pageNumber < 1 {
		if pageNumber < 1 {
			return nil, lib.NewError(40006, "pageNumber must be more than 1.")
		} else {
			return nil, lib.NewError(40006, "resultPerPage must be between 1 and 50.")
		}
	}

	req := map[string]interface{}{
		"msisdn": msisdn,
	}

	if req["pageNumber"] = 1; pageNumber != 0 {
		req["pageNumber"] = pageNumber
	}

	if req["resultPerPage"] = 10; resultPerPage != 0 {
		req["resultPerPage"] = resultPerPage
	}

	privRedeemHistory, errPrivRedeemHistory := s.campaignRepository.GetPrivRedeemHistory(req)
	if errPrivRedeemHistory != nil {
		lib.LogInfoSQL(fmt.Sprintf("errPrivRedeemHistory: %v", errPrivRedeemHistory))
		return nil, lib.NewError(50001, "Have error occur while query data.")
	}

	privArr := []PrivilegeRedeemHistoryArrResponse{}
	for _, detail := range privRedeemHistory {
		privArr = append(privArr, PrivilegeRedeemHistoryArrResponse{
			BrandNameTh: detail.BrandNameTh.String,
			BrandNameEn: detail.BrandNameEn.String,
			BarcodeType: detail.BarcodeType.String,
			DescEn:      detail.Headline_En.String,
			DescTh:      detail.Headline_Th.String,
			DefaultImg:  detail.Default_Img.String,
			Msisdn:      detail.Msisdn.String,
			TranDate:    detail.RegTime.String,
			MsgBarcode:  detail.Msg_Barcode.String,
			MsgReply:    detail.MsgReply.String,
		})
	}

	res := PrivRedeemHistoryResponse{
		PageNumber:      req["pageNumber"].(int),
		ResultPerPage:   req["resultPerPage"].(int),
		TotalPage:       len(privArr) / req["resultPerPage"].(int),
		ResultAvailable: len(privArr),
		SystemTime:      time.Now().Format("2006/01/02 15:04:05"),
	}

	if (len(privArr) % req["resultPerPage"].(int)) != 0 {
		res.TotalPage++
	}

	startRow := (resultPerPage * (pageNumber - 1))
	endRow := (resultPerPage * pageNumber)

	if endRow > len(privArr) {
		endRow = len(privArr)
	}

	if startRow < endRow {
		res.PrivilegeRedeemHistoryArr = privArr[startRow:endRow]
		res.TotalResultReturned = endRow - startRow
	}

	return &res, nil
}

func (s campaignService) GetCampaignRecommend(msisdn, categoryType string) (res []CampaignRecommendResponse, err error) {
	req := map[string]interface{}{
		"msisdn":       msisdn,
		"categoryType": categoryType,
	}

	campaignRecommend, errCampaignRecommend := s.campaignRepository.GetCampaignRecommend(req)
	if errCampaignRecommend != nil {
		lib.LogInfoSQL(fmt.Sprintf("errCampaignRecommend: %v", errCampaignRecommend))
		return nil, lib.NewError(50001, "Have error occur while query data.")
	}

	for _, detail := range campaignRecommend {
		res = append(res, CampaignRecommendResponse{
			PrivilegeInfoID: detail.PrivilegeInfoID.String,
			Points:          detail.Points.String,
			BrandNameEn:     detail.BrandNameEn.String,
			BrandNameTh:     detail.BrandNameTh.String,
			HeadlineEn:      detail.HeadlineEn.String,
			HeadlineTh:      detail.HeadlineTh.String,
			DefaultImg:      detail.DefaultImg.String,
		})
	}

	return res, nil
}

func (s campaignService) GetSerenadeExclusive(msisdn, categoryType string) (res []SerenadeExclusiveResponse, err error) {
	req := map[string]interface{}{
		"msisdn":       msisdn,
		"categoryType": categoryType,
	}

	SerenadeExclusive, errSerenadeExclusive := s.campaignRepository.GetSerenadeExclusive(req)
	if errSerenadeExclusive != nil {
		lib.LogInfoSQL(fmt.Sprintf("errSerenadeExclusive: %v", errSerenadeExclusive))
		return nil, lib.NewError(50001, "Have error occur while query data.")
	}

	for _, detail := range SerenadeExclusive {
		res = append(res, SerenadeExclusiveResponse{
			PrivilegeInfoID: detail.PrivilegeInfoID.String,
			Points:          detail.Points.String,
			BrandNameEn:     detail.BrandNameEn.String,
			BrandNameTh:     detail.BrandNameTh.String,
			HeadlineEn:      detail.HeadlineEn.String,
			HeadlineTh:      detail.HeadlineTh.String,
			DefaultImg:      detail.DefaultImg.String,
		})
	}

	return res, nil
}

func (s campaignService) GetNearByPrivilege(name, catType, latitude, longitude string, radius float64, pageNumber, resultPerPage int) (res []NearByPrivilegeResponse, err error) {

	caseQuery := 0
	if name != "" {
		if catType != "" {
			caseQuery = 1
		} else {
			caseQuery = 2
		}
	} else {
		if catType != "" {
			caseQuery = 3
		} else {
			caseQuery = 4
		}
	}

	latitudeCon, _ := strconv.ParseFloat(latitude, 64)
	longitudeCon, _ := strconv.ParseFloat(longitude, 64)

	req := map[string]interface{}{
		"BrandName":    "%" + name + "%",
		"CategoryType": catType,
		"Latitude":     latitudeCon,
		"Longitude":    longitudeCon,
		"Cases":        caseQuery,
	}

	if req["Radius"] = 3; radius != 0 {
		req["Radius"] = radius
	}

	nearByPrivilege, errNearByPrivilege := s.campaignRepository.GetNearByPrivilege(req)
	if errNearByPrivilege != nil {
		lib.LogInfoSQL(fmt.Sprintf("error near by privilege: %v", errNearByPrivilege))
		return nil, lib.NewError(50001, "Have error occur while query data.")
	}
	if len(nearByPrivilege) < 10 {
		i := 2.00
		for {
			nearByPrivilege, errNearByPrivilege = s.campaignRepository.GetNearByPrivilege(req)
			if errNearByPrivilege != nil {
				lib.LogInfoSQL(fmt.Sprintf("error near by privilege: %v", errNearByPrivilege))
				return nil, lib.NewError(50001, "Have error occur while query data.")
			}
			if len(nearByPrivilege) >= 10 {
				break
			} else {
				req["Radius"] = radius * i
				i++
			}
		}
	}

	for _, detail := range nearByPrivilege {
		segmentArr := []NearByPrivilegeDesArrResponse{}
		sec := strings.Split(detail.Segment.String, ",")
		for _, segmentDetail := range sec {
			switch segmentDetail {
			case "All":
				segmentArr = append(segmentArr, NearByPrivilegeDesArrResponse{
					Segment: "All",
					DescEN:  detail.DescEn.String,
					DescTH:  detail.DescTh.String,
				})
			case "Classic":
				segmentArr = append(segmentArr, NearByPrivilegeDesArrResponse{
					Segment: "Mass",
					DescEN:  detail.DescEn.String,
					DescTH:  detail.DescTh.String,
				})
			case "SerenadePriv":
				segmentArr = append(segmentArr, NearByPrivilegeDesArrResponse{
					Segment: "SerenadePrivilege",
					DescEN:  detail.SerenadePrivilegeEn.String,
					DescTH:  detail.SerenadePrivilegeTh.String,
				})
			case "Emerald":
				segmentArr = append(segmentArr, NearByPrivilegeDesArrResponse{
					Segment: "Emerald",
					DescEN:  detail.SerenadePrivilegeEn.String,
					DescTH:  detail.SerenadePrivilegeTh.String,
				})
			case "Gold":
				segmentArr = append(segmentArr, NearByPrivilegeDesArrResponse{
					Segment: "Gold",
					DescEN:  detail.SerenadePrivilegeEn.String,
					DescTH:  detail.SerenadePrivilegeTh.String,
				})
			case "Platinum":
				segmentArr = append(segmentArr, NearByPrivilegeDesArrResponse{
					Segment: "Gold",
					DescEN:  detail.SerenadePrivilegeEn.String,
					DescTH:  detail.SerenadePrivilegeTh.String,
				})
			case "SerenadeAll":
				all := []string{"SerenadePrivilege", "Emerald", "Gold", "Platinum"}
				for _, addSec := range all {
					segmentArr = append(segmentArr, NearByPrivilegeDesArrResponse{
						Segment: addSec,
						DescEN:  detail.SerenadePrivilegeEn.String,
						DescTH:  detail.SerenadePrivilegeTh.String,
					})
				}
			}
		}

		res = append(res, NearByPrivilegeResponse{
			BrandID:         detail.Latitude.String,
			Latitude:        detail.Latitude.String,
			Longitude:       detail.Longitude.String,
			Distance:        detail.Distance.String,
			Radius:          detail.Radius.String,
			UssdNo:          detail.UssdNo.String,
			Points:          detail.Points.String,
			CategoryEN:      detail.Caten.String,
			CategoryTH:      detail.Catth.String,
			SubCategoryEN:   detail.SubCategoryEn.String,
			SubCategoryTH:   detail.SubCategoryTh.String,
			CategoryType:    detail.CategoryType.String,
			HeadLineEN:      detail.HeadlineEn.String,
			HeadLineTH:      detail.HeadlineTh.String,
			BrandNameEN:     detail.BrandNameEn.String,
			BrandNameTH:     detail.BrandNameTh.String,
			BrandLogo:       detail.BrandLogo.String,
			LocationEN:      detail.LocationUsageEn.String,
			LocationTH:      detail.LocationUsageTh.String,
			ConditionEN:     detail.ConditionEn.String,
			ConditionTH:     detail.ConditionTh.String,
			Segment:         detail.Segment.String,
			FeatureEN:       detail.FeatureEn.String,
			FeatureTH:       detail.FeatureTh.String,
			PrivImg1:        detail.PrivImg1.String,
			PrivImg2:        detail.PrivImg2.String,
			PrivImg3:        detail.PrivImg3.String,
			PrivImg4:        detail.PrivImg4.String,
			PrivImg5:        detail.PrivImg5.String,
			PrivImg6:        "",
			PrivImg7:        "",
			PrivImg8:        "",
			PrivImg9:        "",
			PrivImg10:       "",
			DefaultImg:      detail.DefaultImg.String,
			URL:             detail.URL.String,
			BranchID:        detail.BranchID.String,
			BranchNameEN:    detail.BranchNameEn.String,
			BranchNameTH:    detail.BranchNameTh.String,
			WongnaiID:       detail.WongnaiID.String,
			PrivilegeInfoID: detail.PrivilegeInfoID.String,
			DisplayChannel:  detail.DisplayChannel.String,
			RedeemChannel:   detail.RedeemChannel.String,
			CustomerDescEN:  detail.DescEn.String,
			CustomerDescTH:  detail.DescTh.String,
			RibbonMsgEN:     detail.RibbonMessageEn.String,
			RibbonMsgTH:     detail.RibbonMessageTh.String,
			RibbonMsgColor:  detail.RibbonColor.String,
			Voucher:         detail.Voucher.String,
			TotalDesc:       detail.Segment.String,
			DescriptionArr:  segmentArr,
		})
	}

	return res, nil
}
