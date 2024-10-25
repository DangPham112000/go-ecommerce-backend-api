package account

import (
	"fmt"

	"github.com/DangPham112000/go-ecommerce-backend-api/global"
	"github.com/DangPham112000/go-ecommerce-backend-api/internal/model"
	"github.com/DangPham112000/go-ecommerce-backend-api/internal/service"
	"github.com/DangPham112000/go-ecommerce-backend-api/pkg/response"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

// management controller Login user
var UserLogin = new(cUserLogin)

type cUserLogin struct{}

// Verify User Login OTP
// @Summary      Verify User Login OTP
// @Description  Verify User Login OTP
// @Tags         account management
// @Accept       json
// @Produce      json
// @Param        payload body model.VerifyInput true "payload"
// @Success      200  {object}  response.ResponseData
// @Failure      500  {object}  response.ErrorResponseData
// @Router       /user/verify_account [post]
func (c *cUserLogin) VerifyOTP(ctx *gin.Context) {
	var params model.VerifyInput
	err := ctx.ShouldBindJSON(&params)
	if err != nil {
		response.ErrorResponse(ctx, response.ErrCodeInvalidParam, err.Error())
		return
	}
	result, err := service.UserLogin().VerifyOTP(ctx, &params)
	if err != nil {
		fmt.Printf("Error:: %v\n", err)
		response.ErrorResponse(ctx, response.ErrCodeInvalidOTP, err.Error())
		return
	}
	response.SuccessResponse(ctx, response.ErrCodeSuccess, result)
}

func (c *cUserLogin) Login(ctx *gin.Context) {
	err := service.UserLogin().Login(ctx)
	if err != nil {
		response.ErrorResponse(ctx, response.ErrCodeInvalidParam, err.Error())
		return
	}
	response.SuccessResponse(ctx, response.ErrCodeSuccess, nil)
}

// User Registration documentation
// @Summary      User Registration
// @Description  When user is registered send OTP to email
// @Tags         account management
// @Accept       json
// @Produce      json
// @Param        payload body model.RegisterInput true "payload"
// @Success      200  {object}  response.ResponseData
// @Failure      500  {object}  response.ErrorResponseData
// @Router       /user/register [post]
func (c *cUserLogin) Register(ctx *gin.Context) {
	var param model.RegisterInput
	err := ctx.ShouldBindJSON(&param)
	if err != nil {
		response.ErrorResponse(ctx, response.ErrCodeInvalidParam, err.Error())
		return
	}

	codeStatus, err := service.UserLogin().Register(ctx, &param)
	if err != nil {
		global.Logger.Error("Error registration user OTP", zap.Error(err))
		response.ErrorResponse(ctx, codeStatus, err.Error())
		return
	}

	response.SuccessResponse(ctx, codeStatus, nil)
}
