package middleware

import (
	"GoLog/commen"
	"GoLog/utils"

	"github.com/gin-gonic/gin"
)

// TODO: 基于jwt，创建token，解析token，并处理解析的各种情况，如过期，无法匹配等等
// 作为中间件，获取token，确定用户登录状态，未登录，则让前端显示，并跳转到登录界面；如果已登录，则ctx下一个流程

func JWTAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 首先获取token，并校验token的有效性
		token := c.Request.Header.Get("x-token")
		if token == "" {
			commen.FailedWithDetailed(gin.H{"reload": true}, "未登录", c)
			c.Abort()
			return
		}
		// 解析token包含内容，并看是否过期
		j := utils.JWT{SigningKey: []byte(commen.GVA_CONFIG.JWT.SigningKey)}
		claims, err := j.ParseToken(token)
		if err != nil {
			if err == utils.TokenExpired {
				commen.FailedWithDetailed(gin.H{"reload": true}, "授权已过期", c)
				c.Abort()
				return
			}
			commen.FailedWithDetailed(gin.H{"reload": true}, err.Error(), c)
		}
		c.Set("claims", claims) // 设置状态，可以跨中间件取值，设置值
		//TODO: JWT需要设置一个缓冲时间，当距离过期时间小于缓冲时间时，重新生成token
		// 在gin-vue-admin中采用的是singleFlight包避免并发问题
		c.Next()
	}
}
