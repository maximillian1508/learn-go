package main

import (
	"fmt"
	"main/average"
)

func main() { // Go can infer the type of the variable if you don't specify it
	//Variable declaration
	var a int = 10

	var (
		b = 2
		c = 2.1
	) // this one can be used to declare variables with different types if we explicitly specify the type of the variable

	// Short Declaration, Syntax Shorthand, cant be used outside a function, cant declare the type
	d := 3

	// Multiple Declaration
	e, f := 4, 5 // cannot declare different types in the same line, will generate compiler error - type mismatch

	fmt.Printf("a: %8T %[1]v\n", a) // %T VERB is the type of the variable, %v VERB is the value of the variable
	fmt.Printf("b: %8T %[1]v\n", b) // The 8 is the width of the field, it will be padded with spaces to the right
	fmt.Printf("c: %8T %[1]v\n", c) // The [1] is the index of the variable, it will be padded with spaces to the right
	fmt.Printf("d: %8T %[1]v\n", d) // Printf's family uses 1-based indexing
	fmt.Printf("e: %8T %[1]v\n", e)
	fmt.Printf("f: %8T %[1]v\n", f)

	a = int(c)                      // Type cast
	fmt.Printf("a: %8T %[1]v\n", a) // %T VERB is the type of the variable, %v VERB is the value of the variable
	c = float64(a)                  // Type cast
	fmt.Printf("c: %8T %[1]v\n", c)

	// Initialization
	// Every variable in go is initialized
	// All numerical types are initialized to 0
	// All boolean types are initialized to false
	// All string types are initialized to ""
	// For aggregate types, like arrays and structs, all the elements are initialized to their "zero" value
	// Other than that, everything is initialized to nil

	// bool
	// string
	// int, int8, int16, int32, int64
	// uint, uint8, uint16, uint32, uint64, uintptr

	// byte: alias for uint8
	// rune: alias for int32 - represents a Unicode code point

	// float32, float64
	// complex64, complex128

	// Constants
	// In go, constants can only be numbers, strings, and booleans
	fmt.Println("The average is", average.Average())

	// Numeric Constants
	// Numeric constants are high precision values that are not stored in memory, they are evaluated at compile time
	// Therefore, they are not affected by the precision of the floating point numbers
	// const Big = 1 << 100 // 1 shifted 100 times to the left
	// const Small = Big >> 99 // 1 shifted 99 times to the right
	// Go constants are computed at compile time with unlimited precision, then checked against the target type only when assigned or passed to a function.
	// This gives you flexibility to do intermediate calculations that exceed normal integer limits.

	// Type conversion
	// var i int = 42
	// var f float64 = float64(i)
	// var u uint = uint(f)

}
