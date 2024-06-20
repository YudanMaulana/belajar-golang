package main

import (
	"fmt"
	"go-calculator/calculator"
	"log"
)

func main() {
	var op string
	var a, b float64

	fmt.Println("Masukan angka pertama: ")
	_, err := fmt.Scanln(&a)
	if err != nil {
		log.Fatalf("Nomor tidak valid: %v", err)
	}

	fmt.Println("Masukan angka kedua: ")
	_, err = fmt.Scanln(&b)
	if err != nil {
		log.Fatalf("Nomor tidak valid: %v", err)
	}

	fmt.Println("Silahkan pilih operator penjumlahan(+, -, *, /): ")
	_, err = fmt.Scanln(&op)
	if err != nil {
		log.Fatalf("Nomor tidak valid: %v", err)
	}

	result, err := calculator.Calculate(a, b, op)
	if err != nil {
		log.Fatalf("Error dalam penghitungan: %v", err)
	}

	fmt.Printf("Hasil Penjumlahan: %v\n", result)
}
