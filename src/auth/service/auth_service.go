package service

import (
	"fmt"
	"privilege-api-myais/lib"
	"privilege-api-myais/src/auth/repository"
	"time"

	"github.com/gofiber/utils"
	"github.com/golang-jwt/jwt/v4"
)

type authService struct {
	authRepository repository.AuthRepository
}

func NewAuthService(connect repository.AuthRepository) AuthService {
	return authService{connect}
}

// GenerateToken implements AuthService
func (s authService) GenerateToken(user string) (*UserResponse, error) {
	connectDB, err := s.authRepository.GetUser(user)
	if err != nil {
		lib.LogInfoSQL(fmt.Sprintf("GetUser: %v", err))
		return nil, lib.NewError(50001, "user not found")
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"uuid":        utils.UUID(),
		"name":        connectDB.UserName.String,
		"api_user_id": connectDB.AuthenAPIUserID.Int64,
		"create_date": time.Now(),
	})
	hashToken, err := token.SignedString([]byte(connectDB.Password.String))
	if err != nil {
		lib.LogInfoSQL(fmt.Sprintf("JWT Error: %v", err))
		return nil, lib.NewError(50001, "jwt error")
	}

	generate := map[string]interface{}{
		"UserName": user,
		"Token":    hashToken,
	}

	res, err := s.authRepository.GenerateToken(generate)
	if err != nil {
		lib.LogInfoSQL(fmt.Sprintf("GenerateToken: %v", err))
		return nil, lib.NewError(50002, "can not update data in database")
	}

	response := UserResponse{
		UserName:    res.UserName.String,
		Token:       res.Token.String,
		Description: "",
	}

	return &response, nil
}

// RevokeToken implements AuthService
func (s authService) RevokeToken(user string) (*UserResponse, error) {
	connectDB, err := s.authRepository.GetUser(user)
	if err != nil {
		lib.LogInfoSQL(fmt.Sprintf("GetUser: %v", err))
		return nil, lib.NewError(50001, "user not found")
	}

	res, err := s.authRepository.RevokeToken(connectDB.UserName.String)
	if err != nil {
		lib.LogInfoSQL(fmt.Sprintf("RevokeToken: %v", err))
		return nil, lib.NewError(50002, "can not update data in database")
	}

	response := UserResponse{
		UserName:    res,
		Token:       "",
		Description: "success revoke token",
	}
	return &response, nil
}

// GetCommand implements AuthService
func (s authService) GetCommand(token string) (res []CommandResponse, err error) {
	connectDB, err := s.authRepository.GetCommand(token) // Find all command by token in database
	if err != nil {
		lib.LogInfoSQL(fmt.Sprintf("Validate token: %v", err))
		return nil, lib.NewError(50001, "user not found")
	}

	for _, detail := range connectDB {
		res = append(res, CommandResponse{
			Command: detail.Command.String,
		})
	}
	return res, nil
}
