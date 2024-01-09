package handler

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/traPtitech/go-traq"
)

func (b *BotHandler) BotSimplePost(channelID string, content string) (messageid string) {
	q, r, err := b.bot.API().
		MessageApi.
		PostMessage(context.Background(), channelID).
		PostMessageRequest(traq.PostMessageRequest{
			Content: content,
		}).
		Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	return q.Id
}

func (b *BotHandler) BotSimpleEdit(messageid string, content string) {
	_, err := b.bot.API().
		MessageApi.EditMessage(context.Background(), messageid).PostMessageRequest(traq.PostMessageRequest{
		Content: content,
	}).Execute()
	if err != nil {
		log.Println("Internal error:", err.Error())
	}
	log.Println("Done")
}

// User名をUUIDから返す
func (b *BotHandler) BotGetUserName(postUserID string) (userName string) {
	userdetail, httpres, err := b.bot.API().UserApi.GetUser(context.Background(), postUserID).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", httpres)
	}
	return userdetail.GetName()
}

// group名からグループUUIDを取得
func (b *BotHandler) BotGetGroupUUID(groupName string) (groupuuid string) {
	groups, httpres, err := b.bot.API().GroupApi.GetUserGroups(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", httpres)
	}
	for _,g:= range groups {
		if g.Name==groupName {
			return g.Id
		}
	}
	return ""
}
// groupUUIDからグループ所属者を返す
func (b *BotHandler) BotGetGroupMembers(groupid string) (groupmembersids []string){
	usergroupmember, httpres, err := b.bot.API().GroupApi.GetUserGroupMembers(context.Background(),groupid).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", httpres)
	}
	for _,member := range usergroupmember{
		groupmembersids = append(groupmembersids,member.GetId())
	}
	return groupmembersids

}
