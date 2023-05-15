package main

import (
	"fmt"
	"math/big"
	"math/rand"
	"time"
)

func main() {

	// random number in crypto

	// myRandomNumber := rand.Intn(5)

	rand.Seed(time.Now().Unix())

	myRandomNumber := rand.Intn(5) + 1
	fmt.Println(myRandomNumber)

	// fmt.Println(myRandomNumber)

	// random number from crypto

	myRandNum := rand.Int(rand.Reader, big.NewInt(7))

	fmt.Println(myRandNum)
}
