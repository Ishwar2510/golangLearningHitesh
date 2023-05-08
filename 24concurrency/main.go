package main

import (
	"fmt"
	"net/http"
	"sync"
)

// concurrency and parallelism
// in go everythig is Type remember thius thing

// explanation
// suppose u r eating food and at that time a notification pops in and u r feelinbg hot and u want to turn on ac
// with concurrency it means u r allowed to perform multiple task but not all at smae time
// u r eating  --> u left eating--> u checkcekd notification--> droped the phone --> again start eating
// u r doing all the wroks but not all the works at same time

// where as in parallelism
// you do all the thing simultaniously

// go routines
// it is the way in which u achive parallelism

// do not communicate by sharing memory instead share memory by communicating


// race conditions solve by using mutex

// channels 
// one simple go routine one go routine donot have any idea of another goroutine
// with the help of channels gorotine can communicate with eachother
//

var wg sync.WaitGroup
var mu sync.Mutex // pointer    var mu &sync.Mutex{}   // * used while declaration and &used while send the data or receiveinf the data
var signal = []string {"test"}
func main() {
	// go greeter("hello") // fire up  a thread  but we rae not waiting  to handle this we have deictaed pkg sync
	// greeter("world")
	
	websiteLsit := []string{
		"https://lco.dev",
		"https://google.com",
		"https://github.com",
	}
	for _, value := range websiteLsit {
		// getstatusCode(value)   // this will be  the blocking code  // we have to convert it into async await by creating multiple threads
		go getstatusCode(value) // we will use sync wait group   // go  makes this call as async call
		wg.Add(1)

	}
	fmt.Println("print 1")
	fmt.Println("print 2")
	fmt.Println("print 3")
	fmt.Println("print 4")
	wg.Wait()  // resposible for not allowing main method to finish untill all are done
	fmt.Println(signal)

}
func greeter(s string) {
	for i := 0; i < 3; i++ {
		// time.Sleep(3 * time.Second)
		fmt.Println(s)
	}
}

func getstatusCode(endpoint string) {
	defer wg.Done()
	res, err := http.Get(endpoint)
	if err != nil {
		fmt.Println(err)
	}
	// since the threads are managed by goroutines in go and not my os threads
	// we need to lock the shared memory, mutex comes into picture 
	mu.Lock()
	signal = append(signal,fmt.Sprintf("%v-->%v",endpoint,res.Status)) // differnt threads are working on the same variable which will cuase an error  u cananot write multiple data at the same time // this is whwer mutex comes in
	mu.Unlock()
	fmt.Println(res.StatusCode, res.Status)
	
}
