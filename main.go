package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

type User struct {
	id int
	name string
}

func main() {
	// データベース接続
    db, _ := sql.Open("mysql", "root:root@tcp(mysql:3306)/go_database")

    // deferで処理終了前に必ず接続をクローズする
    defer db.Close()

    // 接続確認
	// cmd := "INSERT INTO users(name) VALUES(?)"
	// _, err := db.Exec(cmd, "kotoha")

    //特定のuserの取得
	// cmd := "select * from users where id = ?"
	// //レコードの取得
	// row := db.QueryRow(cmd, 3)
	// var user User
	// //Scanで構造体にデータ挿入する。
	// err := row.Scan(&user.id, &user.name)

	//一覧を取得
	cmd := "select * from users"
	rows, _ := db.Query(cmd)
	defer rows.Close()
	var users []User

	for rows.Next() {
		var user User
		err := rows.Scan(&user.id, &user.name)
		if err != nil {
			log.Println(err)
		}
		users = append(users, user)
	}

	for _, user := range users {
		fmt.Println(user.id, user.name)
	}
}
