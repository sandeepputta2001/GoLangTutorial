package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"strconv"
	"time"

	"github.com/gorilla/mux"
)

type Course struct {
	CourseId    string  `json:"courseid"`
	CourseName  string  `json:"coursename"`
	CoursePrice int     `json:"price"`
	Author      *Author `json:"author"` //* author means that it is of type pointer

}

type Author struct {
	FullName string `json:"fullname"`
	Website  string `json:"website"`
}

func main() {

	fmt.Println("welcome to API build")

	r := mux.NewRouter()

	// seeding the data : it means that initialising the data

	var Courses []Course

	Courses = append(Courses, Course{CourseId: "2", CourseName: "ReactJS", Author: &Author{FullName: "sandeep putta ", Website: "sandeep.putta.com"}})
	Courses = append(Courses, Course{CourseId: "4", CourseName: "MERN", Author: &Author{FullName: "hemanth putta  ", Website: "hemanth.putta.com"}})

	//routing
	r.HandleFunc("/", ServeHome).Methods("GET")
	r.HandleFunc("/courses", GetAllCourses).Methods("GET")
	r.HandleFunc("/course/{id}", GetOneCourse).Methods("GET")
	r.HandleFunc("/course/{id}", CreateOneCourse).Methods("POST")
	r.HandleFunc("/course/{id}", UpdateOneCourse).Methods("PUT")
	r.HandleFunc("/course/{id}", deleteOneCourse).Methods("DELETE")

	//listen to a port

	log.Fatal(http.ListenAndServe(":8001", r))
}

func ServeHome(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("<h1>welcome to API by sandeep putta </h1>"))

}

func GetAllCourses(w http.ResponseWriter, R *http.Request) {
	fmt.Println("get all courses ")
	w.Header().Set("content-type", "application/json")
	json.NewEncoder(w).Encode(Courses)
}

func getOneCourse(w http.ResponseWriter, r *http.Request) string {

	fmt.Println("get one course ")
	w.Header().Set("content-type", "application/json")

	// grab id from request
	params := mux.Vars(r) // mux package mapping requests and responses ; vars fucntion extracts all the variables in the request(r)
	fmt.Printf("the type of params is :%t\n", params)

	// loop through the courses , find the matching id and return the response

	for _, course := range Courses {
		if course.CourseId == params["id"] {
			json.NewEncoder(w).Encode(course)
			return
		}
	}
	json.NewEncoder(w).Encode("No course found with given id ")
	return
}

// create a file to send in to database
func CreateOneCourse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("create one course to send into database ")
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
	//append course into courses
	rand.Seed(time.Now().UnixNano())
	course.CourseId = strconv.Itoa(rand.Intn(100))
	courses = append(courses, course)
	json.NewEncoder(w).Encode(course)
	return

}

// updating a course component to the database
func UpdateOneCourse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("updating a course component to the database ")
	w.Header().Set("content-type", "application/json")

	// first grab id from request

	params := mux.Vars(r)

	//loop through the body
	// get the id from the body
	// remove the course from the body
	//add with my id to the database

	for index, course := range Courses {

		if course.CourseId == params["id"] {
			courses = append(courses[:index], courses[index+1:]...)
			var course Course
			json.NewDecoder(r.Body).Decode(&course)
			_ = course.CourseId == params["id"] // if course id is equal to the id coming from the request
			courses = append(courses, course)
			json.NewEncoder(w).Encode(course)
			return

		}

	}
	//TODO: send a response when id is not found

}

func deleteOneCourse(w httpp.ResponseWriter, r *http.Request) {
	fmt.Println("code to delete one file from the struct type ")
	w.Header().Set("content-type", "application/json")

	params := mux.Vars(r)

	// loop through the datatype (struct)
	// find the id of the course u want to remove
	// remove the course using slice operation(variatic operation)

	for index, course := range Courses {
		if course.CourseId == params["id"] {
			Courses = append(Courses[:index], Courses[index+1:])
			break

		}
	}
	response := json.NewEncoder(w).Encode(&Courses)
	fmt.Println("we have successfully deleted the course ", response)
}
