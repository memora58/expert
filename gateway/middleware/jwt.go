package middleware

import (
	"gateway/common/response"
	"gateway/common/util"
	"github.com/gin-gonic/gin"
	"time"
)

// JWT token验证中间件
func JWT() gin.HandlerFunc {
	return func(c *gin.Context) {
		var code int
		var data interface{}
		code = 200
		token := c.GetHeader("Authorization")
		if token == "" {
			code = 404
		} else {
			claims, err := util.ParseToken(token)
			if err != nil {
				code = response.ErrorAuthCheckTokenFail
			} else if time.Now().Unix() > claims.ExpiresAt {
				code = response.ErrorAuthCheckTokenTimeout
			}
		}
		if code != response.SUCCESS {
			c.JSON(200, gin.H{
				"status": code,
				"msg":    response.GetMsg(uint(code)),
				"data":   data,
			})
			c.Abort()
			return
		}
		c.Next()
	}
}
