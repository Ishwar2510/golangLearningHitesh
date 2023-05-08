package main

import "fmt"

func main() {
	num := 10
	if num > 10 {
		fmt.Println("num er grt than 10")
	}else{
		fmt.Println("less than 10")
	} 
	// there is one more synatax which u will see a lot in web api

	// intilaize and check in same line
	// first intialiozed then ; then check // mostly 
	if number:=20; number>20{
		fmt.Println("grts than 20")
	}else{
		fmt.Println("lesser than 20")
	}

	// real life usecase for above syntax

	if isAl := isalive(); isAl {
		fmt.Println("it is True")
	}else{
		fmt.Println("it is false")
	}
}
func isalive() bool{
	return false
}
