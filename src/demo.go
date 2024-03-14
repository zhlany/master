package main

import (
	"fmt"
	"runtime"
	"sync"
)

func main() {
	runtime.GOMAXPROCS(1)
	var wg sync.WaitGroup
	for v := range 10 {
		wg.Add(1)
		go func() {
			fmt.Print(v, "-")
			wg.Done()
		}()
	}
	wg.Wait()
	fmt.Println("\nall goroutines finished")
}
