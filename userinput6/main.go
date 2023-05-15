package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	welcome := "welcome to user input "
	fmt.Println(welcome)

	fmt.Println("enter the rating for our pizza")
	reader := bufio.NewReader(os.Stdin) //libraries where functions are created
	input, _ := reader.ReadString('\n') //comma ok or comma error syntax
	fmt.Println("Thanks for rating ", input)
	fmt.Printf("The type of the input is %T", input)
}
