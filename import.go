package main

// IMPORT
// secara standar, file golang hanya bisa mengakses file golang lainnya yang berada di dalam package yang sama
// jika kita mau akses file golang yang berada di package yang berbeda, kita bisa pakai import
import (
	// pastikan direktori nya benar, bisa cek di file "go.mod"
	"belajar-golang-dasar/helper"
	"fmt"
)

func main(){
	// lalu panggil function yang ada di file helped tadi yang sudah di import
	result := helper.SayHello(" Muhammad Ridho Febrian")
	fmt.Println(result)
}