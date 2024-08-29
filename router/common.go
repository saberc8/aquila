package router

import (
	"github.com/gin-gonic/gin"
)

func InitCommonRouter(Router *gin.RouterGroup) {
	resisterRouter := Router.Group("common")
	{
		resisterRouter.GET("captcha", func(ctx *gin.Context) {
			ctx.JSON(200, gin.H{
				"message": "captcha",
			})
		}) // 获取验证码

	}
}
