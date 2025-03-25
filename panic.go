package main

import "fmt"

// panic adalah function yang biasa digunakan untuk menghentikan program
// biasanya di panggil ketika terjadi panic pada saat program berjalan
// saat di panggil, program akan terhenti, namun defer function akan tetap di jalankan
func endApp(){
	fmt.Println("Aplikasi Berhenti")
}

func runApp(error bool){
	defer endApp()
	// gunakan syntax panic
	if error{
		// contoh jika terjadi error
		panic("Aplikasi Error")
	}
}

func main(){
	runApp(false)
	runApp(true)
}