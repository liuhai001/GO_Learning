package main

import (
	"strconv"
	"fmt"
	"reflect"
)

func main() {
	f,_ := strconv.ParseFloat("1.234",64)
	fmt.Printf("%#v,%v\n",f,reflect.TypeOf(f))

	i,_ := strconv.ParseInt("123",0,64)//0 表示自动推断进制类型
	fmt.Printf("%#v,%v\n",i,reflect.TypeOf(i))

	d,_ := strconv.ParseInt("0x123",0,64)//0 表示自动推断进制类型
	fmt.Printf("%#v,%v\n",d,reflect.TypeOf(d))

	k,_ := strconv.Atoi("1234")
	fmt.Printf("%#v,%v\n",k,reflect.TypeOf(k))

	_,error := strconv.Atoi("liuhai")
	fmt.Println(error)

	str:= strconv.Itoa(1234)
	fmt.Printf("%#v,%v\n",str,reflect.TypeOf(str))

}
