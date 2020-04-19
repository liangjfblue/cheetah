package errno

import (
	"fmt"
)

var (
	_codes = map[int32]struct{}{}
)

func New(e int32) int32 {
	if e <= 0 {
		panic("business ecode must greater than zero")
	}
	return add(e)
}

func add(e int32) int32 {
	if _, ok := _codes[e]; ok {
		panic(fmt.Sprintf("ecode: %d already exist", e))
	}
	_codes[e] = struct{}{}
	return e
}

type Errno struct {
	Code int32       `json:"Code"`
	Msg  string      `json:"Msg"`
	Data interface{} `json:"Data,omitempty"`
}

func (e Errno) Error() string {
	return e.Msg
}
