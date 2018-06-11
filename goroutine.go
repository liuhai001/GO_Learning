package main

import (
	"fmt"
	"time"
)
//非缓冲型的通道是阻塞的，写了之后要等读出去才能执行下一行代码，读的操作也会一直在等待写入；如果写了没读或者没写一直在读，都会导致goroutine死锁deadlock！

func main() {
	messages := make(chan string)
	//messages <- "liuhai"
	go func() {
		msg := "hi"
		select {
		case messages <- msg:
			fmt.Println("sent message", msg)
		}
	}()

	select {
	case msg := <-messages:
		fmt.Println("received message", msg)
	}

	time.Sleep(time.Second * 2)
}
