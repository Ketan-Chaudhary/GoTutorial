package main

import (
	"fmt"
	"io"
	"net/http"
)

const URL = "https://rufrent.com"

func main() {
	fmt.Println("Handling Web Request in Go ")

	response, err := http.Get(URL)

	if err != nil {
		panic(err)
	}

	fmt.Printf("Type of the response is : %T\n", response)

	defer response.Body.Close()

	byteData, err := io.ReadAll(response.Body)

	if err != nil {
		panic(err)
	}

	fmt.Println("Data :", string(byteData))
}
