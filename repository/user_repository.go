package repository

import (
	"context"

	"github.com/budougumi0617/go-sql-sample/entity"
)

// FindUser gets user from repository.
func (repo *Repo) FindUser(ctx context.Context, id int64) (*entity.User, error) {
	return nil, nil
}

// AddUser adds user to repository.
func (repo *Repo) AddUser(ctx context.Context, u *entity.User) error {
	return nil
}
