package main

import "fmt"

func main() {
	jobs := make(chan int, 5)
	done := make(chan bool)

	go func() {
		for {
			j, more := <-jobs
			if more {
				fmt.Printf("received job:%v\n", j)
			} else {
				fmt.Println("received all jobs")
				done <- true
				return
			}
		}
	}()

	for i := 0; i < 5; i++ {
		jobs <- i
		fmt.Printf("sent job:%v\n", i)
	}
	close(jobs)

	<-done

}
