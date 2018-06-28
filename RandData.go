package main

import (
	"fmt"
	"math/rand"
)

func main() {
	//每次调用都返回的是一样的随机数，没有做到真正的随机
	fmt.Print(rand.Intn(100), ",") //[0,100]
	fmt.Print(rand.Intn(100))
	fmt.Println()
	fmt.Println(rand.Float64()) // [0.0,1.0]

	//确定性生成随机数，给明确种子
	s1 := rand.NewSource(42)
	r1 := rand.New(s1)
	fmt.Print(r1.Intn(100), ",")
	fmt.Print(r1.Intn(100))
	fmt.Println()

	s2 := rand.NewSource(42)
	r2 := rand.New(s2)
	fmt.Print(r2.Intn(100), ",")
	fmt.Print(r2.Intn(100))
	fmt.Println()

	s3 := rand.NewSource(10)
	r3 := rand.New(s3)
	fmt.Print(r3.Intn(100), ",")
	fmt.Print(r3.Intn(100))
	fmt.Println()
}
