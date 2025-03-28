package main

import (
	"belajar-golang-dasar/database"   //otomatis akan langsung menjalankan function init, yaitu koneksi ke database
	_ "belajar-golang-dasar/internal" //contoh kalau tidak di gunakan maka akan terjadi error

	// bisa gunakan blank identifier ketika ingin tetap memanggil func init yang ada di package tersebut, tanpa menggunakan function itu sendiri (_)
	"fmt"
)

func main(){
	fmt.Println(database.GetDatabase())
}