package db

import (
	"github.com/jmoiron/sqlx"
)

func NewDBHandler(db *sqlx.DB) *DB {
	return &DB{db: db}
}
