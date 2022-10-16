package router

import (
	"GoLog/controller"

	"github.com/gin-gonic/gin"
)

type CommentRouter struct{}

func (cr *CommentRouter) InitCommentRouter(router gin.RouterGroup) {
	commentRouter := router.Group("/comment")

	commentRouter.POST("/postGetComments", controller.AllControllerGroup.PostGetComments)
	commentRouter.POST("/postCreateComments", controller.AllControllerGroup.PostCreateComment)
	commentRouter.GET("/getUserComments", controller.AllControllerGroup.GetUserComments)
}
