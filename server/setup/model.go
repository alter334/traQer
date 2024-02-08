package setup

import (
	"traQer/bot"
	"traQer/service"
)

type Server struct {
	bot     *bot.Bot
	service *service.Service
}
