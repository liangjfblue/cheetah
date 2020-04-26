/*
@Time : 2020/4/1 22:06
@Author : liangjiefan
*/
package _select

import (
	"errors"
	"math/rand"
)

type Select interface {
	Index(members []string) (string, error)
}

//RandSelect rand select algorithm
type RandSelect struct{}

func NewRandSelect() *RandSelect {
	return &RandSelect{}
}

func (r *RandSelect) Index(nodeInfos []string) (string, error) {
	if len(nodeInfos) <= 0 {
		return "", errors.New("srv no node")
	}

	max := len(nodeInfos)
	idx := rand.Intn(max)
	return nodeInfos[idx], nil
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
func (r *RotaSelect) Index(nodeInfos []string) (string, error) {
	max := len(nodeInfos)
	if max <= 0 {
		return "", errors.New("srv no node")
	}

	r.index++
	if r.index >= max {
		r.index = 0
	}

	return nodeInfos[r.index], nil
}
