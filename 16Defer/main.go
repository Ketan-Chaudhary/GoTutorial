package main

import "fmt"

func main() {
	defer fmt.Println("Last Line to print")
	defer fmt.Println("LIFO")
	fmt.Println("First Line to be printed")
	fmt.Println("What happens now")
	myDefer()
}

func myDefer() {
	for i := 0; i < 5; i++ {
		defer fmt.Print(i)
	}
}
