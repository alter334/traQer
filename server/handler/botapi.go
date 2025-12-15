package handler

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
	"time"
	"unicode/utf8"

	"github.com/traPtitech/go-traq"
)

type Messagewithlen struct {
	Len int
	Id  string
}

// ----------------------------------------------------------------
// traQ操作系
// ----------------------------------------------------------------
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
		return ""
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

//----------------------------------------------------------------
// DM送信
//----------------------------------------------------------------

func (b *BotHandler) BotDM(userid string, content string) {
	_, r, err := b.bot.API().
		MessageApi.
		PostDirectMessage(context.Background(), userid).
		PostMessageRequest(traq.PostMessageRequest{
			Content: content,
		}).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}

//----------------------------------------------------------------
// traQAPI叩く系
//----------------------------------------------------------------

// User名をUUIDから返す
func (b *BotHandler) BotGetUserName(postUserID string) (userName string) {
	userdetail, httpres, err := b.bot.API().UserApi.GetUser(context.Background(), postUserID).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", httpres)
	}
	return userdetail.GetName()
}

// MessageをUUIDから返す
func (b *BotHandler) BotGetMessage(messageuuid string) (message *traq.Message) {
	message, httpres, err := b.bot.API().MessageApi.GetMessage(context.Background(), messageuuid).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", httpres)
	}
	return message
}

// user名からユーザUUIDを取得
func (b *BotHandler) BotGetUserUUID(userName string) (useruuid string) {
	users, httpres, err := b.bot.API().UserApi.GetUsers(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", httpres)
	}
	for _, u := range users {
		if u.Name == userName {
			return u.Id
		}
	}
	return ""
}

// group名からグループUUIDを取得
func (b *BotHandler) BotGetGroupUUID(groupName string) (groupuuid string) {
	groups, httpres, err := b.bot.API().GroupApi.GetUserGroups(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", httpres)
	}
	for _, g := range groups {
		if g.Name == groupName {
			return g.Id
		}
	}
	return ""
}

// groupUUIDからグループ所属者を返す
func (b *BotHandler) BotGetGroupMembers(groupid string) (groupmembersids []string) {
	usergroupmember, httpres, err := b.bot.API().GroupApi.GetUserGroupMembers(context.Background(), groupid).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", httpres)
	}
	for _, member := range usergroupmember {
		groupmembersids = append(groupmembersids, member.GetId())
	}
	return groupmembersids

}

// UserのホームチャンネルUUIDを返す
func (b *BotHandler) GetUserHome(userID string) (homeUUID string) {
	userdetail, httpres, err := b.bot.API().UserApi.GetUser(context.Background(), userID).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", httpres)
	}
	return userdetail.GetHomeChannel()
}

// スタンプUUIDからスタンプ名を返す
func (b *BotHandler) BotGetStampName(stampID string) (stampName string) {
	stamp, httpres, err := b.bot.API().StampApi.GetStamp(context.Background(), stampID).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", httpres)
		return stampID // エラーの場合はUUIDをそのまま返す
	}
	return stamp.GetName()
}

// ----------------------------------------------------------------
// メッセージ収集関連
// ----------------------------------------------------------------

// user名から指定の文字数より長いメッセージのuuid配列とメッセージ内容を返す
func (b *BotHandler) BotGetLongMessages(username string, length int) (messageuuids []string, content string) {
	ct := 0
	maxlen := 0
	userid := b.BotGetUserUUID(username) //uuidの取得
	if userid == "" {
		return messageuuids, "Please insert valid username"
	}
	collections := "" //取得したメッセージのリンクと文字数を記録する
	longmessages := []Messagewithlen{}
	for i := 0; ; i += 100 {
		messages, err := b.BotGetUserMessages(userid, i)
		if err != nil {
			return messageuuids, "Internal Error:" + err.Error()
		}

		// 受信したメッセージの読み取り 長さが規定文字数以上かどうか判定する
		// メッセージを格納するmap:keyを文字数,elemをuuidに

		for _, message := range messages.Hits {
			len := utf8.RuneCountInString(message.Content)
			if len > maxlen {
				maxlen = len
			}
			if len >= length {
				longmessages = append(longmessages, Messagewithlen{Len: len, Id: message.Id})
				messageuuids = append(messageuuids, message.Id)
				ct++
			}
		}
		if len(messages.Hits) < 100 || i == 9900 {
			break
		}
	}

	//取得したメッセージの降順取得
	sort.Slice(longmessages, func(i, j int) bool { return longmessages[i].Len > longmessages[j].Len })

	//メッセージの作成
	for i, message := range longmessages {
		collections += "文字数:" + strconv.Itoa(message.Len) + "\n" + "https://q.trap.jp/messages/" + message.Id + "\n"
		if i == 4 {
			break
		}
	}
	content += "## :@" + username + ": " + username + "の長文投稿一覧\n" +
		"(指定された文字数:" + strconv.Itoa(length) + ",該当する投稿数:" + strconv.Itoa(ct) +
		",最大文字数:" + strconv.Itoa(maxlen) + ")\n" +
		collections

	return messageuuids, content
}

