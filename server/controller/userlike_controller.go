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
	entityType := c.Query("entity_type")
	userIdStr := c.Query("user_id")
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

// 对实体进行点赞
func (uls *UserLikeController) PostLikeEntity(c *gin.Context) {
	userId := utils.GetUserID(c)
	entityType := c.PostForm("entity_type")
	entityIdStr := c.PostForm("entity_id")
	entityId, err := strconv.ParseInt(entityIdStr, 10, 64)
	if err != nil {
		commen.GVA_LOG.Error("解析entityId失败!", zap.Error(err))
		commen.FailedWithMsg(err.Error(), c)
		return
	}
	if entityType == model.EntityArticle {
		err = service.AllServiceApp.ArticleLike(commen.GVA_DB, userId, entityId)
		if err != nil {
			commen.GVA_LOG.Error("对文章点赞失败，请检查!", zap.Error(err))
			commen.FailedWithMsg(err.Error(), c)
			return
		}
	} else if entityType == model.EntityComment {
		err = service.AllServiceApp.CommentLike(commen.GVA_DB, userId, entityId)
		if err != nil {
			commen.GVA_LOG.Error("对评论点赞失败，请检查", zap.Error(err))
			commen.FailedWithMsg(err.Error(), c)
			return
		}
	} else {
		commen.GVA_LOG.Error("entityType错误，请检查！")
		commen.FailedWithMsg("entityType错误，请检查！", c)
		return
	}
	commen.OKWithMsg("对实体点赞成功！", c)
}

// 对实体取消点赞
func (uls *UserLikeController) PostUnlikeEntity(c *gin.Context) {
	userId := utils.GetUserID(c)
	entityType := c.PostForm("entity_type")
	entityIdStr := c.PostForm("entity_id")
	entityId, err := strconv.ParseInt(entityIdStr, 10, 64)
	if err != nil {
		commen.GVA_LOG.Error("解析entityId失败!", zap.Error(err))
		commen.FailedWithMsg(err.Error(), c)
		return
	}
	if entityType == model.EntityArticle {
		err = service.AllServiceApp.ArticleUnLike(commen.GVA_DB, userId, entityId)
		if err != nil {
			commen.GVA_LOG.Error("对文章取消点赞失败，请检查!", zap.Error(err))
			commen.FailedWithMsg(err.Error(), c)
			return
		}
	} else if entityType == model.EntityComment {
		err = service.AllServiceApp.CommentUnLike(commen.GVA_DB, userId, entityId)
		if err != nil {
			commen.GVA_LOG.Error("对评论取消点赞失败，请检查", zap.Error(err))
			commen.FailedWithMsg(err.Error(), c)
			return
		}
	} else {
		commen.GVA_LOG.Error("entityType错误，请检查！")
		commen.FailedWithMsg("entityType错误，请检查！", c)
		return
	}
	commen.OKWithMsg("对实体取消点赞成功！", c)
}
