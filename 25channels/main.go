package main

import (
	"fmt"
	"sync"
)

// channels
// in go routine onme go function i.e async function doesnot have any info of the other
// go function .
// with channels ine goroutine can actually communicate eoth other goroutine


func main(){
	mych := make(chan int )
	// mych <- 5
	// fmt.Println(<-mych)
	wg := &sync.WaitGroup{}
	wg.Add(2)
	go func(ch chan int,wg *sync.WaitGroup){
		
		wg.Done()
	}(mych,wg)
	go func(ch chan int,wg *sync.WaitGroup){
		mych<-5
		wg.Done()
	}(mych,wg)
	wg.Wait()
	
}