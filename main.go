package main

import "fmt"

//スライス

func main() {
	var sl []int
	fmt.Println(sl)

	//明示的宣言
	var sl2 []int = []int{100, 200}
	fmt.Println(sl2)

	s3 := []string{"A", "B"}
	fmt.Println(s3)

	s4 := make([]int, 5)
	fmt.Println(s4)

	ssl5 := []int{100, 200}
	fmt.Println(ssl5)

	ssl5 = append(ssl5, 400)
	fmt.Println(ssl5)

	sl6 := make([]int, 5)
	fmt.Println(sl6)

	fmt.Println(len(sl6))

	fmt.Println(cap(sl6))

	sl7 := []int{100, 200}

	sl8 := sl7

	sl8[0] = 1000
	fmt.Println(sl7)
	fmt.Println(sl8)

	sl9 := []int{1, 2, 3, 4, 5}
	sl10 := make([]int, 5, 10)
	fmt.Println(sl10)
	n := copy(sl10, sl9)
	fmt.Println(n, sl10)

	sls := []string{"A", "B", "C"}
	fmt.Println(sls)

	for i, v := range sls {
		fmt.Println(i, v)
	}
}
