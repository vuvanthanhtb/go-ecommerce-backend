package response

const (
	ErrCodeSuccess       = 2001
	ErrCodeParamInvalid  = 2003
	ErrInvalidToken      = 3001
	ErrCodeUserHasExists = 5001
)

var msg = map[int]string{
	ErrCodeSuccess:       "success",
	ErrCodeParamInvalid:  "Email is invalid",
	ErrInvalidToken:      "Token is invalid",
	ErrCodeUserHasExists: "Email is existed",
}
