package initialize

import (
	"github.com/gin-gonic/gin"
	"go_gin/middlewares"
	"go_gin/router"
)

func Routers() *gin.Engine {
	Router := gin.Default()
	// 注册zap相关中间件
	Router.Use(middlewares.GinLogger(), middlewares.GinRecovery(true))
	ApiGroup := Router.Group("/v1/")
	// 设置跨域中间件
	Router.Use(middlewares.Cors())
	//路由分组
	router.UserRouter(ApiGroup)
	router.InitBaseRouter(ApiGroup)

	return Router
}
