package main

import (
	"fmt"
	//"time"
)

//interface
type stringfy interface {
	ToString() string
}

type Person struct {
	Name string
	Age int
}

func (p *Person) ToString()string {
	return fmt.Sprintf("Name=%v, Age=%v", p.Name, p.Age)
}

type Car struct {
	Number string
	Model string
}

func (c *Car) ToString() string {
	return fmt.Sprintf("Number=%v Model=%v", c.Number, c.Model)
}

//カスタムエラー
type MyError struct {
	Message string
	ErrorCode int
}

func (e *MyError) Error() string {
	return e.Message
}

func RaiseError() error {
	return &MyError{
		Message: "カスタムエラーが発生しました",
		ErrorCode: 1234 ,
	}
}
func main() {
	//interfaceを定義することでスライスに異なる型を定義できる
	vm := []stringfy {
		&Person{
			Name: "mami",
			Age: 32,
		},
		&Car{
			Number: "123-111",
			Model: "toyota",
		},
	}


	for _, v := range vm {
		fmt.Println(v)
	}

	err := RaiseError()
	fmt.Println(err.Error())
}
