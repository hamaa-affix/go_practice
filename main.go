package main

import (
	"fmt"
	"newmod/alib"
	"newmod/foo"
)

//test
func IsOne(i int) bool {
	if i == 1 {
		return true
	} else {
		return false
	}
}

//package  のインポート
func main() {
		fmt.Println(foo.Max)
		fmt.Println(foo.ReturnMin())

		s := []int{1, 2, 3, 4, 5}
		fmt.Println(alib.Average(s))
}
