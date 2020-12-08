package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

func main() {

	var wg sync.WaitGroup
	var counter int64

	n := 100

	wg.Add(n)

	for i := 0; i < n; i++ {
		go func() {
			atomic.AddInt64(&counter, 1)
			r := atomic.LoadInt64(&counter)
			fmt.Println(r)
			wg.Done()
		}()
	}

	wg.Wait()

	fmt.Println("Counter", counter)

}
