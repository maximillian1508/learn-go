package main

import (
	"fmt"
	search "main/searchandreplace"
)

func main() {
	s := "élite"

	// byte: alias for uint8
	// rune: alias for int32

	// In go, strings are immutable sequence of "characters" and can share the underlying storage
	// Physically, strings are a sequence of bytes (UTF-8). This defines the length of the string.
	// The length of the string is the number of bytes required to represent the unicode characters, not the number of unicode characters
	// Logically, strings are a sequence of (unicode) runes

	fmt.Printf("%T %[1]v %d\n", s, len(s)) // 6
	fmt.Printf("%T %[1]v %d\n", []rune(s), len([]rune(s))) // []int32{233, 108, 105, 116, 101} 5
	fmt.Printf("%T %[1]v %d\n", []byte(s), len([]byte(s))) // []uint8{195, 169, 108, 105, 116, 101} 6

	// String Structure
	// The internal string representation is a pointer and a length
	// String descriptor describes something, it has a pointer and additional information (number of bytes that makes up the string)
	// In go, the length of the string is encoded in this descriptor
	st := "hello, world"
	ts := st        // ts is a copy of st, it will have the same length and a pointer leading to the same bytes (st and ts have different address pointer but points to the same storage), but ts is a different descriptor
	hello := st[:5] // points to the same characters in memory as the original string

	f := "the quick brown fox"

	a := len(f)                 // 19
	b := f[:3]                  // "the"
	c := f[4:9]                 // "quick"
	d := f[:4] + "slow" + f[9:] //replaces "quick", then d will point to a completely different memory location

	//f[5] = "x" // this will not work because string is immutable and we cant change a part of it
	f += "es" // this will create a new string and point to a different memory location, the original "the quick brown fox" doesnt change but s will point to a different new location
	fmt.Println(ts, hello, f, a, b, c, d)

	// Go is a garbage collected language, if the "the quick brown fox" is not referenced anywhere else, it will be garbage collected
	search.SearchAndReplace()

	// A string can point to another string
}
