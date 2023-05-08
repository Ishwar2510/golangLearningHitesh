package main

import "fmt"

func main() {
	days := []string{"monday", "tues", "wed", "thurst", "friday"}
	fmt.Println(days)
	// to iterate it 
	for i:=0; i<len(days);i++{
		fmt.Println(days[i]);
	}
	 //or 
	for data := range days{
		fmt.Println(data)
	}// or with index and value
	for index,value  := range days{
		fmt.Println( index,value)
	}
	
}