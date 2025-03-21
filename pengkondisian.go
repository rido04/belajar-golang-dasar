package main

import "fmt"

func main(){
	name:= "Ridho"

	// menggunakan if
	if name == "Ridho"{
		// bisa juga menggunakan perbandingan yang lain
		fmt.Println("Halooo ma frennn",name)
		// menggunakan else if
	}else if name == "Bagas"{
		fmt.Println("Si Bagas Ngapain dah")
	} else {
	fmt.Println("Lu siapa bjir")
	}

	// contoh menggunakan short statement
	// jika ingin menggunakan nya di beberapa pengkondisian bisa di inisialisasi dulu di awal
	// length := len(name)
	// jika hanya ingin menggunakan nya di satu kondisi bisa langsung di masukkan ke dalam if
	if length := len(name); length > 5 {
		fmt.Println("Namanya kepanjangan brader")
	} else {
		fmt.Println("Nah Pas udah segini")
	}
}