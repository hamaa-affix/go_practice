package main

import "fmt"

//定数

/*
	定数はグローバルに定義数ことがデフォ
*/
const Pi = 3.14 //頭文字を大文字にすると他のパッケージから呼び出せる

/*
	複数の定数を適宜数
*/
const (
	Url = "http://xxx.com"
	SiteName = "test"
)

func main() {
	//後から値が変更できない
	fmt.Println(Pi)
	fmt.Println(Url, SiteName)
}
