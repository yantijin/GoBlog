package router

import (
	"GoLog/controller"

	"github.com/gin-gonic/gin"
)

type ArticleRouter struct{}

// 用户相关路由注册
func (u *ArticleRouter) InitArticleRouter(router gin.RouterGroup) {
	ArticleRouter := router.Group("/article")

	ArticleRouter.GET("/getArticle", controller.AllControllerGroup.GetArticle)
	ArticleRouter.POST("/postPublishArticle", controller.AllControllerGroup.PostPublishArticle)
	ArticleRouter.GET("/getEditArticle", controller.AllControllerGroup.GetEditArticle)
	ArticleRouter.POST("/postEditArticle", controller.AllControllerGroup.PostEditArticle)
	ArticleRouter.DELETE("/deleteArtile/:articleId", controller.AllControllerGroup.DelArticle)
	ArticleRouter.GET("/getUserArticles", controller.AllControllerGroup.GetUserAllArticles)
}
