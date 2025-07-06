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
	var arrayD [3]int

	fmt.Println(arrayA, arrayB, arrayC, arrayD)

	// arrayD = arrayB // Type mismatch

	// Slices: Variable length array, Slice has a descriptor similar to a string
	var sliceA []int
	var sliceB = []int{1, 2} // it has a length and capacity

	sliceA = append(sliceA, 1) // append takes a slice and an element and puts that element at the end of the slice
	// same as a string, the sliceA descriptor will change, it will have a new pointer to a new location in memory and a new length
	sliceB = append(sliceB, 3)
	sliceA = sliceB // in this case, we just copied the sliceB descriptor to sliceA, so both will point to the same underlying array
	sliceC := make([]int, 2) // make takes a type, a length and a capacity. result is same as []int{0, 0}
	sliceD := sliceA // alias, same descriptor as sliceA
	fmt.Println(sliceC[:2])
	fmt.Println(len(sliceD))
	// sliceD[0] == sliceB[0] // true because they point to the same underlying array 

	// Slices are passed by reference, thus elements are not copied, meaning that if i change something in sliceD, it will also reflect in sliceA
    // Slices are constructed of pointer to an array, length and capacity, pass by reference in this case is the pointer to the underlying array is the one that is copied
	// Most go apis use slices as an input
	// Generally not safe to use the equality operator to compare slices
	// Slices are generally used as a function parameter, because they are passed by reference


	// Array Slice Example
	var arrayW = [...]int{1, 2, 3}
	var sliceX = []int{0, 0, 0}

	y := arraySliceExample(arrayW, sliceX)

	fmt.Println(arrayW, sliceX, y)

	// Maps: map of key-value pairs
	// Maps are almost certainly like a hash table behind the scenes
	// We can read from a nil map, but we cannot write to it as it will panic
	var mapA map[string]int // map[keyType]valueType, this one will be nil with no storage
	mapB := make(map[string]int) // map[keyType]valueType, this one will be initialized menaing not nil with no storage

	fmt.Println(mapA, mapB)
}

func arraySliceExample(a [3]int, b []int) []int {
	//a = b // Syntax error because we cant assign a slice to an array
	a[0] = 4 // array variable outside the function scope (parameter - arrayW) will not change, only the a variable will be change because a copies the whole values of the parameters 
	b[0] = 3 // this will change the slice variable outside the function scope (parameter - sliceX) because b is a pointer to the underlying array
	
	c := make([]int, 5) // []int{0, 0, 0, 0, 0}
	c[4] = 42 // []int{0, 0, 0, 0, 42}

	copy(c, b) // copy the values of b to c. int[]{3, 0, 0, 0, 42}
	
	return c
}
