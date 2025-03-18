package main

import "fmt"

func main() {
	// type declaration
	// type bisa di sebut sebagai alias dari tipe data yang sudah ada
	type NoKtp string

	var contoh string = "33221124124"
	var contohKtp NoKtp = "12121212121"

	fmt.Println(contoh)
	fmt.Println(contohKtp)

}