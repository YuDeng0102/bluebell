package mysql

import (
	"go.uber.org/zap"
	"web-app/model"
)

func CreatePost(p *model.Post) error {
	result := db.Create(p)
	if result.Error != nil {
		zap.L().Error("create post failed", zap.Error(result.Error))
		return result.Error
	}
	return nil
}

func GetPost(id int64) (*model.Post, error) {
	post := new(model.Post)
	result := db.Where("id=?", id).First(post)
	if result.Error != nil {
		zap.L().Error("get post failed", zap.Error(result.Error))
		return nil, result.Error
	}
	//postDetail := new(model.ApiPostDetail)
	//postDetail.Post = post
	return post, nil
}

func GetPostList(pageNum, pageSize int) (list []model.Post, err error) {
	result := db.Limit(pageSize).Offset((pageNum - 1) * pageSize).Find(&list)
	if result.Error != nil || result.RowsAffected == 0 {
		zap.L().Error("get post list failed", zap.Error(result.Error))
		return nil, result.Error
	}
	return
}
