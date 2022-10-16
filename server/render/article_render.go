package render

import (
	"GoLog/commen"
	"GoLog/model"
	"GoLog/service"
	"GoLog/utils"

	"go.uber.org/zap"
)

// 这个包的作用是将model中的struct做额外的加工，形成最后给前端的一些数据结构

func BuildArticle(article *model.Article) *model.ArticleResponse {
	if article == nil {
		return nil
	}

	rsp := model.ArticleResponse{}
	rsp.ArticleId = article.ID
	rsp.CreateTime = article.CreatedAt.Unix()
	rsp.Title = article.Title
	rsp.ViewCount = article.ViewCount
	rsp.LikeCount = article.LikeCount
	rsp.CommentCount = article.CommentCount
	rsp.CommentTime = article.LastCommentTime
	// rsp.
	user, err := service.AllServiceApp.FindUser(commen.GVA_DB, article.UserId)
	if err != nil {
		commen.GVA_LOG.Error("article信息有误，无法获取对应user信息，请检查！", zap.Error(err))
	}

	rsp.UserInfo = &model.UserInfo{
		ID:         article.UserId,
		NickName:   user.NickName,
		Avatar:     user.Avatar,
		CreateTime: user.CreatedAt.Unix(),
	}

	// 对于Content，需要根据ContentType来将content转换为对应类型
	if article.ContentType == model.ContentTypeMarkDown {
		content := utils.ToHtml(article.Content)
		rsp.Content = utils.HandleHtmlContent(content)
	} else if article.ContentType == model.ContentTypeHtml {
		rsp.Content = utils.HandleHtmlContent(article.Content)
	}
	return &rsp
}

func BuildArticles(articles []model.Article) []model.ArticleResponse {
	if len(articles) == 0 {
		return nil
	}
	var responses []model.ArticleResponse
	for _, article := range articles {
		responses = append(responses, *BuildArticle(&article))
	}
	return responses
}
