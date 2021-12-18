package router

import (
	"net/http"

	"github.com/allbuleyu/chat/api/handler"
	"github.com/gin-gonic/gin"
)

func Register() *gin.Engine {
	r := gin.Default()

	r.Use(CorsMiddleware())

	initUserRouter(r)
	initPushRouter(r)

	return r
}

func initUserRouter(r *gin.Engine) *gin.RouterGroup {
	user := r.Group("/user")
	{
		user.POST("/login", handler.Login)
		user.POST("/register", handler.Register)
		user.Use(handler.CheckToken())
		{
			user.POST("/checkAuth", handler.CheckAuth)
			user.POST("/logout", handler.Logout)
		}
	}

	return user
}

func initPushRouter(r *gin.Engine) *gin.RouterGroup {
	push := r.Group("/push")
	push.Use(handler.CheckToken())
	{
		push.POST("/push", handler.Push)
		push.POST("/pushRoom", handler.PushRoom)
		push.POST("/count", handler.Count)
		push.POST("/getRoomInfo", handler.GetRoomInfo)
	}

	return push
}

func CorsMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		method := c.Request.Method
		var openCorsFlag = true
		if openCorsFlag {
			c.Header("Access-Control-Allow-Origin", "*")
			c.Header("Access-Control-Allow-Headers", "Origin, X-Requested-With, Content-Type, Accept")
			c.Header("Access-Control-Allow-Methods", "GET, OPTIONS, POST, PUT, DELETE")
			c.Set("content-type", "application/json")
		}
		if method == "OPTIONS" {
			c.JSON(http.StatusOK, nil)
		}
		c.Next()
	}
}
