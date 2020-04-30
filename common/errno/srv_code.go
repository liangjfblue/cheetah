package errno

//	[10000, 1000000)
var (
	ErrUserRegister = &Errno{Code: New(100000), Msg: "web user register error"}
	ErrUserInfo     = &Errno{Code: New(100001), Msg: "web user get error"}
	ErrUserLogin    = &Errno{Code: New(100002), Msg: "web user login error"}
	ErrUserAuthMid  = &Errno{Code: New(100003), Msg: "web user auth mid error"}
	ErrUserList     = &Errno{Code: New(100004), Msg: "web user list error"}
	ErrUserNotLogin = &Errno{Code: New(100005), Msg: "web user not login"}
)
