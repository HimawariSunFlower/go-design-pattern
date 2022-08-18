package decorator

import "testing"

func TestDecorator(t *testing.T) {
	p := Person{name: "xiaocai", finery: []string{}}
	p.Decorate(TShirtFunc)
	p.Decorate(TrouserFunc)
	p.Show()
}
