package main

import (
	"fmt"
	//"time"
)

//pointer
func Dubole(i int) {
	i = i * 2
	fmt.Println(&i)
}
//ポインタで受け取る仕様の関数を定義する必要がある
func Dubolev2(i *int) {
	*i = *i * 2
}

func main() {
	var n int = 100
	fmt.Println(n)

	//memoryのアドレスを確認する
	fmt.Println(&n)

	//下記だと値渡しになっているから下記の関数での変数のアドレスは違うもになる
	Dubole(n)

	//参照渡しにする為にはポインタ型で値を渡してやる必要がある
	var p *int = &n
	Dubolev2(p)
	fmt.Println(n)
}
