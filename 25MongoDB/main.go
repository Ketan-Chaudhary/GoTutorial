package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/ketan-chaudhary/mongodb/router"
)

func main() {
	fmt.Println("MongoDB connection")

	r := router.Router()

	log.Fatal(http.ListenAndServe(":8000", r))
	fmt.Println("SERVER STARTS")
}
