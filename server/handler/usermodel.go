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

type RecentMessageCountbyChannel struct {
	Channelid string `json:"channelid" db:"channelid"`
	Count     int    `json:"count" db:"COUNT(*)"`
}
