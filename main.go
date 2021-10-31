package main

import (
	"fmt"
	"newmod/alib"
	"newmod/foo"
	"log"
	"os"
	"sort"
)

//test
func IsOne(i int) bool {
	if i == 1 {
		return true
	} else {
		return false
	}
}

type Entry struct {
		Name string
		Value int
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

		//sort
		i := []int{2, 4 ,5, 1, 3}
		sort.Ints(i)
		fmt.Println(i)

		string := []string{"a", "y", "u", "b"}
		sort.Strings(string)
		fmt.Println(string)

		//sliceテーブル
		el := []Entry{
			{"A", 20},
			{"t", 30},
			{"b", 40},
			{"c", 50},
			{"t", 20},
			{"a", 40},
		}

		fmt.Println(el)
		sort.Slice(el, func(i, j int) bool { return el[i].Name < el[j].Name})
		fmt.Println(el)
}
