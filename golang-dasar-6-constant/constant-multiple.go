package main

import "fmt"

func main() {
	const (
		fullname       = "Jhon Deer"
		age      uint8 = 28
		married  bool  = true
	)
	fmt.Printf("Nama %s, umur %d, sudah menikah? %t\n", fullname, age, married)
}
