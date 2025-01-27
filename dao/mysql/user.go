package mysql

import (
	"crypto/md5"
	"encoding/hex"
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

func Login(loginUser *model.User) error {
	user := new(model.User)
	result := db.Where("username = ?", loginUser.Username).First(&user)
	if result.RowsAffected == 0 {
		return ErrorUserNotExists
	} else if user.Password != encryptPassword(loginUser.Password) {
		//log.Printf("%v's password: %v invalid", loginUser.Username, loginUser.Password)
		return ErrorInvalidPassword
	} else if result.Error != nil {
		return result.Error
	}
	return nil
}

func encryptPassword(password string) string {
	h := md5.New()
	h.Write([]byte(secrect))
	return hex.EncodeToString(h.Sum([]byte(password)))
}
