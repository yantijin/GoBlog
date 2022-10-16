package router

import (
	"GoLog/controller"

	"github.com/gin-gonic/gin"
)

type UserLikeRouter struct{}

func (ulr *UserLikeRouter) InitUserLikeRouter(router gin.RouterGroup) {
	userLikeRouter := router.Group("/user_like")
	userLikeRouter.GET("/getUserLike", controller.AllControllerGroup.GetUserLike)
	userLikeRouter.POST("/postLike", controller.AllControllerGroup.PostLikeEntity)
	userLikeRouter.POST("/postUnlike", controller.AllControllerGroup.PostUnlikeEntity)
}
