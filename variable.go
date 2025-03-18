package main

import "fmt"

// Di golang, Kalau Variable tidak digunakan maka akan error

func main() {
	// buat variable dengan tipe data string
	// var name string

	// name = "Muhammad"
	// fmt.Println(name)
	// // ketika variable yang sudah di isi lalu di isi dengan data yang berbeda, akan di override

	// name = "Muhammad Ridho "
	// fmt.Println(name)

	// name = "Muhammad Ridho Febrian"
	// fmt.Println(name)

	// versi yang sangat simple
	// tanda := hanya untuk menginisialisai deklrasai awal atau variabel utama
	name:= "Muhammad"
	fmt.Println(name)

	name = "Muhammad Ridho"
	fmt.Println(name)

	name = "Muhammad Ridho Febrian"
	fmt.Println(name)


	// Contoh deklarasi menggunakan multi variable
	var(
		firstName = "Muhammad"
		middleName = "Ridho"
		lastName = "Febrian"
	)

	fmt.Println(firstName)
	fmt.Println(middleName)
	fmt.Println(lastName)

}