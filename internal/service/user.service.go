package service

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/DangPham112000/go-ecommerce-backend-api/global"
	"github.com/DangPham112000/go-ecommerce-backend-api/internal/repo"
	"github.com/DangPham112000/go-ecommerce-backend-api/internal/utils/crypto"
	"github.com/DangPham112000/go-ecommerce-backend-api/internal/utils/random"
	"github.com/DangPham112000/go-ecommerce-backend-api/pkg/response"
	"github.com/segmentio/kafka-go"
)

// Interface
type IUserService interface {
	Register(email string, purpose string) int
}

type userService struct {
	userRepo     repo.IUserRepository
	userAuthRepo repo.IUserAuthRepository
}

func NewUserService(userRepo repo.IUserRepository, userAuthRepo repo.IUserAuthRepository) IUserService {
	return &userService{
		userRepo:     userRepo,
		userAuthRepo: userAuthRepo,
	}
}

// Register implements IUserService.
func (us *userService) Register(email string, purpose string) int {
	// 1. Hash email: we will store OTP to redis, and hash email to protect this OTP when redis is leaked, hacker dont know OTP belong to which email
	hashEmail := crypto.GetHash(email)
	fmt.Printf("Hashed email::%s\n", hashEmail)

	// 2. Check email exists
	if us.userRepo.GetUserByEmail(email) {
		return response.ErrCodeUserHasExist
	}

	// 3. Gen OTP
	otp := random.GenerateSixDigitOTP()
	if purpose == "TEST_USER" {
		otp = 123456
	}
	fmt.Printf("OTP:::%d\n", otp)

	// 4. Save OTP with expiration time
	err := us.userAuthRepo.AddOTP(hashEmail, otp, int64(10*time.Minute))
	if err != nil {
		return response.ErrCodeInvalidOTP
	}

	// 5. Send Email OTP
	// err = sendto.SendTemplateEmailOTP(
	// 	[]string{email},
	// 	"dangpham112000@gmail.com",
	// 	"otp-auth.html",
	// 	map[string]interface{}{"otp": strconv.Itoa(otp)},
	// )
	// if err != nil {
	// 	return response.ErrCodeFailEmailOTP
	// }

	// 5. Request JAVA send OTP
	// err = sendto.SendEmailToJavaAPI(strconv.Itoa(otp), email, "authentication")
	// if err != nil {
	// 	return response.ErrCodeFailEmailOTP
	// }

	// 5. Send OTP to Kafka then Java will receive
	body := make(map[string]interface{})
	body["otp"] = otp
	body["email"] = email

	bodyRequest, _ := json.Marshal(body)

	message := kafka.Message{
		Key:   []byte("otp-auth"),
		Value: []byte(bodyRequest),
		Time:  time.Now(),
	}

	err = global.KafkaProducer.WriteMessages(context.Background(), message)
	if err != nil {
		fmt.Printf("Err kafka send otp %v\n", err)
		return response.ErrCodeFailEmailOTP
	}

	return response.ErrCodeSuccess
}
