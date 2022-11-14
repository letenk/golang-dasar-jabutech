package main

import (
	"fmt"
)

func main() {
	var nilaiPositif uint8 = 81
	var nilaiNegatif = -2147483648

	// Print nilai
	fmt.Printf("Bilangan positif %d\n", nilaiPositif)
	fmt.Printf("Bilangan negative %d\n", nilaiNegatif)

	// Cek tipe data
	fmt.Printf("Tipe data var nilaiPositif: %T\n", nilaiPositif)
	fmt.Printf("Tipe data var nilaiNegatif: %T\n", nilaiNegatif)
}
