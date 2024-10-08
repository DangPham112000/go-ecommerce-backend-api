package user

import (
	"github.com/DangPham112000/go-ecommerce-backend-api/internal/wire"
	"github.com/gin-gonic/gin"
)

type UserRouter struct{}

func (ur *UserRouter) InitUserRouter(router *gin.RouterGroup) {
	// public router
	userController, _ := wire.InitUserRouterHandler()
	userRouterPublic := router.Group("/user")
	{
		userRouterPublic.POST("/register", userController.Register)
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