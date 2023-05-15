package main

import (
	"fmt"
	"io"
	"net/http"
	"strings"
)

func main() {
	fmt.Println("welcome to tut of post request in go ")
	PerformPostRequestInJson()
}

func PerformPostRequestInJson() {

	const myUrl = "http://localhost:5000/post"

	requestBody := strings.NewReader(`{
		"name":"sandeep putta",
		"age":21,
		"isMarried":false
	}
	`)

	response, err := http.Post(myUrl, "application/json", requestBody)
	if err != nil {
		panic(err)
	}
	defer response.Body.Close()
	content, _ := io.ReadAll(response.Body)
	fmt.Println("the post data is :", string(content))
}
