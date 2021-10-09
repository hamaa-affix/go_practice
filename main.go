package main

import (
	"fmt"
)

func main() {
	var m = map[string]int{"A": 100, "B": 200}
	fmt.Println(m)

	m1 := map[string]int{"B": 300, "C": 400}
	fmt.Println(m1)

	m2 := make(map[int]string)
	fmt.Println(m2)
	m2[1] = "JAPAN"
	m2[2] = "USA"
	fmt.Println(m2)

	_, isValue := m2[3]
	if !isValue {
		 fmt.Println("error")
	}

	for i, v := range m2 {
		fmt.Println(i, v)
	}
}
