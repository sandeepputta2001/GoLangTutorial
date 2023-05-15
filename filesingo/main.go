package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

func main() {

	content := "Hello sandeep "
	file, err := os.Create("./filesingo.txt")
	checkNilError(err)
	length, err := io.WriteString(file, content)
	checkNilError(err)
	fmt.Println("The length of the file is ", length)
	defer file.Close()
	readFile("./filesingo.txt")

}

func readFile(filename string) {

	dataByte, err := ioutil.ReadFile(filename)

	checkNilError(err)
	fmt.Println("The text inside the file is ", string(dataByte))
}

func checkNilError(err error) {
	if err != nil {
		panic(err)
	}
}
