package main

import "fmt"

func main() {

	ch1 := make(chan int)
	ch2 := make(chan int)
	quit := make(chan bool)

	go send(ch1, ch2, quit)

	receive(ch1, ch2, quit)

	fmt.Println("about to exit")

}

func receive(ch1, ch2 <-chan int, quit <-chan bool) {

	for {
		select {
		case v := <-ch1:
			fmt.Println("from ch1 channel", v)
		case v := <-ch2:
			fmt.Println("from ch2 channel", v)
		case v, ok := <-quit:

			if !ok {
				fmt.Println("from comma ok bit", v, ok)
				return
			} else {
				fmt.Println("from comma ok bit", v)
			}

		}

	}

}

func send(ch1, ch2 chan<- int, quit chan<- bool) {

	for i := 0; i < 10; i++ {
		if i%2 == 0 {
			ch1 <- i
		} else {
			ch2 <- i
		}
	}
	close(quit)

}
