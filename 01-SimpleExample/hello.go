package hello

import (
	"fmt"
	"strings"
)

// any function that starts with a capital letter can be exported, meaning it can be used outside the package (Private and Public functions)
// in go, functions can return multiple values - return x, y, z... - a, b := test()
func Say(name string) string {
	return fmt.Sprintf("Hello, %s!", name)
}

func SayMany(names []string) string {
	if len(names) == 0 {
		names = []string{"world"}
	}

	return "Hello, " + strings.Join(names, ", ") + "!"
}
