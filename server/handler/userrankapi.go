package handler

import (
	"log"
	"net/http"
	"time"

	"github.com/traPtitech/go-traq"
)

//------------------------------------------------
//ユーザー毎メッセージ数取得系
//------------------------------------------------

// ユーザ毎traQ投稿数DB記録補正:全ユーザstats取得 高負荷のため1日に1回実施
func (h *Handler) GetUserPostCount() {

	//ユーザリストの取得
	v, _, err := h.client.UserApi.
		GetUsers(h.auth).
		Execute()
	if err != nil {
		log.Println("Internal error:", err.Error())
		return
	}
	log.Println("Userlistavailable:", len(v))

	var userMessages []UserMessages

	//各ユーザ毎投稿数の取得(Botは無視)
	for i, user := range v {
		if user.Bot {
			log.Println("isBot:", i)
			continue
		}
		userStats, _, err := h.client.UserApi.GetUserStats(h.auth, user.Id).Execute()
		if err != nil {
			log.Println("Internal error:", err.Error())
			return
		}
		var message UserMessages
		message.User = user
		message.TotalMessageCount = userStats.TotalMessageCount
		userMessages = append(userMessages, message)
		log.Println("traQing:", i, "mescount:", message.TotalMessageCount)
	}

	//収集完了時刻を最終調査時刻とする
	h.lasttrack = time.Now().UTC()
	//最終探索時間のdb記録
	_, err = h.db.Exec("UPDATE `trackinginfo` SET `lasttracktime`=?", time.Now().UTC())
	if err != nil {
		log.Println("Internal error:", err.Error())
		return
	}

	//ユーザデータのdb反映
	for _, message := range userMessages {
		_, err = h.db.Exec("INSERT INTO `messagecounts`(`totalpostcounts`,`userid`) VALUES(?,?) ON DUPLICATE KEY UPDATE `totalpostcounts`=VALUES(totalpostcounts)", message.TotalMessageCount, message.User.Id)
		if err != nil {
			log.Println("Internal error:", err.Error())
			return
		}

	}
	//ハンドラに情報を持たせる
	h.MessageCountsBind(false)

	log.Println("done")
}

//----------------------------------------------------------------

// ユーザ毎一定期間traQ投稿数取得(差分取得のため若干誤差の可能性有 基本的にはこの探索法を利用)
func (h *Handler) SearchMessagesRunner() {
	from := h.lasttrack.Add(-time.Minute) // メッセージ反映にある1分のラグを捕捉する
	to := time.Now().UTC()
	h.lasttrack = time.Now().UTC()
	//最終探索時間のdb記録
	_, err := h.db.Exec("UPDATE `trackinginfo` SET `lasttracktime`=?", time.Now().UTC())
	if err != nil {
		log.Println("Internal error:", err.Error())
		return
	}

	//記録用mapの作成
	messageCountperUser := map[string]int{}
	var tmplastmessageuuid string

	//オフセットを100ずつ増やすことで100件しかメッセージが格納されない問題を解決する
	for i := 0; ; i += 100 {
		messages, err := h.CorrectUserMessageDiff(from, to, i)
		if err != nil {
			log.Println("Internal error:", err.Error())
			break
		}
		//取得したメッセージをmap型に記録

		for j, message := range messages.Hits {
			if i == 0 && j == 0 {
				//取得した最新のメッセージのuuidを取得
				tmplastmessageuuid = message.Id
			}
			if message.Id == h.lastmessageuuid {
				break
			}
			messageCountperUser[message.UserId]++

		}
		if len(messages.Hits) < 100 {
			break
		}

	}
	h.lastmessageuuid = tmplastmessageuuid
	//最新メッセージuuidの更新
	//最終探索時間のdb記録
	_, err = h.db.Exec("UPDATE `trackinginfo` SET `lasttrackmessageid`=?", tmplastmessageuuid)
	if err != nil {
		log.Println("Internal error:", err.Error())
		return
	}

	//mapに応じてsqlを発行
	for userId, messageCount := range messageCountperUser {
		log.Println("ユーザUUID:", userId, "実追加数:", messageCount)
		_, err := h.db.Exec("INSERT INTO `messagecounts`(`totalpostcounts`,`userid`) VALUES(?,?) ON DUPLICATE KEY UPDATE `totalpostcounts`=`totalpostcounts`+VALUES(totalpostcounts)", messageCount, userId)
		if err != nil {
			log.Println("Internal error:", err.Error())
			return
		}
	}
	//ハンドラに情報を持たせる
	h.MessageCountsBind(false)

}

