package response

const (
	ErrCodeSuccess      = 2001
	ErrCodeParamInvalid = 2003
	ErrInvalidToken     = 3001
)

var msg = map[int]string{
	ErrCodeSuccess:      "success",
	ErrCodeParamInvalid: "Email is invalid",
	ErrInvalidToken:     "Token is invalid",
}
