package utils

import (
	"GoLog/commen"

	"golang.org/x/crypto/bcrypt"
)

// 对密码进行加密,对密码进行验证

func GetHashPwd(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		commen.GVA_LOG.Error("密码加密错误,请检查")
		return "", err
	}
	return string(bytes), nil
}

func CheckPwd(password string, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}
