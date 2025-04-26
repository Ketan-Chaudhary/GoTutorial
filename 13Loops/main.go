package main

import (
	"fmt"
)

func main() {
	fmt.Println("Loops in GoLang")

	// 	for initialization; condition; post{
	// 		// statements....
	//  }

	for i := 0; i < 5; i++ {
		fmt.Println(i)
	}

	lcl := 1
	for lcl < 3 {
		fmt.Println(lcl)
		lcl++
	}

	slc := []string{"is", "there", "is", "no", "char", "type", "in", "golang"}
	for i := range slc {
		fmt.Println(slc[i]) // returning index only
	}

	for index, value := range slc {
		fmt.Printf("Index %v : Value %v\n", index, value)
	}
}
