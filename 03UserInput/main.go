package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	text := "Enter input from the user"
	fmt.Println(text)

	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Enter some input ")
	// comma ok syntax // error ok syntax (similar to try/catch block)
	input, _ := reader.ReadString('\n')
	// input, err := reader.ReadString('\n')
	// fmt.Println("If error occurs :",err)
	fmt.Println("The input is :", input)
	fmt.Printf("Type of input %T", input)

}
