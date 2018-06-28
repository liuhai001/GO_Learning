package main

import (
	"fmt"
)

func main() {
	//---function variable
	fn := func() { fmt.Println("Hello,World!") }
	fn()
	//---function collection
	fns := [](func(x int) int){
		func(x int) int { return x + 1 },
		func(x int) int { return x + 2 },
	}
	fmt.Println(fns[1](100))
	//---function as field
	d := struct {
		fn func() string
	}{
		fn: func() string { return "Hello,World!" },
	}

	fmt.Println(d.fn())
	//-------channel of function
	fc := make(chan func() string, 2)
	fc <- func() string { return "Hello,LiuHai!" }
	fmt.Println((<-fc)())
}
