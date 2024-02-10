package qapi

import "log"

func (q *QapiHandler) GetUserPostCount() {

	//ユーザリスト取得
	v, _, err := q.client.UserApi.GetUsers(q.auth).Execute()

	if err != nil {
		log.Println("Internal error-Qapi error:", err.Error())
		return
	}

	log.Println("Userlistavailable:", len(v))

	var userMessages []UserMessages

	for i, user := range v {
		if user.Bot {
			log.Println("isBot:", i)
			continue
		}
		userStats, _, err := q.client.UserApi.GetUserStats(q.auth, user.Id).Execute()
		if err != nil {
			log.Println("Internal error-Qapi error:", err.Error())
			return
		}
		var message UserMessages
		message.User = user
		message.TotalMessageCount = userStats.TotalMessageCount

		userMessages = append(userMessages, message)
		log.Println("traQing:", i, "mescount:", message.TotalMessageCount)
	}

	
}
