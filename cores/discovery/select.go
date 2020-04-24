/*
@Time : 2020/4/1 22:06
@Author : liangjiefan
*/
package discovery

import (
	"errors"
	"math/rand"
)

type Select interface {
	Index(members []*Member) (*Member, error)
}

//RandSelect rand select algorithm
type RandSelect struct{}

func NewRandSelect() *RandSelect {
	return &RandSelect{}
}

func (r *RandSelect) Index(members []*Member) (*Member, error) {
	if len(members) <= 0 {
		return nil, errors.New("srv no node")
	}

	max := len(members)
	idx := rand.Intn(max)
	return members[idx], nil
}

//RotaSelect rota select algorithm
type RotaSelect struct {
	index int
}

func NewRotaSelect() *RotaSelect {
	return &RotaSelect{
		index: 0,
	}
}
func (r *RotaSelect) Index(members []*Member) (*Member, error) {
	max := len(members)
	if max <= 0 {
		return nil, errors.New("srv no node")
	}

	r.index++
	if r.index >= max {
		r.index = 0
	}

	return members[r.index], nil
}
