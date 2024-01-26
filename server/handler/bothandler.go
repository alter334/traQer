package handler

import (
	"database/sql"
	"errors"
	"log"
	"strconv"
)

func (h *Handler) BotCollectUserRank(groupName string) (x string) {
	log.Println("Init")

	res := ""
	//グループ指定なしなら全ユーザランク付
	if groupName == "" {
		res = "全ユーザー投稿数ランキング\n|順位|ユーザー|投稿数|\n|---|---|---|\n"
		for i, data := range h.nowhavingdata {
			res += ("|" + strconv.Itoa(i+1) + "|:@" + data.Name + ": " + data.Name + "|" + strconv.Itoa(int(data.TotalMessageCount)) + "|\n")
			if i == 99 {
				break
			}
		}
		return res
	}

	res = (groupName + " 所属投稿数ランキング\n|順位|ユーザー|投稿数|\n|---|---|---|\n")
	//グループ指定ありのランク グループIDを取得する
	groupid := h.b.BotGetGroupUUID(groupName)
	if groupid == "" {
		return "Such a group does not exist"
	}
	//グループメンバをmapのkey化 後の探索での高速化
	groupmembersids := h.b.BotGetGroupMembers(groupid)
	groupmembermap := map[string]struct{}{}
	for _, member := range groupmembersids {
		groupmembermap[member] = struct{}{}
	}

	//全ユーザーに対してグループ存在するか探索 100件拾ったら終了
	ct := 0
	for _, data := range h.nowhavingdata {
		_, exist := groupmembermap[data.Id]
		if exist {
			ct++
			res += ("|" + strconv.Itoa(ct) + "|:@" + data.Name + ": " + data.Name + "|" + strconv.Itoa(int(data.TotalMessageCount)) + "|\n")
			if ct == 100 {
				break
			}

		}
	}
	return res

}

func (h *Handler) BotDMSubscribe(userid string, notifyflowamount int) string {
	dmsubscriber := DMSubscriber{}
	err := h.db.Get(&dmsubscriber, "SELECT * FROM `dmsubscribers` WHERE userid=?", userid)
	if errors.Is(err, sql.ErrNoRows) {
		_, err = h.db.Exec("INSERT INTO `dmsubscribers`(`userid`,`notifyflowamount`) VALUES(?,?)", userid, notifyflowamount)
		if err != nil {
			return "DB insert error"
		}
		return "traQerDMの購読が完了しました **投稿数" + strconv.Itoa(notifyflowamount) + "/15分** で通知します。\n 通知基準量変更時は`@BOT_traQer \\dm\\enroll {設定したい値}`,購読解除時は`@BOT_traQer \\dm\\unsubscribe`コマンドで"
	} else if err != nil {
		return "DB select error"
	}
	_, err = h.db.Exec("UPDATE `dmsubscribers` SET `notifyflowamount`=? WHERE `userid`=?", notifyflowamount, userid)
	if err != nil {
		return "DB update error"
	}
	return "設定が完了しました **投稿数" + strconv.Itoa(notifyflowamount) + "/15分** で通知します。\n 通知基準量変更時は`@BOT_traQer \\dm\\enroll {設定したい値}`,購読解除時は`@BOT_traQer \\dm\\unsubscribe`コマンドで"
}

func (h *Handler) BotDMUnSubscribe(userid string) string {
	_, err := h.db.Exec("DELETE FROM `dmsubscribers` WHERE `userid`=?", userid)
	if err != nil {
		return "DB delete error"
	}
	return "traQerDMの購読を解除しました"
}
