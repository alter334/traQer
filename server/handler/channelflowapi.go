package handler

import "log"

func (h *Handler) GetChannelNameWithParents(channelid string, children string) (channelname string, err error) {
	log.Println(channelname)
	channel, _, err := h.client.ChannelApi.GetChannel(h.auth, channelid).Execute()
	if err != nil {
		log.Println("GetChannelError:", err.Error())
		return "", err
	}

	parent, _ := channel.GetParentIdOk()

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

func (h *Handler) GetViewers(channelid string) (users []string, err error) {
	userstats, _, err := h.client.ChannelApi.GetChannelViewers(h.auth, channelid).Execute()
	if err != nil {
		log.Println("GetViewersError:", err.Error())
		return users, err
	}
	for _, stat := range userstats {
		users = append(users, stat.UserId)
	}

	return users, nil
}
