package model

import (
	"GoLog/commen"

	"github.com/gofrs/uuid"
)

type User struct {
	commen.GVA_MODEL
	UUID     uuid.UUID `gorm:"comment:用户UUID" json:"uuid"`
	UserName string    `gorm:"comment:用户登录名" json:"userName"`
	Password string    `gorm:"comment:用户登录密码" json:"-"`
	NickName string    `gorm:"default:系统用户;comment:用户昵称" json:"nickName"`
	Email    string    `gorm:"comment:用户邮箱" json:"email"`
}
