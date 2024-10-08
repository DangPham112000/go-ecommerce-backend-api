package user

import "github.com/gin-gonic/gin"

type UserRouter struct{}

func (ur *UserRouter) InitUserRouter(router *gin.RouterGroup) {
	// public router
	userRouterPublic := router.Group("/user")
	{
		userRouterPublic.POST("/register")
		userRouterPublic.POST("/otp")
	}
	// private router
	userRouterPrivate := router.Group("/user")
	{
		// userRouterPrivate.Use(Limiter())
		// userRouterPrivate.Use(Authen())
		// userRouterPrivate.Use(Permission())
		userRouterPrivate.GET("get_info")
	}
}
