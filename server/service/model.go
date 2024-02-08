package service

import (
	"traQer/service/channel"
	"traQer/service/user"
)

type Service struct {
	user    *user.UserHandler
	channel *channel.ChannelHandler
}
