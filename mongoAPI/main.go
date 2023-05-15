package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/sandeeputta2001/mongoAPI/router"
)

func main() {
	fmt.Println("mongdb api")
	fmt.Println("server is getting started ")
	r:=router.Router()
	log.Fatal(http.ListenAndServe(":4000",r))
	fmt.Println("Listening at port 4000")
}
