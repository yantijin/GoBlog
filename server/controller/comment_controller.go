package controller

import (
	"GoLog/commen"
	"GoLog/model"
	"GoLog/render"
	"GoLog/service"
	"GoLog/utils"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type CommentController struct{}

// 获取comments，这里是登录状态，才能保证返回前端的是对应的点赞数据
func (cc *CommentController) GetComments(c *gin.Context) {
	userId := utils.GetUserID(c)
	currentUser, err := service.AllServiceApp.FindUser(commen.GVA_DB, userId)
	if err != nil {
		commen.GVA_LOG.Error("解析用户出错", zap.Error(err))
		commen.FailedWithMsg(err.Error(), c)
		return
	}
	entityType := c.PostForm("entityType")
	entityId_str := c.PostForm("entityId")
	if len(entityType) == 0 {
		commen.GVA_LOG.Error("entityType解析出错或为空，请检查")
		commen.FailedWithMsg("entityType解析出错或为空，请检查", c)
		return
	}
	entityId, err := strconv.ParseInt(entityId_str, 10, 64)
	if err != nil {
		commen.GVA_LOG.Error("entityId解析出错或为空，请检查")
		commen.FailedWithMsg("entityId解析出错或为空，请检查", c)
		return
	}
	// 调用commentService中GetComments的服务
	comments := service.AllServiceApp.GetComments(commen.GVA_DB, entityType, entityId)
	commentResp := render.BuildComments(comments, currentUser)
	commen.OkWithDetailed(commentResp, "获取评论成功", c)
}

// 发布评论，需要jwt验证
func (cc *CommentController) PostCreateComment(c *gin.Context) {
	userId := utils.GetUserID(c)
	var cj model.CommentJson
	err := c.ShouldBindJSON(&cj)
	if err != nil {
		commen.GVA_LOG.Error("绑定数据失败，请检查传输数据格式", zap.Error(err))
		commen.FailedWithMsg(err.Error(), c)
		return
	}
	comment := &model.Comments{
		UserId:          userId,
		EntityType:      cj.EntityType,
		EntityId:        cj.EntityId,
		Content:         cj.Content,
		LikeCount:       0,
		CommentCount:    0,
		LastCommentTime: time.Now().Unix(),
	}
	// 发布评论，并对对应实体的comment count + 1
	err = service.AllServiceApp.PublishComment(comment)
	if err != nil {
		commen.GVA_LOG.Error("评论入库失败，请检查数据库连接", zap.Error(err))
		commen.FailedWithMsg(err.Error(), c)
		return
	}
	commen.OKWithMsg("创建成功", c)
	return
}

// 编辑评论,对content内容进行编辑，这里需要考虑怎么获取评论，因为commentId这个是无法写在前端的
// func (cc *CommentController) PostEditComment(c *gin.Context) {
// userId := utils.GetUserID(c)

// }

// 删除评论,如果有二级评论怎么办？应该需要逻辑删除,而不是真的删除
// func (cc *CommentController) DelComment(c *gin.Context) {}
