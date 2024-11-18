package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	// Seed the random number generator
	rand.Seed(time.Now().UnixNano())

	// Generate a random number between 1 and 100
	num := rand.Intn(100) + 1

	// Print the random number
	fmt.Printf("Generated Number: %d\n", num)

	// Check if the number is even or odd
	if num%2 == 0 {
		fmt.Println("The number is Even!")
	} else {
		fmt.Println("The number is Odd!")
	}
}
