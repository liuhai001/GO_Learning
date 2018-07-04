package main

import (
	"fmt"
	"runtime"
	"sync"
)

func main() {
	wg := new(sync.WaitGroup)
	wg.Add(2)

	go func() {
		defer wg.Done()
		for i := 0; i < 6; i++ {
			fmt.Println(i)
			if i == 3 {
				runtime.Gosched()
			}
		}
	}()

	go func() {
		defer wg.Done()
		for i := 0; i < 5; i++ {
			if i == 1 {
				runtime.Gosched()
			}
			fmt.Println(i, "Hello,World!")
		}

	}()

	wg.Wait()
}
