package repository

import (
	"github.com/jmoiron/sqlx"
)

type categoryRepositoryDB struct {
	db *sqlx.DB
}

func NewCategoryRepositoryDB(connect *sqlx.DB) CategoryRepository {
	return categoryRepositoryDB{connect}
}

func (r categoryRepositoryDB) GetCategory() (connectDB []Category, err error) {
	query := `
SELECT * 
FROM PRIVUSER.CATEGORY `

	if err = r.db.Select(&connectDB, query); err != nil {
		return nil, err
	}

	return connectDB, nil
}
