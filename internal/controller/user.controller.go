package controller

import (
	"github.com/DangPham112000/go-ecommerce-backend-api/internal/service"
	"github.com/DangPham112000/go-ecommerce-backend-api/pkg/response"
	"github.com/gin-gonic/gin"
)

type UserController struct {
	userService *service.UserService
}

func NewUserController() *UserController {
	return &UserController{
		userService: service.NewUserService(),
	}
}

func (uc *UserController) GetUserByID(c *gin.Context) {
	// name := c.DefaultQuery("name", "dangpham")
	// // c.shouldBindJSON()
	// uid := c.Query("uid")
	// c.JSON(http.StatusOK, gin.H{
	// 	"message": uc.userService.GetInfoUser(),
	// 	"users":   []string{"123", "alo", "test"},
	// })

	// response.ErrorResponse(c, 2003, "No need!")

	response.SuccessResponse(c, 2001, []string{"123", "alo", "test"})
}
