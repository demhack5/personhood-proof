package user

import (
	"github.com/jmoiron/sqlx"
)

const insertUserQuery = ``

type UserRepository struct {
	db *sqlx.DB
}

func NewUserRepository(db *sqlx.DB) *UserRepository {
	return &UserRepository{
		db: db,
	}
}
