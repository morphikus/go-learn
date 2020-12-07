package main

import (
	"fmt"
	"runtime"
	"sync"
)

func main() {

	fmt.Println("NumGoroutine", runtime.NumGoroutine())
	fmt.Println("NumCPU", runtime.NumCPU())

	var wg sync.WaitGroup

	wg.Add(2)

	go bar(&wg)
	go foo(&wg)

	fmt.Println("NumGoroutine", runtime.NumGoroutine())

	wg.Wait()

	fmt.Println("Exit")

}

func foo(wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Println("foo")
}

func bar(wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Println("bar")
}
