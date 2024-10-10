package controller

import (
	"fmt"

	"github.com/DangPham112000/go-ecommerce-backend-api/internal/service"
	"github.com/DangPham112000/go-ecommerce-backend-api/internal/vo"
	"github.com/DangPham112000/go-ecommerce-backend-api/pkg/response"
	"github.com/gin-gonic/gin"
)

type UserController struct {
	userService service.IUserService
}

func NewUserController(userService service.IUserService) *UserController {
	return &UserController{
		userService: userService,
	}
}

func (uc *UserController) Register(c *gin.Context) {
	var params vo.UserRegistratorRequest
	err := c.ShouldBindJSON(&params)
	if err != nil {
		response.ErrorResponse(c, response.ErrCodeInvalidParam, err.Error())
		return
	}
	fmt.Printf("Email params: %s\n", params.Email)
	result := uc.userService.Register(params.Email, params.Purpose)
	response.SuccessResponse(c, result, nil)
}

// func (uc *UserController) GetUserByID(c *gin.Context) {
// 	// name := c.DefaultQuery("name", "dangpham")
// 	// // c.shouldBindJSON()
// 	// uid := c.Query("uid")
// 	// c.JSON(http.StatusOK, gin.H{
// 	// 	"message": uc.userService.GetInfoUser(),
// 	// 	"users":   []string{"123", "alo", "test"},
// 	// })

// 	// response.ErrorResponse(c, 2003, "No need!")

// 	response.SuccessResponse(c, 2001, []string{"123", "alo", "test"})
// }
