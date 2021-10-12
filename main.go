package main

import (
	"fmt"
	//"time"
)

type User struct {
	Name string
	Age  int
}

type Job struct {
	Name string
}

type T struct {
	User
	Job
}

	//constructer
func NewUser(name string, age int) *User  {
	return &User{
		Name: name,
		Age: age,
	}
}

//スライスを定義し代入する
type Users []*User

func intUser1(user User) {
	user.Name = "user1です"
	user.Age = 25
}

func intUser(user *User) {
	user.Name = "userですよ"
	user.Age = 20
}

func (u User) sayName() {
	fmt.Println(u.Name)
}

//setter 参照渡しになる
func (u  *User) setName(name string) {
	u.Name = name
}

func main() {
	var user1 User
	fmt.Println(user1)
	user1.Name = "user1"
	user1.Age = 30
	fmt.Println(user1)

	//暗黙的な宣言
	user2 := User{}
	user2.Name = "user2"
	user2.Age = 30
	fmt.Println(user2)

	user3 := User{Name: "user3", Age: 30}
	fmt.Println(user3)

	user4 := new(User)
	fmt.Println(user4)

	//値渡しではデータの更新はされない
	intUser1(user3)
	fmt.Println(user3)

	//参照渡しではデータは更新される
	intUser(&user3)
	fmt.Println(user3)

	//レシバーの定義をする
	user5 := User{Name: "user5"}
	user5.sayName()

	//setterを利用する
	user5.setName("chengName")
	user5.sayName()

	//構造体の埋め込み
	user6 := T{User: User{Name: "user6", Age: 30}}
	fmt.Println(user6.User.Name)

	user7 := T{
		User: User{
			Name: "user7",
			Age: 30,
		},
		Job: Job{
			Name: "Ns",
		},
	}
	fmt.Println(user7.Job.Name)

	//constructer demo
	user8 := NewUser("user8", 10)
	fmt.Println(user8)
	fmt.Println(user8.Name)

	user9 := User{"a", 1}
	user10 := User{"a", 1}
	user11 := User{"a", 1}

	//sliceへ格納する
	users := Users{}
	//ポインターの値を代入しないといけない
	users = append(users, &user9, &user10, &user11)
	fmt.Println(users)
	//forで回すときは配列、スライス、マップでないとだめ
	for _, u := range users {
		fmt.Println(*u)
	}

	//mapで定義する
	m := map[int]User{
		1: {Name: "b", Age: 2},
		2: {Name: "c", Age: 3},
	}

	for _, i := range m {
		fmt.Println(i)
	}
	//mapのkeyにUser型を定義することも可能
	m2 := map[User]string {
		{Name: "d", Age: 10}: "tokyo",
		{Name: "e", Age: 12}: "USA",
	}

	for i, m := range m2 {
		if i.Name == "d" {
			continue
		}
		fmt.Println(m)
	}
}
