package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	fmt.Println("welcome to the modules concept in golang ")
	Greeter()
	r := mux.NewRouter()
	r.HandleFunc("/", ServeHome).Methods("GET")
	log.Fatal(http.ListenAndServe(":8000", r)) //if an error shows fatal func logs the error // http.ListenAndServe creates server using golang
}

func Greeter() {
	fmt.Println("hello mod users ")

}

func ServeHome(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("<h1>welcome to golang series on YT</h1>"))

}
