package main

import (
	"fmt"
	"net/url"
)

const urlb = "https://lco.dev"

func main() {
	// response, err := http.Get(urlb)
	// defer response.Body.Close() // user responsibility to close the connection

	// if err != nil{
	// 	panic(err)
	// }
	// fmt.Printf("the response type is %T\n",response)
	// fmt.Printf("the response  is",response.Body)
	// databytes,err := ioutil.ReadAll(response.Body)
	// fmt.Println(databytes)
	// fmt.Printf("the read dat is of type : %T\n",databytes)
	// fmt.Println(string(databytes))
	// fmt.Print(response.Status)
	// fmt.Print(response.StatusCode)
	fet()

}

const myurl = "https://lco.dev/learn?course=react.js&id=20"

func fet() {
	fmt.Println(myurl)

	// parsing url

	parsedurl, err := url.Parse(myurl)
	if err != nil{
		fmt.Println(err)
	}
	fmt.Printf("%T\n",parsedurl)
	fmt.Println(parsedurl.RawQuery)
	fmt.Println(parsedurl.Path)

	// extractig query params

	queryparams := parsedurl.Query()
	fmt.Println(queryparams)

	couseaname := queryparams["course"]
	id := queryparams["id"]
	fmt.Println(couseaname,id);

}
