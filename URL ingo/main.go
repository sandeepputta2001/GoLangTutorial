package main

import (
	"net/url"
)

const myUrl = "https://local.dev:3000/learn?cousername=reactjs&paymentid=gbh 1234"

func main() {

	result, _ := url.Parse(myUrl) //url package is used to handle url related and it is in net library
	// fmt.Println(result.Scheme)
	// fmt.Println(result.Host)
	// fmt.Println(result.Port())
	// fmt.Println(result.Path)
	// fmt.Println(result.RawQuery)

	// qparams := result.Query()
	// //

	// for key, val := range qparams {
	// 	fmt.Println(key, val)
	// }

	partsOfUrl := &url.URL{
		Scheme: "https",
	}
}
