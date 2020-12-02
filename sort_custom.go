package main

import (
	"fmt"
	"sort"
)

type person struct {
	first string
	age   int
}

type peopleByAge []person

func (p person) String() string {
	return fmt.Sprintf("%s: %d", p.first, p.age)
}

func (a peopleByAge) Len() int           { return len(a) }
func (a peopleByAge) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a peopleByAge) Less(i, j int) bool { return a[i].age < a[j].age }

func main() {

	p1 := person{"An", 39}
	p2 := person{"Gala", 48}
	p3 := person{"Vika", 28}
	p4 := person{"Dima", 31}

	people := []person{p1, p2, p3, p4}

	fmt.Println(people)

	sort.Sort(peopleByAge(people))
	fmt.Println()
	fmt.Println(people)

}
