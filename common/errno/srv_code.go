package errno

//	[10000, 1000000)
var (
	ErrUserRegister = &Errno{Code: New(100000), Msg: "user register error"}
	ErrUserInfo     = &Errno{Code: New(100001), Msg: "user info error"}
	ErrUserLogin    = &Errno{Code: New(100002), Msg: "user login error"}
	ErrUserAuthMid  = &Errno{Code: New(100003), Msg: "user user mid error"}
)
