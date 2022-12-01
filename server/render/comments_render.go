package render

import (
	"GoLog/commen"
	"GoLog/model"
	"GoLog/service"
	"GoLog/utils"
)

func BuildComment(comment *model.Comments, currentUser *model.User) *model.CommentResponse {
	if comment == nil {
		return nil
	}
	// build comment response
	cr := &model.CommentResponse{
		CommentId:    comment.ID,
		EntityType:   comment.EntityType,
		EntityId:     comment.EntityId,
		LikeCount:    comment.LikeCount,
		Content:      comment.Content,
		CommentCount: comment.CommentCount,
		CreateTime:   comment.CreatedAt.Unix(),
	}
	// userinfo
	cr.UserInfo = &model.UserResponse{
		ID:       currentUser.ID,
		NickName: currentUser.NickName,
		Avatar:   currentUser.Avatar,
		Email:    currentUser.Email,
		UserName: currentUser.UserName,
		UUID:     currentUser.UUID,
		// CreateTime: currentUser.CreatedAt.Unix(),
	}
	// 判断当前登录用户对此条评论是否点过赞，这里不合适，因为一条评论就要查一次数据库，太频繁了，要一次性把某用户和某些commentId是否点赞过返回
	// _, cr.Liked = service.AllServiceApp.ExistsUserLike(commen.GVA_DB, currentUser.ID, comment.EntityType, comment.EntityId)
	return cr
}

func BuildComments(comments []model.Comments, currentUser *model.User) []model.CommentResponse {
	if len(comments) == 0 {
		return nil
	}
	// 获取commentIds
	var commentIds []int64
	for _, comment := range comments {
		commentIds = append(commentIds, comment.ID)
	}
	likedCommentIds := service.AllServiceApp.IsLiked(commen.GVA_DB, currentUser.ID, model.EntityComment, commentIds)

	var mrs []model.CommentResponse
	for _, comment := range comments {
		item := BuildComment(&comment, currentUser)
		item.Liked = utils.Contains(comment.ID, likedCommentIds)
		mrs = append(mrs, *item)
	}
	return mrs
}
