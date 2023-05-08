package main

import "fmt"

// in majority we dont use array we use slices (arrayList ho hota hai)
func main(){
	// int array [] = new int [5]
	// declare first then intialize later
	var array [5] int 
	var straray[5]string
	array[0] = 1
	straray[1] = "hello"
	fmt.Println(array,straray,len(array))

	// declare and intilaize at same time

	var arr  = [5]int{1,2,3}
	fmt.Println(arr)
	aar:=[6]string{"hello","how"}
	fmt.Println(aar)


}