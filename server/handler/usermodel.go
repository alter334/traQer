package handler

import (
	"time"

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
	Id                string `json:"id" db:"userid"`
	DisplayName       string `json:"displayname" db:"displayname"`
	Name              string `json:"name" db:"username"`
	Homechannel       string `json:"homechannel" db:"homechannelid"`
	TotalMessageCount int64  `json:"totalmessagecount" db:"totalpostcounts"`
	DailyMessageCount int64  `json:"dailymessagecount" db:"dailypostcounts"`
}

// メッセージ取得用stract
type MessageCountuuid struct {
	Userid       string `json:"userid" db:"userid"`
	MessageCount int    `json:"messagecount" db:"totalpostcounts"`
}

// 直近メッセージのid関連を読み取るstruct
type RecentMessages struct {
	Messageid string    `json:"messageid" db:"messageid"`
	Userid    string    `json:"userid" db:"userid"`
	Channelid string    `json:"channelid" db:"channelid"`
	Posttime  time.Time `json:"posttime" db:"posttime"`
}

// メッセージflow用
type RecentMessageCountbyChannel struct {
	Channelid string `json:"channelid" db:"channelid"`
	Count     int    `json:"count" db:"COUNT(*)"`
}

type DMSubscriber struct {
	Userid           string `json:"userid" db:"userid"`
	NotifyFlowAmount int    `json:"notifyflowamount" db:"notifyflowamount"`
}

//--------------------------------------------------------

// メッセージ集計mapのソート
type MessageCountPair struct {
	Key   string
	Value int
}

type MessageCountPairList []MessageCountPair

func (p MessageCountPairList) Len() int           { return len(p) }
func (p MessageCountPairList) Less(i, j int) bool { return p[i].Value < p[j].Value }
func (p MessageCountPairList) Swap(i, j int)      { p[i], p[j] = p[j], p[i] }
