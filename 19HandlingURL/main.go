package main

import (
	"fmt"
	"net/url"
)

const MYURL string = "https://rufrent.com:5000/property/myaparna?availability=immediately&size=1BHK"

func main() {
	fmt.Println("Handling URLs in GoLang")
	fmt.Println(MYURL)

	// Parsing Url
	result, err := url.Parse(MYURL)

	if err != nil {
		panic(err)
	}

	fmt.Printf("Type of the result %T", result)
	fmt.Println("Scheme :", result.Scheme)
	fmt.Println(result.Host)
	fmt.Println(result.Path)
	fmt.Println(result.Port())
	fmt.Println(result.RawQuery)

	qparams := result.Query() // key value pair
	fmt.Printf("Type of query params : %T\n", qparams)
	fmt.Println("Query :", qparams)

	for _, val := range qparams {
		fmt.Println(val)
	}

	// Creating a url

	NewURL := &url.URL{
		Scheme:   "https",
		Host:     "rufrent.com:5000",
		Path:     "/property/mynew",
		RawQuery: "User=Ketan&password=123",
	}

	fmt.Println("New URL :", NewURL.String())

}
