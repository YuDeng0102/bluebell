package model

type Post struct {
	AuthorId    int64  `json:"author_id" gorm:"not null"`
	CommunityId int64  `json:"category_id" gorm:"not null" binding:"required"`
	Title       string `json:"title" gorm:"not null;type:varchar(255)" binding:"required"`
	Content     string `json:"content" gorm:"not null;type:text" binding:"required"`
	Status      int8   `json:"status" gorm:"not null;type:tinyint(1)"`
	BaseModel
}

type ApiPostDetail struct {
	AuthorName    string `json:"author_name"`
	CommunityName string `json:"community_name"`
	*Post
}
