package main

import (
	"fmt"
)

// recover adalah function yang bisa kita gunakan untuk menangkap panic
// dengan recover, proses panic akan berhenti, sehingga program akan tetap berjalan

func endApp(){
	// contoh pemanggilan recover yang benar
	fmt.Println("Aplikasi Selesai Berjalan")
	// jadi langsung di cek di function yang di defer apakah ada error atau tidak
	message:= recover()
	fmt.Println("ups, terjadi kesalahan", message)
	// jadi ketika di recover, panicnya tidak akan berhenti
}

func runApp(error bool){
	defer endApp()
	if error {
		// jika sudah menemukan kata kunci panic, maka akan berhenti di situ
		panic("ERROR")
	}
	// contoh recover yang salah, yaitu memanggil recover nya di bawah
	// message:= recover()
	// fmt.Println("ups Eror", message)
}

func main(){
	runApp(true)
	fmt.Println("Muhammad Ridho Febrian")
	// akan tetap di jalankan ketika terjadi panic karena recover
}