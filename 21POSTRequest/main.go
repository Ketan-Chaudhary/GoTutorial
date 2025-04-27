package main

import (
	"fmt"
	"io"
	"net/http"
	"strings"
)

func main() {
	fmt.Println("POST Request")
	PostRequestHandler()
}

func PostRequestHandler() {
	const MyUrl = "http://localhost:8000/post"

	// fake json payload

	requestBody := strings.NewReader(`
		{
			"Ketan":"Chaudhary",
			"Learning":"GoLang"
		}
	`)

	response, _ := http.Post(MyUrl, "application/json", requestBody)
	defer response.Body.Close()

	var DataResponse strings.Builder
	cnt, _ := io.ReadAll(response.Body)
	byteData, _ := DataResponse.Write(cnt)
	fmt.Println("Length of byteData", byteData)
	fmt.Println("Data :", DataResponse.String())
}
