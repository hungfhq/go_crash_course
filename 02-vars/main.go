package main

import "fmt"

// var name = "Brad"

func main() {
	// MAIN TYPES
	// string
	// bool
	// int
	// int int8 int16 int32 int64
	// uint uint8 uint16 uint32 unint64 uintptr
	// byte - alias for uint8
	// rune - alias for int32
	// float32 float64
	// complex64 complex128

	// Using var
	// var name = "Brad"
	var age int32 = 37
	const isCool = true
	var size float32
	size = 1.3

	// name := "Brad"
	// email := "brad@gmail.com"

	// Shorthand
	name, email := "Brad", "brad@gmail.com"

	fmt.Println(name, age, isCool)
	fmt.Printf("%T, %T, %T\n", name, age, isCool)
	fmt.Printf("%T\n", size)
}
