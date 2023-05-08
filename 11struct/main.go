package main

import "fmt"

// structs are same as classes
// no inheritance  in golang
// no patrent no super
type User struct {
	name    string
	age     int
	email   string
	isAlive bool
}

func NewUser(name string, age int, email string, isAlive bool) *User {
	return &User{name,age,email,isAlive}
}

func main() {
	var user1 *User = NewUser("ishwar",20,"@mai",true)
	fmt.Println(*user1)
	fmt.Println("hello")
	

}
