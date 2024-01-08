package handler

import (
	"context"
	"time"

	"github.com/jmoiron/sqlx"
	"github.com/traPtitech/go-traq"
	traqwsbot "github.com/traPtitech/traq-ws-bot"
)

type Handler struct {
	db              *sqlx.DB
	b               *BotHandler
	client          *traq.APIClient
	auth            context.Context
	lasttrack       time.Time
	lastmessageuuid string
	nowhavingdata   []UserDetailWithMessageCount
}

type BotHandler struct {
	bot *traqwsbot.Bot
}

// ハンドラ作成
func NewHandler(db *sqlx.DB, bot *traqwsbot.Bot, client *traq.APIClient, auth context.Context, lasttrack time.Time, lastmessageuuid string) *Handler {
	return &Handler{db: db, b: NewBotHandler(bot), client: client, auth: auth, lasttrack: lasttrack, lastmessageuuid: lastmessageuuid}
}

func NewBotHandler(bot *traqwsbot.Bot) *BotHandler {
	return &BotHandler{bot: bot}
}
