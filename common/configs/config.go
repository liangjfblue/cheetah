package configs

import "os"

var (
	TraceContext string
	TraceAddr    string
	TraceParam   string

	TokenKey  = "jhf987y01h1j1h89"
	TokenTime = 3600
)

func init() {
	TraceAddr = os.Getenv("CONFIGOR_TRACE_ADDRESS")
	TraceContext = os.Getenv("CONFIGOR_TRACE_TRACECONTEXT")
	TraceParam = os.Getenv("CONFIGOR_TRACE_REQPARAM")
	if TraceAddr == "" {
		TraceAddr = "127.0.0.1:6831"
	}
	if TraceContext == "" {
		TraceContext = "trace_ctx"
	}
	if TraceParam == "" {
		TraceParam = "req_param"
	}
}
