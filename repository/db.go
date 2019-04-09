package repository

import (
	"database/sql"
)

// Repo is a repository for a DB.
type Repo struct {
	db *sql.DB
}

// NewRepo creates a new Repo.
func NewRepo(db *sql.DB) *Repo {
	return &Repo{db}
}
