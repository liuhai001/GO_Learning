package main

import (
	"fmt"
)

type User struct {
	id   int
	name string
}

func (self *User) String() string {
	return fmt.Sprintf("user:%d,%s", self.id, self.name)
}

func main() {
	var o interface{} = &User{1, "liuhai"}
	if i, ok := o.(fmt.Stringer); ok {
		fmt.Println(i)
	}

	u := o.(*User)
	fmt.Println(u)

	var _ fmt.Stringer = (*User)(nil)

}
