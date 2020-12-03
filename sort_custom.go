package main

import (
	"fmt"
	"sort"
)

type person3 struct {
	first string
	age   int
}

type peopleByAge []person3
type peopleByFirst []person3

func (p person3) String() string {
	return fmt.Sprintf("%s: %d", p.first, p.age)
}

func (a peopleByAge) Len() int           { return len(a) }
func (a peopleByAge) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a peopleByAge) Less(i, j int) bool { return a[i].age < a[j].age }

func (a peopleByFirst) Len() int           { return len(a) }
func (a peopleByFirst) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a peopleByFirst) Less(i, j int) bool { return a[i].first < a[j].first }

func main() {

	p1 := person3{"An", 39}
	p2 := person3{"Gala", 48}
	p3 := person3{"Vika", 28}
	p4 := person3{"Dima", 31}

	people := []person3{p1, p2, p3, p4}

	fmt.Println(people)

	sort.Sort(peopleByAge(people))
	fmt.Println()
	fmt.Println(people)

	sort.Sort(peopleByFirst(people))
	fmt.Println()
	fmt.Println(people)

}
