package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	fmt.Println("JSON Request")
	//EncodeJson()
	DecodeJson()
}

type course struct {
	Name     string `json:"coursename"`
	Price    int
	Platform string
	Password string   `json:"-"`              // remove or hide this
	Tags     []string `json:"tags,omitempty"` // if nil remove that
}

func EncodeJson() {
	myCourses := []course{
		{"ReactJs", 299, "mydev.gov.in", "new123", []string{"web-dev", "JS"}},
		{"AngularJs", 199, "mydev.gov.in", "newP123", []string{"web-dev", "Angular"}},
		{"MERN stack", 399, "mydev.gov.in", "newB123", nil},
	}

	// pack this data as JSON data

	//finalJSON, err := json.Marshal(myCourses)
	finalJSON, err := json.MarshalIndent(myCourses, "", "\t")

	if err != nil {
		panic(err)
	}

	fmt.Printf("%s \n", finalJSON)
}

func DecodeJson() {
	jsonData := []byte(`
		{
                "coursename": "ReactJs",
                "Price": 299,
                "Platform": "mydev.gov.in",
                "tags": ["web-dev","JS"]
        }
	`)

	var myCourses course

	valid := json.Valid(jsonData)
	if valid {
		fmt.Println("Valid JSON")
		json.Unmarshal(jsonData, &myCourses)
		fmt.Printf("%#v \n", myCourses)
	} else {
		fmt.Println("InValid JSON")
	}
}
