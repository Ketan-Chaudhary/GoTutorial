package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	fmt.Println("File Handling")
	content := "Adding this content in a file"

	file, err := os.Create("./newfile.txt")

	if err != nil {
		panic(err)
	}

	length, err := io.WriteString(file, content)
	if err != nil {
		panic(err)
	}
	fmt.Println("Length is :", length)
	//defer file.Close() // using defer is recommended
	file.Close() // to resolve the file opened by another program while deleting the file

	readFile(file.Name())
	writeFile(file.Name())
	defer deleteFile(file.Name())
}

func readFile(filename string) {
	data, err := os.ReadFile(filename)
	if err != nil {
		panic(err)
	}
	fmt.Println("Name/Path of the file :", filename)
	fmt.Println("Length of data :", len(data))
	fmt.Println("Data :", string(data)) // data comes in byte slice so we need to convert it into String
}

func writeFile(filename string) {
	content := "New content"
	file, err := os.OpenFile(filename, os.O_APPEND|os.O_WRONLY, 0644) // 0644 is a file permission mode in octal notation — it's a Unix-style permission code that defines who can read or write the file.

	if err != nil {
		panic(err)
	}

	file.WriteString(content)

	// 0   -> Indicates octal literal
	// 6   -> Owner:    read + write (4 + 2)
	// 4   -> Group:    read only
	// 4   -> Others:   read only

	defer file.Close()

}

func deleteFile(filename string) {

	err := os.Remove(filename)
	if err != nil {
		panic(err)
	}
}
