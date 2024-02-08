package api

import (
	"time"
	"traQer/api/db"
	"traQer/api/qapi"

	"github.com/traPtitech/go-traq"
)

type ApiHandler struct {
	db         *db.DBHandler     // db関連
	qapi       *qapi.QapiHandler // traqApi関連
	serverData ServerData // サーバの持つデータ
}

type ServerData struct {
	lastTrackMessage traq.Message // 最後に取得したメッセージ
	lastTrackTime    time.Time    // 最後の取得日時
	// 増えたらここに書く
}
