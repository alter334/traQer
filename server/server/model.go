package server

import (
	"time"
	"traQer/handler/bot"
	"traQer/handler/db"
	"traQer/handler/qapi"

	"github.com/traPtitech/go-traq"
)

type Server struct {
	bot        *bot.Bot
	db         *db.DB
	qapi       *qapi.Qapi
	serverData serverData
}

type serverData struct {
	lastTrackMessage traq.Message // 最後に取得したメッセージ
	lastTrackTime    time.Time    // 最後の取得日時
	// 増えたらここに書く
}
