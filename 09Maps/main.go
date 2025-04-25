package main

import (
	"fmt"
)

func main() {
	fmt.Println("Maps in Go")

	// Maps are reference to a hash table

	//Keys must be unique and of a comparable type, such as int, float64, rune, string, arrays, structs, or pointers. However, types like slices and non-comparable arrays or structs cannot be used as keys.
	//Values, on the other hand, can be of any type, including another map, pointers, or even reference types.

	ages := map[string]int{
		"A": 12,
		"B": 13,
		"C": 14,
	}
	fmt.Println("Age map :", ages)

	// Using make
	languages := make(map[string]string)
	languages["JS"] = "JavaScript"
	languages["PY"] = "Python"
	languages["GO"] = "GoLang"

	for keys, value := range languages {
		fmt.Printf("The key is %v with the value of %v\n", keys, value)
	}

	delete(languages, "PY")

	fmt.Println("Updated values", languages)

	value, exists := languages["PY"]
	if !exists {
		fmt.Println("Value is not present")
	} else {
		fmt.Println("Value is there :", value)
	}
}
