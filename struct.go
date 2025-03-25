package main

import "fmt"

/*struct adalah sebuah template data yang di gunakan
untuk menggabungkan nol atau lebih tipe data
lainnya dalam satu kesatuan*/
// biasanya representasi data dalam program aplikasi yang kita buat
// data struct di simpan dalam field
// sederhananya struct adalah kumpulan dari field

// contoh
// buat dengan menggunakan type
// biasanya nama sebuah struct di awali huruf besar
type Customer struct{
	// lalau definisikan struct customer ini punya field atau atribut apa aja
	// biasanya field atau atribut di struct di tulis dalam huruf besar atau kapital di awal huruf(pascal cage)
	Name, Address string
	Age int
}
// struct adalah template data atau prototype data
// struct tidak bisa langsung digunakan
// namun kita bisa membuat object/data dari struct yang telah kita buat

func main(){
	var ridho Customer
	// ketika membuat var, akan di buat objek ridho dari struct cutomer
	// jika data nya belum di buat atau di isi, akan mengembalikan default valunya yaitu string kosong {0}
	fmt.Println(ridho)
	// cara membuat data dari struct
	// nama variabelnya, lalu titik nama fieldnya
	ridho.Name = "Muhammad Ridho Febrian"
	ridho.Address = "Pakulonan Barat"
	ridho.Age = 21
	fmt.Println(ridho)
	// bisa juga akses datanya satu persatu dengan
	fmt.Println(ridho.Name)
	fmt.Println(ridho.Address)
	fmt.Println(ridho.Age)

	// adapaun struct literal
	// contoh pertama
	sandi := Customer{
		Name : "Sandi Arba Putra",
		Address : "Tanah tinggi",
		Age : 20,
	}
	fmt.Println(sandi)
	
	// contoh ke dua struct literal
	// harus sesuai dengan urutannya kalau cara kedua ini
	dhaniel := Customer{"Dhaniel", "Tigaraksa", 25}
	fmt.Println(dhaniel)
}