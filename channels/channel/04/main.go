package main

import "fmt"

func  main()  {

	ch1 := make(chan int)
	ch2 := make(chan int)
	ch3 := make(chan int)

	go send(ch1, ch2, ch3)

	receive(ch1, ch2, ch3)

	fmt.Println("about to exit")

}

func receive(ch1, ch2, ch3 <-chan int)  {

	for{
		select {
			case v := <- ch1:
				fmt.Println("from ch1 channel", v)
			case v := <- ch2:
				fmt.Println("from ch2 channel", v)
			case v := <-ch3:
				fmt.Println("from ch3 channel", v)
				return
		}

	}

}

func send(ch1, ch2, ch3 chan<- int)  {

	for i := 0; i < 10; i++{
		if i % 2 == 0 {
			ch1 <- i
		} else {
			ch2 <- i
		}
	}
	ch3 <- 0

}
