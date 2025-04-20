package main

import "fmt"

func main() {
	// Types and Variables Declaration in Golang

	// Data Types
	// Basic type: Numbers, strings, and booleans come under this category.
	// Aggregate type: Array and structs come under this category.
	// Reference type: Pointers, slices, maps, functions, and channels come under this category.
	// Interface type

	// Basic type :
	// Integers:
	// 		int 8, int 16, int 32, int 64, unit8(unsigned), ... , int(either 32 or 64 bit), unit(similar to int with unsigned integer), byte(unit8), rune(int32), uintptr(unsigned interger type with no defined width hold all bits of a pointer value)

	var number int8 = 127
	fmt.Println("value :", number)
	fmt.Printf("Type of the variable %T \n ", number)

	var number2 int64 = 255
	fmt.Println("value :", number2)
	fmt.Printf("Type of the variable %T \n ", number2)

	var num int = 24563
	fmt.Println(num)
	fmt.Printf("Datatype %T \n", num)

	// other ways to declare & initialize variables
	num2 := 234 // work inside main or other function didn't work for global declaration of variables
	fmt.Println(num2)

	var num3 = 3444 // Lexer will automatically assign the correct data type
	fmt.Println(num3)

	// Multiple variables of same type
	var multinum1, multinum2, multinum3 int = 1, 2, 3
	fmt.Println(multinum1, multinum2, multinum3)

	// Floating-Point Numbers:
	// 		float32, float64

	var flt1 float32 = 23.34
	fmt.Println(flt1)

	flt2 := 34.34
	fmt.Println(flt2)

	//var flt3 float = 323.334 // incorrect declaration

	// Complex Number:
	// 		complex64(complex number which contains float32 as a real and imaginary component)
	//		complex128(complex number which contains flaot64 as a real and imaginary component)
	var a complex128 = complex(6, 3)
	fmt.Println(a)

	// Boolean : True/False
	var boolVal bool = true
	fmt.Println(boolVal)

	// String
	var str string = "Ketan Chaudhary"
	fmt.Println(str)

	str2 := "Go lang"
	fmt.Println(str2)

}
