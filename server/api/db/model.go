package db

import "github.com/jmoiron/sqlx"

type DBHandler struct {
	db *sqlx.DB
}
