package main

import "fmt"

func main() {
	name := "Jhon Deer"
	length := len(name)
	stringIndex1 := name[1]

	fmt.Printf("Hallo %s\n", name)
	fmt.Printf("Panjang string %d\n", length)
	fmt.Printf("Panjang string %q\n", stringIndex1)
}
