package main

import "fmt"

func main() {

	var myMap map[string]uint8 = make(map[string]uint8)
	fmt.Println(myMap)

	var myMap2 = map[string]uint8{"Nittya": 23, "Aaditya": 45}
	fmt.Println(myMap2["Nittya"])

	var age, ok = myMap2["Jason"]
	if ok {
		fmt.Printf("The age is %v", age)
		fmt.Println("Invalid Name")

		for name, age := range myMap2 {
			fmt.Printf("Name: %v, Age: %v\n", name, age)
		}

	}
}
