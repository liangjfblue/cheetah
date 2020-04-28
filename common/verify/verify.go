/*
@Time : 2020/4/24 17:57
@Author : liangjiefan
*/
package verify

import (
	"reflect"
	"strings"

	"gopkg.in/go-playground/validator.v9"
)

func Validate(v interface{}) error {
	validate := validator.New()
	validate.RegisterTagNameFunc(func(fld reflect.StructField) string {
		name := strings.SplitN(fld.Tag.Get("json"), ",", 2)[0]
		if name == "-" {
			return ""
		}
		return fld.Tag.Get("msg") + " " + "[" + name + "]"
	})
	return validate.Struct(v)
}

func TranslateErr2MsgTag(err error) string {
	for _, err := range err.(validator.ValidationErrors) {
		return err.Field()
	}
	return err.Error()
}
