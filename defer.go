package main

import "fmt"

// defer function adalah function yang bisa kita jadwalkan untuk dieksekusi setelah sebuah function selesai di eksekusi
// defer function akan selalu diekskusi walaupun terjadi error di function yang di eksekusi
func logging(){
	fmt.Print("Memanggil Logging")
}
 func runApplication(){
	// gunakan syntax defer
	defer logging()
	// maka function yang akan di jalankan duluan adalah runApplication
	fmt.Println("Aplikasi Berjalan")
 }

//  note : pastika defer di jalankan d function urutan terakhir agar bekerja
 func main(){
	runApplication()
 }