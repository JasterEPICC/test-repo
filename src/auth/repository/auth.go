package repository

import (
	"database/sql"
)

type AuthRepository interface {
	GetUser(user string) (*GetUser, error)
	GenerateToken(map[string]interface{}) (*GenerateToken, error)
	RevokeToken(user string) (string, error)
	GetCommand(string) ([]GetCommand, error)
}

type (
	GetUser struct {
		UserName        sql.NullString `db:"USER_NAME"`
		Password        sql.NullString `db:"PASSWORD"`
		AuthenAPIUserID sql.NullInt64  `db:"AUTHEN_API_USER_ID"`
	}

	GenerateToken struct {
		UserName sql.NullString `db:"USER_NAME"`
		Token    sql.NullString `db:"TOKEN"`
	}

	GetCommand struct {
		Token              sql.NullString `db:"TOKEN"`
		AuthToken          sql.NullString `db:"AUTH_TOKEN"`
		AuthenAPIUserID    sql.NullInt64  `db:"AUTHEN_API_USER_ID"`
		AuthenAPICommandID sql.NullInt64  `db:"AUTHEN_API_COMMAND_ID"`
		Command            sql.NullString `db:"COMMAND"`
	}
)
