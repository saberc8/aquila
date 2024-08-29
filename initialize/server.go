package initialize

import (
	"aquila/global"
	"aquila/middleware"
	"aquila/router"
	"fmt"
	"github.com/gin-gonic/gin"
)

type server interface {
	ListenAndServe() error
}

func InitServer() {
	// 初始化路由
	routers := Routers()
	address := fmt.Sprintf(":%d", global.AQUILA_CONFIG.App.Port)
	// 启动服务
	if err := routers.Run(address); err != nil {
		panic(err)
	}
}

// Routers 配置全局的路由
func Routers() *gin.Engine {
	Router := gin.Default()
	// 注册中间件
	Router.Use(
		middleware.CorsMiddleWare(),
		middleware.LoggerMiddleWare(),
		middleware.RecoverMiddleWare(),
		//限流中间件 1秒中放100个令牌
		//middleware.RateLimitMiddleware(time.Second*1, 1000),
	)
	// 配置全局路径
	ApiGroup := Router.Group("/api")
	// 注册路由
	router.InitLoginRouter(ApiGroup)  // 用户登录
	router.InitCommonRouter(ApiGroup) // 公共路由
	return Router
}
