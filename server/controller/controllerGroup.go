package controller

type AllControllers struct {
	UserController
	ArticleController
	CommentController
}

var AllControllerGroup = new(AllControllers)
