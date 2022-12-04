package controller

import (
	"GoLog/commen"
	"GoLog/model"
	"GoLog/render"
	"GoLog/service"
	"GoLog/utils"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type ArticleController struct{}

//获取文章，这里不需要登录，也不需要检查用户登录状态
func (ac *ArticleController) GetArticle(c *gin.Context) {
	s_id, status := c.GetQuery("id")
	if !status {
		commen.GVA_LOG.Error("未解析到文章id参数，请检查")
		commen.FailedWithMsg("未解析到文章id参数，请检查", c)
		return
	}
	// 类型转换
	id, err := strconv.ParseInt(s_id, 10, 64)
	if err != nil {
		commen.GVA_LOG.Error("文章id为非整数值，请检查！")
		commen.FailedWithMsg("文章id为非整数值，请检查！", c)
	}
	article, err := service.AllServiceApp.FindArticle(commen.GVA_DB, id)
	if err != nil {
		commen.GVA_LOG.Error("获取文章错误，请检查", zap.Error(err))
		commen.FailedWithMsg(err.Error(), c)
		return
	}
	// 获取文章成功，调用浏览量+1的服务
	err = service.AllServiceApp.PlusViewCount(commen.GVA_DB, id)
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
	userId := utils.GetUserID(c)
	if userId == 0 {
		commen.FailedWithMsg("尚未登录或登录过期，请先登录", c)
		return
	}
	// 下面就是将解析的内容绑定到model.Article上
	now := time.Now()
	article := model.Article{
		UserId:          userId,
		Title:           ar.Title,
		Content:         ar.Content,
		ViewCount:       0,
		LikeCount:       0,
		ContentType:     ar.ContentType,
		CommentCount:    0,
		LastCommentTime: now.Unix(),
	}
	// 对article入库
	err = service.AllServiceApp.CreateArticle(commen.GVA_DB, &article)
	if err != nil {
		commen.GVA_LOG.Error("文章入库失败，请检查！", zap.Error(err))
		commen.FailedWithMsg(err.Error(), c)
	}
	// 入库之后，给出成功的请求
	response := render.BuildArticle(&article)
	commen.OkWithDetailed(response, "发表文章成功", c)
}

// 编辑文章分两步，首先确定用户id和article id，然后获取文章内容，之后再给前端得到对应的数据
func (ac *ArticleController) GetEditArticle(c *gin.Context) {
	// 首先得到用户的ID
	s_id, status := c.GetQuery("articleId")
	if !status {
		commen.GVA_LOG.Error("未解析到文章id参数，请检查")
		commen.FailedWithMsg("未解析到文章id参数，请检查", c)
		return
	}
	// 类型转换
	articleId, err := strconv.ParseInt(s_id, 10, 64)
	if err != nil {
		commen.GVA_LOG.Error("文章id为非整数值，请检查！")
		commen.FailedWithMsg("文章id为非整数值，请检查！", c)
	}
	userId := utils.GetUserID(c)
	// 根据articleId获取对应的Article结构体
	article, err := service.AllServiceApp.FindArticle(commen.GVA_DB, articleId)
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
func (ac *ArticleController) PostEditArticle(c *gin.Context) {
	// 首先解析前端发过来的请求
	// 首先得到用户的ID
	var aer model.ArticleEditRequest
	err := c.ShouldBindJSON(&aer)
	if err != nil {
		commen.GVA_LOG.Error("绑定结构体错误，请检查！", zap.Error(err))
		commen.FailedWithMsg(err.Error(), c)
		return
	}
	userId := utils.GetUserID(c)
	// 根据articleId获取对应的Article结构体
	article, err := service.AllServiceApp.FindArticle(commen.GVA_DB, aer.ArticleId)
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
	err = service.AllServiceApp.ArticleService.UpdateColumns(commen.GVA_DB, aer.ArticleId, map[string]interface{}{
		"title":   aer.Title,
		"content": aer.Content,
	})
	if err != nil {
		commen.GVA_LOG.Error("更新文章失败，请检查！", zap.Error(err))
		commen.FailedWithMsg(err.Error(), c)
		return
	}
	commen.OKWithMsg("入库成功", c)
}

// 删除文章,需要进行jwt验证，之后再判断用户id和article中的userid相同，最后是删除
func (ac *ArticleController) DelArticle(c *gin.Context) {
	id := c.Param("articleId")
	// 类型转换
	articleId, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		commen.GVA_LOG.Error("文章id为非整数值，请检查！")
		commen.FailedWithMsg("文章id为非整数值，请检查！", c)
	}
	userId := utils.GetUserID(c)
	article, err := service.AllServiceApp.FindArticle(commen.GVA_DB, articleId)
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
	err = service.AllServiceApp.DelArtilce(commen.GVA_DB, articleId)
	if err != nil {
		commen.GVA_LOG.Error("删除文章失败，请检查数据库连接", zap.Error(err))
		commen.FailedWithMsg(err.Error(), c)
	}
	commen.OKWithMsg("删除成功", c)
}

// 获取某用户所有的文章,首先获取要查询的用户ID,之后根据用户ID获取所有，文章太多可能有问题，需要limit 和 分页
func (ac *UserController) GetUserAllArticles(c *gin.Context) {
	userIdStr, status := c.GetQuery("id")
	if !status {
		commen.GVA_LOG.Error("未解析到用户id参数，请检查")
		commen.FailedWithMsg("未解析到用户id参数，请检查", c)
		return
	}
	// 类型转换
	id, err := strconv.ParseInt(userIdStr, 10, 64)
	if err != nil {
		commen.GVA_LOG.Error("用户id为非整数值，请检查！")
		commen.FailedWithMsg("用户id为非整数值，请检查！", c)
	}
	// 绑定成功后，调用用户文章查询服务
	articleList := service.AllServiceApp.GetUserArticles(commen.GVA_DB, id)
	// 返回给response
	articleListResponse := render.BuildArticles(articleList)
	commen.OkWithDetailed(articleListResponse, "获取成功", c)
}
