package handler

import (
	"privilege-api-myais/src/campaign/service"
)

type templatePrivToday struct {
	TransactionID string                      `json:"transactionID"`
	Status        string                      `json:"status"`
	Description   string                      `json:"description"`
	PrivTodayArr  []service.PrivTodayResponse `json:"privTodayArr"`
}

func ResponsePrivToday(id string, arg []service.PrivTodayResponse) *templatePrivToday {
	response := new(templatePrivToday)
	response.TransactionID = id
	if len(arg) != 0 {
		response.Status = "20000"
		response.Description = "Success"
	} else {
		response.Status = "20001"
		response.Description = "Data not found"
	}
	response.PrivTodayArr = arg

	return response

}

type ()

func ResponseCampaign(id string, arg *service.GetCampaignResponse) *service.GetCampaignResponse {

	if arg != nil {
		arg.TransactionID = id
		arg.Status = "2000"
		arg.Description = "SUCCESS"
		return arg
	} else {
		res := service.GetCampaignResponse{
			TransactionID: id,
			Status:        "20001",
			Description:   "Data not found",
		}
		return &res
	}

}

func ResponsePrivRedeemHistory(id string, arg *service.PrivRedeemHistoryResponse) *service.PrivRedeemHistoryResponse {

	if arg.ResultAvailable != 0 {
		arg.TransactionID = id
		arg.HttpStatus = "200"
		arg.Status = "20000"
		arg.Description = "SUCCESS"
	} else {
		arg.TransactionID = id
		arg.HttpStatus = "200"
		arg.Status = "20000"
		arg.Description = "Data not found"
	}

	return arg
}

type templateCampaignRecommend struct {
	TransactionID string                              `json:"transactionID"`
	Status        string                              `json:"status"`
	Description   string                              `json:"description"`
	Count         int                                 `json:"count"`
	PrivilegeArr  []service.CampaignRecommendResponse `json:"privilegeArr"`
}

func ResponseCampaignRecommend(id string, arg []service.CampaignRecommendResponse) *templateCampaignRecommend {
	response := new(templateCampaignRecommend)
	response.TransactionID = id
	if len(arg) != 0 {
		response.Status = "20000"
		response.Description = "SUCCESS"
	} else {
		response.Status = "20001"
		response.Description = "Data not found."
	}
	response.Count = len(arg)
	response.PrivilegeArr = arg

	return response
}

type templateSerenadeExclusive struct {
	TransactionID string                              `json:"transactionID"`
	Status        string                              `json:"status"`
	Description   string                              `json:"description"`
	Count         int                                 `json:"count"`
	PrivilegeArr  []service.SerenadeExclusiveResponse `json:"privilegeArr"`
}

func ResponseSerenadeExclusive(id string, arg []service.SerenadeExclusiveResponse) *templateSerenadeExclusive {
	response := new(templateSerenadeExclusive)
	response.TransactionID = id
	if len(arg) != 0 {
		response.Status = "20000"
		response.Description = "SUCCESS"
	} else {
		response.Status = "20001"
		response.Description = "Data not found."
	}
	response.Count = len(arg)
	response.PrivilegeArr = arg

	return response
}

type templateNearByPrivilege struct {
	TransactionID       string                            `json:"transactionID"`
	HttpStatus          string                            `json:"httpStatus"`
	Status              string                            `json:"status"`
	Description         string                            `json:"description"`
	PageNumber          int                               `json:"pageNumber"`
	TotalPage           int                               `json:"totalPage"`
	ResultPerPage       int                               `json:"resultPerPage"`
	ResultAvailable     int                               `json:"resultAvailable"`
	TotalResultReturned int                               `json:"totalResultReturned"`
	NearbyLocationArr   []service.NearByPrivilegeResponse `json:"nearbyLocationArr"`
}

func ResponseNearByPrivilege(id string, pageNumber, resultPerPage int, arg []service.NearByPrivilegeResponse) *templateNearByPrivilege {

	response := new(templateNearByPrivilege)
	response.TransactionID = id
	if len(arg) != 0 {
		response.HttpStatus = "200"
		response.Status = "20000"
		response.Description = "Success"
		response.PageNumber = int(pageNumber)
		response.TotalPage = len(arg) / int(resultPerPage)
		response.ResultPerPage = int(resultPerPage)
		response.ResultAvailable = len(arg)
		startRow := (resultPerPage * (pageNumber - 1))
		endRow := (resultPerPage * pageNumber)

		if (len(arg) % int(resultPerPage)) != 0 {
			response.TotalPage++
		}

		if endRow > len(arg) {
			endRow = len(arg)
		}

		if startRow < endRow {
			response.NearbyLocationArr = arg[startRow:endRow]
			response.TotalResultReturned = endRow - startRow
		}

	} else {
		response.HttpStatus = "200"
		response.Status = "20002"
		response.Description = "Data not found"
		response.PageNumber = int(pageNumber)
		response.TotalPage = 1
		response.ResultPerPage = int(resultPerPage)
		response.ResultAvailable = len(arg)
		response.TotalResultReturned = len(arg)
	}

	return response
}
