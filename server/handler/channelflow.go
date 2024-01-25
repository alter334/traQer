package handler

import (
	"log"
	"time"
)

// after以降のメッセージ数をチャンネル毎に取得する
func (h *Handler) RecentMessageCollector(after time.Time) ([]RecentMessageCountbyChannel, error) {
	messagecounts := []RecentMessageCountbyChannel{}
	err := h.db.Select(&messagecounts, "SELECT `channelid`,COUNT(*) FROM recentmessages WHERE `posttime` > ? GROUP BY `channelid`", after)
	if err != nil {
		log.Println("Error getting recent messages:", err.Error())
		return messagecounts, err
	}
	log.Println("Success:", len(messagecounts))
	return messagecounts, nil
}
