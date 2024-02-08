package channel

import "traQer/api"

func NewChannelHandler(api *api.ApiHandler) *ChannelHandler {
	return &ChannelHandler{api: api}
}
