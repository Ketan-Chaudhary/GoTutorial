package main

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"
)

func main() {
	fmt.Println("POST Request - Form data input")
	PostFormHandler()
}

func PostFormHandler() {
	const myURL = "http://localhost:8000/postform"

	// formdata

	data := url.Values{}
	data.Add("firstname", "ketan")
	data.Add("Lastname", "chaudhary")
	data.Add("email", "test@gov.in")

	response, _ := http.PostForm(myURL, data)
	defer response.Body.Close()

	var DataResponse strings.Builder
	cont, _ := io.ReadAll(response.Body)
	DataBytes, _ := DataResponse.Write(cont)

	fmt.Println("DataByte:", DataBytes)
	fmt.Println("Data :", DataResponse.String())

}
