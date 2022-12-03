// 用户相关的前端响应体
export interface UserInfoData {
  id: number;
  nickname: string;
  username: string;
  avatar: string;
  email: string;
  uuid: string;
  // CreateTime: number;
}

export interface UserResponse {
  username: string;
  nickname: string;
  email: string;
  avatar: string;
}

// 文章相关的前端响应结构体
export interface ArticleResponse {
  articleId: number;
  title: string;
  content: string;
  viewCount: number;
  createTime: number;
  user: UserInfoData;
  likeCount: number;
  commentCount: number;
  commentTime: number;
}

// 评论相关的前端响应结构体
export interface CommentResponse {
  user: UserInfoData;
  entityType: string;
  entityId: number;
  commentId: number;
  content: string;
  likeCount: number;
  commentCount: number;
  createTime: number;
  liked: boolean;
}
