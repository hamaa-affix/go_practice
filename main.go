package main

import (
	"fmt"
	"time"
)

func sub() {
	for {
		fmt.Println("sub loop")
		time.Sleep(100 * time.Microsecond)
	}
}

func main() {
	go sub()

	for {
		fmt.Println("main loop")
		time.Sleep(200 * time.Microsecond)
	}
}
