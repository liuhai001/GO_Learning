package main

import (
	"fmt"
	"net/url"
)

func main() {
	urlstr := "https://www.shiyanlou.com/courses/485/labs/1616/document"
	u, err := url.Parse(urlstr)
	if err != nil {
		panic(err)
	}
	fmt.Println(u.Scheme)
	fmt.Println(u.User)
	if u.User !=nil{
		fmt.Println(u.User.Username())
		p,_ := u.User.Password()
		fmt.Println(p)
	}

	fmt.Println(u.Host)
	fmt.Println(u.Path)
	fmt.Println(u.Fragment)
	fmt.Println(u.RawQuery)



}
