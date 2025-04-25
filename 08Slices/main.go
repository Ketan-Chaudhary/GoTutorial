package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println("Slices in Go")
	// Slices are abstraction of the arrays
	// flexible and efficient way to represent arrays
	// Dynmaic size
	// slices reference to portion of a array
	// mostly used

	var slc = []string{}
	fmt.Printf("Type of the slice is : %T\n", slc)
	slc = append(slc, "Ketan", "Chaudhary", "Learning", "a", "new", "language")
	fmt.Println("List Contains:", slc)

	fmt.Println(len(slc))
	slc = append(slc[1:])
	fmt.Println("New value :", slc)

	slc = append(slc[:len(slc)-1])
	fmt.Println("New :", slc)

	// Using Make keyword
	newData := make([]int, 2)
	fmt.Println("Default:", newData)

	newData[0] = 1
	newData[1] = 2
	//newData[3]=3 Out of Bound error

	newData = append(newData, 3, 4, 5, 6) // works fine dymanically allocates the memory
	fmt.Println("NewData:", newData)

	sort.Ints(newData)
	fmt.Println("Sorted or not :", sort.IntsAreSorted(newData))
}
