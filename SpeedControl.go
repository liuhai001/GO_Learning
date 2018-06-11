package main

import (
	"fmt"
	"time"
)

func main() {
	requests := make(chan int, 5)
	for i := 1; i <= 5; i++ {
		requests <- i
	}
	close(requests)
	limiter := time.Tick(time.Second)

	for k := range requests {
		<-limiter
		fmt.Println("req", k, time.Now())
	}

	burstyLimiter := make(chan time.Time, 2)
	for i := 1; i <= 2; i++ {
		burstyLimiter <- time.Now()
	}
	go func() {
		for t := range time.Tick(time.Second) {
			burstyLimiter <- t
		}
	}()

	burstyrequests := make(chan int, 5)
	for i := 1; i <= 5; i++ {
		burstyrequests <- i
	}
	close(burstyrequests)
	for k := range burstyrequests {
		<-burstyLimiter
		fmt.Println("burstyrequests", k, time.Now())
	}

}
