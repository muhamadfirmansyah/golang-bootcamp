package main

import (
	"fmt"
)

func main() {
	var angka1 int
	var angka2 int
	var operator string

	fmt.Printf("Masukan angka 1: ")
	fmt.Scan(&angka1)

	fmt.Printf("Masukan angka 2: ")
	fmt.Scan(&angka2)

	fmt.Printf("Masukan operator (+, -, *, /): ")
	fmt.Scan(&operator)


	switch {
		case operator == "+":
			fmt.Printf("Hasil dari %d tambah %d adalah %d\n", angka1, angka2, angka1 + angka2)
		case operator == "-":
			fmt.Printf("Hasil dari %d kurang %d adalah %d\n", angka1, angka2, angka1 - angka2)
		case operator == "*":
			fmt.Printf("Hasil dari %d kali %d adalah %d\n", angka1, angka2, angka1 * angka2)
		case operator == "/":
			fmt.Printf("Hasil dari %d bagi %d adalah %d\n", angka1, angka2, angka1 / angka2)
		default:
			fmt.Printf("Data yang dimasukan tidak valid!")
	}
}
