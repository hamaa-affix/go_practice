package main

import (
	"fmt"
	"strconv"
)

func ReturnFunc() func() {
	return func ()  {
		fmt.Println("I am fucntion")
	}
}

func CallFunction(f func()) {
	f()
}

 //クロージャー
  func Later() func(string) string {
		var store string
		return func (next string) string {
				s := store
				store = next
				return s
		}
	}

	//ジェネレーター なんらかのルールをもとに連続した値を返す。go ではジェネレーターはないがクロージャーで実装可能
	func integers() func () int {
		i := 0
		return func() int {
			  i ++
				return i
		}
	}

func main() {

	//無記名関数 クロージャー
	f := func (x, y int) int {
			return x + y
	}
	i := f(1, 2)
	fmt.Println(i)

	//関数を返す関数
	f2 := ReturnFunc()
	f2()

	//引数に関数をとる。クロージャ的
	CallFunction(func ()  {
		fmt.Println("Iam argu")
	})

	f4 := Later()
	fmt.Println(f4("hello"))

	f5 := integers()
	fmt.Println(f5())
	fmt.Println(f5())
	fmt.Println(f5())
	fmt.Println(f5())

	//if文
	a := 0
	if a == 2 {
		fmt.Println("two")
	} else {
		fmt.Println("i don't know")
	}
	//errorハンドリング
	var t string = "100"

	i, err := strconv.Atoi(t)
	if err != nil {
		println(err)
	}
	println("t = %T\n", t)

	o := 0
	for {
		o ++
		if o == 3 {
			break
		}
		fmt.Println("hi")
	}

	point := 0
	for point < 10 {
		fmt.Println(point)
		point++
	}

	// for i := 0; i < 10; i++ {
	// 	if i == 3 {
	// 		continue
	// 	}
	// 	fmt.Println(i)
	// }

	 arr := [3]int{1, 2, 3}
	// for i:=0; i < len(arr); i++ {
	// 	fmt.Println(arr[i])
	// }

	for i, v := range arr {
		fmt.Println(i, v)
	}
}
