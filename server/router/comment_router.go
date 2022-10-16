package router

import (
	"GoLog/controller"

	"github.com/gin-gonic/gin"
)

type CommentRouter struct{}

func (cr *CommentRouter) InitCommentRouter(router gin.RouterGroup) {
	commentRouter := router.Group("/comment")

	commentRouter.POST("/getcomments", controller.AllControllerGroup.PostGetComments)
	commentRouter.POST("/createcomments", controller.AllControllerGroup.PostCreateComment)
	commentRouter.GET("/getusercomments", controller.AllControllerGroup.GetUserComments)
}
