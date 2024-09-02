package main

import (
	"fmt"
)

func main() {
	// Initialize the array
	thing1 := [5]float64{1, 2, 3, 4, 5}

	// Print the memory location of the thing1 array
	fmt.Printf("\nThe memory location of the thing1 array is: %p\n", &thing1)

	// Call the square function and store the result
	result := square(thing1)

	// Print the result array
	fmt.Printf("\nThe result is: %v\n", result)
}

// Function to square the elements of an array
func square(thing2 [5]float64) [5]float64 {
	// Print the memory location of the thing2 array
	fmt.Printf("\nThe memory location of the thing2 array is: %p\n", &thing2)

	// Square each element in the array
	for i := range thing2 {
		thing2[i] = thing2[i] * thing2[i]
	}
	return thing2
}
