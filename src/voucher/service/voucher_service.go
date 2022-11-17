package service

import (
	"fmt"
	"privilege-api-myais/lib"
	"privilege-api-myais/src/voucher/repository"
)

type voucherService struct {
	voucherRepository repository.VoucherRepository
}

func NewVoucherService(connect repository.VoucherRepository) VoucherService {
	return voucherService{connect}
}

func (s voucherService) GetVoucher(msisdn string) (res []VoucherResponse, err error) {
	voucher, errVoucher := s.voucherRepository.GetVoucher(msisdn)
	if errVoucher != nil {
		lib.LogInfoSQL(fmt.Sprintf("errVoucher: %v", errVoucher))
		return nil, lib.NewError(50001, "Have error occur while query data.")
	}

	for _, detail := range voucher {
		res = append(res, VoucherResponse{
			Msisdn:                detail.Msisdn.String,
			VoucherType:           detail.VoucherType.String,
			VoucherID:             detail.VoucherID.String,
			ExpireDate:            detail.ExpireDate.String,
			CategoryEn:            detail.CategoryEn.String,
			CategoryTh:            detail.CategoryTh.String,
			HeadlineEn:            detail.HeadlineEn.String,
			HeadlineTh:            detail.HeadlineTh.String,
			DescEn:                detail.DescEn.String,
			DescTh:                detail.DescTh.String,
			DescEmeraldEn:         detail.DescEmeraldEn.String,
			DescEmeraldTh:         detail.DescEmeraldEn.String,
			DescGoldEn:            detail.DescGoldEn.String,
			DescGoldTh:            detail.DescGoldTh.String,
			DescPlatinumEn:        detail.DescPlatinumEn.String,
			DescPlatinumTh:        detail.DescPlatinumTh.String,
			BrandNameEn:           detail.BrandNameEn.String,
			BrandNameTh:           detail.BrandNameTh.String,
			Condition_En:          detail.Condition_En.String,
			Condition_th:          detail.Condition_th.String,
			ActivateDate:          detail.ActivateDate.String,
			DeactivateDate:        detail.DeactivateDate.String,
			PrivilegeInfoID:       detail.PrivilegeInfoID.String,
			BrandLogo:             detail.BrandLogo.String,
			PrivilegeInfoImageAll: detail.PrivilegeInfoImageAll.String,
		})
	}

	return res, nil
}

func (s voucherService) GetRedeemVoucher(msisdn, voucherType, voucherId string) (res []RedeemVoucherResponse, err error) {
	req := map[string]interface{}{
		"msisdn":      msisdn,
		"voucherType": voucherType,
		"voucherId":   voucherId,
	}

	redeemVoucher, errRedeemVoucher := s.voucherRepository.GetRedeemVoucher(req)
	if errRedeemVoucher != nil {
		lib.LogInfoSQL(fmt.Sprintf("errRedeemVoucher: %v", errRedeemVoucher))
		return nil, lib.NewError(50001, "Have error occur while query data.")
	}
	for _, detail := range redeemVoucher {
		res = append(res, RedeemVoucherResponse{
			RefundStatus: detail.RefundStatus.String,
			MsgReply:     detail.MsgReply.String,
			Msg_Barcode:  detail.Msg_Barcode.String,
			BarcodeType:  detail.BarcodeType.String,
			ExpireDate:   detail.ExpireDate.String,
			Msg:          detail.Msg.String,
		})
	}

	return res, nil
}

func (s voucherService) GetVoucherToday(pageNumber, resultPerPage int) (res []VoucherTodayResponse, err error) {
	req := map[string]interface{}{
		"pageNumber":    ((pageNumber - 1) * resultPerPage),
		"resultPerPage": pageNumber * resultPerPage,
	}
	voucherToday, errVoucherToday := s.voucherRepository.GetVoucherToday(req)
	if errVoucherToday != nil {
		lib.LogInfoSQL(fmt.Sprintf("error: %v", errVoucherToday))
		return nil, lib.NewError(50001, "Have error occur while query data.")
	}

	for _, detail := range voucherToday {
		res = append(res, VoucherTodayResponse{
			PrivilegeInfoID: int(detail.PrivilegeInfoID.Int64),
			Points:          int(detail.Points.Int64),
			BrandNameEn:     detail.BrandNameEn.String,
			BrandNameTh:     detail.BrandNameTh.String,
			HeadlineEn:      detail.HeadlineEn.String,
			HeadlineTh:      detail.HeadlineTh.String,
			DefaultImg:      detail.DefaultImg.String,
			DeactivateDate:  detail.DeactivateDate.String,
			Quota:           int(detail.Quota.Int64),
		})
	}

	return res, nil
}

func(s voucherService) GetVoucherActive(msisdn string)(res []VoucherActiveResponse,err error){
	req := map[string]interface{}{
		"msisdn":      msisdn,
	}
	voucherActive , errVoucherActive := s.voucherRepository.GetVoucherActive(req)
	if errVoucherActive != nil{
		lib.LogInfoSQL(fmt.Sprintf("errVoucherActive: %v", errVoucherActive))
		return nil, lib.NewError(50001, "Have error occur while query data.")
	}
	for _, detail := range voucherActive{
		res = append(res,VoucherActiveResponse{
			RegID: detail.RegID.String,
		})
	}
	return res, nil
}