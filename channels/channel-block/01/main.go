package main

import "fmt"

func main() {

	c := make(chan int)

	go foo(c)
	bar(c)

	fmt.Println("Exit")

}

func foo(c chan<- int) {
	c <- 39
}

func bar(c <-chan int) {
	fmt.Println(<-c)
}
