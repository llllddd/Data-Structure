package main

import (
	"fmt"
)

//Return Sum of elements in a array
func SumArray(array []int) int {
	var s int
	for _, item := range array {
		s = s + item
	}
	return s
}

func main() {
	a := []int{1, 2, 3, 4, 5}
	s := SumArray(a)
	fmt.Println(s)
}
