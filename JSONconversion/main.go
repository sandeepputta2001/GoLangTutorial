package main

import (
	"encoding/json"
	"fmt"
)

func main() {

	fmt.Println("welcome to the tut of conversion of json data from any kind of data to ideal json data ")
	EncodeJson()

}

type course struct {
	Name     string `json:"coursename" `
	Price    int
	Platform string   `json:"plat"`
	Password string   `json:"-"`
	Tags     []string `json:"tags,omitempty"`
}

func EncodeJson() {

	lcoCourses := []course{
		{"golang", 1000, "sandeep.putta", "123asg", []string{"web", "js"}},
		{"java ", 1000, "sandeep", "123qwe", []string{"rust ", "firebase "}},
		{"scala ", 234, "putta .hemanth ", "123asg", nil},
	}

	// finalJson, err := json.Marshal(lcoCourses)
	finalJson, err := json.MarshalIndent(lcoCourses, "", "\t")
	if err != nil {
		panic(err)
	}

	fmt.Println(string(finalJson))

}
