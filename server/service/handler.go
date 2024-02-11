package service

import (
	"traQer/api"
	"traQer/service/channel"
	"traQer/service/user"
)

func NewService(user *user.UserHandler, channel *channel.ChannelHandler) *Service {
	return &Service{user: user, channel: channel}
}

func ServiceSetup() *Service {
	api := api.ApiSetup()

	user := user.NewUserHandler(api)
	channel := channel.NewChannelHandler(api)

	service := NewService(user, channel)

	return service
}
