package manage

import "github.com/gin-gonic/gin"

type UserRouter struct{}

func (ur *UserRouter) InitUserRouter(router *gin.RouterGroup) {
	// public router

	// private router
	userRouterPrivate := router.Group("/admin/user")
	{
		// userRouterPrivate.Use(Limiter())
		// userRouterPrivate.Use(Authen())
		// userRouterPrivate.Use(Permission())
		userRouterPrivate.POST("active_user")
	}
}
