package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"strconv"
	"time"

	"slices"

	"github.com/gorilla/mux"
)

// Defining model
type Course struct {
	CourseId    string  `json:"courseid"`
	CourseName  string  `json:"coursename"`
	CoursePrice int     `json:"price"`
	Author      *Author `json:"author"`
}

type Author struct {
	Fullname string `json:"fullname"`
	Website  string `json:"website"`
}

// Sample DB
var courses []Course

// Helper functions
func (c *Course) IsEmpty() bool {
	// return c.CourseId == "" && c.CourseName == ""
	return c.CourseName == ""
}

func main() {
	fmt.Println("API")

	r := mux.NewRouter()

	// sedding data
	courses = append(courses, Course{CourseId: "2",
		CourseName:  "ReactJs",
		CoursePrice: 299,
		Author: &Author{
			Fullname: "Ketan",
			Website:  "ketan.gov.in"}})

	courses = append(courses, Course{CourseId: "1",
		CourseName:  "GoLang",
		CoursePrice: 499,
		Author: &Author{
			Fullname: "Ketan",
			Website:  "ketan.gov.in"}})

	// routing
	r.HandleFunc("/", serveHome).Methods("GET")
	r.HandleFunc("/courses", getAllCourses).Methods("GET")
	r.HandleFunc("/course/{id}", getOneCourse).Methods("GET") // params
	r.HandleFunc("/course", createOneCourse).Methods("POST")
	r.HandleFunc("/course/{id}", updateOneCourse).Methods("PUT")
	r.HandleFunc("/course/{id}", deleteOneCourse).Methods("DELETE")

	// Listen to a port
	log.Fatal(http.ListenAndServe(":8000", r))
}

// Controllers

// serve home route

func serveHome(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("<h1>Welcome to Home page</h1>"))
}

func getAllCourses(w http.ResponseWriter, r *http.Request) {
	fmt.Println("All courses fetched")
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(courses)
}

func getOneCourse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Getting one particular course")
	w.Header().Set("Content-Type", "Application/josn")

	// grab id from request
	params := mux.Vars(r)
	fmt.Println("Param Value:", params)

	// loop through courses, find matching id and return the response
	for _, val := range courses {
		if val.CourseId == params["id"] { // id : defined in routing
			json.NewEncoder(w).Encode(val)
			return
		}
	}

	json.NewEncoder(w).Encode("No Course found with given id")
	// return
}

func createOneCourse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Create one course")
	w.Header().Set("Content-Type", "application/json")

	// what if: body is empty
	if r.Body == nil {
		json.NewEncoder(w).Encode("Please send some data")
	}

	// what if : data is empty like {}

	var course Course
	_ = json.NewDecoder(r.Body).Decode(&course)
	if course.IsEmpty() {
		json.NewEncoder(w).Encode("No Data Inside")
		return
	}

	//TODO: check only if title is duplicate
	//loop -> title matches with course -> Duplicate course

	// How to generate id, -> Convert them to string
	// append course into courses

	src := rand.NewSource(time.Now().UnixNano())
	random := rand.New(src)
	course.CourseId = strconv.Itoa(random.Intn(100))
	courses = append(courses, course)
	json.NewEncoder(w).Encode(course)
	return

}

func updateOneCourse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Updating Value ")
	w.Header().Set("Content-Type", "application/json")

	// first - grab id from request
	params := mux.Vars(r)

	// loop ->id -> remove -> add with my ID
	for index, course := range courses {
		if course.CourseId == params["id"] {
			//courses = append(courses[:index], courses[index+1:]... )
			courses = slices.Delete(courses, index, index+1)
			var course Course
			_ = json.NewDecoder(r.Body).Decode(&course)
			course.CourseId = params["id"]
			courses = append(courses, course)
			json.NewEncoder(w).Encode(course)
			return
		}
	}
}

func deleteOneCourse(w http.ResponseWriter, r *http.Request) {
	fmt.Print("Delete one course")
	w.Header().Set("Content-Type", "application/json")

	// grab the id -> loop
	pramas := mux.Vars(r)

	// loop ->id ->remove

	for index, course := range courses {
		if course.CourseId == pramas["id"] {
			//courses = append(courses[:index], courses[index+1:]...)
			courses = slices.Delete(courses, index, index+1)
			json.NewEncoder(w).Encode("Course Deleted ..")
			break
		}
	}
}
