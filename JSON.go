package main

import (
	"encoding/json"
	"fmt"
)

type Response1 struct {
	Page   int
	Fruits []string
}
type Response2 struct {
	Page   int      `json:"liuhai"`
	Fruits []string `json:"zhouyou"`
}

func main() {
	bolB, _ := json.Marshal(true)
	fmt.Println(string(bolB))
	fmt.Printf("%#v\n", bolB)

	intB, _ := json.Marshal(1)
	fmt.Println(string(intB))
	fmt.Printf("%#v\n", intB)

	fltB, _ := json.Marshal(2.34)
	fmt.Println(string(fltB))
	fmt.Printf("%#v\n", fltB)

	strB, _ := json.Marshal("gopher")
	fmt.Println(string(strB))
	fmt.Printf("%#v\n", strB)

	// 这里是一些切片和 map 编码成 JSON 数组和对象的例子。
	slcD := []string{"apple", "peach", "pear"}
	slcB, _ := json.Marshal(slcD)
	fmt.Println(string(slcB))
	fmt.Printf("%#v\n", slcB)

	mapD := map[string]int{"apple": 5, "lettuce": 7}
	mapB, _ := json.Marshal(mapD)
	fmt.Println(string(mapB))
	fmt.Printf("%#v\n", mapB)

	res1D := &Response1{
		Page :1,
		Fruits: []string{"apple","peach","pear"}}
	res1B,_ := json.Marshal(res1D)
	fmt.Println(string(res1B))
	fmt.Printf("%#v\n",res1B)

	res2D := &Response2{
		Page :1,
		Fruits: []string{"apple","peach","pear"}}
	res2B,_ := json.Marshal(res2D)
	fmt.Println(string(res2B))
	fmt.Printf("%#v\n",res2B)

	//var dat map[string]interface{}
	res2G := &Response2{}

	if err := json.Unmarshal(res2B,&res2G);err!=nil{
		panic(err)
	}
	fmt.Printf("%#v\n",res2G)
}
