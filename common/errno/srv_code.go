package errno

//	[10000, 1000000)
var (
	ErrUserRegister = &Errno{Code: New(100000), Msg: "web register error"}
	ErrUserInfo     = &Errno{Code: New(100001), Msg: "web info error"}
	ErrUserLogin    = &Errno{Code: New(100002), Msg: "web login error"}
	ErrUserAuthMid  = &Errno{Code: New(100003), Msg: "web web mid error"}
)
