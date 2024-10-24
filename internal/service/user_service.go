package service

import "context"

type (
	IUserLogin interface {
		Login(ctx context.Context) error
		Register(ctx context.Context) error
		VerifyOTP(ctx context.Context) error
		UpdatePasswordRegister(ctx context.Context) error
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
