package average

import (
	"fmt"
	"os"
)

// this function can read numbers from files
// go run . < nums.txt
// cat nums.txt | go run .
func Average() float64 {
	var sum float64
	var count int

	for {
		var value float64
	 
		if _, err := fmt.Fscanln(os.Stdin, &value); err != nil { // ctrl + d to exit
			break;
		}

		sum += value
		count++
	}

	if count == 0 {
		fmt.Println("No value provided")
	}

	return sum/float64(count)
}