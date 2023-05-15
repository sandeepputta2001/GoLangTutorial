package main

import "fmt"

const UserName string = "sandeep putta " // variables or methods declared with first letter as capital letter , it means that the method or variable can be accessed can where inside the program

func main() {
	fmt.Println("Hello world")
	fmt.Println("hello sandeep ")
	var isLogged bool = true
	fmt.Printf("The type of the variable is :%T \n", isLogged)
	fmt.Println(isLogged)

	//default values and aliasses

	var anotherVariable string
	fmt.Println(anotherVariable)
	fmt.Printf("The type of the variable is %T \n", anotherVariable)
	fmt.Println(UserName)
}
