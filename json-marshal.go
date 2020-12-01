package main

import (
	"encoding/json"
	"fmt"
)

type person struct {
	First string
	Last  string
	Age   int
}

func main() {

	p1 := person{
		First: "Толик",
		Last:  "К.",
		Age:   39,
	}

	p2 := person{
		First: "Галина",
		Last:  "В.",
		Age:   48,
	}

	people := []person{p1, p2}

	bs, err := json.Marshal(people)

	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(string(bs))

}
