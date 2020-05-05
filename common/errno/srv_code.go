package errno

//	[10000, 1000000)
var (
	ErrUserAdd      = &Errno{Code: New(100000), Msg: "error: web user add"}
	ErrUserGet      = &Errno{Code: New(100001), Msg: "error: web user get"}
	ErrUserLogin    = &Errno{Code: New(100002), Msg: "error: web user login"}
	ErrUserAuthMid  = &Errno{Code: New(100003), Msg: "error: web user auth mid"}
	ErrUserList     = &Errno{Code: New(100004), Msg: "error: web user list"}
	ErrUserNotLogin = &Errno{Code: New(100005), Msg: "error: web user not login"}
	ErrUserDelete   = &Errno{Code: New(100006), Msg: "error: web user delete"}
	ErrUserUpdate   = &Errno{Code: New(100007), Msg: "error: web user update"}
	ErrUserSetRole  = &Errno{Code: New(100008), Msg: "error: web user set role"}

	ErrRoleAdd      = &Errno{Code: New(100050), Msg: "error: web role add"}
	ErrRoleGet      = &Errno{Code: New(100051), Msg: "error: web role get"}
	ErrRoleList     = &Errno{Code: New(100052), Msg: "error: web role list"}
	ErrRoleDelete   = &Errno{Code: New(100053), Msg: "error: web role delete"}
	ErrRoleUpdate   = &Errno{Code: New(100054), Msg: "error: web role update"}
	ErrRoleSetMenus = &Errno{Code: New(100055), Msg: "error: web role set menus"}
	ErrRoleAllMenus = &Errno{Code: New(100056), Msg: "error: web role all menus"}

	ErrMenuAdd         = &Errno{Code: New(100070), Msg: "error: web menu add"}
	ErrMenuGet         = &Errno{Code: New(100071), Msg: "error: web menu get"}
	ErrMenuList        = &Errno{Code: New(100072), Msg: "error: web menu list"}
	ErrMenuDelete      = &Errno{Code: New(100073), Msg: "error: web menu delete"}
	ErrMenuUpdate      = &Errno{Code: New(100074), Msg: "error: web menu update"}
	ErrMenuMenuButtons = &Errno{Code: New(100075), Msg: "error: web menu menu buttons"}
)
