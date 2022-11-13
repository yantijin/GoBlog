// 用户相关的请求体内容
export interface LogInUserData {
  UserName: string;
  Password: string;
}

export interface RegisterUserData {
  UserName: string;
  Password: string;
  NickName: string;
  Email: string;
  Avatar: string;
}

export interface ChangePwdData {
  UserName: string;
  Password: string;
  newPassword: string;
}

export interface SetUserInfoData {
  UserName: string;
  NickName: string;
  Email: string;
  Avatar: string;
}

// 文章相关的请求体内容
export interface PublishArticleData {
  UserId: number;
  Title: string;
  Content: string;
  ContentType: string;
}

export interface EditArticleData {
  UserId: number;
  ArticleId: number;
  Title: string;
  Content: string;
}

// 评论相关的请求体内容
export interface CommentData {
  EntityType: string;
  EntityId: number;
  Content: string;
}

export interface CommentBodyData {
  EntityType: string;
  EntityId: number;
}

// 点赞相关的请求体内容
export interface UserLikeRequest {
  UserId: number;
  EntityType: string;
  EntityId: number;
}
