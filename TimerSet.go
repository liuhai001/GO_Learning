package main

import (
	"fmt"
	"time"
)

func main() {
	timer1 := time.NewTimer(time.Second * 2)
	<-timer1.C
	fmt.Println("timer1 is expired!")

	timer2 := time.NewTimer(time.Second * 3)
	go func() {
		for {
			fmt.Println("going")
			if _, ok := <-timer2.C; ok {
				fmt.Println("timer2 is expired!")
			} else {
				fmt.Println("timer is stop!")
			}
		}

	}()
	time.Sleep(time.Second * 1)
	timer2.Stop()
	time.Sleep(time.Second * 4)
}
