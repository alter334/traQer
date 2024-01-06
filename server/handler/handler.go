package handler

import (
	"context"
	"log"
	"net/http"
	"time"

	"github.com/jmoiron/sqlx"
	"github.com/labstack/echo"
	"github.com/traPtitech/go-traq"
)

//----------------------------------------------------------------

// メッセージ取得用stract
type UserMessages struct {
	User              traq.User `json:"user"`
	TotalMessageCount int64     `json:"TotalMessageCount"`
}

// 取得したメッセージを詳細にstract
type UserDetailWithMessageCount struct {
	UserDetail        traq.UserDetail `json:"userdetail"`
	TotalMessageCount int64           `json:"TotalMessageCount"`
}

// メッセージ取得用stract
type MessageCountuuid struct {
	Userid       string `json:"userid" db:"userid"`
	MessageCount int    `json:"messagecount" db:"totalpostcounts"`
}

//----------------------------------------------------------------

// ハンドラ作成
func NewHandler(db *sqlx.DB, client *traq.APIClient, auth context.Context, lasttrack time.Time, lastmessageuuid string) *Handler {
	return &Handler{db: db, client: client, auth: auth, lasttrack: lasttrack, lastmessageuuid: lastmessageuuid}
}

// 全データベースを読み取るAPI
func (h *Handler) GetMessageCountsParUser(c echo.Context) error {
	var messageCountuuid []MessageCountuuid
	err := h.db.Select(&messageCountuuid, "SELECT * FROM `messagecounts` ORDER BY `totalpostcounts` DESC")
	if err != nil {
		return c.String(http.StatusInternalServerError, err.Error())
	}

	var messageCountwithUser []UserDetailWithMessageCount
	for i, messageCount := range messageCountuuid {
		userdetail, r, err := h.client.UserApi.GetUser(h.auth, messageCount.Userid).Execute()
		log.Println(i+1, ":",messageCount.MessageCount,":", userdetail.DisplayName)
		if err != nil {
			return c.String(r.StatusCode, err.Error())
		}
		messageCountwithUser = append(messageCountwithUser, UserDetailWithMessageCount{UserDetail: *userdetail, TotalMessageCount: int64(messageCount.MessageCount)})
	}
	return c.JSON(http.StatusOK, messageCountwithUser)
}

// 
