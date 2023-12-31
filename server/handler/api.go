package handler

import (
	"context"
	"fmt"
	"os"

	"github.com/traPtitech/go-traq"
	traqwsbot "github.com/traPtitech/traq-ws-bot"
)

func SimplePost(bot *traqwsbot.Bot, channelID string, content string) (x string) {
	q, r, err := bot.API().
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

func SimpleEdit(bot *traqwsbot.Bot, m string, s string) {
	bot.API().
		MessageApi.EditMessage(context.Background(), m).PostMessageRequest(traq.PostMessageRequest{
		Content: s,
	}).Execute()
}

//UserのホームチャンネルUUIDとユーザー名を返す:現在不使用
func GetUserHome(bot *traqwsbot.Bot, postUserID string) (homeUUID string,userName string){
	userdetail, httpres, err := bot.API().UserApi.GetUser(context.Background(), postUserID).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", httpres)
	}
	return userdetail.GetHomeChannel() ,userdetail.GetName()
}

//User名をUUIDから返す
func GetUserName(bot *traqwsbot.Bot, postUserID string) (userName string) {
	userdetail, httpres, err := bot.API().UserApi.GetUser(context.Background(), postUserID).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", httpres)
	}
	return userdetail.GetName()
}

