package main

import "fmt"

func main() {

	intArr := [...]int32{1, 2, 3}
	fmt.Println(intArr)

	intSlice := []int32{4, 5, 6}
	fmt.Printf("The length is %v with capacity %v", len(intSlice), cap(intSlice))

	intSlice = append(intSlice, 7)
	fmt.Printf("\nThe length is %v with capacity %v", len(intSlice), cap(intSlice))
}
