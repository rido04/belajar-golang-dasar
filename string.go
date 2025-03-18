package main

import "fmt"

// Tipe data string
// Tipe data string adalah tipe data yang digunakan untuk menyimpan data teks. Tipe data string dideklarasikan dengan menggunakan tanda petik dua (”).
// Tipe data string memiliki beberapa fungsi yang dapat digunakan untuk memanipulasi data string, seperti len(), untuk menghitung panjang string, dan +, untuk menggabungkan dua string.

func main(){
	fmt.Println("Muhammad")
	fmt.Println("Muhammad Ridho")
	fmt.Println("Muhammad Ridho Febrian")

	fmt.Println(len("Muhammad"))
	fmt.Println("Muhammad Ridho"[0])  // Mengambil dalam bentuk bytee
	fmt.Println("Muhammad Ridho Febrian"[1]) // Mengambil dalam bentuk bytee
	fmt.Println("Muhammad Ridho Febrian"[0:8]) // Mengambil dalam bentuk string
}