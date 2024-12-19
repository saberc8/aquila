package router

import (
	"github.com/gin-gonic/gin"
)

// q: *是什么意思
// a:
func InitLoginRouter(Router *gin.RouterGroup) {
	resisterRouter := Router.Group("login")
	{
		resisterRouter.POST("register", func(ctx *gin.Context) {
			ctx.JSON(200, gin.H{
				"message": "register",
			})
		}) // 用户注册账号

	}
}
