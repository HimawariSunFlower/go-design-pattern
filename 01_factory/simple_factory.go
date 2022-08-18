package factory

import (
	"github.com/HimawariSunFlower/go-design-pattern/util"
)

type Operation[T util.Integer] interface {
	AddParam(A, B T)
	GetResult() T
}

type OperationAdd[T util.Integer] struct {
	A T
	B T
}

func (r *OperationAdd[T]) GetResult() T {
	return r.A + r.B
}
func (r *OperationAdd[T]) AddParam(A, B T) {
	r.A = A
	r.B = B
}

type OperationSub[T util.Integer] struct {
	A T
	B T
}

func (r *OperationSub[T]) GetResult() T {
	return r.A - r.B
}
func (r *OperationSub[T]) AddParam(A, B T) {
	r.A = A
	r.B = B
}

func createOperate[T util.Integer](input string) Operation[T] {
	switch input {
	case "+":
		return new(OperationAdd[T])
	case "-":
		return new(OperationSub[T])
	default:
		//error 未完成对应运算
		return nil
	}
}
