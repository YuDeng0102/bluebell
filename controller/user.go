package controller

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"go.uber.org/zap"
	"web-app/dao/mysql"
	"web-app/logic"
	"web-app/model"
)

// RegisterHanndler 用户注册
func RegisterHanndler(c *gin.Context) {
	var p model.ParamRegister
	if err := c.ShouldBind(&p); err != nil {
		zap.L().Error("SignUp with invalid param", zap.Error(err))
		errs, ok := err.(validator.ValidationErrors)
		if !ok {
			ResponseError(c, CodeInvalidParams)
			return
		}
		ResponseErrorWithMsg(c, CodeInvalidParams, errs.Error())
		return
	}
	if err := logic.Register(&p); err != nil {
		if errors.Is(err, mysql.ErrorUserExist) {
			ResponseError(c, CodeUserExist)
			return
		}
		ResponseError(c, CodeServerBusy)
		return
	}
	ResponseSuccess(c, nil)
}

// LoginHanndler 用户注册
func LoginHanndler(c *gin.Context) {
	var p model.ParamLogin
	if err := c.ShouldBind(&p); err != nil {
		zap.L().Error("SignUp with invalid param", zap.Error(err))
		errs, ok := err.(validator.ValidationErrors)
		if !ok {
			ResponseError(c, CodeInvalidParams)
			return
		}
		ResponseErrorWithMsg(c, CodeInvalidParams, errs.Error())
		return
	}
	if err := logic.Login(&p); err != nil {
		zap.L().Error("service.Login failed.", zap.String("用户:", p.Username), zap.Error(err))
		if errors.Is(err, mysql.ErrorUserNotExists) {
			ResponseError(c, CodeUserNotExist)
			return
		}
		ResponseError(c, CodeInvalidPassword)
		return
	}
	ResponseSuccess(c, nil)
}
