package api

import (
	"traQer/api/db"
	"traQer/api/qapi"
)

func NewApiHandler(db *db.DBHandler, qapi *qapi.QapiHandler) *ApiHandler {
	return &ApiHandler{db: db, qapi: qapi ,serverData: ServerData{}}
}
