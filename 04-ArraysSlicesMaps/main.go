package main

import (
	"fmt"
)

func main() {
	// Composite Types: Arrays, Slices, Maps
	// Arrays: Sequence of elements of the same type, have fixed size at compile time
	// Array is passed by value, thus elements are copied
	// Arrays are usually used for pseudo-constants, like the days of the week, the months of the year, etc.
	var arrayA [2]int
	var arrayB = [2]int{1, 2}
	var arrayC = [...]int{1, 2, 3} // ... is used to tell the compiler to infer the length of the array based on the number of elements in the initializer
	var arrayD [3]int              // [0, 0, 0]

	fmt.Println(arrayA, arrayB, arrayC, arrayD)

	// arrayD = arrayB // Type mismatch

	// Slices: Variable length array, Slice has a descriptor/header similar to a string and it points to an underlying array
	var sliceA []int
	var sliceB = []int{1, 2} // it has a length and capacity
	// length is the number of elements in the slice
	// capacity is the number of elements in the underlying array

	sliceA = append(sliceA, 1) // append takes a slice and an element and puts that element at the end of the slice
	// same as a string, the sliceA descriptor will change, it will have a new pointer to a new location in memory and a new length
	sliceB = append(sliceB, 3)
	sliceA = sliceB          // in this case, we just copied the sliceB descriptor to sliceA, so both will point to the same underlying array
	sliceC := make([]int, 2) // make takes a type, a length and a capacity. result is same as []int{0, 0}
	// we use capacity to preallocate the memory for the slice, so that we dont have to resize the slice everytime we append an element later
	// sliceC := make([]int, 2, 4) // []int{0, 0} with capacity of 4
	// Reallocation is bad because it creates a new slice and copies the elements to the new slice, this is expensive
	// Memory allocation is slow, it uses cpu cycles and time, and it causes memory churn (memory fragmentation)
	// However, dont over optimize, only pre-allocate when you know the size or when performance matters

	// When to use Make with capacity?
	// Scenario 1: When it will produce the same number of outputs as inputs, use make with same length and capacity (make([]int, 2, 2) or make([]int, 2)) we then use index based transformation
	// Scenario 2: When you know it will produce fewer outputs than inputs, use make with 0 length and exact capacity needed (make([]int, 0, 2)), we then use append to add elements
	// Scenario 3: Unknown input and outputs, use make with 0 length and estimated capacity

	sliceD := sliceA // alias, same descriptor as sliceA
	fmt.Println(sliceC[:2])
	fmt.Println(len(sliceD))
	// sliceD[0] == sliceB[0] // true because they point to the same underlying array

	// Slices are passed by value, meaning it passes the full header/descriptor that includes a pointer to the underlying array, length and capacity
	// By including a pointer to the underlying array, changes made to the values/elements inside the slice will be reflected in the caller's slice as well.
	// However, since its passed by value, the header/descriptor (e.g. length and capacity) of the caller's slice will not be affected if we modify the slice using append, etc.
	// Most go apis use slices as an input
	// Generally not safe to use the equality operator to compare slices
	// Slices are generally used as a function parameter, because they are cheap to pass/copy (only the header/descriptor is copied), it accepts variable length of elements

	// Array Slice Example
	var arrayW = [...]int{1, 2, 3}
	var sliceX = []int{0, 0, 0}

	y := arraySliceExample(arrayW, sliceX)

	fmt.Println(arrayW, sliceX, y)

	// Maps: map of key-value pairs
	// Maps are almost certainly like a hash table behind the scenes
	// We can read from a nil map, but we cannot write to it as it will panic
	// Maps points to a hash table behind the scenes
	var mapWithoutMake map[string]int // map[keyType]valueType, this one will be nil with no storage, this will not point to a hash table because it is not initialized yet
	mapMake := make(map[string]int)   // map[keyType]valueType, this one will be initialized meaning not nil with no storage

	// Comma-OK Idiom:
	// m := map[string]int{"a": 1, "b": 0}
	// if we want to check if a key exists in the map, we can use the comma-ok idiom
	// v, ok := m["a"] // v=1, ok=true because "a" is in the map
	// v, ok := m["b"] // v=0, ok=true because "b" is in the map
	// v, ok := m["c"] // v=0, ok=false because "c" is not in the map

	fmt.Println(mapWithoutMake, mapMake)

	mapA := mapMake["the"]        // returns 0, any time we read a key that doesnt exists, we get the default value of the map type
	mapB := mapWithoutMake["the"] // same thing
	//mapWithoutMake["the"] = 1 // this will panic, because we cannot write to a nil map
	mapWithoutMake = mapMake // copies the mapMake descriptor to mapWithoutMake, so both will point to the same hash table
	mapWithoutMake["and"]++  // mapWithoutMake will utilise the same hash table as mapMake
	mapC := mapMake["and"]   // mapC will be 1, because mapMake["and"]++ will increment the value of the key "and" in the hash table
	// the type used for the key must (be comparable) have == and != operators defined for it, e.g. int, string, bool, arrays (not slices, maps, functions)
	// maps cant be compared to one another, it can be compared to nil
	fmt.Println(mapA, mapB, mapC)
}

func arraySliceExample(a [3]int, b []int) []int {
	//a = b // Syntax error because we cant assign a slice to an array
	a[0] = 4 // array variable outside the function scope (parameter - arrayW) will not change, only the a variable will be change because a copies the whole values of the parameters
	b[0] = 3 // this will change the slice variable outside the function scope (parameter - sliceX) because b's header contains a pointer to the underlying array

	c := make([]int, 5) // []int{0, 0, 0, 0, 0}
	c[4] = 42           // []int{0, 0, 0, 0, 42}

	copy(c, b) // copy the values of b to c. []int{3, 0, 0, 0, 42}

	return c
}
