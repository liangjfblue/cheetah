/*
@Time : 2020/4/19 14:28
@Author : liangjiefan
*/
package token

import (
	"errors"
	"time"

	"github.com/liangjfblue/cheetah/common/comConfigs"

	"github.com/dgrijalva/jwt-go"
)

type Context struct {
	Uid string `json:"uid"`
}

type Token struct {
	JwtKey  string
	JwtTime int
}

var (
	_token = Token{
		JwtKey:  comConfigs.TokenKey,
		JwtTime: comConfigs.TokenTime,
	}
)

func secretFunc(secret string) jwt.Keyfunc {
	return func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, jwt.ErrSignatureInvalid
		}

		return []byte(secret), nil
	}
}

func Parse(tokenString string, secret string) (*Context, error) {
	ctx := &Context{}

	token, err := jwt.Parse(tokenString, secretFunc(secret))
	if err != nil {
		return ctx, err
	} else if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		ctx.Uid = claims["uid"].(string)
		return ctx, nil
	} else {
		return ctx, err
	}
}

func ParseRequest(token string) (*Context, error) {
	if len(token) == 0 {
		return &Context{}, errors.New("`Authorization` header token is 0")
	}

	return Parse(token, _token.JwtKey)
}

func SignToken(c Context) (tokenString string, err error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"uid": c.Uid,
		"nbf": time.Now().Unix(),
		"iat": time.Now().Unix(),
		"exp": time.Now().Add(time.Second * time.Duration(_token.JwtTime)).Unix(),
	})

	tokenString, err = token.SignedString([]byte(_token.JwtKey))
	return
}
