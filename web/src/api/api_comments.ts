import { CommentBodyData, CommentData } from "@/model/request";
import service from "@/utils/request";

// @Summary: 发表评论
// @Produce application/json
// @Param data body: model.request.CommentData
// @Router /article/postCreateComments [post]
export const createComment = (data: CommentData) => {
  return service({
    url: "/comment/postCreateComments",
    method: "post",
    data: data,
  });
};

// @Summary: 根据实体类型和id获取对应的评论
// @Produce application/json
// @Param data body: model.request.CommentBodayData
// @Success: data: nidek.response.CommentResponse
// @Router /article/postGetComments [post]
export const getComment = (data: CommentBodyData) => {
  return service({
    url: "/comment/postGetComments",
    method: "post",
    data: data,
  });
};

// @Summary: 根据用户id获取其所有评论
// @Produce application/json
// @Success: data: list of model.response.CommentResponse
// @Router /article/getUserComments [get]
export const getUserComments = (id: string) => {
  return service({
    url: "/comment/getUserComments?id=" + id,
    method: "get",
  });
};
