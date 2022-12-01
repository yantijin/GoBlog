package model

import (
	"GoLog/commen"

	uuid "github.com/satori/go.uuid"
)

type User struct {
	commen.GVA_MODEL
	UUID     uuid.UUID `gorm:"index;comment:用户UUID" json:"uuid"`
	UserName string    `gorm:"column:username;index;comment:用户登录名" json:"username"`
	Password string    `gorm:"column:password;comment:用户登录密码" json:"password"`
	NickName string    `gorm:"column:nickname;default:nickname;comment:用户昵称" json:"nickname"`
	Email    string    `gorm:"comment:用户邮箱" json:"email"`
	Avatar   string    `gorm:"comment:头像;default:https://avatars.githubusercontent.com/u/28244675?s=400&u=cb7f41621699813e9a352e8974e7c32dc4ba9220&v=4" json:"avatar"`
}

type RegisterUser struct {
	UserName string `json:"username"`
	Password string `json:"password"`
	NickName string `json:"nickname"`
	Email    string `json:"email"`
	Avatar   string `json:"avatar" gorm:"default:https://avatars.githubusercontent.com/u/28244675?s=400&u=cb7f41621699813e9a352e8974e7c32dc4ba9220&v=4"`
}

type LogInUser struct {
	UserName string `json:"username"`
	Password string `json:"password"`
}

type ChangePassword struct {
	ID          uint   `json:"-"`
	Password    string `json:"password"`
	NewPassword string `json:"newPassword"`
}

type LogInUserResponse struct {
	User      UserResponse `json:"user"`
	Token     string       `json:"token"`
	ExpiresAt int64        `json:"expiresAt"`
}

// 响应时的给前端的用户信息
type UserInfo struct {
	ID         int64  `json:"id"`
	NickName   string `json:"nickname"`
	Avatar     string `json:"avatar"`
	CreateTime int64  `json:"createTime"`
}

type UserRequest struct {
	UserName string `json:"userName"`
	NickName string `json:"nickname"`
	Email    string `json:"email"`
	Avatar   string `json:"avatar"`
}

type UserResponse struct {
	UserName string    `json:"userName"`
	NickName string    `json:"nickname"`
	Email    string    `json:"email"`
	Avatar   string    `json:"avatar"`
	ID       int64     `json:"id"`
	UUID     uuid.UUID `json:"uuid"`
}
