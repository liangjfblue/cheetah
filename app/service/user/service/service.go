package service

import (
	"github.com/liangjfblue/cheetah/common/token"
)

type Service struct {
	Token *token.Token
}

func New(token *token.Token) *Service {
	return &Service{
		Token: token,
	}
}
