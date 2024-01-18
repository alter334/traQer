package handler

import (
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


