package main

import (
	"fmt"
	"sync"
)

func main() {

	fmt.Println("race condition in golang ")
	wg := &sync.WaitGroup{}
	mut := &sync.Mutex{}
	mutrw := &sync.RWMutex{}
	var score = []int{0}

	wg.Add(3) // single add wait group can be written at the top for all the goroutines existing in the code

	go func(wg *sync.WaitGroup, m *sync.Mutex) {

		fmt.Println("FIRST goROUTINE")
		mut.Lock()
		score = append(score, 1)
		mut.Unlock()

		wg.Done()

	}(wg, mut)
	go func(wg *sync.WaitGroup, m *sync.Mutex) {

		fmt.Println("second  goROUTINE")

		mut.Lock()
		score = append(score, 2)
		mut.Unlock()

		wg.Done()

	}(wg, mut)

	go func(wg *sync.WaitGroup, m *sync.Mutex) {

		fmt.Println("third  goROUTINE")
		mut.Lock()
		score = append(score, 3)

		mut.Unlock()

		wg.Done()

	}(wg, mut)

	wg.Wait()

	fmt.Println(score)

	// for RWmutex the following is the process
	//in RWmutex the write operation process is same to the mutex and to use the RLock() and RUnlock() we have to use these read methods only when the resource (resource = variable ) is read

	go func(wg *sync.WaitGroup, m *sync.RWMutex) {

		mutrw.RLock()
		fmt.Println(score)
		mutrw.RUnlock() // the two methods are used here since the resource value is read here. There is no point of using read mutex methods with the resource

		wg.Done()

	}(wg, mutrw)

}
