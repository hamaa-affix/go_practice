package main

import (
	"fmt"
	//"time"
)

func receiver(c chan int) {
	for {
		i := <- c
		fmt.Println(i)
	}
}

func main() {

	//channel
	//go routien間でデータの受け渡しを行う為に設計されたデータ構造
	//宣言、操作

	//双方向のチャネル
	var ch1 chan int

	//受信専用channel
	// var ch2 <-chan int

	// //送信専用channel
	// var ch3 chan<- int

	 //これでチャネルの生成と読み書きができる状態である
	 ch1 = make(chan int)
	 fmt.Println(ch1)

	 //この定義でもOK,上記を満たす
	 ch2 := make(chan int)
	 fmt.Println(ch2)

	 //バッファーサイズ(要領)を指定して作成することも可能
	 chan3 := make(chan int, 5)

	 //チャネルに送信
	 chan3 <- 1
	 chan3 <- 2
	 chan3 <- 3
	 fmt.Println(len(chan3), chan3)

	 i := <-chan3
	 fmt.Println(i)
	 i2 := <-chan3
	fmt.Println(i2)
	i3 := <-chan3
	fmt.Println(i3)

	//channelとgo routine
	chan4 := make(chan int)
	chan5 := make(chan int)

	go receiver(chan4)
	go receiver(chan5)

	//チャネルにデータを送る
	//t := 0
	// for t < 100 {
	// 	chan4 <- t
	// 	chan5 <- t
	// 	time.Sleep(50 * time.Millisecond)
	// 	t ++
	// }

	chan6 := make(chan int, 2)
	chan7 := make(chan string, 2)

	chan7 <- "A"
	select {
	case v1 := <-chan6:
		fmt.Println(v1 + 1000)
	case v2 := <-chan7:
		fmt.Println(v2 + "!?")
	}
}
