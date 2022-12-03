// 用户相关的请求体内容
export type LogInUserData = {
  username: string;
  password: string;
};

export type RegisterUserData = {
  username: string;
  password: string;
  nickname: string;
  email: string;
  avatar: string;
};

export type ChangePwdData = {
  username: string;
  password: string;
  newPassword: string;
};

export type SetUserInfoData = {
  username: string;
  nickname: string;
  email: string;
  avatar: string;
};

// 文章相关的请求体内容
export type PublishArticleData = {
  userId: number;
  title: string;
  content: string;
  contentType: string;
};

export type EditArticleData = {
  userId: number;
  articleId: number;
  title: string;
  content: string;
};

// 评论相关的请求体内容
export type CommentData = {
  entityType: string;
  entityId: number;
  content: string;
};

export type CommentBodyData = {
  entityType: string;
  entityId: number;
};

// 点赞相关的请求体内容
export type UserLikeRequest = {
  userId: number;
  entityType: string;
  entityId: number;
};
