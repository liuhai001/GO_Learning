package main

import (
	"fmt"
	//"math/rand"
)

func bubbleSort(a []int) {
	n := len(a)
	for i := 0; i < n; i++ {
		for j := n - 1 - i; j > 0; j-- {
			if a[j-1] > a[j] {
				a[j-1], a[j] = a[j], a[j-1]
			}
		}
	}

}

func main() {
	arr := []int{2, 1, 3, 4, 2, 6, 1, 8, 11}
	fmt.Println(arr)
	bubbleSort(arr)
	fmt.Println(arr)

}
