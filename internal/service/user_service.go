package service

import (
	"context"

	"github.com/DangPham112000/go-ecommerce-backend-api/internal/model"
)

type (
	IUserLogin interface {
		Login(ctx context.Context, in *model.LoginInput) (codeResult int, out model.LoginOutput, err error)
		Register(ctx context.Context, in *model.RegisterInput) (codeResult int, err error)
		VerifyOTP(ctx context.Context, in *model.VerifyInput) (out model.VerifyOTPOutput, err error)
		UpdatePasswordRegister(ctx context.Context, token string, password string) (userId int, err error)

		IsTwoFactorEnable(ctx context.Context, userId int) (codeResult int, rs bool, err error)

		SetupTwoFactorAuth(ctx context.Context, in *model.SetupTwoFactorAuthInput) (codeResult int, err error)

		VerifyTwoFactorAuth(ctx context.Context, in *model.TwoFactorVerificationInput) (codeResult int, err error)
	}

	IUserInfo interface {
		GetInfoByUserId(ctx context.Context) error
		GetAllUser(ctx context.Context) error
	}

	IUserAdmin interface {
		RemoveUser(ctx context.Context) error
		FindOneUser(ctx context.Context) error
	}
)

var (
	localUserAdmin IUserAdmin
	localUserInfo  IUserInfo
	localUserLogin IUserLogin
)

func UserAdmin() IUserAdmin {
	if localUserAdmin == nil {
		panic("Implement localUserAdmin not found for interface IUserAdmin")
	}

	return localUserAdmin
}

func InitUserAdmin(i IUserAdmin) {
	localUserAdmin = i
}

func UserInfo() IUserInfo {
	if localUserInfo == nil {
		panic("Implement localUserInfo not found for interface IUserInfo")
	}

	return localUserInfo
}

func InitUserInfo(i IUserInfo) {
	localUserInfo = i
}

func UserLogin() IUserLogin {
	if localUserLogin == nil {
		panic("Implement localUserLogin not found for interface IUserLogin")
	}

	return localUserLogin
}

func InitUserLogin(i IUserLogin) {
	localUserLogin = i
}
