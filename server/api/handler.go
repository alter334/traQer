package api

import (
	"traQer/api/db"
	"traQer/api/qapi"
	"traQer/model"
)

type ApiHandler struct {
	db         *db.DBHandler     // db関連
	qapi       *qapi.QapiHandler // traqApi関連
	serverData model.ServerData  // サーバの持つデータ
}

func NewApiHandler(db *db.DBHandler, qapi *qapi.QapiHandler) *ApiHandler {
	return &ApiHandler{db: db, qapi: qapi, serverData: model.ServerData{}}
}

func ApiSetup() *ApiHandler {
	db := db.DBSetup()
	qapi := qapi.QapiSetup()

	api := NewApiHandler(db, qapi)
	return api
}
