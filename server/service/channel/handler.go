package channel

import "traQer/api"

type ChannelHandler struct {
	api *api.ApiHandler
}

func NewChannelHandler(api *api.ApiHandler) *ChannelHandler {
	return &ChannelHandler{api: api}
}
