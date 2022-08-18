package strategy

//策略模式，定义一族算法类，分别封装，相互之间可以互相替换，不互相影响

type CashContext interface {
	GetResult(float64) float64
}

func NewCashContext(t CashContext) CashContext {
	//do something
	return t
}

type CashNormal struct {
}

func (r CashNormal) GetResult(input float64) float64 {
	return input
}

type CashReturn struct {
	A float64
	B float64
}

func (r CashReturn) GetResult(input float64) float64 {
	ret := float64(0)
	for input > r.A {
		input -= r.A
		ret += r.A - r.B
	}
	return ret + input
}
