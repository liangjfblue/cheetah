package utils

import (
	"fmt"
	"strconv"
)

func Int2Bool(v interface{}) bool {
	switch v.(type) {
	case string:
		if v == "0" {
			return false
		}
	default:
		ii := fmt.Sprint(v)
		i, _ := strconv.Atoi(ii)
		if i <= 0 {
			return false
		}
	}
	return true
}
