package db

import "github.com/jmoiron/sqlx"

type DB struct {
	db *sqlx.DB
}
