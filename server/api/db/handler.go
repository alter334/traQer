package db

import (
	"github.com/jmoiron/sqlx"
)

func NewDBHandler(db *sqlx.DB) *DBHandler {
	return &DBHandler{db: db}
}
