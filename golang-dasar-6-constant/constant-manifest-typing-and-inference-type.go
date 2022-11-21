package main

import "fmt"

func main() {
	const firstName = "john" // Manifest typing
	const age uint = 27      // Inference typing
	// Print out
	fmt.Printf("Hi, saya %s umur saya %d\n", firstName, age)
}
