package main

import (
	"encoding/json"
	"fmt"
)

// sendind json data
func main() {
	// // creating get request in go lang

	// response, err := http.Get("http://localhost:3000/kuchbhi")

	// if err != nil {
	// 	fmt.Println(err)
	// }
	// // to read data in request body we need to use ioutil
	// data, err := ioutil.ReadAll(response.Body)
	// fmt.Println("the dats is byte format ", string(data))

	// // creating a post request
	// // we need to send a json data or formencoded url
	// // sendin jsondata we need to
	// // creting a json data
	// jsondata := strings.NewReader(`{
	// 	"name":"ishewr",
	// 	"age" :"20"
	// }`)
	// responseb, err := http.Post("http://localhost:3000/kuchbhi", "application/json", jsondata)
	// if err != nil {
	// 	fmt.Println(err)
	// }

	// defer responseb.Body.Close()
	// content, _ := ioutil.ReadAll(responseb.Body)
	// fmt.Println(string(content))

	// url encoded form sending post data

	// // form data
	// mydata := url.Values{}
	// mydata.Add("name","ishear")
	// mydata.Add("age","20")

	// responsec,_ := http.PostForm("http://localhost:3000/kuchbhi",mydata)

	// defer responsec.Body.Close()
	// contentn,_ := ioutil.ReadAll(responsec.Body)
	// fmt.Println(string(contentn))

	// creating json data in go lang
	// in js  we do json.stringify(data)

	loadData := []course{
		{"java", 200, 5},
		{"js", 201, 6},
	}
	// loadData is normal object
	fmt.Println(loadData)
	// how to stringify this objet
	jsonData, err := json.Marshal(loadData)
	if err != nil {
		panic(err)
	}

	fmt.Println((jsonData))


	// parsing the json data consuming json data
	

}

type course struct {
	name   string `json:name`
	year   int	`json:year`
	rating int `json:year`
}
