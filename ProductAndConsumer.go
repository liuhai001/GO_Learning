package main

import (
	"fmt"
	"time"
)

func Producer(id int, ch chan int) {
	for i := 1; i <= 10; i++ {
		ch <- i
		fmt.Printf("Producter:%d is producting, value:%d\n", id, i)
	}
}

func Consumer(id int, ch chan int, done chan bool) {
	for {
		value, ok := <-ch
		if ok {
			fmt.Printf("Consumer:%d is Consuming, value:%d\n", id, value)
		} else {
			fmt.Printf("Consumer: %d, closed\n", id)
			break
		}
	}
	done <- true
}

func main() {
	ch := make(chan int, 3)

	coNum := 3
	PoNum := 5
	done := make(chan bool, coNum)

	for i := 1; i <= coNum; i++ {
		go Consumer(i, ch, done)
	}

	for i := 1; i <= PoNum; i++ {
		go Producer(i, ch)
	}

	time.Sleep(time.Second)
	close(ch)

	for i := 1; i <= coNum; i++ {
		<-done
	}
}
