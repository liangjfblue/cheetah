package errno

//	[10000, 1000000)
var (
	ErrUserRegister = &Errno{Code: New(100000), Msg: "user register error"}
	ErrUserInfo     = &Errno{Code: New(100001), Msg: "user info error"}
	ErrUserLogin    = &Errno{Code: New(100002), Msg: "user login error"}
	ErrUserAuthMid  = &Errno{Code: New(100003), Msg: "user user mid error"}

	ErrCoinGet = &Errno{Code: New(100100), Msg: "coin get error"}
	ErrCoinAdd = &Errno{Code: New(100101), Msg: "coin add error"}

	ErrArticlePost = &Errno{Code: New(100200), Msg: "post article error"}
	ErrArticleGet  = &Errno{Code: New(100201), Msg: "get article error"}
	ErrArticleDel  = &Errno{Code: New(100202), Msg: "del article error"}

	ErrCommentAdd  = &Errno{Code: New(100220), Msg: "comment add error"}
	ErrCommentDel  = &Errno{Code: New(100221), Msg: "comment del error"}
	ErrCommentList = &Errno{Code: New(100222), Msg: "comment list error"}
)
