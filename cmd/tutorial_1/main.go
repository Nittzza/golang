package main

import (
	"fmt"
	"math/rand"
	"time"
)

var dbData = []string{"id1", "id2", "id3", "id4", "id5"}

func main() {
	rand.Seed(time.Now().UnixNano()) // Seed the random number generator
	t0 := time.Now()                 // Start time

	for i, data := range dbData {
		dbCall(i, data)
	}

	fmt.Printf("\nTotal execution time: %v\n", time.Since(t0))
}

func dbCall(i int, data string) {
	// Simulate DB call delay between 1000 and 3000 ms
	delay := rand.Intn(2000) + 1000
	time.Sleep(time.Duration(delay) * time.Millisecond)

	fmt.Printf("The result from the database (ID %d): %s\n", i+1, data)
}
