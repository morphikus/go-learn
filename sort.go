package main

import (
	"fmt"
	"sort"
)

func main() {

	xi := []int{56, 67, 98, 3, 1, 0, -5}
	xi = append(xi, -100)

	xs := []string{"Формула", "Бабушка", "Ю", "І", "99", "А", "D", "Z", "I"}

	fmt.Println(xi)
	fmt.Println(xs)

	sort.Ints(xi)
	sort.Strings(xs)

	fmt.Println()

	fmt.Println(xi)
	fmt.Println(xs)

}
