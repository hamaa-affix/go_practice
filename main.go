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

	func anything(a interface{}) {
		switch v := a.(type) {
		case string:
			fmt.Println(v + "!?")
		case int:
			fmt.Println(v + 10000)
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

	//スイッチ文
	// n := 1
	// switch n {
	// case 1, 2:
	// 	fmt.Println("1 or 2")
	// case 3, 4:
	// 	fmt.Println("3 or 4")
	// default:
	// 	fmt.Println("I don't know")
	// }
	// switch n :=2; n {
	// case 1, 2:
	// 	fmt.Println("1 or 2")
	// case 3, 4:
	// 	fmt.Println("3 or 4")
	// default:
	// 	fmt.Println("I don't know")
	// }
	n := 6
	switch {
	case n > 0 && n < 4:
		fmt.Println("0 < n < 4")
	case n > 3 && n < 7:
		fmt.Println("3 < n < 7")
	}

	anything("aaa")
	anything(1)

	//型アサーション
	var x interface{} = 3
	it := x.(int)
	fmt.Println(it + 2)
}
