package main

import "fmt"

//関数
func Plus(x int, y int) int {
	return x + y
}
//複数戻り値がある場合
func Div(x int, y int) (int, int) {
	plus := x + y
	minus := x - y
	return plus, minus
}

//戻り値の変数を指定することができる
//return の値が省略されているが、戻り値で変数をしているので省略することができる
func Double(price int) (result int) {
	result = price  * 2
	return
}

//返り値がない型
func NoReturn() {
	fmt.Println("no return")
	return
}

func main() {
	i := Plus(2, 4)
	fmt.Println(i)//6
	fmt.Println(Plus(1, 4)) //5

	i2, i3 := Div(3, 2)
	fmt.Println(i2, i3)

	i4, _ := Div(4, 5) //変数をアンダーバーで定義すると戻り値を破棄することができる
	fmt.Println(i4)

	i5 := Double(300)
	fmt.Println(i5)

	NoReturn()
}
