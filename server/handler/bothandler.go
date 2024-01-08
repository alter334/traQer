package handler

import (
	"log"
	"strconv"
)

func (h *Handler) BotCollectUserRank(group string) (x string) {
	log.Println("Init")
	var collectdata []MessageCountuuid
	err := h.db.Select(&collectdata, "SELECT * FROM `messagecounts` ORDER BY `totalpostcounts` DESC")
	if err != nil {
		log.Println("Internal error: " + err.Error())
		return ("Internal error: " + err.Error())
	}

	res := "|順位|ユーザー|投稿数|\n|---|---|---|\n"
	//グループ指定なしなら全ユーザランク付
	if group == "" {

		for i, data := range collectdata {
			log.Println(strconv.Itoa(i))
			username := h.b.BotGetUserName(data.Userid)
			res += ("|" + strconv.Itoa(i+1) + "|:@" + username + ": " + username + "|" + strconv.Itoa(data.MessageCount) + "|\n")
		}
	}
	log.Println(res)
	return res

}
