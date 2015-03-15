package main

import (
	"fmt"
	st "github.com/wackonline/structrecord"
)

type User struct {
	id    int `relat:"OneToOne" target:"Profile"`
	name  string
	age   int
	phone string
}

type Profile struct {
	uid     int `relat:"OneToOne" target:"User"`
	address string
	city    string
}

func main() {
	var u User
	u.age = 12
	u.id = 1
	u.name = "joe"
	u.phone = "110"
	st.Run().Save(u)

	fmt.Println("hello,save!!!")
}
