package handler

import (
	"log"
	"net/http"
	"strconv"

	"github.com/jmoiron/sqlx"
	"github.com/labstack/echo"
	traqwsbot "github.com/traPtitech/traq-ws-bot"
	"github.com/traPtitech/traq-ws-bot/payload"
)

type Handler struct {
	db  *sqlx.DB
	bot *traqwsbot.Bot
}

type AttackTo struct {
	Channelid        string `json:"channelid" db:"channelid"`
	Channnelusername string `json:"channelusername" db:"channelusername"`
}

type Enroll struct {
	Channelid        string `json:"channelid" db:"id"`
	Channnelusername string `json:"channelusername" db:"name"`
}

func NewHandler(db *sqlx.DB, bot *traqwsbot.Bot) *Handler {
	return &Handler{db: db, bot: bot}
}

// エントリー:テロ会員でない場合(db上に存在しなかった場合)はここにきてエントリーメッセージを投稿する
func (h *Handler) Entry(p *payload.MessageCreated) {
	_, err := h.db.Exec("INSERT INTO `users`(`name`,`id`,`attack`,`rate`) VALUES(?,?,0,0)", p.Message.User.Name, p.Message.User.ID)
	if err != nil {
		SimplePost(h.bot, p.Message.ChannelID, "Internal error: "+err.Error())
		log.Println("Internal error: " + err.Error())
		return
	}
	SimplePost(h.bot, p.Message.ChannelID, ":@"+p.Message.User.Name+":さん\n"+"## ようこそtraP飯テロ部へ\n")
	SimplePost(h.bot, "baaf247d-125a-47e4-82a8-ffcccab5f0b8", ":@"+p.Message.User.Name+"::sansen_1::sansen_2::sansen_3:")
}

// 投稿先候補の追加:過去にメンションされたことのないチャンネルIDの場合投稿先候補に加わる
func (h *Handler) MonitorInsert(p *payload.MessageCreated) {
	_, err := h.db.Exec("INSERT INTO `places`(`channelid`,`channelusername`) VALUES(?,?)", p.Message.ChannelID, p.Message.User.Name)
	if err != nil {
		SimplePost(h.bot, p.Message.ChannelID, "Internal error: "+err.Error())
		log.Println("Internal error: " + err.Error())
		return
	}
	monitorMessageId := SimplePost(h.bot, p.Message.ChannelID, "**このチャンネルが飯テロ対象チャンネルに登録されました。**\n登録者::@"+p.Message.User.Name+":")
	SimplePost(h.bot, "baaf247d-125a-47e4-82a8-ffcccab5f0b8", "https://q.trap.jp/messages/"+monitorMessageId+"\n:sansen_1::sansen_2::sansen_3:")
}

// 通常攻撃:db上に存在するユーザーから1人を選んで爆撃します
func (h *Handler) Attack(p *payload.MessageCreated, meshiurl string, attackNum int) {
	var attackTo AttackTo

	//初の攻撃なら自分のtimesに飛ぶ
	if attackNum == 0 {
		attackTo.Channelid = "402a1c2c-878e-40ef-ae14-011354394e36"
		attackTo.Channnelusername = "Alt--er"
		log.Println("InitAttack実行")
	} else {
		//ランダム選択1名
		err := h.db.Get(&attackTo, "SELECT `channelid`,`channelusername` FROM `places` ORDER BY RAND() LIMIT 1")
		if err != nil {
			SimplePost(h.bot, p.Message.ChannelID, "Internal error: "+err.Error())
			log.Println("Internal error: " + err.Error())
			return
		}
		log.Println("Attack実行")

	}

	attackNum++
	attackNumstr := strconv.Itoa(attackNum)
	attackMessageId := SimplePost(h.bot, attackTo.Channelid, ":@"+p.Message.User.Name+":"+":oisu-1::oisu-2::oisu-3::oisu-4yoko:"+meshiurl)
	SimplePost(h.bot, p.Message.ChannelID, ":@"+attackTo.Channnelusername+":"+"に爆撃しました。\n累積攻撃回数:"+attackNumstr+"回\n"+"https://q.trap.jp/messages/"+attackMessageId)
	_, err := h.db.Exec("UPDATE `users` SET `attack`=? WHERE `id`=?", attackNum, p.Message.User.ID)
	if err != nil {
		SimplePost(h.bot, p.Message.ChannelID, "Internal error: "+err.Error())
		log.Println("Internal error: " + err.Error())
		return
	}
	log.Println("Attack完了")
	//bot_playgroundチャンネルに飛ばす
	SimplePost(h.bot, "baaf247d-125a-47e4-82a8-ffcccab5f0b8", ":@"+p.Message.User.Name+":"+":oisu-1::oisu-2::oisu-3::oisu-4yoko:"+meshiurl)

}

//--------------------------------
// 開発用:既存ユーザのホームチャンネルを攻撃対象に

func (h *Handler) EnrollExistingUserHometoPlace() {
	var channeluuids []Enroll
	err := h.db.Select(&channeluuids, "SELECT `id`,`name` FROM users")
	if err != nil {
		log.Println("Internal error: " + err.Error())
		return
	}
	for _, u := range channeluuids {
		u.Channelid, _ = GetUserHome(h.bot, u.Channelid)
		_, err := h.db.Exec("INSERT INTO `places`(`channelid`,`channelusername`) VALUES(?,?)", u.Channelid, u.Channnelusername)
		if err != nil {
			log.Println("Internal error: " + err.Error())
		}
	}
	log.Println("Finished")
}

// テスト:自爆
func (h *Handler) SelfAttack(p *payload.MessageCreated, meshiurl string) {
	log.Println("SelfAttack実行")
	attackId, _ := GetUserHome(h.bot, p.Message.User.ID)
	SimplePost(h.bot, attackId, ":@"+p.Message.User.Name+":"+":oisu-1::oisu-2::oisu-3::oisu-4yoko:"+meshiurl)
	log.Println("SelfAttack完了")
}

//traQ投稿数取得

func (h *Handler) GetUserPostCount(c echo.Context) error {
	return c.String(http.StatusOK, "get")
}
