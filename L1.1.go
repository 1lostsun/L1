package main

import "fmt"

func main() {
	hum := Human{
		name:  "Yaroslav",
		age:   19,
		phone: "+79999999999",
	}

	act := Action{hum}

	act.Walk(act.name)
}

type Human struct {
	name  string
	age   int
	phone string
}

type Action struct {
	Human
}

func (h Human) Walk(name string) {
	fmt.Printf("human %s is walking", name)
}
