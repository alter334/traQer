package db

import "traQer/model"

func (d *DBHandler) UserInsert(userMessages *[]model.UserDetailWithMessageCount) error{

	for _, message := range &(userMessages) {
		_, err = h.db.Exec("INSERT INTO `messagecounts`(`totalpostcounts`,`userid`,`dailypostcounts`) VALUES(?,?,?) ON DUPLICATE KEY UPDATE `totalpostcounts`=VALUES(totalpostcounts)", message.TotalMessageCount, message.User.Id, 0)
		if err != nil {
			log.Println("Internal error:", err.Error())
			return
		}
	}
	return nil
}
