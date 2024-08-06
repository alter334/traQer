package handler

import (
	"log"
	"os"
	"strconv"
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
			if cmd[1] == "\\dm\\enroll" {
				h.b.BotDM(p.Message.User.ID, h.BotDMSubscribe(p.Message.User.ID, 10))
				h.b.BotSimplePost(p.Message.ChannelID, "DMに詳細を送付しました")
				break
			} else if cmd[1] == "\\dm\\unsubscribe" {
				h.b.BotDM(p.Message.User.ID, h.BotDMUnSubscribe(p.Message.User.ID))
				h.b.BotSimplePost(p.Message.ChannelID, "DMに詳細を送付しました")
				break
			} else if cmd[1] == ":w:" {
				h.b.BotWUserrank("", "", p.Message.ChannelID)
				break
			} else if cmd[1] == "\\tag" {
				message := h.b.BotSimplePost(p.Message.ChannelID, "Nowcollecting...")
				h.b.BotSimpleEdit(message, h.BotCollectTagRank(""))
				break
			} else if cmd[1] == "\\tagper" {
				message := h.b.BotSimplePost(p.Message.ChannelID, "Nowcollecting...")
				h.b.BotSimpleEdit(message, h.BotCollectTagRateRank("", 0))
				break
			}
			message := h.b.BotSimplePost(p.Message.ChannelID, "Nowcollecting...")
			h.b.BotSimpleEdit(message, h.BotCollectUserRank(cmd[1]))
		case 3:
			if cmd[1] == "\\dm\\enroll" {
				amount, err := strconv.Atoi(cmd[2])
				if err != nil || amount < 0 {
					h.b.BotSimplePost(p.Message.ChannelID, "Insert valid number")
					break
				}
				h.b.BotDM(p.Message.User.ID, h.BotDMSubscribe(p.Message.User.ID, amount))
				h.b.BotSimplePost(p.Message.ChannelID, "DMに詳細を送付しました")

			} else if cmd[1] == "\\long" {
				message := h.b.BotSimplePost(p.Message.ChannelID, "Nowcollecting...")
				_, content := h.b.BotGetLongMessages(cmd[2], 1000)
				h.b.BotSimpleEdit(message, content)
			} else if cmd[1] == ":w:" {
				h.b.BotWUserrank(cmd[2], "", p.Message.ChannelID) //after のみ
				break
			} else if cmd[1] == "\\tag" {
				message := h.b.BotSimplePost(p.Message.ChannelID, "Nowcollecting...")
				h.b.BotSimpleEdit(message, h.BotCollectTagRank(cmd[2]))
				break
			} else if cmd[1] == "\\tagper" {
				message := h.b.BotSimplePost(p.Message.ChannelID, "Nowcollecting...")
				h.b.BotSimpleEdit(message, h.BotCollectTagRateRank(cmd[2], 0))
				break
			} else if cmd[1] == "\\tagperwithamount" {
				message := h.b.BotSimplePost(p.Message.ChannelID, "Nowcollecting...")
				amount, err := strconv.Atoi(cmd[2])
				if err != nil || amount < 0 {
					h.b.BotSimplePost(p.Message.ChannelID, "Insert valid number")
					break
				}
				h.b.BotSimpleEdit(message, h.BotCollectTagRateRank("", amount))
				break
			} else {
				h.b.BotSimplePost(p.Message.ChannelID, "Insert valid commands")
			}
		case 4:
			if cmd[1] == "\\long" {
				cmdint, err := strconv.Atoi(cmd[3])
				if err != nil {
					h.b.BotSimplePost(p.Message.ChannelID, "Insert valid commands")
				} else {
					message := h.b.BotSimplePost(p.Message.ChannelID, "Nowcollecting...")
					_, content := h.b.BotGetLongMessages(cmd[2], cmdint)
					h.b.BotSimpleEdit(message, content)
				}
			} else if cmd[1] == ":w:" {
				h.b.BotWUserrank(cmd[2], cmd[3], p.Message.ChannelID) //after before
				break
			} else if cmd[1] == "\\tagperwithamount" {
				message := h.b.BotSimplePost(p.Message.ChannelID, "Nowcollecting...")
				amount, err := strconv.Atoi(cmd[3])
				if err != nil || amount < 0 {
					h.b.BotSimplePost(p.Message.ChannelID, "Insert valid number")
					break
				}
				h.b.BotSimpleEdit(message, h.BotCollectTagRateRank(cmd[2], amount))
				break
			} else {
				h.b.BotSimplePost(p.Message.ChannelID, "Insert valid commands")
			}
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
