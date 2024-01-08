package handler

import (
	"log"
	"os"
	"strings"
	"time"

	traqwsbot "github.com/traPtitech/traq-ws-bot"
	"github.com/traPtitech/traq-ws-bot/payload"
)

func NewBot() *traqwsbot.Bot {
	bot, err := traqwsbot.NewBot(&traqwsbot.Options{
		AccessToken: os.Getenv("BOT_TOKEN"), // Required
	})
	if err != nil {
		panic(err)
	}
	return bot
}

func (h *Handler) BotHandler() {

	log.Println(time.Now())

	h.b.bot.OnError(func(message string) {
		log.Println("Received ERROR message: " + message)
	})

	h.b.bot.OnMessageCreated(func(p *payload.MessageCreated) {
		log.Println("Received MESSAGE_CREATED event: " + p.Message.Text)
		cmd := strings.Fields(p.Message.Text)

		//コマンドなし->通常モード(attackコマンドでも同様)
		switch len(cmd) {
		case 1:
			message := h.b.BotSimplePost(p.Message.ChannelID, "Nowcollecting...")
			h.b.BotSimpleEdit(message, h.BotCollectUserRank(""))
		case 2:
			message := h.b.BotSimplePost(p.Message.ChannelID, "Nowcollecting...")
			h.b.BotSimpleEdit(message, h.BotCollectUserRank(cmd[1]))

		// case 2:
		// 	switch cmd[1] {
		// 	case "dbenroll":
		// 		if p.Message.User.Name != "Alt--er" {
		// 			handler.SimplePost(bot, p.Message.ChannelID, "This command isn't allowed")
		// 			return
		// 		}
		// 		h.EnrollExistingUserHometoPlace()
		// 	}
		default: //現在はコマンド機能は導入していないので
		}

	})

	if err := h.b.bot.Start(); err != nil {
		panic(err)
	}

}
