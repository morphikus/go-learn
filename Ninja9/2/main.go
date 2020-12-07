package main

import "fmt"

type person struct {
	first string
	last  string
	age   int
}

type human interface {
	speak()
}

func (p *person) speak() {
	fmt.Println("Hello, I'm", p.first, p.last, p.age)
}

func saySomething(h human) {
	h.speak()
}

func main() {

	p := person{
		first: "An",
		last:  "Koch",
		age:   39,
	}

	p.speak()

	saySomething(&p)
	//saySomething(p) //wrong

}
