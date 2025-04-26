package main

import "fmt"

func main() {
	fmt.Println("Functions in GoLang")

	myfunction()
	add(1, 3)
	fmt.Println(addNice(3, 4))
	fmt.Println(addPro(1, 2, 3, 4, 5, 6, 7, 8))

	msg, val := addProMax(1, 3, 4)
	fmt.Println("Msg :", msg)
	fmt.Println("Val :", val)
}

func myfunction() {
	fmt.Println("Message for my function")
}

func add(val1, val2 int) {
	fmt.Println(val1 + val2)
}

func addNice(val1, val2 int) int {
	return val1 + val2
}

func addPro(values ...int) int {
	total := 0
	for _, val := range values {
		total += val
	}
	return total
}

func addProMax(values ...int) (int, string) {
	total := 0
	for _, val := range values {
		total += val
	}
	msg := "message for proMax function"
	return total, msg
	// return total,"New message from function"
}
