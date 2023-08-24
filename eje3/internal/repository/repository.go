package repository

import (
	"context"
	"ejercicio3/internal/entity"

	"github.com/jmoiron/sqlx"
)

type Repository interface {
	SaveUser(ctx context.Context, email, name, password string) error
	GetUserByEmail(ctx context.Context, email string) (entity.User, error)
}
type repo struct {
	db *sqlx.DB
}

func New(db *sqlx.DB) Repository {
	return &repo{
		db: db,
	}
}
