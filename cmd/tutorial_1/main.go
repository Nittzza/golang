package main

import "fmt"

func main() {

	var myString = []rune("résumé")

	var indexed = myString[1]

	fmt.Printf("%v, %T\n", indexed, indexed)

	for i, v := range myString {
		fmt.Println(i, v)
	}
}
