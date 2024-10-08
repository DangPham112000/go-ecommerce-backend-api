package response

const (
	ErrCodeSuccess      = 2001 // Success
	ErrCodeEmailInvalid = 2003 // Email is invalid

	ErrCodeInvalidToken = 3001 // Token is invalid
	// Register code
	ErrCodeUserHasExist = 5001 // User has already registered
)

var msg = map[int]string{
	ErrCodeSuccess:      "Success",
	ErrCodeEmailInvalid: "Email is invalid",
	ErrCodeInvalidToken: "Token is invalid",
	ErrCodeUserHasExist: "User has already registered",
}
