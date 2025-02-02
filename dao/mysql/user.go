package mysql

import (
	"crypto/md5"
	"encoding/hex"
	"log"
	"web-app/model"
)

const secrect = "yudeng"

func CheckUserExist(username string) error {
	user := new(model.User)
	result := db.Where("username = ?", username).First(&user)
	if result.RowsAffected > 0 {
		return ErrorUserExist
	}
	return nil
}

func InsertUser(user *model.User) error {
	user.Password = encryptPassword(user.Password)
	result := db.Create(user)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func Login(user *model.User) error {
	QueryUser := new(model.User)
	result := db.Where("username = ?", user.Username).First(&QueryUser)
	//zap.L().Info("userID", zap.Int64("userID", user.UserID))
	if result.RowsAffected == 0 {
		return ErrorUserNotExists
	} else if QueryUser.Password != encryptPassword(user.Password) {
		log.Printf("%v's password: %v invalid", user.Username, user.Password)
		return ErrorInvalidPassword
	} else if result.Error != nil {

		return result.Error
	}
	user.UserID = QueryUser.UserID
	return nil
}

func GetAuthorName(userID int64) (string, error) {
	user := new(model.User)
	result := db.Where("user_id = ?", userID).First(&user)
	if result.Error != nil || result.RowsAffected == 0 {
		return "", ErrorUserNotExists
	}
	return user.Username, nil
}

func encryptPassword(password string) string {
	h := md5.New()
	h.Write([]byte(secrect))
	return hex.EncodeToString(h.Sum([]byte(password)))
}
