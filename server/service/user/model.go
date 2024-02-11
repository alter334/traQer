package user

import "traQer/api"

// Handler
type UserHandler struct {
	api *api.ApiHandler
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
