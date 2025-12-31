package bot

import traqwsbot "github.com/traPtitech/traq-ws-bot"

func NewBotHandler(bot *traqwsbot.Bot) *Bot {
	return &Bot{bot: bot}
}
