package service

type AuthService interface {
	GenerateToken(user string) (*UserResponse, error)
	RevokeToken(user string) (*UserResponse, error)
	GetCommand(token string) ([]CommandResponse, error)
}

type (
	UserResponse struct {
		UserName    string `json:"user"`
		Token       string `json:"token,omitempty"`
		Description string `json:"description,omitempty"`
	}

	CommandResponse struct {
		Command string `json:"command"`
	}
)
