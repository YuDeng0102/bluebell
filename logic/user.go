package logic

import (
	"web-app/dao/mysql"
	"web-app/model"
	snowflake "web-app/pkg/snowfloke"
)

func Register(register *model.ParamRegister) error {
	if err := mysql.CheckUserExist(register.Username); err != nil {
		return err
	}
	userId := snowflake.GenerateID()
	user := model.User{
		Username: register.Username,
		Password: register.Password,
		UserID:   userId,
	}
	if err := mysql.InsertUser(&user); err != nil {
		return err
	}
	return nil
}

func Login(user *model.User) error {
	if err := mysql.Login(user); err != nil {
		return err
	}
	return nil
}
