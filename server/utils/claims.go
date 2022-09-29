package utils

import (
	"GoLog/commen"
	"GoLog/model"

	"github.com/gin-gonic/gin"
	uuid "github.com/satori/go.uuid"
)

// 主要基于前端请求,解析出token对应的claims,得到相关的信息
func GetClaims(c *gin.Context) (*model.CustomClaims, error) {
	token := c.Request.Header.Get("x-token")
	j := JWT{SigningKey: []byte(commen.GVA_CONFIG.JWT.SigningKey)}
	claims, err := j.ParseToken(token)
	if err != nil {
		commen.GVA_LOG.Error("解析token出错或请求头中x-token信息与claims对应结构不符,请检查!")
	}

	return claims, err
}

// 从前端请求,解析出claims对应的ID;claims可以直接获取或从token中解析出来
func GetUserID(c *gin.Context) uint {
	if cl, exists := c.Get("claims"); !exists {
		if claims, ok := GetClaims(c); ok != nil {
			return 0
		} else {
			return claims.ID
		}
	} else {
		resClaims := cl.(*model.CustomClaims)
		return resClaims.ID
	}

}

// 从前端请求,解析出claims中的UUID,claims可以直接获取或从token中解析出来
func GetUserUUID(c *gin.Context) uuid.UUID {
	if cl, exists := c.Get("claims"); !exists {
		if claims, ok := GetClaims(c); ok != nil {
			return uuid.UUID{}
		} else {
			return claims.UUID
		}
	} else {
		resClaims := cl.(*model.CustomClaims)
		return resClaims.UUID
	}

}
