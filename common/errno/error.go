package errno

import "errors"

var (
	ErrVerifyEmail = errors.New("verify email error")
	ErrVerifyPhone = errors.New("verify phone error")
)
