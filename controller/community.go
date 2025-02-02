package controller

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"strconv"
	"web-app/dao/mysql"
)

func CommunityHandler(c *gin.Context) {
	communityList, err := mysql.GetCommunityList()
	if err != nil {
		zap.L().Error("logic.GetCommunityList failed", zap.Error(err))
		ResponseError(c, CodeServerBusy)
		return
	}
	ResponseSuccess(c, communityList)
}

func CommunityDetailHandler(c *gin.Context) {
	idstr := c.Param("id")
	id, err := strconv.ParseInt(idstr, 10, 64)
	if err != nil {
		ResponseError(c, CodeInvalidParams)
		return
	}
	community, err := mysql.GetCommunityById(id)
	if err != nil {
		zap.L().Error("logic.GetCommuntyDetail failed", zap.Error(err))
		ResponseError(c, CodeInvalidParams)
		return
	}
	ResponseSuccess(c, community)
}
