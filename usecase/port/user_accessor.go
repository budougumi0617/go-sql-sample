package port

import (
	"context"

	"github.com/budougumi0617/go-sql-sample/entity"
)

// UserAccessor is a set of reader and writer for User in a data store.
type UserAccessor interface {
	UserReader
	UserWriter
}

// UserReader retrieves User data from a data store.
type UserReader interface {
	FindUser(context.Context, int64) (*entity.User, error)
}

// UserWriter stores User data into a data store.
type UserWriter interface {
	AddUser(context.Context, *entity.User) error

