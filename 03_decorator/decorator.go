package decorator

import "fmt"

type fineryFunc func(person *Person)

type Person struct {
	name   string
	finery []string
}

func (p Person) Show() {
	for _, v := range p.finery {
		fmt.Print(v, " ")
	}
	fmt.Println("的", p.name)
}
func (p *Person) Decorate(f fineryFunc) {
	f(p)
}

var TShirtFunc = func(p *Person) {
	p.finery = append(p.finery, "T恤")
}

var TrouserFunc = func(p *Person) {
	p.finery = append(p.finery, "裤子")
}
