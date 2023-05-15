package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

func main() {
	fmt.Println("welcome to how to get request in golang")
	PerformGetRequest()
}

func PerformGetRequest() {
	const myUrl = "http://localhost:5000/get"

	response, err := http.Get(myUrl)

	myError(err)

	defer response.Body.Close()
	fmt.Println(response.StatusCode)
	fmt.Println(response.ContentLength)
	fmt.Println(response.Header)
	fmt.Println(response.Request.URL)
	fmt.Println(response.Proto)
	fmt.Println(response.Cookies())
	fmt.Println(response.TransferEncoding)
	fmt.Println(response.TLS)

	content, err := ioutil.ReadAll(response.Body)
	myError(err)

	var responseString strings.Builder

	byteCount, _ := responseString.Write(content)
	fmt.Println("byteCount is :", byteCount)

	fmt.Println("the data inside the file is :", responseString.String())
	// fmt.Println(string(content))
}

func myError(err error) {
	if err != nil {
		panic(err)
	}
}
