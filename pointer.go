package main

import "fmt"

// PASS BY VALUE
// secara default, di golang semua variabel itu di passing by value, bukan reference
// artinya jika mengirim sebuah variabel ke dalam function, method atau variabel lain, sebenarnya yang dikirim adlaah duplikasi valuenya

// POINTER
// Adalah kemampuan membuat reference ke lokasi data di memory yang sama, tanpa mengcopy data yang sudah ada
// maksudnya, ini adalah pass by reference, jadi dia melakukan referensi, bukan duplikasi
// untuk membuatnya bisa menggunakan OPERATOR & di ikuti nama variabelnya
type Address struct{
	City, Province, Country string
}

type Orang struct {
	Nama, Alamat, Tanggal string
}

func main(){
	address1 := Address{"Subang", "Jakarta", "Tangerang"}
	// di sini seakan address 2 akan mengcopy data address 1, jika di bahasa pemrograman lain biasanya itu melakukan reference, di golang berbeda
	address2 := address1

	// lalu coba ubah address 2 citynya
	address2.City = "Bandung"

	// address 1 tidak akan berubah isi datanya
	fmt.Println(address1)
	// address 2 berubah, jadi seakan di duplikasi, lalu di ubah datanya
	// jadi perubahan yang terjadi di address 2 tidak berdampak ke address 1
	fmt.Println(address2)


	// Contoh jika menggunakan pointer
	orang1 := Orang{"Ridho Febrian", "Tangerang", "14 Februari"}
	// gunakan syntax/operator & untuk menggunakan pointer
	orang2 := &orang1
	// maka data dari variabel orang1 dan orang2 akan sama, karena bukan diduplikasi, melainkan pass by reference
	fmt.Println(orang1)
	fmt.Println(orang2)
	// contoh jika kita buat lagi
	orang3:= &orang1
	fmt.Println(orang3)
	// lalu ubah lagi isi datanya
	orang3.Nama = "Ujang"
	
	// jika isi datanya di ubah, maka dua variabel tersebut akan terkena dampak perubahan data nya
	// jadi aslinya dia merubah orang1, karena melakukan pointer
	orang2.Alamat = "Jakarta"
	fmt.Println(orang1)
	fmt.Println(orang3)
	fmt.Println(orang2)
	// data yang berubah adalah data terbaru yang di ubah
}