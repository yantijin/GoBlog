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

// 获取comments
func (cc *CommentController) PostGetComments(c *gin.Context) {
	var cj model.CommentJson
	err := c.ShouldBindJSON(&cj)
	if err != nil {
		commen.GVA_LOG.Error("绑定结构体失败，请检查", zap.Error(err))
		commen.FailedWithMsg(err.Error(), c)
		return
	}
	comments := service.AllServiceApp.GetComments(commen.GVA_DB, cj.EntityType, cj.EntityId)
	userId := utils.GetUserID(c)
	currentUser := &model.User{
		UserName: "",
	}
	// userId=0 => 用户没有登录
	if userId != 0 {
		currentUser, err = service.AllServiceApp.FindUser(commen.GVA_DB, userId)
		if err != nil {
			commen.GVA_LOG.Error("获取当前用户信息失败", zap.Error(err))
			commen.FailedWithMsg(err.Error(), c)
			return
		}
	}

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
		ContentType:     cj.ContentType,
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
}

// 编辑评论,对content内容进行编辑，这里需要考虑怎么获取评论，因为commentId这个是无法写在前端的
// func (cc *CommentController) PostEditComment(c *gin.Context) {
// userId := utils.GetUserID(c)

// }

// 删除评论,如果有二级评论怎么办？应该需要逻辑删除,而不是真的删除
// func (cc *CommentController) DelComment(c *gin.Context) {}

// 获取某个用户的Comment
func (cc *CommentController) GetUserComments(c *gin.Context) {
	userId, ok := c.GetQuery("id")
	if !ok {
		commen.GVA_LOG.Error("没有userId字段或网络错误,请检查")
		commen.FailedWithMsg("没有userId字段或网络错误,请检查", c)
		return
	}
	uid, err := strconv.ParseInt(userId, 10, 64)
	if err != nil {
		commen.GVA_LOG.Error("userId应该为数字,请检查")
		commen.FailedWithMsg("userId应该为数字,请检查", c)
		return
	}

	cUserId := utils.GetUserID(c)
	currentUser := &model.User{
		UserName: "",
	}
	if cUserId != 0 {
		currentUser, err = service.AllServiceApp.FindUser(commen.GVA_DB, cUserId)
		if err != nil {
			commen.GVA_LOG.Error("解析用户出错", zap.Error(err))
			commen.FailedWithMsg(err.Error(), c)
			return
		}
	}
	// 调用查找服务
	comments := service.AllServiceApp.GetUserComments(commen.GVA_DB, uid)
	// build response
	commentResp := render.BuildComments(comments, currentUser)
	commen.OkWithDetailed(commentResp, "获取评论成功", c)
}
