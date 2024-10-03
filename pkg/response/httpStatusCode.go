package response

const (
	ErrCodeSuccess      = 2001 // Success
	ErrCodeEmailInvalid = 2003 // Email is invalid
)

var msg = map[int]string{
	ErrCodeSuccess:      "Success",
	ErrCodeEmailInvalid: "Email is invalid",
}