// 差分取得実施(limit等条件のため)
func (h *Handler) CorrectUserMessageDiff(from time.Time, to time.Time, offset int) (message *traq.MessageSearchResult, err error) {
	messages, _, err := h.client.MessageApi.SearchMessages(h.auth).
		Bot(false).After(from).Before(to).Limit(100).Offset(int32(offset)).Sort(`createdAt`).Execute()
	if err != nil {
		return messages, err
	}
	log.Println("取得数:", len(messages.Hits))
	log.Println("取得mes:", messages.TotalHits)
	return messages, nil

}

//----------------------------------------------------------------

// グループ取得の実施 h.nowhavingdata のグループ毎振り分けを行う []UserDetailWithMessageCount,err を返す
func (h *Handler) GetGroupMembers(groupid string) (res []UserDetailWithMessageCount, httpres *http.Response, err error) {
	// groupメンバの取得
	usergroupmember, httpres, err := h.client.GroupApi.GetUserGroupMembers(h.auth, groupid).Execute()
	if err != nil {
		log.Println("Internal error:", err.Error())
		return res, httpres, err
	}

	// メンバのuuid取得
	groupmembersids := []string{}
	for _, member := range usergroupmember {
		groupmembersids = append(groupmembersids, member.GetId())
	}

	// グループメンバをmapkey化 探索高速化
	groupmembermap := map[string]struct{}{}
	for _, member := range groupmembersids {
		groupmembermap[member] = struct{}{}
	}

	//全ユーザーに対してグループ存在するか探索
	for _, data := range h.nowhavingdata {
		_, exist := groupmembermap[data.Id]
		if exist {
			res = append(res, data)
		}
	}

	return res, httpres, nil
}

//----------------------------------------------------------------

// DB読み取り実施, usetraqAPI? traQAPI:手元 より情報取得してハンドラ(traQAPIからの場合はdbにも)に情報を持たせる
func (h *Handler) MessageCountsBind(usetraqAPI bool) {

	var dbuserdata []UserDetailWithMessageCount
	err := h.db.Select(&dbuserdata, "SELECT * FROM `messagecounts` ORDER BY `totalpostcounts` DESC")
	if err != nil {
		log.Println("Internal error:", err.Error())
		return
	}

	if !usetraqAPI {
		//API使用しない(DisplayName等更新しない)なら既存のデータだけ返す
		h.nowhavingdata = dbuserdata
		return
	}

	var nowcollectingdata []UserDetailWithMessageCount

	for i, messageCount := range dbuserdata {
		userdetail, _, err := h.client.UserApi.GetUser(h.auth, messageCount.Id).Execute()
		if i <= 2 {
			log.Println(i+1, ":", messageCount.TotalMessageCount, ":", userdetail.DisplayName)
		}
		if err != nil {
			log.Println("Internal error:", err.Error())
			return
		}
		home := userdetail.GetHomeChannel()
		nowcollectingdata = append(nowcollectingdata,
			UserDetailWithMessageCount{Id: userdetail.Id,
				DisplayName:       userdetail.DisplayName,
				Name:              userdetail.Name,
				Homechannel:       home,
				TotalMessageCount: int64(messageCount.TotalMessageCount)})
		// db更新
		_, err = h.db.Exec("UPDATE `messagecounts` SET `displayname`=?, `username`=?, `homechannelid`=? WHERE `userid`=?", userdetail.DisplayName, userdetail.Name, home, userdetail.Id)
		if err != nil {
			log.Println("Internal error:", err.Error())
			return
		}
	}
	h.nowhavingdata = nowcollectingdata

}

//--------------------------------------------------------
// チャンネル毎メッセージ数取得系
//--------------------------------------------------------
