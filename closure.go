package main

import (
	"fmt"
)

func test() {
	x, y := 10, 20
	defer func(i int) {
		fmt.Println("defer:", i, y) //y是闭包引用
	}(x) //x被复制
	x += 10
	y += 100
	fmt.Println("x=", x, "y=", y)
}

func main() {
	test()
}
