package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("pls rate the app betwen 1 and 5");
	reader := bufio.NewReader(os.Stdin)
	input,_:= reader.ReadString('\n')
	// numRating,err := strconv.ParseFloat(input,10,64)
	numRating,err := strconv.ParseInt(strings.TrimSpace(input),10,64)
	if err !=nil{
		fmt.Println("erro is --->>:",err)
		fmt.Println(input)
	}else{
		numRating++
		fmt.Println(numRating)
	}

}