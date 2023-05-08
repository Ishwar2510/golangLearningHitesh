package main

import "fmt"

func main() {
	var number int = 20  // defualt is zero
	fmt.Println(number)
	fmt.Printf("The number is : %d of type %T",number,number)
	fmt.Println()

	numbers := 30
	fmt.Println("This declration can be done inside function only")
	fmt.Println(numbers)
	fmt.Printf("The number is : %d of type %T",numbers,numbers)

	var name string ="ishwar" // default is  "" works same as java
	var myname string;
	fmt.Println(name)
	fmt.Println(myname)

	var money float64 = 52.456666
	fmt.Println(money)

	var isLogged bool = true
	var isLogg bool     // default value is false
	fmt.Print(isLogged)
	fmt.Println(isLogg)

	

}