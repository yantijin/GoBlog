package controller

import (
	"GoLog/commen"
	"GoLog/render"
	"GoLog/service"
	"GoLog/utils"
	"strconv"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type CommentController struct{}

// 获取comments，
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
