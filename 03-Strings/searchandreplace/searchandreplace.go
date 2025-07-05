package searchandreplace

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func SearchAndReplace() {
	var t string
	var s []string
	if len(os.Args) < 3 {
		fmt.Fprintln(os.Stderr, "not enough arguments")
		os.Exit(1)
	}

	old, new := os.Args[1], os.Args[2]
	scanner := bufio.NewScanner(os.Stdin)

	for scanner.Scan() {
		s = strings.Split(scanner.Text(), old)
		t = strings.Join(s, new)
		fmt.Println(t)
	}
}
