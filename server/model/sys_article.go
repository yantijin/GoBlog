package model

import "GoLog/commen"

type Article struct {
	commen.GVA_MODEL
	UserId          int64  `gorm:"index:article_user_id" json:"userId"`                // 文章所属用户的ID
	Title           string `gorm:"size:128;not null" json:"title"`                     // 题目
	Content         string `gorm:"type:longtext;not null" json:"content"`              // 文章内容
	ViewCount       int64  `gorm:"not null; index: idx_view_count" json:"viewCount"`   // 文章浏览次数
	ContentType     string `gorm:"type:varchar(32);not null" json:"contentType"`       // 内容类型：markdown还是html
	LikeCount       int64  `gorm:"not null;default:0" json:"likeCount"`                // 评论点赞人数
	CommentCount    int64  `gorm:"not null;default:0" json:"commentCount"`             // 文章被人评论的次数
	LastCommentTime int64  `gorm:"index:idx_last_comment_time" json:"lastCommentTime"` // 文章最后被回复的时间
}

type Comments struct {
	commen.GVA_MODEL
	UserId          int64  `gorm:"index:comment_user_id" json:"userId"`                // 评论对应的用户ID
	EntityType      string `gorm:"not null;default:article" json:"entityType"`         // 评论对应的主体：文章 / 用户评论
	EntityId        int64  `gorm:"not null;index:comment_entity_id" json:"entityId"`   // 评论对应的主体的ID,如article ID/ comment ID
	Content         string `gorm:"type:longtext;not null" json:"content"`              // 评论内容
	LikeCount       int64  `gorm:"not null;default:0" json:"likeCount"`                // 评论点赞人数
	CommentCount    int64  `gorm:"not null;default:0" json:"commentCount"`             // 评论的回复人数
	LastCommentTime int64  `gorm:"index:idx_last_comment_time" json:"lastCommentTime"` // 评论最后被回复的时间
	Status          int    `gorm:"index:status;default:1" json:"status"`               // 当前评论状态, 0:删除,1:正常
}

type UserLike struct {
	commen.GVA_MODEL
	UserId     int64  `gorm:"index:ul_user_id" json:"userId"`                  // 评论对应的用户ID
	EntityType string `gorm:"not null;default:article" json:"entityType"`      // 喜欢的的主体：文章 / 用户评论
	EntityId   int64  `gorm:"not null;index:commen_entity_id" json:"entityId"` // 喜欢的主体的ID,如article ID/ comment ID
}

// 响应时，尽量不要把全部的article信息返回，也可以加上其他信息，比如user，tags等等
type ArticleResponse struct {
	ArticleId    int64         `json:"articleId"`
	Title        string        `json:"title"`
	Content      string        `json:"content"` // 此时已全部转化为html的格式
	ViewCount    int64         `json:"viewCount"`
	CreateTime   int64         `json:"createTime"`
	UserInfo     *UserResponse `json:"user"`
	LikeCount    int64         `json:"likeCount"`
	CommentCount int64         `json:"commenCount"` // 文章当前的评论数
	CommentTime  int64         `json:"commentTime"` // 最后一次评论的时间
}

type ArticleRequest struct {
	UserId      int64  `josn:"userId"`
	Title       string `json:"title"`
	Content     string `json:"content"`
	ContentType string `json:"contentType"`
	// ViewCount   int64  `json:"viewCount"`
	// LikeCount   int64  `json:"lieCount"`
}

type ArticleEditRequest struct {
	UserId    int64  `json:"userId"`
	ArticleId int64  `json:"articleId"`
	Title     string `json:"title"`
	Content   string `json:"content"`
}

type CommentResponse struct {
	UserInfo     *UserResponse `json:"user"`
	EntityType   string        `json:"entityType"`
	EntityId     int64         `json:"entityId"`
	CommentId    int64         `json:"commentId"`
	Content      string        `json:"content"`
	LikeCount    int64         `json:"likeCount"`
	CommentCount int64         `json:"commentCount"`
	CreateTime   int64         `json:"createTime"`
	Liked        bool          `json:"liked"` // 此评论是否被当前登录用户点赞
}

type CommentJson struct {
	EntityType string `json:"entityType"`
	EntityId   int64  `json:"entityId"`
	Content    string `json:"content"`
}

type UserLikeRequest struct {
	UserId     int64  `json:"userId"`
	EntityType string `json:"entityType"`
	EntityId   int64  `json:"entityId"`
}
