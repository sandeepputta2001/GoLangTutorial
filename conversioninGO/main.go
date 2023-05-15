package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {

	welcome := "welcome to our rating system "
	fmt.Println(welcome)

	reader := bufio.NewReader(os.Stdin) // method to create the new reader variable

	fmt.Println("enter your rating ")

	input, _ := reader.ReadString('\n') //method to read the variable
	fmt.Println("Thanks for rating ", input)

	numRating, err := strconv.ParseFloat(strings.TrimSpace(input), 64)

	if err != nil {
		fmt.Println(err)

	} else {
		fmt.Println("Added 1 to your string ", numRating+1)
	}
}
