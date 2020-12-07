package main

import (
	"fmt"
	"sync"
)

func main()  {

	wg := sync.WaitGroup{}

	wg.Add(2)

	go bar(&wg)
	go foo(&wg)

	wg.Wait()

}

func foo(wg *sync.WaitGroup)  {
	defer wg.Done()
	fmt.Println("foo")
}

func bar(wg *sync.WaitGroup)  {
	defer wg.Done()
	fmt.Println("bar")
}