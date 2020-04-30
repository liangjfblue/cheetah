package balancer

import "errors"

var (
	ErrNotFoundInstance = errors.New("not found instance")
	ErrEmptyInstance    = errors.New("empty instance")
)
