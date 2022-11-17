package lib

import (
	"fmt"
	"strings"

	"github.com/go-playground/validator/v10"
)

var (
	Validate *validator.Validate = validator.New()
)

type (
	TemplateError struct {
		TransactionID string `json:"transactionID,omitempty"`
		HttpStatus    string `json:"httpStatus,omitempty"`
		Status        string `json:"status,omitempty"`
		Description   string `json:"description,omitempty"`
	}

	AppError struct {
		Code    int
		Message string
	}
)

func CheckReqParamValidate(errValidate, errBodyParser error, id string) *TemplateError {
	caseID := 0
	if errBodyParser != nil {
		LogInfoValidate(fmt.Sprintf("errBodeParser: %v", errBodyParser))
		typeStruct := strings.Split(fmt.Sprintf("%v", errBodyParser), " ")
		if typeStruct[0] == "json:" {
			caseID = 4
		} else if typeStruct[0] == "invalid" {
			caseID = 2
		} else if typeStruct[0] == "Unprocessable" {
			caseID = 1
		}
	}
	fieldErr := ""
	if errValidate != nil {
		if _, ok := errValidate.(*validator.InvalidValidationError); ok {
			LogInfoValidate(fmt.Sprintf("errValidate: %v", errBodyParser))
		}
		for _, err := range errValidate.(validator.ValidationErrors) {
			// fmt.Println(err.Field())
			fieldErr = err.Field()
			if err.Tag() != "" && caseID == 0 {
				caseID = 3
				break
			}
		}
	}
	errCode := 20000
	switch caseID {
	case 1:
		errCode = 40001
	case 2:
		errCode = 40002
	case 3:
		errCode = 40003
	case 4:
		errCode = 40004
	default:
		return nil
	}
	return ErrorRes(id, NewError(errCode, fieldErr))
}

func ErrorRes(id string, err error) *TemplateError {
	response := new(TemplateError)
	if id == "" {
		response.TransactionID = ""
	} else {
		response.TransactionID = id
	}

	responseErr, _ := err.(AppError)
	if responseErr.Code >= 20000 && responseErr.Code < 30000 {
		return nil
	} else if responseErr.Code >= 40000 && responseErr.Code < 50000 {
		response.HttpStatus = "400"
		switch responseErr.Code {
		case 40001:
			response.Status = "40001"
			response.Description = "Empty post body."
		case 40002:
			response.Status = "40002"
			response.Description = "Invalid json post body."
		case 40003:
			response.Status = "40003"
			response.Description = "Mandatory[ " + responseErr.Message + " ]."
		case 40004:
			response.Status = "40004"
			response.Description = "Invalid[ " + responseErr.Message + " ]."
		case 40301:
			response.HttpStatus = "403"
			response.Status = "40301"
			response.Description = "User unauthorized."
		}
	} else if responseErr.Code >= 50000 && responseErr.Code < 60000 {
		response.HttpStatus = "500"
		switch responseErr.Code {
		case 50001:
			response.Status = "50001"
			response.Description = "Have error occur while query data."
		case 50002:
			response.Status = "50002"
			response.Description = "Have error occur while updating data"
		case 51001:
			response.Status = "51001"
			response.Description = "Have error occur while calling CORE"
		}
	}
	return response
}

func (e AppError) Error() string {
	return e.Message
}

func NewError(code int, message string) error {
	return AppError{
		Code:    code,
		Message: message,
	}
}
