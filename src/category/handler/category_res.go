package handler

import "privilege-api-myais/src/category/service"

type templateCategory struct {
	TransactionID string                        `json:"transactionID"`
	Status        string                        `json:"status"`
	Description   string                        `json:"description"`
	Count         int                           `json:"count"`
	CategoryArr   []service.GetCategoryResponse `json:"categoryArr"`
}

func ResponseCategory(id string, arg []service.GetCategoryResponse) *templateCategory {

	response := new(templateCategory)
	response.TransactionID = id
	if len(arg) != 0 {
		response.Status = "20000"
		response.Description = "SUCCESS"
	} else {
		response.Status = "20001"
		response.Description = "Data not found."
	}
	response.CategoryArr = arg
	response.Count = len(arg)

	return response
}
