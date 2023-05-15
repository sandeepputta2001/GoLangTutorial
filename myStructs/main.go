package main

import "fmt"

func main() {

	type User struct {
		Name   string
		Email  string
		Status bool
		Age    int
	}

	sandeep := User{"sandeep", "puttasandeep4@gmail.com", true, 21}

	fmt.Println(sandeep)
	fmt.Printf("sandeep's details are : %+v\n", sandeep)
	fmt.Printf("The name of the user is %v and email is %v , the status is %v  and the age of the user is %v \n ", sandeep.Name, sandeep.Email, sandeep.Status, sandeep.Age)
}
