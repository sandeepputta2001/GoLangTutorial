package main

import "fmt"

func main() {

	Sandeep := User{"Sandeep", "puttaSandeep4@gmail.com", true, 21}

	Sandeep.GetStatus()

	fmt.Println(Sandeep)
	fmt.Printf("Sandeep's details are : %+v\n", Sandeep)
	fmt.Printf("The name of the user is %v and email is %v , the status is %v  and the age of the user is %v \n ", Sandeep.Name, Sandeep.Email, Sandeep.Status, Sandeep.Age)
}

type User struct {
	Name   string
	Email  string
	Status bool
	Age    int
}

func (u User) GetStatus() {
	fmt.Println("the status of user is: ", u.Status)
}

func (u User) GetEmail() {
	u.Email = "SANDEEP1@gmail.com"
	fmt.Println("the new email of the user is ", u.Email)
}
