package router

import (
	"aquila/api/common"
	"github.com/gin-gonic/gin"
)

func InitCommonRouter(Router *gin.RouterGroup) {
	resisterRouter := Router.Group("common")
	auth := common.Auth{}
	resisterRouter.GET("/captcha", auth.Captcha)
	resisterRouter.POST("/login", auth.Login)
	resisterRouter.POST("/logout", auth.Logout)
	resisterRouter.POST("/register", auth.Register)
}
