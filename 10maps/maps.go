package main

import "fmt"

func main() {
	var mymap = make(map[string]string)
	mymap["first"] = "A"
	mymap["last"] = "ishwar"
	mymap["number"] = "1234555"
	fmt.Println(mymap)
	delete(mymap,"number")
	fmt.Println(mymap)
	// to iterate over my map like for each loop 
	for key, value := range mymap{
		fmt.Printf("the key is %v and value is %v\n",key,value)
	}




}