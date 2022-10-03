package controller

import (
	"GoLog/commen"
	"GoLog/model"
	"GoLog/render"
	"GoLog/service"
	"GoLog/utils"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type ArticleController struct{}

//获取文章，这里不需要登录，也不需要检查用户登录状态
func (ac *ArticleController) GetArticle(c *gin.Context, id int64) {
	article, err := service.AllServiceApp.FindArticle(id)
	if err != nil {
		commen.GVA_LOG.Error("获取文章错误，请检查", zap.Error(err))
		commen.FailedWithMsg(err.Error(), c)
		return
	}
	// 获取文章成功，调用浏览量+1的服务
	err = service.AllServiceApp.PlusViewCount(id)
	if err != nil {
		commen.GVA_LOG.Error("增加浏览量失败，请检查", zap.Error(err))
		commen.FailedWithMsg(err.Error(), c)
		return
	}
	response := render.BuildArticle(article)
	commen.OKWithData(response, c)
}

// 发表文章，需要校验用户登录状态！
func (ac *ArticleController) PostPublishArticle(c *gin.Context) {
	var ar model.ArticleRequest
	err := c.ShouldBindJSON(&ar)
	if err != nil {
		commen.GVA_LOG.Error("解析文章失败", zap.Error(err))
		commen.FailedWithMsg("解析文章失败", c)
	}
	// 下面就是将解析的内容绑定到model.Article上
	article := model.Article{
		UserId:      ar.UserId,
		Title:       ar.Title,
		Content:     ar.Content,
		ViewCount:   0,
		LikeCount:   0,
		ContentType: ar.ContentType,
	}
	// 对article入库
	err = service.AllServiceApp.CreateArticle(&article)
	if err != nil {
		commen.GVA_LOG.Error("文章入库失败，请检查！", zap.Error(err))
		commen.FailedWithMsg(err.Error(), c)
	}
	// 入库之后，给出成功的请求
	commen.OkWithDetailed(article, "发表文章成功", c)
}

// 编辑文章分两步，首先确定用户id和article id，然后获取文章内容，之后再给前端得到对应的数据
func (ac *ArticleController) GetEditArticle(c *gin.Context, articleId int64) {
	// 首先得到用户的ID
	userId := utils.GetUserID(c)
	// 根据articleId获取对应的Article结构体
	article, err := service.AllServiceApp.FindArticle(articleId)
	if err != nil {
		commen.GVA_LOG.Error("获取文章失败，请检查articleId！", zap.Error(err))
		commen.FailedWithMsg(err.Error(), c)
		return
	}
	// article 对应的数据
	if article.UserId != userId {
		commen.GVA_LOG.Error("token解析的用户ID与article中的用户ID不符，拒绝请求！", zap.Error(err))
		commen.FailedWithMsg(err.Error(), c)
		return
	}
	// 都没问题了，将数据返回到前端
	response := render.BuildArticle(article)
	commen.OkWithDetailed(response, "请求成功", c)
}

// 编辑文章最后提交修改
func (ac *ArticleController) PostEditArticle(c *gin.Context, articleId int64) {
	// 首先解析前端发过来的请求
	// 首先得到用户的ID
	userId := utils.GetUserID(c)
	// 根据articleId获取对应的Article结构体
	article, err := service.AllServiceApp.FindArticle(articleId)
	if err != nil {
		commen.GVA_LOG.Error("获取文章失败，请检查articleId！", zap.Error(err))
		commen.FailedWithMsg(err.Error(), c)
		return
	}
	// article 对应的数据
	if article.UserId != userId {
		commen.GVA_LOG.Error("token解析的用户ID与article中的用户ID不符，拒绝请求！", zap.Error(err))
		commen.FailedWithMsg(err.Error(), c)
		return
	}
	//都没问题了，调用修改article服务
	err = service.AllServiceApp.UpdateArticle(article)
	if err != nil {
		commen.GVA_LOG.Error("更新文章入库失败，请检查数据库连接", zap.Error(err))
		commen.FailedWithMsg(err.Error(), c)
		return
	}
	commen.OKWithMsg("入库成功", c)
}

// 删除文章,需要进行jwt验证，之后再判断用户id和article中的userid相同，最后是删除
func (ac *ArticleController) DelArticle(c *gin.Context, articleId int64) {
	userId := utils.GetUserID(c)
	article, err := service.AllServiceApp.FindArticle(articleId)
	if err != nil {
		commen.GVA_LOG.Error("获取文章失败，请检查articleId！", zap.Error(err))
		commen.FailedWithMsg(err.Error(), c)
		return
	}
	if article.UserId != userId {
		commen.GVA_LOG.Error("token解析用户ID与article对应的用户ID不符，拒绝请求!", zap.Error(err))
		commen.FailedWithMsg(err.Error(), c)
		return
	}
	// 都没问题，调用文章删除服务
	err = service.AllServiceApp.DelArtilce(articleId)
	if err != nil {
		commen.GVA_LOG.Error("删除文章失败，请检查数据库连接", zap.Error(err))
		commen.FailedWithMsg(err.Error(), c)
	}
	commen.OKWithMsg("删除成功", c)
}

// 获取某用户所有的文章,首先获取要查询的用户ID,之后根据用户ID获取所有，文章太多可能有问题，需要limit 和 分页
func (ac *UserController) GetUserAllArticles(c *gin.Context) {
	// 首先绑定articleRequest结构体，然后通过UserId来查找所有的文章
	var ar model.ArticleRequest
	err := c.ShouldBindJSON(&ar)
	if err != nil {
		commen.GVA_LOG.Error("解析数据失败，请检查！", zap.Error(err))
		commen.FailedWithMsg(err.Error(), c)
		return
	}
	// 绑定成功后，调用用户文章查询服务
	articleList := service.AllServiceApp.GetUserArticles(ar.UserId)
	// 返回给response
	articleListResponse := render.BuildArticles(articleList)
	commen.OkWithDetailed(articleListResponse, "获取成功", c)
}