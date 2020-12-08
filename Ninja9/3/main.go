package main

import (
	"fmt"
	"runtime"
	"sync"
)

var counter int
var wg sync.WaitGroup
var mu sync.Mutex

func main() {

	n := 100
	counter = 0

	wg.Add(n)

	for i := 0; i < n; i++ {

		go func() {

			mu.Lock()
			defer mu.Unlock()

			defer wg.Done()

			c := counter
			c++
			counter = c

			runtime.Gosched()

		}()

	}

	wg.Wait()

	fmt.Println("Counter", counter)
	fmt.Println("Exit")

}
