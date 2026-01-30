package main

// factored import statement, this is a good practice to keep the imports organized and easy to read
import (
	"fmt"
	"hello"
	"os"
)

func main() { // the main function is the entry point of go program and it needs to be in the main package
	// if len(os.Args) > 1 {
	// 	fmt.Println(hello.Say(os.Args[1])) // os.Args is a slice of strings that contains the command line arguments passed to the program
	// } else {
	// 	fmt.Println(hello.Say("World"))
	// }

	fmt.Println(hello.SayMany(os.Args[1:]))
}