package main

import (
	"bufio"
	"fmt"
	"os"
)
 // value comma err syntax is similar to try catch 
 // value, err ===  try {} catch{}
func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("enter a number")
	data, _ := reader.ReadString('\n')
	fmt.Printf("tHe data is %v",data)

}