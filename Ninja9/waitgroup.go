package main

import (
	"fmt"
	"runtime"
	"sync"
)

var wg sync.WaitGroup

func main() {

	fmt.Println("OS\t\t", runtime.GOOS)
	fmt.Println("ARCH\t\t", runtime.GOARCH)
	fmt.Println("CPUs\t\t", runtime.NumCPU())
	fmt.Println("Goroutines\t", runtime.NumGoroutine())

	wg.Add(1)
	go foo()

	wg.Add(1)
	go bar()

	fmt.Println("Goroutines\t", runtime.NumGoroutine())

	wg.Wait()

}

func foo() {
	defer wg.Done()
	for i := 0; i < 10; i++ {
		fmt.Println("foo:", i)
	}

}

func bar() {
	defer wg.Done()
	for i := 0; i < 10; i++ {
		fmt.Println("bar:", i)
	}
}
