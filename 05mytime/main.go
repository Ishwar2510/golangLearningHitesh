package main

import (
	"fmt"
	"time"
)

func main() {
	presentTime := time.Now()
	fmt.Println(presentTime)
	formatedTime := presentTime.Format("01-02-2006")
	fmt.Print(formatedTime)
	
}
// type go env  in cli 
// ctrl + l shortcut to lear cli 
