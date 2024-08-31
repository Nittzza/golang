package main

import (
	"fmt"
)

func main() {
	var thing1 = [7]float64{5, 10, 15, 20, 25, 30, 35} // Changed array size and values

	fmt.Printf("\nThe memory location of the thing1 array is: %p", &thing1)

	var result [7]float64 = processArray(thing1)

	fmt.Printf("\nThe result is: %v", result)
	fmt.Printf("\nThe value of thing1 is: %v", thing1)
}

func processArray(thing2 [7]float64) [7]float64 {
	fmt.Printf("\nThe memory location of the thing2 array is: %p", &thing2)

	for i := range thing2 {
		if thing2[i] > 20 {
			thing2[i] = thing2[i] + 10 // If the value is greater than 20, add 10
		} else {
			thing2[i] = thing2[i] - 5 // Otherwise, subtract 5
		}
	}

	return thing2
}
