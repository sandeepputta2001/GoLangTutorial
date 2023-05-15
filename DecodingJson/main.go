package main

import (
	"encoding/json"
	"fmt"
)

type course struct {
	Name     string `json:"coursename" `
	Price    int
	Platform string   `json:"plat"`
	Password string   `json:"-"`
	Tags     []string `json:"tags,omitempty"`
}

func main() {

	fmt.Println("welcome to the tutorial of decoding the json data ")
	DecodeJson()
}

func DecodeJson() {

	JsonDataFromWeb := []byte(`{
	
			"Name": "golang",
			"Price": 1000,
			"Platform": "sandeep.putta",
			"Password": "123asg",
			"Tags": [
					"web",
					"js"
			]
	}
	`)

	// var lcocourse course
	// check := json.Valid(JsonDataFromWeb)
	// if check {
	// 	fmt.Println("JSON was valid ")
	// 	response := json.Unmarshal(JsonDataFromWeb, &lcocourse)
	// 	fmt.Printf("%v#\n", response)
	// }

	var myOnlineData map[string]interface{}
	json.Unmarshal(JsonDataFromWeb, &myOnlineData)
	fmt.Printf("%v\n", myOnlineData)

	for k, v := range myOnlineData {
		fmt.Printf("The Key  is %v and value is %v and type is %T\n", k, v, v)
	}
}
