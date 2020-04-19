/*
@Time : 2020/4/19 14:28
@Author : liangjiefan
*/
package token

import (
	"errors"
	"time"

	"github.com/dgrijalva/jwt-go"
)

type Context struct {
	Uid string `json:"uid"`
}

type Token struct {
	JwtKey  string
	JwtTime int
}

func New(jwtKey string, jwtTime int) *Token {
	return &Token{
		JwtKey:  jwtKey,
		JwtTime: jwtTime,
	}
}

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

func (t *Token) ParseRequest(token string) (*Context, error) {
	if len(token) == 0 {
		return &Context{}, errors.New("`Authorization` header token is 0")
	}

	return Parse(token, t.JwtKey)
}

func (t *Token) SignToken(c Context) (tokenString string, err error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"uid": c.Uid,
		"nbf": time.Now().Unix(),
		"iat": time.Now().Unix(),
		"exp": time.Now().Add(time.Second * time.Duration(t.JwtTime)).Unix(),
	})

	tokenString, err = token.SignedString([]byte(t.JwtKey))
	return
}
