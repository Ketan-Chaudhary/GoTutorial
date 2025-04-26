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
	ketan.getStatus()
	ketan.UpdateEmail()
	fmt.Println("User :", ketan)
	ketan.UpdateEmailActually()
	fmt.Println("User:", ketan)
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

func (u User) getStatus() {
	fmt.Println("Is User Active :", u.IsActive)
}

func (u User) UpdateEmail() { // didn't pass the actual object just a copy so the updated entry is not reflected
	u.Email = "new@gmail.com"
	fmt.Println("Updated Mail :", u.Email)
}
func (u *User) UpdateEmailActually() { // didn't pass the actual object just a copy so the updated entry is not reflected
	u.Email = "new@gmail.com"
	fmt.Println("Updated Mail :", u.Email)
}
