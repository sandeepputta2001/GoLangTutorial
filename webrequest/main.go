package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

const url string = "https://lco.dev" //url is declared globally to get request and responses from other urls

func main() {
	fmt.Println("LCO web request ")

	response, err := http.Get(url) //http module is used for web requests

	if err != nil {
		panic(err)
	}

	fmt.Printf("The datatype of the response is %T", response)
	defer response.Body.Close() // caller's responsibility to close the connection

	dataBytes, err := ioutil.ReadAll(response.Body)
	if err != nil {
		panic(err)
	}
	content := string(dataBytes) ///(imp)  read data should be compulsorily converted to string since read data always get returned in databytes format
	fmt.Println(content)

}
