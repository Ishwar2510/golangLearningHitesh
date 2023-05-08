package main

import (
	"fmt"
	
)

func main() {
	fmt.Println("hello bhai soethinf fishy")
	var arrayList = []string{}  // with slice dontot metion anything in the sqaure bracketr and also init it using {}
	arrayList = append(arrayList,"hello","yar","kaise","ho","bhai","log")  // append returns a new arraylist it doesnt do opeartion on same arraylist
	fmt.Println(arrayList)
	// fmt.Println(arrayList[1:3],arrayList[:2],arrayList[2:])
	
	// dleteing from an index
	index:=2
	arrayList = append(arrayList[:index],arrayList[index+1:]...)
	fmt.Println(arrayList)
	

	

	


	
}