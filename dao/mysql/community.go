package mysql

import (
	"go.uber.org/zap"
	"web-app/model"
)

func GetCommunityList() (list []model.Community, err error) {
	result := db.Find(&list)
	if result.Error != nil {
		return nil, result.Error
	}
	if result.RowsAffected == 0 {
		zap.L().Warn("no community found")
	}
	return list, nil
}

func GetCommunityById(id int64) (community *model.Community, err error) {
	community = &model.Community{}
	result := db.Where("id=?", id).First(community)
	if result.Error != nil || result.RowsAffected == 0 {
		return nil, ErrorInvalidID
	}
	return community, nil
}
