import { UserLikeRequest } from "@/model/request";
import service from "@/utils/request";

// @Summary: 根据实体类型和用户id获取点赞的所有内容
// @Produce application/json
// @Success: data: list of model.response.CommentResponse/ArticleResponse
// @Router /article/getUserLike [get]
export const getUserLike = (entityType: string, id: number) => {
  return service({
    url: "/user_like/getUserLike?entity_type=" + entityType + "&user_id=" + id,
    method: "get",
  });
};

// @Summary: 对某实体[文章/评论]点赞
// @Produce application/json
// @Param data body: model.request.UserLikeRequest
// @Router /article/postLike [post]
export const postLike = (data: UserLikeRequest) => {
  return service({
    url: "/user_like/postLike",
    method: "post",
    data: data,
  });
};

// @Summary: 对某实体[文章/评论]取消点赞
// @Produce application/json
// @Param data body: model.request.UserLikeRequest
// @Router /article/postUnlike [post]
export const postUnlike = (data: UserLikeRequest) => {
  return service({
    url: "/user_like/postUnlike",
    method: "post",
    data: data,
  });
};
