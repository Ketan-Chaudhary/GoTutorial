package main

import (
	"fmt"
	"io"
	"net/http"
	"strings"
)

func main() {
	fmt.Println("GET Request")
	GetRequestHandle()

}

func GetRequestHandle() {
	const myUrl = "http://localhost:8000/get"

	response, err := http.Get(myUrl)
	if err != nil {
		fmt.Println(response.Status)
		panic(err)
	}

	fmt.Println("Status Code :", response.StatusCode)
	fmt.Println("Status :", response.Status)
	fmt.Println("Length of the response :", response.ContentLength)

	var responseString strings.Builder
	content, err := io.ReadAll(response.Body)
	byteCount, _ := responseString.Write(content)

	if err != nil {
		panic(err)
	}

	fmt.Println("Content from the server:", string(content))
	fmt.Println("Byte length :", byteCount)
	fmt.Println("Data :", responseString.String())
}
