package repo

import (
	"im/internal/consts"
	"im/internal/models"

	"gorm.io/gorm"
)

// 查詢好友列表
func GetFrindListRepo(db *gorm.DB, userId int) []models.User {
	friends := make([]models.User, 0)
	subQuery1 := db.Table("friends").Select("UserID2").
		Where("UserID1 = ? AND FriendshipStatus = ?", userId, consts.FriendshipStatusAccepted)
	db.Table("users").Where("id IN (?)", subQuery1).Find(&friends)

	// 重新設計, 好友與好友之間關係改為兩筆資料,方便確認阻擋以及被阻擋
	// subQuery2 := db.Table("friends").Select("UserID1").
	// 	Where("UserID2 = ? AND FriendshipStatus = ?", userId, consts.FriendshipStatusAccepted)
	// db.Table("users").Where("id IN (?) OR id IN (?)", subQuery1, subQuery2).Find(&friends)

	return friends
}
