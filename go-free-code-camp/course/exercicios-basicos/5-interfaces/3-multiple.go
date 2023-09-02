package main

import (
	"fmt"
)

func (e email) cost() float64 {

	if !e.isSubscribed {
		return 0.05 * float64(len(e.body))
	} else {
		return 0.01 * float64(len(e.body))
	}

}

func (e email) print() {
	fmt.Println(e.body)
}

type expense interface {
	cost() float64
}

type printer interface {
	print()
}

type email struct {
	isSubscribed bool
	body         string
}

func print(p printer) {
	p.print()
}

func test(e expense, p printer) {
	fmt.Printf("Imprimindo os custos: $%.2f ...\n", e.cost())
	p.print()
	fmt.Println("====================================")
}

func main() {
	e := email{
		isSubscribed: true,
		body:         "Olá mundo",
	}
	test(e, e)
	e = email{
		isSubscribed: false,
		body:         "quero meu dinheiro de volta",
	}
	test(e, e)
	e = email{
		isSubscribed: true,
		body:         "ta disponivel pra conversar?",
	}
	test(e, e)
	e = email{
		isSubscribed: false,
		body:         "Esse mett podia ter sido um email",
	}
	test(e, e)
}
