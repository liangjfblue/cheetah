package errno

//	[1, 10000)
var (
	Success = &Errno{Code: New(1), Msg: "ok"}

	ErrBind             = &Errno{Code: New(10), Msg: "bind json error"}
	ErrParams           = &Errno{Code: New(11), Msg: "params empty error"}
	ErrTraceNoContext   = &Errno{Code: New(12), Msg: "tracer no context error"}
	ErrTraceIntoContext = &Errno{Code: New(13), Msg: "tracer into context error"}
	ErrCopy             = &Errno{Code: New(14), Msg: "copy data error"}
	ErrNoTokenUid       = &Errno{Code: New(15), Msg: "no token uid error"}
	ErrTooManyReqyest   = &Errno{Code: New(16), Msg: "too many request error"}
)
