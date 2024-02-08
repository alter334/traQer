package service

import (
	"traQer/service/channel"
	"traQer/service/user"
)

func NewService(user *user.UserHandler, channel *channel.ChannelHandler) *Service {
	return &Service{user: user, channel: channel}
}
