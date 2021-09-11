package routes

import (
	"github.com/debuginn/GoShortLink/controller"
	"github.com/debuginn/GoShortLink/middleware"
	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	var r = gin.Default()
	r.Use(middleware.CORSMiddleware())                                    // CORS 同源策略 中间件
	r.POST("/api/auth/register", controller.Register)                     // 用户注册
	r.POST("/api/auth/login", controller.Login)                           // 用户登录
	r.GET("/api/auth/info", middleware.AuthMiddleware(), controller.Info) // 用户信息

	return r
}