// 収集したメッセージ群から情報を求める
func (b *BotHandler) BotWUserrank(after string, before string, channeltopost string) string {

	messageCountperUser := map[string]int{} //ランキング用mapの生成
	var pl MessageCountPairList
	response := "searching... :loading:"                     //結果表示用文字列
	responseuuid := b.BotSimplePost(channeltopost, response) //返信用メッセージの作成
	channelid := "5b7b8143-7c0d-4ade-8658-3a8d8ce4dd83"      //gps/trend/w のチャンネルid
	botpikatestid := b.BotGetUserUUID("BOT_pika_test")       //BOT_pika_testのユーザid
	if after == "" {
		after = "20200101" //空白処理
	}
	if before == "" {
		before = time.Now().AddDate(0, 0, 2).Format("20060102") //空白処理
	}

	aftertime, err := time.Parse("20060102", after)
	if err != nil {
		return "ATime parsing Error:" + err.Error()
	}
	beforetime, err := time.Parse("20060102", before)
	if err != nil {
		return "BTime parsing Error:" + err.Error()
	}

	for i := 0; ; i += 10 {
		messages, err := b.BotGetChannelMessagesWithQuote(channelid, i)
		if err != nil {
			return "Polling Error:" + err.Error()
		}

		// 受信したメッセージの読み取り 規定条件を満たすか確認
		// "URL数が10","Bot_pika_testによるもの","指定日付以前以降"

		for _, message := range messages.Hits {
			if message.UserId != botpikatestid {
				continue
			}
			if !message.CreatedAt.After(aftertime) {
				continue
			}
			if !message.CreatedAt.Before(beforetime) {
				break
			}
			urls := strings.Fields(message.Content)
			if len(urls) != 10 {
				continue
			}

			//メッセージの取り出しと集計
			for _, messageurl := range urls {
				messageuuid := strings.Split(messageurl, "/")[4]
				message := b.BotGetMessage(messageuuid)
				if message != nil {
					messageCountperUser[message.UserId]++
				}
			}

		}

		// ランキング集計状況の更新
		//mapのpair化
		pl = make(MessageCountPairList, len(messageCountperUser))
		j := 0
		for k, v := range messageCountperUser {
			pl[j] = MessageCountPair{k, v}
			j++
		}

		sort.Sort(sort.Reverse(pl)) //pair化したmapのソート

		//pairを元に返信の書き出し
		response = "collecting..." + strconv.Itoa(i) + "\nsearching... :loading:\n| rank | username | total |\n| - | - | - |\n" //基礎
		for x, list := range pl {
			response += ("|" + strconv.Itoa(x) + "|:@" + b.BotGetUserName(list.Key) + ":|" + strconv.Itoa(list.Value) + "|\n")
			if x >= 50 {
				break
			}
		}
		b.BotSimpleEdit(responseuuid, response)

		if len(messages.Hits) < 10 || i == 9990 {
			break
		}

	}

	response = "| rank | username | total |\n| - | - | - |\n" //基礎
	for i, list := range pl {
		response += ("|" + strconv.Itoa(i+1) + "|:@" + b.BotGetUserName(list.Key) + ":|" + strconv.Itoa(list.Value) + "|\n")
		if i >= 49 {
			break
		}
	}
	b.BotSimpleEdit(responseuuid, response)

	return response
}

// 選択したuserからメッセージを100件取得する(Bot版)
func (b *BotHandler) BotGetUserMessages(userid string, offset int) (message *traq.MessageSearchResult, err error) {
	messages, _, err := b.bot.API().MessageApi.SearchMessages(context.Background()).
		From(userid).Limit(100).Offset(int32(offset)).
		Sort(`createdAt`).Execute()
	if err != nil {
		return messages, err
	}
	log.Println("取得数:", len(messages.Hits))
	log.Println("取得mes:", messages.TotalHits)
	return messages, nil
}

