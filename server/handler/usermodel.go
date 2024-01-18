package handler

import (
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
}

// メッセージ取得用stract
type MessageCountuuid struct {
	Userid       string `json:"userid" db:"userid"`
	MessageCount int    `json:"messagecount" db:"totalpostcounts"`
}
