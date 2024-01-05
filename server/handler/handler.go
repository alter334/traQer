package handler

import (
	"context"
	"log"
	"time"

	"github.com/jmoiron/sqlx"
	"github.com/traPtitech/go-traq"
)

type Handler struct {
	db              *sqlx.DB
	client          *traq.APIClient
	auth            context.Context
	lasttrack       time.Time
	lastmessageuuid string
}

// ハンドラ作成
func NewHandler(db *sqlx.DB, client *traq.APIClient, auth context.Context, lasttrack time.Time) *Handler {
	return &Handler{db: db, client: client, auth: auth, lasttrack: lasttrack}
}

//------------------------------------------------
//メッセージ数取得系
//------------------------------------------------

type UserMessages struct {
	User              traq.User `json:"user"`
	TotalMessageCount int64     `json:"TotalMessageCount"`
}

// ユーザ毎traQ投稿数DB記録補正:高負荷のため1日に1回実施
func (h *Handler) GetUserPostCount() {
	//最終探索時間
	h.lasttrack = time.Now().UTC()
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

	//ユーザデータのdb反映
	for _, message := range userMessages {
		_, err = h.db.Exec("INSERT INTO `messagecounts`(`totalpostcounts`,`userid`) VALUES(?,?) ON DUPLICATE KEY UPDATE `totalpostcounts`=VALUES(totalpostcounts)", message.TotalMessageCount, message.User.Id)
		if err != nil {
			log.Println("Internal error:", err.Error())
			return
		}

	}
	log.Println("done")
}

// ユーザ毎一定期間traQ投稿数取得(差分取得のため誤差有)
func (h *Handler) SearchMessagesRunner() {
	from := h.lasttrack.Add(-time.Minute) // メッセージ反映にある1分のラグを捕捉する
	to := time.Now().UTC()
	h.lasttrack = time.Now().UTC()

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
		for _, message := range messages.Hits {
			if message.Id == h.lastmessageuuid {
				break
			}
			messageCountperUser[message.UserId]++
			tmplastmessageuuid = message.Id
		}
		if len(messages.Hits) < 100 {
			break
		}
		//取得した最後のメッセージのuuidを取得
		h.lastmessageuuid = tmplastmessageuuid
	}

	//mapに応じてsqlを発行
	for userId, messageCount := range messageCountperUser {
		log.Println("ユーザUUID:", userId, "取得数:", messageCount)
		_, err := h.db.Exec("INSERT INTO `messagecounts`(`totalpostcounts`,`userid`) VALUES(?,?) ON DUPLICATE KEY UPDATE `totalpostcounts`=`totalpostcounts`+VALUES(totalpostcounts)", messageCount, userId)
		if err != nil {
			log.Println("Internal error:", err.Error())
			return
		}
	}

}

// 差分取得実施(limit等条件のため)
func (h *Handler) CorrectUserMessageDiff(from time.Time, to time.Time, offset int) (message *traq.MessageSearchResult, err error) {
	messages, _, err := h.client.MessageApi.SearchMessages(h.auth).
		Bot(false).After(from).Before(to).Limit(100).Offset(int32(offset)).Execute()
	if err != nil {
		return messages, err
	}
	log.Println("取得数:", len(messages.Hits))
	log.Println("取得mes:", messages.TotalHits)
	return messages, nil

}

//差分取得実施(limit等条件のため)
