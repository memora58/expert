package middleware

import "github.com/gin-gonic/gin"

// 接受服务实例，并存到gin.Key中
func InitMiddleware(service []interface{}) gin.HandlerFunc {
	return func(c *gin.Context) {
		// 将实例存在gin.Keys中

		// 将实例存在gin.Keys中
		c.Keys = make(map[string]interface{})

		c.Keys["user"] = service[0]
		c.Keys["task"] = service[1]
		c.Next()
	}
}
