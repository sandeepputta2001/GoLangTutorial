package main

import "fmt"

func main() {

	variable := make(map[string]string, 150)
	fmt.Println(variable)
	variable["name"] = "sandeep"
	variable["rollno"] = "160819"
	fmt.Println(variable)
	for i := 0; i < 10; i++ {

		for key, value := range variable {

			fmt.Printf("The values of keys and values are :\n")
			fmt.Printf("%v   %v\n", key, value)

		}
	}

	delete(variable, "name")

	fmt.Println(variable)
}
