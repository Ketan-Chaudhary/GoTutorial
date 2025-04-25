package main

import "fmt"

func main() {
	fmt.Println("Pointers")

	// creting a pointer
	var ptr *int
	// Default value pointer
	fmt.Println("Defalult value of a pointer is :", ptr)

	myReference := 21
	// Pointing a pre initialized variable
	var ptr2 = &myReference
	// & - reference
	fmt.Println("The value of the ptr2 is :", ptr2)
	fmt.Println("The actual value it is holding is :", *ptr2)

}
