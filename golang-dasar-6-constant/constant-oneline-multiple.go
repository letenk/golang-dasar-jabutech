package main

import "fmt"

func main() {
	const name, married, age = "Dicky Rush", true, 23

	fmt.Println(name, married, age)
	fmt.Printf("Type: %T, %T, %T\n", name, married, age)
}
