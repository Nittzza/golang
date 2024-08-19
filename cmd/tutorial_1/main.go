package main

import "fmt"

func main() {

	var myMap map[string]uint8 = make(map[string]uint8)
	fmt.Println(myMap) // Prints the empty map

	var myMap2 = map[string]uint8{"Adam": 23, "Sarah": 45}
	fmt.Println(myMap2["Adam"])

	var age, ok = myMap2["Jason"]
	if ok {
		fmt.Printf("The age is %v", age)
	} else {
		fmt.Println("Invalid Name")
	}

	for name := range myMap2 {
		fmt.Printf("Name: %v\n", name)
	}
}
