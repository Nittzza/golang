package main

import (
	"fmt"
)

func main() {
	var thing1 = [5]float64{10, 20, 30, 40, 50} // Changed array values

	fmt.Printf("\nThe memory location of the thing1 array is: %p", &thing1)

	var result [5]float64 = square(thing1)

	fmt.Printf("\nThe result is: %v", result)
	fmt.Printf("\nThe value of thing1 is: %v", thing1)
}

func square(thing2 [5]float64) [5]float64 {
	fmt.Printf("\nThe memory location of the thing2 array is: %p", &thing2)

	for i := range thing2 {
		thing2[i] = thing2[i] * 2 // Changed logic to multiply by 2 instead of squaring
	}

	return thing2
}
