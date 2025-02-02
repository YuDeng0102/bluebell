package logic

import (
	"fmt"
	"go.uber.org/zap"
	"web-app/dao/mysql"
	"web-app/dao/redis"
	"web-app/model"
	"web-app/pkg/snowfloke"
)

func CreatePost(p *model.Post) error {
	// 1 生成postId
	p.ID = snowflake.GenerateID()
	if err := mysql.CreatePost(p); err != nil {
		return err
	}
	community, err := mysql.GetCommunityById(p.CommunityId)
	if err != nil {
		zap.L().Error("mysql.GetCommunityNameByID failed", zap.Error(err))
		return err
	}
	if err := redis.CreatePost(
		fmt.Sprint(p.ID),
		fmt.Sprint(p.AuthorId),
		p.Title,
		p.Content,
		community.CommunityName,
	); err != nil {
		zap.L().Error("redis.CreatePost failed", zap.Error(err))
		return err
	}
	return nil
}

func GetPost(id int64) (*model.ApiPostDetail, error) {
	post, err := mysql.GetPost(id)
	if err != nil {
		zap.L().Error("mysql.GetPost(id) failed", zap.Error(err))
		return nil, err
	}

	AuthorName, err := mysql.GetAuthorName(post.AuthorId)
	if err != nil {
		zap.L().Error("mysql.GetAuthorName(post.AuthorId) failed", zap.Error(err))
		return nil, err
	}

	community, err := mysql.GetCommunityById(post.CommunityId)
	if err != nil {
		zap.L().Error("mysql.GetCommunityById(post.CommunityId) failed", zap.Error(err))
		return nil, err
	}

	return &model.ApiPostDetail{
		Post:          post,
		AuthorName:    AuthorName,
		CommunityName: community.CommunityName,
	}, nil
}

func GetPostList(pageNum, size int) ([]*model.ApiPostDetail, error) {
	posts, err := mysql.GetPostList(pageNum, size)
	if err != nil {
		return nil, err
	}
	list := make([]*model.ApiPostDetail, 0, len(posts))
	for _, post := range posts {
		AuthorName, err := mysql.GetAuthorName(post.AuthorId)
		if err != nil {
			zap.L().Error("mysql.GetAuthorName(post.AuthorId) failed", zap.Error(err))
			return nil, err
		}
		community, err := mysql.GetCommunityById(post.CommunityId)
		if err != nil {
			zap.L().Error("mysql.GetCommunityById(post.CommunityId) failed", zap.Error(err))
			return nil, err
		}
		list = append(list, &model.ApiPostDetail{
			Post:          &post,
			AuthorName:    AuthorName,
			CommunityName: community.CommunityName,
		})
	}
	return list, nil
}
