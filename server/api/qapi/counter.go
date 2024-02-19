package qapi

import (
	"log"
	"traQer/model"
)

// traQAPI(/users/{userid}/stats)を叩きまくる *高負荷かつ長時間
func (q *QapiHandler) GetUserPostCount() *[]model.QUserMessages {

	//ユーザリスト取得
	v, _, err := q.client.UserApi.GetUsers(q.auth).Execute()

	if err != nil {
		log.Println("Internal error-Qapi error:", err.Error())
		return nil
	}

	log.Println("Userlistavailable:", len(v))

	// traqAPIのUser型に対応したstructに
	var userMessages []model.QUserMessages

	// BOTは排除
	for i, user := range v {
		if user.Bot {
			log.Println("isBot:", i)
			continue
		}
		userStats, _, err := q.client.UserApi.GetUserStats(q.auth, user.Id).Execute()
		if err != nil {
			log.Println("Internal error-Qapi error:", err.Error())
			return nil
		}
		var message model.QUserMessages
		message.User = user
		message.TotalMessageCount = userStats.TotalMessageCount

		userMessages = append(userMessages, message)
		log.Println("traQing:", i, "mescount:", message.TotalMessageCount)
	}

	return &userMessages

}

func (q *QapiHandler) GetUserDetails(userMessages *[]model.QUserMessages) *[]model.UserDetailWithMessageCount {
	userdetails := []model.UserDetailWithMessageCount{}
	for _, userMessage := range *userMessages {
		user:=userMessage.User
		userdetails = append(userdetails,model.UserDetailWithMessageCount{
			Id: user.Id,
			TotalMessageCount: userMessage.TotalMessageCount,
		})

	}
	return nil
}
