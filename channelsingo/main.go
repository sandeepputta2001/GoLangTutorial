package main

import (
	"fmt"
	"sync"
)

func main() {

	myChan := make(chan int, 5) // buffer channel : it takes multiple values at the same time
	wg := &sync.WaitGroup{}

	wg.Add(3)

	// receive only channel

	go func(wg *sync.WaitGroup, ch <-chan int) {
		// close(myChan)  : we are not allowed to close the channel in receiver channel

		val, isChannelOpen := <-myChan // this code is to check whether the channel is open or closed

		fmt.Println(val)
		fmt.Println(isChannelOpen)

		fmt.Println(myChan)

		wg.Done()

	}(wg, myChan)

	// send only channel

	go func(wg *sync.WaitGroup, ch chan<- int) {

		myChan <- 5
		myChan <- 234

		close(myChan)

		wg.Done()

	}(wg, myChan)

	go func(wg *sync.WaitGroup, ch chan int) {

		wg.Done()

	}(wg, myChan)

	wg.Wait()
}
