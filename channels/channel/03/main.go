package main

import "fmt"

func main() {

	gen := make(chan int)
	b := make(chan int)

	go func() {
		for i := 0; i < 10; i++ {
			gen <- i
		}
		close(gen)
	}()

	go func() {
		for i := range gen {
			b <- i * i
		}
		close(b)
	}()

	for i := range b {
		fmt.Println(i)
	}

}
