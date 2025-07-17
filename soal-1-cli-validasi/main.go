package main

import (
	"fmt"
)

func main() {
	var umur int
	var nama string

	fmt.Print("Masukan nama anda: ")
	fmt.Scan(&nama)
	fmt.Print("Masukan umur anda: ")
	_, err := fmt.Scan(&umur)
	if err != nil {
		fmt.Println("Masukan angka!")
		return
	}

	if umur >= 18 {
		fmt.Printf("Nama	: %s\n", nama)
		fmt.Printf("Umur	: %d\n", umur)
		fmt.Printf(">> Selamat datang, %s!", nama)
	} else {
		fmt.Printf("Nama	: %s\n", nama)
		fmt.Printf("Umur	: %d\n", umur)
		fmt.Printf(">> Error: umur tidak valid (minimal 18 tahun)\n")
	}

}
