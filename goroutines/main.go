package main

import (
	"fmt"
	"log"
	"net/http"
	"sync"
)

var wg sync.WaitGroup // variable is declared to for weight group to carry waitgroup methods and waitgroup is of pointer type

var signals = []string{"test"}

var mut sync.Mutex // mutex is of type pointer since the mutex variable is passed into multiple go routines we have to use pointer for exact reference
func main() {

	// fmt.Println("welcome to the goroutines in the golang")
	// go greeter("is used")
	// time.Sleep(3 * time.Millisecond)
	// greeter("hello go")

	websiteList := []string{
		"https://lco.dev",
		"https://google.com",
		"https://fb.com",
		"https://youtube.com",
	}

	for _, web := range websiteList {
		go getStatusCode(web)
		wg.Add(1)

		fmt.Println(signals)

	}

	wg.Wait()
}

// func greeter(s string) {
// 	for i := 0; i < 3; i++ {
// 		fmt.Println("sync for goroutines ", s)

// 	}
// }

func getStatusCode(endpoint string) {

	defer wg.Done()

	res, err := http.Get(endpoint)
	mut.Lock()

	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%d is the status code of %s\n", res.StatusCode, endpoint)
	signals = append(signals, endpoint)
	mut.Unlock()
}
