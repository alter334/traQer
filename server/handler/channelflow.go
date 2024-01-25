package handler

import (
	"log"
	"strconv"
	"time"
)

// after以降のメッセージ数をチャンネル毎に取得する
func (h *Handler) RecentMessageCollector(since time.Duration) ([]RecentMessageCountbyChannel, error) {
	messagecounts := []RecentMessageCountbyChannel{}
	from := time.Now().UTC().Add(-since)
	err := h.db.Select(&messagecounts, "SELECT `channelid`,COUNT(*) FROM recentmessages WHERE `posttime` > ? GROUP BY `channelid` ORDER BY `COUNT(*)` DESC", from)
	if err != nil {
		log.Println("Error getting recent messages:", err.Error())
		return messagecounts, err
	}
	log.Println("Success:", len(messagecounts))
	h.SendingDMs(messagecounts)
	return messagecounts, nil
}

// DM配信
func (h *Handler) SendingDMs(messagecounts []RecentMessageCountbyChannel) {
	// 流れのいいチャンネルを通知
	// 購読者ごとに追加していきたい
	// 購読者の取得(基準メッセ数が多い順)->メッセージを伸ばしていく感じ->内容ができたら配信
	subscribers := []DMSubscriber{}
	err := h.db.Select(&subscribers, "SELECT * FROM dmsubscribers ORDER BY `notifyflowamount` DESC")
	if err != nil {
		log.Println("Error getting subscribers:", err.Error())
		return
	}

	i := 0 //messagecountsのどこまで消化したかのイテレーション
	content := ""
	for _, s := range subscribers {
		for messagecounts[i].Count >= s.NotifyFlowAmount {
			channelname, err := h.GetChannelNameWithParents(messagecounts[i].Channelid, "")
			if err != nil {
				log.Println("Error getting channelname:", err.Error())
				return
			}
			// とりあえずは15分のみ対応
			content += "### チャンネル: [#" + channelname + "]" +
				"(https://q.trap.jp/channels/" + channelname + ") " +
				"15min投稿数:" + strconv.Itoa(messagecounts[i].Count) + "\n"
			i++
			if i == len(messagecounts) {
				break
			}
		}
		if content != "" {
			h.b.BotDM(s.Userid, content)
		}
	}

}
