package repo

import (
	"fmt"
	"time"

	"github.com/DangPham112000/go-ecommerce-backend-api/global"
)

type IUserAuthRepository interface {
	AddOTP(hash string, otp int, expirationTime int64) error
}

type userAuthRepository struct{}

// AddOTP implements IUserAuthRepository.
func (uar *userAuthRepository) AddOTP(hash string, otp int, expirationTime int64) error {
	key := fmt.Sprintf("usr:%s:otp", hash) // usr:${hash}:otp
	return global.Rbd.SetEx(ctx, key, otp, time.Duration(expirationTime)).Err()
}

func NewUserAuthRepository() IUserAuthRepository {
	return &userAuthRepository{}
}
