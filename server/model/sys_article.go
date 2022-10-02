package model

import "GoLog/commen"

type Article struct {
	commen.GVA_MODEL
	UserId      int64  `gorm:"index:article_user_id" json:"userId"`              // 文章所属用户的ID
	Title       string `gorm:"size:128;not null" json:"title""`                  // 题目
	Content     string `gorm:"type:longtext;not null" json:"content"`            // 文章内容
	ViewCount   int64  `gotm:"not null; index: idx_view_count" json:"viewCount"` // 文章浏览次数
	ContentType string `gorm:"type:varchar(32);not null" json:"contentType"`     // 内容类型：markdown还是html
	LikeCount   int64  `gorm:"not null;default:0" json:"likeCount"`              // 评论点赞人数
}

type Comments struct {
	commen.GVA_MODEL
	UserId     int64  `gorm:"index:comment_user_id" json:"userId"`              // 评论对应的用户ID
	EntityType string `gorm:"not null;default:article" json:"entityType"`       // 评论对应的主体：文章 / 用户评论
	EntityId   int64  `gorm:"not null;index:comment_entity_id" json:"entityId"` // 评论对应的主体的ID,如article ID/ comment ID
	Content    string `gorm:"type:longtext;not null" json:"content"`            // 评论内容
	LikeCount  int64  `gorm:"not null;default:0" json:"likeCount"`              // 评论点赞人数
}

type UserLike struct {
	commen.GVA_MODEL
	UserId     int64  `gorm:"index:ul_user_id" json:"userId"`                  // 评论对应的用户ID
	EntityType string `gorm:"not null;default:article" json:"entityType"`      // 喜欢的的主体：文章 / 用户评论
	EntityId   int64  `gorm:"not null;index:commen_entity_id" json:"entityId"` // 喜欢的主体的ID,如article ID/ comment ID
}

// 响应时，尽量不要把全部的article信息返回，也可以加上其他信息，比如user，tags等等
type ArticleResponse struct {
	ArticleId  int64     `json:"articleId"`
	Title      string    `json:"title"`
	Content    string    `json:"content"` // 此时已全部转化为html的格式
	ViewCount  int64     `json:"viewCount"`
	CreateTime int64     `json:"createTime"`
	UserInfo   *UserInfo `json:"user"`
	LikeCount  int64     `json:"likeCount"`
}

type ArticleRequest struct {
	UserId      int64  `josn:"userId"`
	Title       string `json:"title"`
	Content     string `json:"content"`
	ContentType string `json:"ContentType"`
	// ViewCount   int64  `json:"viewCount"`
	// LikeCount   int64  `json:"lieCount"`
}
