package main

import "fmt"

func main() {

	fmt.Println("welcome to the for loops concept in go ")

	// putta := []string{"sandeep ", "putta ", "pawan kalyan", "ustaad "}

	// for i := 0; i < len(putta); i++ {
	// 	fmt.Println(putta[i])

	// }

	// for i := range putta {
	// 	fmt.Println(putta[i])
	// }

	// for index, value := range putta {
	// 	fmt.Printf("the index and value of the is %v an %v\n ", index, value)
	// }
	// i := 0

	// for i < 10 {
	// 	// if i == 5 {
	// 	// 	break
	// 	}
	// 	if i == 7 {
	// 		continue
	// 	}
	// 	fmt.Println(i)
	// 	i++
	// }

p:
	fmt.Println("hello sandeep ")

	goto p

}
