package main

import (
	"fmt"
	"math/rand"
	"time"
)

var dbData = []string{"id1", "id2", "id3", "id4", "id5"}

func main() {
	t0 := time.Now() // Start time

	for i := 0; i < len(dbData); i++ {
		dbCall(i)
	}

	fmt.Printf("\nTotal execution time: %v", time.Since(t0))
}

func dbCall(i int) {
	// Simulate DB call delay
	var delay float32 = rand.Float32()*2000 + 1000 // Generate a random delay between 1000 and 3000 ms
	time.Sleep(time.Duration(delay) * time.Millisecond)

	fmt.Println("The result from the database is:", dbData[i])
}
