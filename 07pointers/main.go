package main

import "fmt"

// * -> creating a pointer
// & -> getting a pointer for variable which has already been defined
// pass by value and pass by refrence 
// if u want to pass any value by refrence use &
// if u want to access that value inside that pointer use  *
// * to access hwta inside te poimnter
// var num *int  its says that num is pointer which will hold a data of int type
func main(){
	// pass by value
	num :=0
	// fmt.Println(num)
	// incre(num)
	// fmt.Println(num)
		// pass bu refrence 
	increp(&num) // send pointer of num varibale
	fmt.Println(num)
}
func incre(num int){
	num ++
	fmt.Println(num)

}
func increp(num *int){  // receive pointer
	*num  = *num +10;
	fmt.Println("access to numj pointer",num)
	fmt.Println("acces to num",*num)

}

