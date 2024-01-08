package handler

import (
	"context"
	"net/http"
	"strconv"
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
	Id                string   `json:"id"`
	DisplayName       string   `json:"displayname"`
	Name              string   `json:"name"`
	Groups            []string `json:"groups"`
	Homechannnel      string   `json:"homechannnel"`
	TotalMessageCount int64    `json:"totalmessagecount"`
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

// 全データベースを読み取るAPI ただし制御機構付
func (h *Handler) GetMessageCounts(c echo.Context) error {
	//クエリパラメタでページ制御 50*page
	pagestr := c.QueryParam("page")
	if pagestr == "" {
		return c.JSON(http.StatusOK, h.nowhavingdata)
	}
	page, err := strconv.Atoi(pagestr)
	if err != nil {
		return c.String(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, h.nowhavingdata[page*50:(page+1)*50])
}

//
