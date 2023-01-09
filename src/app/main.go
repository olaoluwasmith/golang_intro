package main

import "fmt"

func main() {
	// Variables in Go
	var firstName string = "Linda" // method 1
	var middleName string          // method 2
	middleName = "Happi"
	lastName := "Smith" // method 3
	fmt.Println(firstName + " " + middleName + " " + lastName)

	// Constants
	const myConst int = 42
	fmt.Printf("%v, %T\n", myConst, myConst)

	// Data Types in Go
	/* Strings (string),
	   Integer (int, int8, int16, int32, int64, uint8..., rune, byte...),
	   Boolean (true or false),
	   Floating-point numbers (float32, float64),
	   Complex numbers (complex64, complex128)
	*/

	// Float
	a := 20.45
	b := 34.89
	c := b - a
	fmt.Printf("%v, %T\n", c, c)

	// Complex No
	var n complex64 = 1 + 2i
	fmt.Printf("%v, %T\n", n, n)
}
