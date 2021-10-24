package main

import (
	"fmt"
	"newmod/alib"
	"newmod/foo"
	"log"
	"os"
)

//test
func IsOne(i int) bool {
	if i == 1 {
		return true
	} else {
		return false
	}
}

//logの出力

//package  のインポート
func main() {
		fmt.Println(foo.Max)
		fmt.Println(foo.ReturnMin())

		s := []int{1, 2, 3, 4, 5}
		fmt.Println(alib.Average(s))

		//標準出力を行う
		log.SetOutput(os.Stdout)//logの出力先を標準出力に変える
		log.Print("Log\n")
		log.Println("Log2")
		log.Printf("Log%d\n", 3)

		//任意のファイルにログを出力させる
		f, err := os.Create("test.log")
		if err != nil {
			return
		}
		//作成したio.Writer型のファイルを出力先に設定する
		log.SetOutput(f)
		log.Println("ファイルに書き込む")
}
