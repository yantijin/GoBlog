package controller

type AllControllers struct {
	UserController
}

var AllControllerGroup = new(AllControllers)
