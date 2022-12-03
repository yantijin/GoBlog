// 用户相关的请求体内容
export interface LogInUserData {
  username: string;
  password: string;
}

export interface RegisterUserData {
  username: string;
  password: string;
  nickname: string;
  email: string;
  avatar: string;
}

export interface ChangePwdData {
  username: string;
  password: string;
  newPassword: string;
}

export interface SetUserInfoData {
  username: string;
  nickname: string;
  email: string;
  avatar: string;
}

// 文章相关的请求体内容
export interface PublishArticleData {
  userId: number;
  title: string;
  content: string;
  contentType: string;
}

export interface EditArticleData {
  userId: number;
  articleId: number;
  title: string;
  content: string;
}

// 评论相关的请求体内容
export interface CommentData {
  entityType: string;
  entityId: number;
  content: string;
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
