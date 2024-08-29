package middleware

import (
	"aquila/enum"
	"aquila/utils"
	"fmt"
	"github.com/gin-gonic/gin"
)

// AuthMiddleWare 中间件校验token登录
func AuthMiddleWare() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 从请求头中获取token
		tokeString := c.GetHeader("token")
		fmt.Println(tokeString, "当前token")
		if tokeString == "" {
			utils.Response(c, enum.NoTokenEnum, "必须传递token", nil)
			c.Abort()
			return
		}
		c.Next()
	}
}
