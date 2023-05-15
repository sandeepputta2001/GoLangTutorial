package main

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
)

func main() {

	fmt.Println("welcome to the tut of send data to the postform directory ")
	PerformRequestOnPostForm()
}

func PerformRequestOnPostForm() {
	const myUrl = "http://localhost:5000"

	data := url.Values{}
	data.Add("firstName", "sandeep")
	data.Add("lastName", "putta ")
	data.Add("email", "puttasandeep4@gmail.com")

	response, err := http.PostForm(myUrl, data)
	if err != nil {
		panic(err)
	}
	defer response.Body.Close()
	content, err := io.ReadAll(response.Body)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(content))

}
