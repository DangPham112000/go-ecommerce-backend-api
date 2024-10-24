package account

import (
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

func (c *cUserLogin) Login(ctx *gin.Context) {
	err := service.UserLogin().Login(ctx)
	if err != nil {
		response.ErrorResponse(ctx, response.ErrCodeInvalidParam, err.Error())
		return
	}
	response.SuccessResponse(ctx, response.ErrCodeSuccess, nil)
}

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
