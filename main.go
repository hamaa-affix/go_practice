package main

import "fmt"

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
}
