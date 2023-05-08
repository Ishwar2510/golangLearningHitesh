package main

import "fmt"

// defer key word makes execution at the end
// whenver a defer keyword is encountered it is pushed into the stack
// ust before the completion of the function the stack is empitied
// i.e defer follows th LIFO
func main(){

	defer fmt.Print("hello")
	defer fmt.Print(" World") 
	// will print world hello
	prinLoop()
}
func prinLoop(){
	for i:=0 ;i<5;i++{
		defer fmt.Println(i)
	}
}