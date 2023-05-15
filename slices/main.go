package main

import (
	"fmt"
	"sort"
)

func main() {

	mySlice := make([]string, 10)

	fmt.Println(len(mySlice))

	mySlice = append(mySlice, "", "", "sandep", "putta ", "goud ", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "")

	fmt.Println(mySlice, len(mySlice))
	sort.Strings(mySlice)
	myBoolean := sort.StringsAreSorted(mySlice)
	fmt.Println(myBoolean)

}
