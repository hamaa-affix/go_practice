package main

import "fmt"

//変数

func outer() {
		s4 := "outer"
		fmt.Println(s4)
}

func main() {
		//明示的な定義
		//var 変数名 型 = 値
		var i int = 100
		var s string = "hello Go"
		var t,f bool = true, false
		//異なる型で定義する場合
		var (
			i2 int = 200
			s2 string = "hello"
		)
		//変数だけ定義する
		var i3 int
		var s3 string

		i3 = 300
		s3 = "GO"

		//再代入
		i = 150

		fmt.Println(i)
		fmt.Println(s)
		fmt.Println(t, f)
		fmt.Println(i2, s2)
		fmt.Println(i3, s3)

		//暗黙的な定義(型宣言をしなくても型注釈してくれる)
		//変数名 := 値
		i4 := 400
		fmt.Println(i4)
		//再代入も可能
		i4 = 500
		fmt.Println(i4)
		//再定義はできない
		// i4 :=500

		//一度定義された型は再代入時に型を変えることはできない
		//i4 = "stringへ" //これはできない。int型で宣言指定いるのでint型の値しか入れれない
		outer()
}
