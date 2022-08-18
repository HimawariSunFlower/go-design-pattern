package factory

import _ "github.com/HimawariSunFlower/go-design-pattern"

type Operation[T Ordered] interface {
	GetResult() T
}

type GetResult func() int

type OperationAdd struct {
	A int
	B int
}

func (r OperationAdd) GetResult() int {
	return r.A + r.B
}
