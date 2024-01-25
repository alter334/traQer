package handler

import "log"

func (h *Handler) GetChannelNameWithParents(channelid string, children string) (channelname string, err error) {
	log.Println(channelname)
	channel, _, err := h.client.ChannelApi.GetChannel(h.auth, channelid).Execute()
	if err != nil {
		log.Println("GetChannelError:", err.Error())
		return "", err
	}

	parent, ok := channel.GetParentIdOk()
	log.Println(ok)

	if parent == nil {
		channelname = channel.GetName() + children
		return channelname, nil
	}
	childname := "/" + channel.GetName() + children
	channelname, err = h.GetChannelNameWithParents(*parent, childname)
	if err != nil {
		log.Println("GetChannelError:", err.Error())
		return "", err
	}
	return channelname, nil

}
