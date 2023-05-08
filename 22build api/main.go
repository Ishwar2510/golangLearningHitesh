package main

import (
	"encoding/json"
	"math/rand"
	"net/http"
	"strconv"
	"time"

	"github.com/gorilla/mux"
)

func main() {
	// seeding data into the course
	courses = append(courses, Course{"2", "React", 45, &Author{FullName: "ishwar", Website: ".com"}})
	courses = append(courses, Course{"3", "js", 100, &Author{FullName: "kumar", Website: ".com"}})
	r := mux.NewRouter()
	r.HandleFunc("/", serveHome).Methods("GET")
	r.HandleFunc("/courses", getAllCourses).Methods("GET")
	r.HandleFunc("/course/{id}", getOneCourse).Methods("GET")
	r.HandleFunc("/course", createOneCourse).Methods("POST")
	r.HandleFunc("/course/{id}", updateOneCourse).Methods("PUT")
	// running the server
	http.ListenAndServe(":4000", r)
}

type Course struct {
	Courseid    string  `json:"courseid`
	CourseName  string  `json:"coursename"`
	CoursePrice int     `json:"courseprice"`
	Author      *Author `json:"author"`
}
type Author struct {
	FullName string `json:"fullname"`
	Website  string `json:"website"`
}

// fake data base
var courses []Course

// middleware
func (c *Course) isEmpty() bool {
	return c.Courseid == "" && c.CourseName == ""
}

// routes

// controllers

func serveHome(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("<h1> Hello boss</h1>"))
}

func getAllCourses(w http.ResponseWriter, r *http.Request) {
	// we need to send all the course data which are in json format
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(courses)
}

func getOneCourse(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "applictaion/json")
	params := mux.Vars(r)
	id := params["id"]

	for _, course := range courses {
		if course.Courseid == id {

			json.NewEncoder(w).Encode(course)
			return
		}
	}
	json.NewEncoder(w).Encode("No course found")

}

func createOneCourse(w http.ResponseWriter, r *http.Request) {
	// creatig data i.e mean you will get a json data in the request
	// we need to decode the json

	// whta if body is nil
	if r.Body == nil {
		json.NewEncoder(w).Encode("pls send some data")
	}

	var coursee Course

	json.NewDecoder(r.Body).Decode(&coursee)
	rand.Seed(time.Now().UnixNano())
	coursee.Courseid = strconv.Itoa(rand.Intn(100))
	courses = append(courses, coursee)
	json.NewEncoder(w).Encode(coursee)

}
func updateOneCourse(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id := params["id"]

	var data Course
	json.NewDecoder(r.Body).Decode(&data)

	for index, value := range courses {
		if value.Courseid == id {
			courses = append(courses[:index], courses[index+1:]...)
			break
		}
	}
	courses = append(courses, data)
}
