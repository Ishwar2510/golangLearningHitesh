package main

import "fmt"

// methods are funtions
// but when we brig in classes i.e  struct in go lang these are said to be methods

type User struct{
	name string
	age int
}
func (u User) getName() string{
	return u.name
}
func (u User) getAge() int {
	return u. age
}
func (u User) incrAge(num int ){
	u.age = u.age+num
}
func main(){
	var user1 User = User{"Ishear",20}
	var user2 User = User{"Navin",10}
	fmt.Println("user 1-->", user1)
	fmt.Println("user 2-->", user2)
	user1.incrAge(20)      // will it chnage the main user1 object ,the ans is no . because in go everything is pass by value;you have to convert it into pass by refrence 
	fmt.Println("user 1-->", user1)




}