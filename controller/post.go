package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"strconv"
	"web-app/dao/mysql"
	"web-app/dao/redis"
	"web-app/logic"
	"web-app/model"
)

func CreatePostHandler(c *gin.Context) {
	post := new(model.Post)
	if err := c.ShouldBindJSON(post); err != nil {
		ResponseError(c, CodeInvalidParams)
		return
	}
	userID, ok := c.Get(ContextUserIDKey)
	if !ok {
		ResponseError(c, CodeNotLogin)
		return
	}
	post.AuthorId = userID.(int64)
	communties, err := mysql.GetCommunityList()
	if err != nil {
		ResponseError(c, CodeInvalidParams)
		return
	}
	ok = false
	for _, communty := range communties {
		if communty.ID == post.CommunityId {
			ok = true
			break
		}
	}
	if !ok {
		ResponseError(c, CodeInvalidParams)
		return
	}
	if err := logic.CreatePost(post); err != nil {
		zap.L().Error("logic.CreatePost failed", zap.Error(err))
		ResponseError(c, CodeServerBusy)
		return
	}

	ResponseSuccess(c, nil)
}

func GetPostHandler(c *gin.Context) {
	postID, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		ResponseError(c, CodeInvalidParams)
		return
	}
	post, err := logic.GetPost(postID)
	if err != nil {
		zap.L().Error("logic.GetPost(postID) failed", zap.Int64("postId", postID), zap.Error(err))
		ResponseError(c, CodeInvalidParams)
		return
	}
	ResponseSuccess(c, post)
}

func GetPostListHandler(c *gin.Context) {
	pageSizeStr := c.DefaultQuery("pageSize", "10")
	pageNumStr := c.DefaultQuery("pageNum", "1")
	pageSize, err := strconv.Atoi(pageSizeStr)
	if err != nil {
		ResponseError(c, CodeInvalidParams)
		return
	}
	pageNum, err := strconv.Atoi(pageNumStr)
	if err != nil {
		ResponseError(c, CodeInvalidParams)
		return
	}
	data, err := logic.GetPostList(pageNum, pageSize)
	if err != nil {
		zap.L().Error("logic.GetPostList failed", zap.Error(err))
		ResponseError(c, CodeInvalidParams)
		return
	}
	ResponseSuccess(c, data)
}

func GetPostList2Handler(c *gin.Context) {
	order, _ := c.GetQuery("order")
	pageStr, ok := c.GetQuery("page")
	if !ok {
		pageStr = "1"
	}
	pageNum, err := strconv.ParseInt(pageStr, 10, 64)
	if err != nil {
		pageNum = 1
	}
	posts := redis.GetPost(order, pageNum)
	fmt.Println(len(posts))
	ResponseSuccess(c, posts)

}
