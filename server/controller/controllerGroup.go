package controller

type AllControllers struct {
	UserController
	ArticleController
}

var AllControllerGroup = new(AllControllers)
