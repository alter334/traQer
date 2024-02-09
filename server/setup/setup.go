package setup

import (
	"traQer/bot"
	"traQer/service"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"github.com/robfig/cron/v3"
)

func Setup() *Server {
	bot := bot.BotSetup()
	service := service.ServiceSetup()

	return &Server{bot: bot, service: service}
}

func (s *Server) CronSetup(c *cron.Cron) {
	
}

func (s *Server) EchoSetup(e *echo.Echo) {
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
}
