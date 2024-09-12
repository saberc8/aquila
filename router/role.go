package router

import (
	"aquila/api/system"
	"aquila/middleware"
	"github.com/gin-gonic/gin"
)

func InitRoleRouter(Router *gin.RouterGroup) {
	resisterRouter := Router.Group("role", middleware.AuthMiddleWare())
	role := system.Role{}
	resisterRouter.POST("", role.CreateRoleApi)
	resisterRouter.GET("", role.GetRoleApi)
	resisterRouter.GET("list", role.GetRolePageApi)
	resisterRouter.POST("update", role.UpdateRoleApi)
	resisterRouter.POST("bindMenu", role.BindMenuApi)
}
