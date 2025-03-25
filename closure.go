package main

import "fmt"

// closure adalah kemampuan sebuah function berinteraksi dengan data disekitarnya dalam scope yang sama
// pastikan gunakan closure dengan bijak karena kalau terlalu banyak menggunakan closure akan membingungkan kode program kita

// contoh
func main(){
	counter := 0
	// lalu kita buat anonymous function di dalam variabel increment
	increment := func ()  {
		fmt.Println("Increment")
		// lalu panggil variabel counter yang ada di atas, lalu naikkan atau tambah counternya
		counter++
	}
	// lalu kita panggil variabel increment yang berisi function untuk menambah variabel counter tsb
	increment()
	fmt.Println(counter)
	increment()
	// lalu prin variabel counter yang sudah di tambah
	fmt.Println(counter)
}