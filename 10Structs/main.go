package main

import "fmt"

func main() {
	fmt.Println("Struct in Go")
	// Alternatives of Classes
	// no inheritance in golang
	// user-defined type
	// can be termed as a lightweight class that does not support inheritance but supports composition

	ketan := User{"Ketan", "ketan@gmail.com", 21, true}
	fmt.Println("User :", ketan)
	fmt.Printf("Detailed Info :%+v\n", ketan)
	fmt.Printf("Particular Info: User Name %v with Age %v\n", ketan.Name, ketan.Age)
}

// type User struct {
// 	Name     string
// 	Email    string
// 	Age      int
// 	IsActive bool
// }
type User struct {
	Name, Email string
	Age         int
	IsActive    bool
}
