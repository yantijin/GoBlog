package router

import "github.com/gin-gonic/gin"

type AllRouters struct {
	UserRouter
}

var ARouters = new(AllRouters)

func SetRouters() *gin.Engine {
	Router := gin.Default()
	// 通过调用函数,来对路由组进行初始化
	routerGroup := Router.Group("")
	ARouters.InitUserRouter(*routerGroup)
	return Router
}
