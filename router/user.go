package router

import (
	"aquila/api/system"
	"aquila/middleware"
	"github.com/gin-gonic/gin"
)

func InitUserRouter(Router *gin.RouterGroup) {
	resisterRouter := Router.Group("user", middleware.AuthMiddleWare())
	user := system.User{}
	resisterRouter.POST("", user.CreateUserApi)
}
