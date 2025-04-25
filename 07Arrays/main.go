package main

import "fmt"

func main() {
	fmt.Println("Arrays in Golang")

	var arr [3]string
	arr[0] = "new"
	arr[1] = "array"
	arr[2] = "defined"

	fmt.Println("Array Contains value :", arr)

	arr2 := [4]string{"Short", "Hand", "notation", "of Array"}
	fmt.Println("Array 2 contains :", arr2)

	// MultiDimensional Array
	arr3 := [2][2]int{{1, 2}, {3, 4}}
	fmt.Println("MultiDimensional Array:", arr3)

	// Length of the array
	fmt.Println("Length of the array:", len(arr3))

	// Creating an array whose size is determined
	// by the number of elements present in it
	// Using ellipsis
	arr4 := [...]int{1, 3, 5, 6, 8, 0}
	fmt.Println("Length of the ellipsis Array:", len(arr4))

	arr5 := [6]int{1, 3, 5, 6, 8, 0}
	arr6 := [6]int{1, 3, 5, 6, 8, 1}
	//arr7 := [5]int{1, 3, 5, 6, 8}
	// two arrays can be compared directly using == operator but the size and the Type should be same
	// i.e the length and the dataType of the both the arrays should be same
	fmt.Println(arr4 == arr5)
	fmt.Println(arr4 == arr6)
	//fmt.Println(arr4 == arr7)
}
