package controller

import (
	"GoLog/commen"
	"GoLog/model"
	"GoLog/render"
	"GoLog/service"
	"GoLog/utils"
	"strconv"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type UserLikeController struct{}

// 获取某用户点赞的评论/文章，取决于entityType, userId
func (ulc *UserLikeController) GetUserLike(c *gin.Context) {
	entityType := c.PostForm("entity_type")
	userIdStr := c.PostForm("user_id")
	if len(entityType) == 0 || len(userIdStr) == 0 {
		commen.GVA_LOG.Error("请求错误，请检查参数是否为空")
		commen.FailedWithMsg("请求错误，请检查参数是否为空", c)
		return
	}

	userId, err := strconv.ParseInt(userIdStr, 10, 64)
	if err != nil {
		commen.GVA_LOG.Error("解析user_id失败，请检查！", zap.Error(err))
		commen.FailedWithMsg(err.Error(), c)
		return
	}

	// 调用获取点赞服务,返回前端的是对应的entityIds
	entityIds := service.AllServiceApp.FindPersonLikes(commen.GVA_DB, userId, entityType)
	// 根据entityIds 来构建响应
	if len(entityIds) == 0 {
		commen.GVA_LOG.Warn("没有查询到对应的点赞数据！")
		commen.FailedWithMsg("没有查询到对应的点赞数据！", c)
		return
	}

	// 分情况讨论,根据entityIds来找articles
	if entityType == model.EntityArticle {
		cnd := utils.NewSqlCnd().In("id", entityIds)
		articleList := service.AllServiceApp.FindArticles(commen.GVA_DB, cnd)
		articleReps := render.BuildArticles(articleList)
		commen.OkWithDetailed(articleReps, "拉取点赞文章成功！", c)
		return
	} else if entityType == model.EntityComment {
		cnd := utils.NewSqlCnd().In("id", entityIds)
		commentList := service.AllServiceApp.FindComments(commen.GVA_DB, cnd)
		blankUser := model.User{
			UserName: "",
			Password: "",
			NickName: "匿名",
			Email:    "",
			Avatar:   "",
		}
		commentResp := render.BuildComments(commentList, &blankUser)
		commen.OkWithDetailed(commentResp, "拉取点赞评论成功", c)
		return
	} else {
		commen.FailedWithMsg("请检查entityType！", c)
		return
	}
}
