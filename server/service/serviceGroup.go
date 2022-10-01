package service

type AllService struct {
	UserService
	ArticleService
	CommentService
	UserLikeService
}

var AllServiceApp = new(AllService)
