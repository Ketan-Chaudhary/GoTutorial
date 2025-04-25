package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("Enter some input")
	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')

	fmt.Println("Your entered input is :", input)

	// Adding 1 to your input
	//fmt.Println("Adding 1 to the input :", input+1) // invalid operation mismatched types
	// so at first we need to convert it into number (int or float)
	// for that we will be using the library strconv

	//incInput, err := strconv.ParseFloat(input, 64) // convert the string value into 64 bit float
	incInput, err := strconv.ParseFloat(strings.TrimSpace(input), 64) // remove the \r\n from the string (trim the string) using strings package

	if err != nil {
		fmt.Println("Error occured:", err)
	} else {
		fmt.Print("Increased number is :", incInput+1)
	}

}
