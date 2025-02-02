package model

type Community struct {
	CommunityName string `gorm:"index:idx_community_name;type:varchar(255);not null;unique" json:"name"`
	Description   string `gorm:"type:varchar(255);" json:"description"`
	BaseModel
}

type CommunityDetail struct {
	CommunityID   int    `json:"id"`
	CommunityName string `json:"name"`
	Description   string `json:"description"`
}
