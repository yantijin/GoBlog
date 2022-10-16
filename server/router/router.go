package router

import "github.com/gin-gonic/gin"

type AllRouters struct {
	UserRouter
	ArticleRouter
	CommentRouter
	UserLikeRouter
}

var ARouters = new(AllRouters)

func SetRouters() *gin.Engine {
	Router := gin.Default()
	// 通过调用函数,来对路由组进行初始化
	routerGroup := Router.Group("")
	ARouters.InitUserRouter(*routerGroup)
	ARouters.InitArticleRouter(*routerGroup)
	ARouters.InitCommentRouter(*routerGroup)
	ARouters.InitUserLikeRouter(*routerGroup)
	return Router
}
