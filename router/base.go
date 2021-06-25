package router

import (
	"github.com/gin-gonic/gin"
	"go_gin/controller"
)

// InitBaseRouter 图形验证码的路由
func InitBaseRouter(Router *gin.RouterGroup) {
	BaseRouter := Router.Group("base")
	{
		BaseRouter.GET("captcha", controller.GetCaptcha)
	}

}
