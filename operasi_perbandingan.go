package main

import "fmt"

func main() {
	var name1 = "Ridho"
	var name2 = "Ridho"
	// buat variable untuk operasi perbandingan
	var result = name1 == name2
	fmt.Println(result)

	var result2 = name1 != name2
	fmt.Println(result2)

	var number1 = 20
	var number2 = 10

	var result3 = number1 > number2
	fmt.Println(result3)

	var result4 = number1 < number2
	fmt.Println(result4)
}

// operator perbandingan
// operator perbandingan digunakan untuk membandingkan dua buah nilai boolean yang menghasilkan true atau false
// > lebih dari
// < kurang dari
// >= lebih dari sama dengan
// <= kurang dari sama dengan
// == sama dengan
// != tidak sama dengan

// tipe data number bisa pakai semua operator perbandingan
// tipe data string hanya bisa pakai == dan !=

// tipe data string juga bisa menggunakan operator lebih dari dll
// contoh string b > a maka akan menghasilkan nilai true
// karena b lebih besar atau datang setelah a dalam urutan ASCII atau alfabet
