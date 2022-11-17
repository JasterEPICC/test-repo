package handler

import "privilege-api-myais/src/auth/service"

type templateJwt struct {
	Username    string `json:"user_name"`
	Token       string `json:"token,omitempty"`
	Description string `json:"description,omitempty"`
}

func ResponseJwt(res *service.UserResponse) *templateJwt {

	response := new(templateJwt)
	response.Username = res.UserName
	if res.Token != "" {
		response.Token = res.Token
	} else if res.Description != "" {
		response.Token = res.Description
	}

	return response
}