// 指定した期間内のメッセージを100件取得する(Bot版)
func (b *BotHandler) BotGetMessagesBeteween(after string, before string, offset int) (message *traq.MessageSearchResult, err error) {
	if after == "" {
		after = "20200101" //空白処理
	}
	if before == "" {
		before = time.Now().AddDate(0, 0, 2).Format("20060102") //空白処理
	}

	aftertime, err := time.Parse("20060102", after)
	if err != nil {
		return nil, err
	}
	beforetime, err := time.Parse("20060102", before)
	if err != nil {
		return nil, err
	}
	messages, _, err := b.bot.API().MessageApi.SearchMessages(context.Background()).
		After(aftertime).Before(beforetime).Limit(100).Offset(int32(offset)).
		Sort(`createdAt`).Execute()
	if err != nil {
		return messages, err
	}
	log.Println("取得数:", len(messages.Hits))
	log.Println("取得mes:", messages.TotalHits)
	return messages, nil
}

// スタンプ数が特定以上のメッセージを取得する スタンプ合計数,スタンプ合計種類数,指定日時より前/後の設定ができる

func (b *BotHandler) BotGetStampedMessage(total int, kind int, maxmes int, after string, before string, channeltopost string) (jsonString string, err error) {
	response := "searching... :loading:"                     //結果表示用文字列
	responseuuid := b.BotSimplePost(channeltopost, response) //返信用メッセージの作成
	type StampInfo struct {
		Name  string `json:"name"`
		Count int    `json:"count"`
	}

	type MessageInfo struct {
		Name   string      `json:"name"`
		Stamps []StampInfo `json:"stamps"`
	}

	var result []MessageInfo

	for i := 0; ; i += 100 {
		hit, err := b.BotGetMessagesBeteween(after, before, i%10000)
		if err != nil {
			return "", err
		}

		for _, message := range hit.Hits {
			stampTotal := 0

			// スタンプ情報を収集（同じ名前のスタンプをまとめる）
			stampMap := make(map[string]int)
			for _, stamp := range message.Stamps {
				// スタンプUUIDからスタンプ名を取得
				stampName := b.BotGetStampName(stamp.StampId)
				stampMap[stampName]++
			}

			var stampSlice []StampInfo
			for name, count := range stampMap {
				stampSlice = append(stampSlice, StampInfo{name, count})
			}
			sort.Slice(stampSlice, func(i, j int) bool {
				return stampSlice[i].Count > stampSlice[j].Count
			})

			// 上位5件だけ残す
			if len(stampSlice) > 5 {
				stampSlice = stampSlice[:5]
			}

			stampTotal = 0
			for _, stamp := range stampSlice {
				stampTotal += stamp.Count
			}
			stampKinds := len(stampMap)
			if stampTotal >= total && stampKinds >= kind {

				// スライスをStampInfoに変換
				var stamps []StampInfo
				stamps = append(stamps, stampSlice...)

				result = append(result, MessageInfo{
					Name:   message.Id,
					Stamps: stamps,
				})
			}
		}
		if (i % 10000) == 9900 {
			// beforeを最後に取得したメッセージの日時に更新する
			if len(hit.Hits) > 0 {
				before = hit.Hits[len(hit.Hits)-1].CreatedAt.Format("20060102")
				log.Println("Updating before to:", before)
			}
		}
		if len(hit.Hits) < 98 || len(result) >= maxmes {
			// 動作安定化のため98件で判定
			break
		}
	}

	log.Println("Filtered取得数:", len(result))

	// JSONにエンコード
	jsonBytes, err := json.MarshalIndent(result, "", "  ")
	if err != nil {
		return "", err
	}
	resstr := "```\n" + string(jsonBytes) + "\n```"
	for _, collect := range result {
		resstr += "https://q.trap.jp/messages/" + collect.Name + "\n"
	}

	b.BotSimpleEdit(responseuuid, resstr)

	return string(jsonBytes), nil
}

//------------------------------------------------
// チャンネル監視関連
//------------------------------------------------

// 指定したチャンネルからメッセージを100件取得する
func (b *BotHandler) BotGetChannelMessages(channelid string, offset int) (message *traq.MessageSearchResult, err error) {
	messages, _, err := b.bot.API().MessageApi.SearchMessages(context.Background()).
		In(channelid).Limit(10).Offset(int32(offset)).
		Sort(`createdAt`).Execute()
	if err != nil {
		return messages, err
	}
	log.Println("取得数:", len(messages.Hits))
	log.Println("取得mes:", messages.TotalHits)
	return messages, nil
}

// 引用ありのみ
func (b *BotHandler) BotGetChannelMessagesWithQuote(channelid string, offset int) (message *traq.MessageSearchResult, err error) {
	messages, _, err := b.bot.API().MessageApi.SearchMessages(context.Background()).
		In(channelid).Limit(10).Offset(int32(offset)).HasURL(true).
		Sort(`-createdAt`).Execute()
	if err != nil {
		return messages, err
	}
	log.Println("取得数:", len(messages.Hits))
	log.Println("取得mes:", messages.TotalHits)
	return messages, nil
}
