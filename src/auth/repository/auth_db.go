package repository

import (
	"fmt"

	"github.com/jmoiron/sqlx"
)

type authRepositoryDB struct {
	db *sqlx.DB
}

func NewAuthRepositoryDB(connect *sqlx.DB) AuthRepository {
	return authRepositoryDB{connect}
}

// GetAuth implements AuthRepository
func (r authRepositoryDB) GetUser(user string) (*GetUser, error) {
	connectDB := GetUser{}
	query := `SELECT USER_NAME, PASSWORD, AUTHEN_API_USER_ID
				FROM PRIVUSER.AUTHEN_API_USER
				WHERE USER_NAME = ? `
	err := r.db.Get(&connectDB, query, user)
	if err != nil {
		return nil, err
	}
	return &connectDB, nil
}

// GenerateToken implements AuthRepository
func (r authRepositoryDB) GenerateToken(arg map[string]interface{}) (*GenerateToken, error) {
	tx := r.db.MustBegin()
	query := `UPDATE PRIVUSER.AUTHEN_API_USER
				SET TOKEN = ?
				WHERE USER_NAME = ?`
	tx.MustExec(query, arg["Token"], arg["UserName"])
	if err := tx.Commit(); err != nil {
		if err := tx.Rollback(); err != nil {
			return nil, err
		}
		return nil, err
	}

	connectDB := GenerateToken{}
	query = `SELECT USER_NAME, TOKEN
				FROM PRIVUSER.AUTHEN_API_USER
				WHERE USER_NAME = ? `
	err := r.db.Get(&connectDB, query, arg["UserName"])
	if err != nil {
		return nil, err
	}
	return &connectDB, nil
}

// RevokeToken implements AuthRepository
func (r authRepositoryDB) RevokeToken(user string) (string, error) {
	tx := r.db.MustBegin()
	query := ` UPDATE PRIVUSER.AUTHEN_API_USER
			SET TOKEN = NULL
			WHERE USER_NAME = ?`
	tx.MustExec(query, user)
	if err := tx.Commit(); err != nil {
		fmt.Println("err commit: " + fmt.Sprint(err))
		if err := tx.Rollback(); err != nil {
			fmt.Println("err rollback: " + fmt.Sprint(err))
			return "", err
		}
		return "", err
	}
	return user, nil
}

// GetCommand implements AuthRepository
func (r authRepositoryDB) GetCommand(token string) (connectDB []GetCommand, err error) {
	query := ` SELECT u.AUTHEN_API_USER_ID
						,u.TOKEN
						,p.AUTHEN_API_COMMAND_ID
						,c.COMMAND
				FROM PRIVUSER.AUTHEN_API_USER u
				INNER JOIN PRIVUSER.AUTHEN_API_PERMISSION p
					ON (p.AUTHEN_API_USER_ID = u.AUTHEN_API_USER_ID)
				INNER JOIN PRIVUSER.AUTHEN_API_COMMAND c
					ON (c.AUTHEN_API_COMMAND_ID = p.AUTHEN_API_COMMAND_ID)
				WHERE u.TOKEN = ? `
	err = r.db.Select(&connectDB, query, token)
	if err != nil {
		return nil, err
	}

	// >>>>>>>>> ADD API COMMAND <<<<<<<<<<<
	// tx := r.db.MustBegin()
	// query := ` INSERT INTO [PRIVUSER].[AUTHEN_API_PERMISSION]
	// 			VALUES (999, ?)`
	// for i := 1; i <= 99; i++ {
	// 	tx.MustExec(query, i)

	// 	fmt.Println(i)
	// }
	// if err := tx.Commit(); err != nil {
	// 	fmt.Println("err commit: " + fmt.Sprint(err))
	// 	if err := tx.Rollback(); err != nil {
	// 		fmt.Println("err rollback: " + fmt.Sprint(err))
	// 		return nil, err
	// 	}
	// 	return nil, err
	// }

	// >>>>>>> REMOVE API COMMAND <<<<<<<
	// tx := r.db.MustBegin()
	// query := ` DELETE FROM [PRIVUSER].[AUTHEN_API_PERMISSION]
	// 			WHERE AUTHEN_API_USER_ID = ? `
	// tx.MustExec(query, 999)
	// if err := tx.Commit(); err != nil {
	// 	fmt.Println("err commit: " + fmt.Sprint(err))
	// 	if err := tx.Rollback(); err != nil {
	// 		fmt.Println("err rollback: " + fmt.Sprint(err))
	// 		return nil, err
	// 	}
	// 	return nil, err
	// }
	return connectDB, nil
}
