import { EditArticleData, PublishArticleData } from "@/model/request";
import service from "@/utils/request";

// @Summary: 发表文章
// @Produce application/json
// @Param data body: model.request.PublishArticleData
// @Router /article/postPublishArticle [post]
// @Success: data: model.response.ArticleResponse
export const publishArticle = (data: PublishArticleData) => {
  return service({
    url: "/article/postPublishArticle",
    method: "post",
    data: data,
  });
};

// @Summary: 根据文章id获取文章
// @Produce application/json
// @Success: data: model.response.ArticleResponse
// @Router /article/getArticle [get]
export const getArticle = (id: string) => {
  return service({
    url: "/article/getArticle?id=" + id,
    method: "get",
  });
};

// @Summary: 根据用户id删除文章
// @Produce application/json
// @Router /article/deleteArticle [delete]
export const delArticle = (id: string) => {
  return service({
    url: "/article/deleteArticle/" + id,
    method: "delete",
  });
};

// @Summary: 根据文章id获取登录用户的某篇文章，以便后续编辑
// @Produce application/json
// @Success: data: model.response.ArticleResponse
// @Router /article/getEditArticle [get]
export const getEditArticle = (id: string) => {
  return service({
    url: "/article/getEditArticle?id=" + id,
    method: "get",
  });
};

// @Summary: 更新用户编辑后的文章
// @Produce application/json
// @Param data body: model.request.EditArticleData
// @Router /article/postEditArticle [post]
export const postEditArticle = (data: EditArticleData) => {
  return service({
    url: "/article/postEditArticle",
    method: "post",
    data: data,
  });
};

// @Summary: 根据用户id获取用户所有文章
// @Produce application/json
// @Success: data: list of model.response.ArticleResponse
// @Router /article/getUserArticles [get]
export const getUserArticles = (id: string) => {
  return service({
    url: "/article/getUserArticles?id=" + id,
    method: "get",
  });
};
