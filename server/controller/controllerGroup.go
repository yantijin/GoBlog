package controller

type AllControllers struct {
	UserController
	ArticleController
	CommentController
	UserLikeController
}

var AllControllerGroup = new(AllControllers)
