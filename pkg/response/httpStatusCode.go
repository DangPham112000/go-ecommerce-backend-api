package response

const (
	ErrCodeSuccess      = 2001 // Success
	ErrCodeEmailInvalid = 2003 // Email is invalid

	ErrCodeInvalidToken = 3001 // Token is invalid
	ErrCodeInvalidOTP   = 3002
	ErrCodeInvalidParam = 3003
	ErrCodeFailEmailOTP = 3004
	// User Authentication
	ErrCodeAuthFail = 4005
	// Register code
	ErrCodeUserHasExist = 5001 // User has already registered

	// Login code
	ErrCodeOtpNotExist           = 6007
	ErrCodeExistButNotRegistered = 6009
	ErrCodeUserOtpNotExist       = 6008
	ErrCodeCanNotCreatePassword  = 6006
)

var msg = map[int]string{
	ErrCodeSuccess:      "Success",
	ErrCodeEmailInvalid: "Email is invalid",
	ErrCodeInvalidToken: "Token is invalid",
	ErrCodeInvalidOTP:   "OTP is invalid",
	ErrCodeInvalidParam: "Param is invalid",
	ErrCodeFailEmailOTP: "Failed to send email OTP",

	ErrCodeUserHasExist: "User has already registered",

	ErrCodeOtpNotExist:           "OTP not exists",
	ErrCodeExistButNotRegistered: "OTP exists but not registered",
	ErrCodeUserOtpNotExist:       "User OTP not exists",
	ErrCodeAuthFail:              "Authentication failed",
}
