package main

import "fmt"

func main() {

	result, myMessage := proAdder(1, 2, 3, 4, 5, 7, 7, 8, 9, 0, 1, 2, 345, 6778, 1233, 534634756)
	fmt.Println(result, myMessage)
}

func proAdder(values ...int) (int, string) {

	total := 0
	for _, val := range values {
		total += val
	}
	return total, "Hello sandeep dont worry"
}
