package main

import "fmt"

func main(){
	// operasi matematika
	var a = 10
	var b = 10
	var c = 100
	// operasi matematika yang di dahulukan adalah perkalian
	var d = a * b % c
	
	fmt.Println(d)

	var i = 100
	i += 10 // i pada variabel awal di tambah dengan i pada variabel baru
	i += 500
	fmt.Println(i) // 110
	// bisa di deklarasikan ulang atau ditambah ulang
	i += 100
	fmt.Println(i) 

	// unary operator
	var j = 100
	j ++
	fmt.Println(j)
	j ++
	fmt.Println(j)
	j --
	fmt.Println(j)

}

// ada beberapa fungsi di operasi matematika
// + : penjumlahan
// - : pengurangan
// * : perkalian
// / : pembagian
// % : sisa pembagian

// adapun augmanted assignment, yaitu operasi untuk diri sendiri
//  += : penjumlahan
//  -= : pengurangan
//  *= : perkalian
//  /= : pembagian
//  %= : sisa pembagian 

// adapun unary operator
// ++ untuk tambah 1
// -- untuk kurang 1
// - untuk membuat jadi negative
// = untuk membuat jadi positive (default nya sudah positif)
// ! untuk kebalikan nilai boolean
