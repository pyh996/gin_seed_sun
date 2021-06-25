package router

import (
	"github.com/gin-gonic/gin"
	"go_gin/controller"
	"go_gin/middlewares"
)

func UserRouter(Router *gin.RouterGroup) {
	UserRouter := Router.Group("user")
	{
		UserRouter.POST("login",  controller.PasswordLogin)
		UserRouter.POST("list",middlewares.JWTAuth(), middlewares.IsAdminAuth(), controller.GetUserList)
		UserRouter.POST("uploadUserHeaderImage",controller.PutHeaderImage)
	}
}
