package strategy

import (
	"fmt"
	"math/rand"
	"testing"
	"time"
)

func TestStrategy(t *testing.T) {
	typ := 1

	if rand.New(rand.NewSource(time.Now().UnixNano())).Intn(1000) > 500 {
		typ = 2
	}
	switch typ {
	case 1:
		fmt.Println(NewCashContext(CashNormal{}).GetResult(100))
	case 2:
		fmt.Println(NewCashContext(CashReturn{A: 300, B: 100}).GetResult(900))
	}
}
