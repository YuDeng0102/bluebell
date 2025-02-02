package controller

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"strconv"
	"web-app/dao/redis"
)

type VoteData struct {
	PostID    string  `json:"post_id"`
	Direction float64 `json:"type"`
}

func VoteHandler(c *gin.Context) {
	var vote VoteData
	if err := c.ShouldBindJSON(&vote); err != nil {
		ResponseErrorWithMsg(c, CodeInvalidParams, err.Error())
		zap.L().Debug("failed to bind json", zap.Error(err))
		return
	}
	userID, ok := c.Get(ContextUserIDKey)
	if !ok {
		ResponseError(c, CodeNotLogin)
		return
	}
	if err := redis.PostVote(vote.PostID, strconv.FormatInt(userID.(int64), 10), vote.Direction); err != nil {
		ResponseErrorWithMsg(c, CodeServerBusy, err.Error())
		return
	}
	ResponseSuccess(c, nil)
}
