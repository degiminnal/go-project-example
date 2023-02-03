package main

import "fmt"

type user struct {
	name     string
	password string
}

func main() {
	a := user{name: "wang", password: "1024"}
	b := user{"wang", "1024"}
	c := user{name: "wang"}
	c.password = "1024"
	var d user
	d.name = "wang"
	d.password = "1024"

	fmt.Println(a, b, c, d)                 // {wang 1024} {wang 1024} {wang 1024} {wang 1024}
	fmt.Println(checkPassword(a, "haha"))   // false
	fmt.Println(checkPassword2(&a, "haha")) // false
	fmt.Println(d.checkPassword("1024"))
}

func checkPassword(u user, password string) bool {
	return u.password == password
}

// 类成员函数
func (u user) checkPassword(password string) bool {
	return u.password == password
}

func checkPassword2(u *user, password string) bool {
	return u.password == password
}
