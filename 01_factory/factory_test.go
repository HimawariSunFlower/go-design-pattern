package factory

import (
	"fmt"
	"testing"
)

func TestSimpleFactory(t *testing.T) {
	oper1 := createOperate[int]("+")
	oper1.AddParam(1, 2)
	fmt.Println(oper1.GetResult())

	oper2 := createOperate[uint]("-")
	oper2.AddParam(3, 4)
	fmt.Println(oper2.GetResult())
}
