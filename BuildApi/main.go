package main

import (
	"crypto/rand"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/gorilla/mux"
)

// Model for Course -file
type Course struct {
	CourseId   string  `json:"Courseid"`
	CourseName string  `json:"Coursename"`
	Author     *Author `json:"author"` //* author means that it is of type pointer

}

type Author struct {
	FullName string `json:"fullname"`
	Website  string `json:"website"`
}

// fake DB

var Courses []Course //fake data created

// middleware or helper methods  //helper methods are perform special tasks like encrypting passwords or not allowing some users to pull in the database

func (C *Course) IsEmpty() bool {
	return C.CourseId == "" && C.CourseName == ""
}

// func () IsEmpty( C *Course) bool {
// 	return C.CourseId == "" && C.CourseName == ""
// }

func main() {

}

// controllers -file

// serve home route

func ServeHome(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("<h1>welcome to API by sandeep putta </h1>"))

}

func GetAllCourse(w http.ResponseWriter, R *http.Request) {
	fmt.Println("get all Course ")
	w.Header().Set("content-type", "application/json")
	json.NewEncoder(w).Encode(Courses)
}

func getOneCourse(w http.ResponseWriter, r *http.Request) string {

	fmt.Println("get one Course ")
	w.Header().Set("content-type", "application/json")

	// grab id from request
	params := mux.Vars(r) // mux package mapping requests and responses ; vars fucntion extracts all the variables in the request(r)
	fmt.Printf("the type of params is :%t\n", params)

	// loop through the Course , find the matching id and return the response

	for _, course := range Courses {
		if course.CourseId == params["id"] {
			json.NewEncoder(w).Encode(course)
			return
		}
	}
	json.NewEncoder(w).Encode("No Course found with given id ")
	return
}

// create a file to send in to database
func CreateOneCourse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("create one Course to send into database ")
	w.Header().Set("content-type ", "application/json")

	// what if :body is empty

	if r.Body == nil {
		json.NewEncoder(w).Encode("please send some data ")
	}

	// what about -{}
	var course Course
	_ = json.NewDecoder(r.Body).Decode(&course)

	if course.IsEmpty() {
		json.NewEncoder(w).Encode("No Data inside JSON ")
		return
	}

	//generate unique user id and convert it into string
	//append Course into Courses
	rand.Seed(time.Now().UnixNano()) // rand package is used to generte random number
	course.CourseId = strconv.Itoa(rand.Intn(100))
	Courses = append(Courses, course)
	json.NewEncoder(w).Encode(course)
	return

}

// updating a Course component to the database
func UpdateOneCourse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("updating a Course component to the database ")
	w.Header().Set("content-type", "application/json")

	// first grab id from request

	params := mux.Vars(r)

	//loop through the body
	// get the id from the body
	// remove the Course from the body
	//add with my id to the database

	for index, course := range Courses {

		if course.CourseId == params["id"] {
			Courses = append(Courses[:index], Courses[index+1:]...) // removes a element from the golang at a specified index //The ... after the second slice tells Go to "unpack" the slice and pass its individual elements as separate arguments to the append() function.

			var course Course

			_ = json.NewDecoder(r.Body).Decode(&course)
			course.CourseId = params["id"] // if Course id is equal to the id coming from the request
			Courses = append(Courses, course)
			json.NewEncoder(w).Encode(course)
			return
		}

	}
	//TODO: send a response when id is not found

}

func deleteOneCourse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("code to delete one file from the struct type ")
	w.Header().Set("content-type", "application/json")

	params := mux.Vars(r)

	// loop through the datatype (struct)
	// find the id of the Course u want to remove
	// remove the Course using slice operation(variatic operation)

	for index, course := range Courses {
		if course.CourseId == params["id"] {
			Courses = append(Courses[:index], Courses[index+1:]...)
			break

		}
	}
	response := json.NewEncoder(w).Encode(Courses)
	fmt.Println("we have successfully deleted the Course ", response)
}
