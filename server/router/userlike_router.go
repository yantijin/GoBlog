package router

import (
	"GoLog/controller"

	"github.com/gin-gonic/gin"
)

type UserLikeRouter struct{}

func (ulr *UserLikeRouter) InitUserLikeRouter(router gin.RouterGroup) {
	userLikeRouter := router.Group("/user_like")
	userLikeRouter.GET("/get_data", controller.AllControllerGroup.GetUserLike)
	userLikeRouter.POST("/like", controller.AllControllerGroup.PostLikeEntity)
	userLikeRouter.POST("/unlike", controller.AllControllerGroup.PostUnlikeEntity)
}
