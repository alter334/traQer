package api

import (
	"traQer/api/db"
	"traQer/api/qapi"
)

func NewApiHandler(db *db.DBHandler, qapi *qapi.QapiHandler) *ApiHandler {
	return &ApiHandler{db: db, qapi: qapi, serverData: ServerData{}}
}

func ApiSetup() *ApiHandler {
	db := db.DBSetup()
	qapi := qapi.QapiSetup()

	api := NewApiHandler(db, qapi)
	return api
}
