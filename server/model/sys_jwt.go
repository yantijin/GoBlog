package model

import (
	"github.com/golang-jwt/jwt/v4"
	uuid "github.com/satori/go.uuid"
)

// 主要创建自定义claims,自定义token信息

type CustomClaims struct {
	BaseClaims
	jwt.StandardClaims
}

type BaseClaims struct {
	UUID     uuid.UUID
	ID       int64
	UserName string
	NickName string
}
