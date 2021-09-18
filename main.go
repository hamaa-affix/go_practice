package main

import "fmt"

//演算
func main() {
	fmt.Println(1 + 2) //3

	fmt.Println(1 - 2) //-1

	fmt.Println(5 * 2) //10

	fmt.Println(30 / 2) //15

	fmt.Println(9 % 4) //1
	//文字列結合
	fmt.Println("ABC" + "def") //ABCdef

	n := 0
	n += 4 //n = n + 4
	fmt.Println(n)
	n ++ //n = n + 1
	n -- //n = n -1

	s := "ABCD"
	s += "ef" //s = s + "ef"
	fmt.Println(s)

	//比較演算子
	fmt.Println(1 == 1) //true

	fmt.Println(1 == 2) //false

	fmt.Println(4 <= 4) //true

	fmt.Println(4 >= 6) //false

	fmt.Println(4 < 4) //false

	fmt.Println(true != false) //true

	//論理演算子
	fmt.Println(true && (false == true)) //false

	fmt.Println(true || false == true) //true

	fmt.Print(!true) //false

	fmt.Print(!true == true) //false
}
