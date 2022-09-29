package router

import (
	"GoLog/controller"

	"github.com/gin-gonic/gin"
)

type UserRouter struct{}

// 用户相关路由注册
func (u *UserRouter) InitUserRouter(router gin.RouterGroup) {
	userRouter := router.Group("user")

	userRouter.POST("/signup", controller.AllControllerGroup.PostSignUp)
	userRouter.POST("/signin", controller.AllControllerGroup.PostLogIn)
	userRouter.POST("/change_pwd", controller.AllControllerGroup.ChangePwd)
}
