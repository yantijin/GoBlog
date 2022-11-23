// 用户相关的前端响应体
export interface UserInfoData {
  ID: number;
  NickName: string;
  Avatar: string;
  CreateTime: number;
}

export interface UserResponse {
  UserName: string;
  NickName: string;
  Email: string;
  Avatar: string;
}

// 文章相关的前端响应结构体
export interface ArticleResponse {
  ArticleId: number;
  Title: string;
  Content: string;
  ViewCount: number;
  CreateTime: number;
  UserInfo: UserInfoData;
  LikeCount: number;
  CommentCount: number;
  CommnetTime: number;
}

// 评论相关的前端响应结构体
export interface CommentResponse {
  UserInfo: UserInfoData;
  EntityType: string;
  EntityId: number;
  CommentId: number;
  Content: string;
  LikeCount: number;
  CommentCount: number;
  CreateTime: number;
  Liked: boolean;
}
