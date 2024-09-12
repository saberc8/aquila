package router

import (
	"aquila/api/system"
	"aquila/middleware"
	"github.com/gin-gonic/gin"
)

func InitMenuRouter(Router *gin.RouterGroup) {
	resisterRouter := Router.Group("menu", middleware.AuthMiddleWare())
	menu := system.Menu{}
	resisterRouter.POST("", menu.CreateMenuApi)
	resisterRouter.GET("", menu.GetMenuApi)
	resisterRouter.GET("list", menu.GetMenuAllApi)
}
