package impl

import (
	"context"
	"database/sql"
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/DangPham112000/go-ecommerce-backend-api/global"
	consts "github.com/DangPham112000/go-ecommerce-backend-api/internal/const"
	"github.com/DangPham112000/go-ecommerce-backend-api/internal/database"
	"github.com/DangPham112000/go-ecommerce-backend-api/internal/model"
	"github.com/DangPham112000/go-ecommerce-backend-api/internal/utils"
	"github.com/DangPham112000/go-ecommerce-backend-api/internal/utils/crypto"
	"github.com/DangPham112000/go-ecommerce-backend-api/internal/utils/random"
	"github.com/DangPham112000/go-ecommerce-backend-api/internal/utils/sendto"
	"github.com/DangPham112000/go-ecommerce-backend-api/pkg/response"
	"github.com/redis/go-redis/v9"
)

type sUserLogin struct {
	r *database.Queries
}

func NewUserLoginImpl(r *database.Queries) *sUserLogin {
	return &sUserLogin{
		r: r,
	}
}

// Implement the IUserLogin here
func (s *sUserLogin) Login(ctx context.Context) error {
	return nil
}

/*
1. Hash email
2. Check user exists in user base
3. Check OTP exists
4. Gen new OTP
5. Save OTP to redis
6. Send OTP
7. Save OTP to mysql
*/
func (s *sUserLogin) Register(ctx context.Context, in *model.RegisterInput) (codeResult int, err error) {
	fmt.Printf("VerifyKey: %s\n", in.VerifyKey)
	fmt.Printf("VerifyType: %d\n", in.VerifyType)
	hashKey := crypto.GetHash(strings.ToLower(in.VerifyKey))
	fmt.Printf("hashKey: %s\n", hashKey)

	userFound, err := s.r.CheckUserBaseExists(ctx, in.VerifyKey)
	if err != nil {
		return response.ErrCodeUserHasExist, err
	}
	if userFound > 0 {
		return response.ErrCodeUserHasExist, fmt.Errorf("user has already registered")
	}

	userKey := utils.GetUserKey(hashKey)
	otpFound, err := global.Rbd.Get(ctx, userKey).Result()
	switch {
	case err == redis.Nil:
		fmt.Println("Key does not exist in Redis")
	case err != nil:
		return response.ErrCodeOtpNotExist, fmt.Errorf("get fail:: %v", err)
	case otpFound != "":
		return response.ErrCodeExistButNotRegistered, nil
	}

	otpNew := random.GenerateSixDigitOTP()
	if in.VerifyPurpose == "TEST_USER" {
		otpNew = 123456
	}
	fmt.Printf("OTP: %d\n", otpNew)

	err = global.Rbd.SetEx(ctx, userKey, strconv.Itoa(otpNew), time.Duration(consts.TIME_OTP_REGISTER)*time.Minute).Err()
	if err != nil {
		return response.ErrCodeInvalidOTP, err
	}

	switch in.VerifyType {
	case consts.EMAIL:
		err = sendto.SendTextEmailOTP([]string{in.VerifyKey}, consts.HOST_EMAIL, strconv.Itoa(otpNew))
		if err != nil {
			return response.ErrCodeFailEmailOTP, err
		}

		result, err := s.r.InsertOTPVerify(ctx, database.InsertOTPVerifyParams{
			VerifyOtp:     strconv.Itoa(otpNew),
			VerifyKey:     in.VerifyKey,
			VerifyKeyHash: hashKey,
			VerifyType:    sql.NullInt32{Int32: 1, Valid: true},
		})
		if err != nil {
			return response.ErrCodeFailEmailOTP, err
		}

		lastIdVerifyUser, err := result.LastInsertId()
		if err != nil {
			return response.ErrCodeFailEmailOTP, err
		}
		fmt.Printf("lastIdVerifyUser: %d\n", lastIdVerifyUser)
		return response.ErrCodeSuccess, nil
	case consts.MOBILE:
		return response.ErrCodeSuccess, nil
	}

	return response.ErrCodeSuccess, nil
}

func (s *sUserLogin) VerifyOTP(ctx context.Context) error {
	return nil
}

func (s *sUserLogin) UpdatePasswordRegister(ctx context.Context) error {
	return nil
}
