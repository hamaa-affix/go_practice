package main

import (
	"fmt"
	"log"

	"github.com/hamaa-affix/go_practice/config"
	// _ "github.com/go-sql-driver/mysql"
)

type User struct {
	id int
	name string
}

func main() {
	fmt.Println(config.Config.Port)
	fmt.Println(config.Config.SQLDriver)
	fmt.Println(config.Config.DbName)
	fmt.Println(config.Config.LogFile)

	log.Println("test")
}
