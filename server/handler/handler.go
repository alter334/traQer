package handler

import (
	"context"
	"net/http"
	"time"

	"github.com/jmoiron/sqlx"
	"github.com/labstack/echo"
	"github.com/traPtitech/go-traq"
)

// メッセージ取得用stract
type MessageCountWithUser struct {
	User         traq.User
	MessageCount int
}

// メッセージ取得用stract
type MessageCountuuid struct {
	Useruuid     string
	MessageCount int
}

// ハンドラ作成
func NewHandler(db *sqlx.DB, client *traq.APIClient, auth context.Context, lasttrack time.Time, lastmessageuuid string) *Handler {
	return &Handler{db: db, client: client, auth: auth, lasttrack: lasttrack, lastmessageuuid: lastmessageuuid}
}

// 全データベースを読み取るAPI
func (h *Handler) GetMessageCountsParUser(c echo.Context) error {
	var messageCountuuid []MessageCountuuid
	err := h.db.Select(&messageCountuuid, "SELECT * FROM `messagecounts`")
	if err != nil {
		return c.String(http.StatusInternalServerError, err.Error())
	}
	var messageCountwithUser []MessageCountWithUser
	for _, messageCount := range messageCountuuid {

	}
}
