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
	resisterRouter.GET("", user.GetUserApi)
	resisterRouter.POST("update", user.UpdateUserApi)
	resisterRouter.GET("list", user.GetUserPageApi)
	resisterRouter.POST("changePassword", user.ChangePasswordApi)
	resisterRouter.POST("bindRole", user.BindRoleApi)
	resisterRouter.GET("menus", user.GetUserMenuApi)
}
