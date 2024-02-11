package bot

import (
	"os"

	traqwsbot "github.com/traPtitech/traq-ws-bot"
)

func NewBotHandler(bot *traqwsbot.Bot) *Bot {
	return &Bot{bot: bot}
}

func BotSetup() *Bot {
	_bot, err := traqwsbot.NewBot(&traqwsbot.Options{
		AccessToken: os.Getenv("BOT_TOKEN"), // Required
	})
	if err != nil {
		panic(err)
	}

	bot := NewBotHandler(_bot)

	return bot
}
