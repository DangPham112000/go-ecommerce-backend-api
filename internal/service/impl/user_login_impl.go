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

/*
1. Hash key - email
2. Check OTP from Redis
3. Check OTP from Mysql
4. Handle if already used OTP
5. Update OTP status
*/
func (s *sUserLogin) VerifyOTP(ctx context.Context, in *model.VerifyInput) (out model.VerifyOTPOutput, err error) {
	hashKey := crypto.GetHash(strings.ToLower(in.VerifyKey))

	otpFound, err := global.Rbd.Get(ctx, utils.GetUserKey(hashKey)).Result()
	if err != nil {
		return out, err
	}
	if in.VerifyCode != otpFound {
		// TODO: Logic for wrong otp too much in a short time (3times in 1m)
		return out, err
	}
	infoOTP, err := s.r.GetInfoOTP(ctx, hashKey)
	if err != nil {
		return out, err
	}
	if infoOTP.IsVerified.Int32 == 1 {
		// TODO: Handle already used OTP
		return out, fmt.Errorf("OTP is already used")
	}

	err = s.r.UpdateUserVerificationStatus(ctx, hashKey)
	if err != nil {
		return out, err
	}

	out.Token = infoOTP.VerifyKeyHash // TODO: Should create serect key
	out.Message = "success"

	return out, nil
}

/*
1. Check token is verified?
2. Update user base table
3. Add userId of userBase to userInfo table
*/
func (s *sUserLogin) UpdatePasswordRegister(ctx context.Context, token string, password string) (userId int, err error) {
	infoOTP, err := s.r.GetInfoOTP(ctx, token)
	if err != nil {
		return response.ErrCodeUserOtpNotExist, err
	}
	if infoOTP.IsVerified.Int32 == 0 {
		return response.ErrCodeUserOtpNotExist, fmt.Errorf("user OTP is not verified")
	}

	userBase := database.AddUserBaseParams{}
	userBase.UserAccount = infoOTP.VerifyKey
	salt, err := crypto.GenerateSalt(16)
	if err != nil {
		return response.ErrCodeOtpNotExist, err
	}
	userBase.UserSalt = salt
	userBase.UserPassowrd = crypto.HashPassword(password, salt)
	newUserBase, err := s.r.AddUserBase(ctx, userBase)
	if err != nil {
	}

	newUserId, err := newUserBase.LastInsertId()
	if err != nil {
		return response.ErrCodeUserOtpNotExist, err
	}
	newUserInfo, err := s.r.AddUserHaveUserId(ctx, database.AddUserHaveUserIdParams{
		UserID:               uint64(newUserId),
		UserAccount:          infoOTP.VerifyKey,
		UserNickname:         sql.NullString{String: infoOTP.VerifyKey, Valid: true},
		UserAvatar:           sql.NullString{String: "", Valid: true},
		UserState:            1,
		UserMobile:           sql.NullString{String: "", Valid: true},
		UserGender:           sql.NullInt16{Int16: 0, Valid: true},
		UserBirthday:         sql.NullTime{Time: time.Time{}, Valid: false},
		UserEmail:            sql.NullString{String: infoOTP.VerifyKey, Valid: true},
		UserIsAuthentication: 1,
	})
	if err != nil {
		return response.ErrCodeUserOtpNotExist, err
	}
	newUserId, err = newUserInfo.LastInsertId()
	if err != nil {
		return response.ErrCodeUserOtpNotExist, err
	}

	return int(newUserId), nil
}
