package model

import (
	"time"

	"github.com/traPtitech/go-traq"
)

type ServerData struct {
	lastTrackMessage traq.Message // 最後に取得したメッセージ
	lastTrackTime    time.Time    // 最後の取得日時
	// 増えたらここに書く
}
