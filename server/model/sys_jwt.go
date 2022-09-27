package model

import (
	"github.com/dgrijalva/jwt-go"
	uuid "github.com/satori/go.uuid"
)

// 主要创建自定义claims,自定义token信息

type CustomClaims struct {
	BaseClaims
	jwt.StandardClaims
}

type BaseClaims struct {
	UUID     uuid.UUID
	ID       uint
	UserName string
	NickName string
}
