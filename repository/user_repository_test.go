package repository

import (
	"context"
	"reflect"
	"testing"
	"time"

	sqlmock "github.com/DATA-DOG/go-sqlmock"

	"github.com/budougumi0617/go-sql-sample/entity"
)

var o = entity.User{
	ID:        1,
	Name:      "budougumi0617",
	Email:     "budougumi0617@example.com",
	CreatedAt: time.Now(),
	UpdatedAt: time.Now(),
}

var existRows = func() *sqlmock.Rows {
	return sqlmock.NewRows([]string{
		"id",
		"name",
		"email",
		"created_at",
		"updated_at",
	}).AddRow(
		o.ID,
		o.Name,
		o.Email,
		o.CreatedAt,
		o.UpdatedAt,
	)
}

func TestRepo_FindUser(t *testing.T) {
	unknownID := int64(999)

	mockdb, mock, err := sqlmock.New()
	if err != nil {
		t.Fatal(err)
	}
	ctx := context.Background()
	repo := NewRepo(mockdb)

	mock.ExpectQuery(`
		SELECT id, name, email, created_at, updated_at FROM user WHERE id = \?
	`).WithArgs(o.ID).WillReturnRows(existRows())

	mock.ExpectQuery(`
		SELECT id, name, email, created_at, updated_at FROM user WHERE id = \?
	`).WithArgs(unknownID).WillReturnRows(sqlmock.NewRows([]string{}))

	got, err := repo.FindUser(ctx, o.ID)
	if err != nil {
		t.Errorf("want no error, but %v", err)
	}
	if !reflect.DeepEqual(&o, got) {
		t.Errorf("want %v, but %v", o, got)
	}

	got2, err := repo.FindUser(ctx, unknownID)
	if err != nil {
		t.Errorf("want no error, but %v", err)
	}
	if got2 != nil {
		t.Errorf("want nil, but %v", got2)
	}
	if mock.ExpectationsWereMet() != nil {
		t.Errorf("mock has error %v", err)
	}
}

func TestRepo_AddUser(t *testing.T) {
	mockdb, mock, err := sqlmock.New()
	if err != nil {
		t.Fatal(err)
	}
	ctx := context.Background()
	repo := NewRepo(mockdb)

	mock.ExpectPrepare(`
        INSERT INTO user \(name, email, created_at, updated_at\)
        VALUES \(\?, \?, \?, \?\)
    `).ExpectExec().WithArgs(
		o.Name,
		o.Email,
		sqlmock.AnyArg(),
		sqlmock.AnyArg(),
	).WillReturnResult(sqlmock.NewResult(o.ID, 1))

	in := o
	if err := repo.AddUser(ctx, &in); err != nil {
		t.Errorf("expected no error, but %v", err)
	}
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("mock has error %v", err)
	}
}
