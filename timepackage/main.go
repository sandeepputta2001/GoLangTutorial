package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("welcome to the time tutorial in golang ")

	presentTime := time.Now()
	fmt.Println(presentTime)

	fmt.Println(presentTime.Format("01-02-2006  15:04:05 Monday "))

	createDate := time.Date(1999, time.December, 12, 12, 38, 60, 12, time.UTC)
	fmt.Println(createDate)
	fmt.Println(createDate.Format("01-02-2006 Monday 15:04:05"))
}
