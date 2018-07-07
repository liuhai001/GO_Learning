//反射中调用方法
package main

import (
	"fmt"
	"reflect"
	"strconv"
)

type Mytype struct {
	i    int
	Name string
}

func (m *Mytype) SetI(i int) {
	m.i = i
}

func (m *Mytype) SetName(name string) {
	m.Name = name
}

func (m *Mytype) Sprint() string {
	return fmt.Sprintf("%p:", m) + "--name:" + m.Name + "--i:" + strconv.Itoa(m.i)
}

func main() {
	mytype := &Mytype{26, "LIUHAI"}
	fmt.Println(reflect.TypeOf(mytype))
	mtV := reflect.ValueOf(&mytype).Elem()
	fmt.Println(reflect.TypeOf(mtV))
	fmt.Printf("%#v\n",mtV)

	fmt.Println("Before:", mtV.MethodByName("Sprint").Call(nil)[0])
	params := make([]reflect.Value, 1)
	params[0] = reflect.ValueOf(18)
	mtV.MethodByName("SetI").Call(params)
	params[0] = reflect.ValueOf("reflection test")
	mtV.MethodByName("SetName").Call(params)
	fmt.Println("After:", mtV.MethodByName("Sprint").Call(nil)[0])


}
