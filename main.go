package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	// Menampilkan pesan sambutan
	fmt.Println("Selamat datang di Kalkulator Sederhana!")

	// Membaca input dari pengguna
	fmt.Println("Masukkan bilangan pertama:")
	num1, err := getInput()
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}

	fmt.Println("Masukkan bilangan kedua:")
	num2, err := getInput()
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}

	// Menampilkan menu operasi
	fmt.Println("Pilih operasi:")
	fmt.Println("1. Penjumlahan (+)")
	fmt.Println("2. Pengurangan (-)")
	fmt.Println("3. Perkalian (*)")
	fmt.Println("4. Pembagian (/)")

	// Membaca pilihan operasi dari pengguna
	var choice int
	fmt.Scanln(&choice)

	// Memproses operasi sesuai pilihan pengguna
	var result float64
	switch choice {
	case 1:
		result = num1 + num2
	case 2:
		result = num1 - num2
	case 3:
		result = num1 * num2
	case 4:
		if num2 == 0 {
			fmt.Println("Error: Pembagian dengan nol tidak diperbolehkan")
			os.Exit(1)
		}
		result = num1 / num2
	default:
		fmt.Println("Error: Pilihan tidak valid")
		os.Exit(1)
	}

	// Menampilkan hasil operasi
	fmt.Printf("Hasil: %.2f\n", result)
}

// Fungsi untuk membaca input dari pengguna dan mengonversi ke tipe float64
func getInput() (float64, error) {
	var input string
	fmt.Scanln(&input)
	num, err := strconv.ParseFloat(input, 64)
	if err != nil {
		return 0, err
	}
	return num, nil
}
