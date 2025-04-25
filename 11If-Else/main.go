package main

import "fmt"

func main() {
	fmt.Println("If-else statements")

	age := 12
	if age > 18 {
		fmt.Println("Entry Allowed")
	} else if age == 18 {
		fmt.Println("Entry Allowed in supervision")
	} else {
		fmt.Println("Entry Not allowed")
	}

	if age := 21; age > 20 {
		fmt.Println("Can get DL")
	} else {
		fmt.Println("UnderAge for DL")
	}
}
