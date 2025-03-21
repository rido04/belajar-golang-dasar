package main

import "fmt"

func main(){
	person := map[string]string {
		"nama" : "Muhammad Ridho Febrian",
		"kelas" : "D Informatika",
	}
	// contoh map kosong dengan cara lain
	var orang map[string]string = map[string]string{}
	orang["nama"] = "Ridho"
	orang["alamat"] = "pakbar"

	fmt.Println(person)
	fmt.Println(person["nama"])
	fmt.Println(person["kelas"])
	fmt.Println(orang)
	fmt.Println(orang["nama"])
	fmt.Println(orang["alamat"])
	// jika mencoba akses data yang tidak ada maka akan menggunakan default value tergantung tipe data apa yang di gunakan(di sini string, jadi akan menggunakan string kosong)

	// bisa juga menggunakan make
	book := make(map[string]string)
	book["judul"] = "belajar golang"
	book["penulis"] = "Ridho"
	book["tahun terbit"] = "2025"
	book["harga"] = "Rp.200.000"
	
	fmt.Println(book)

	//contoh menggunakan delete, dengan isi string atau data mana yang mau di hapus
	delete(book, "harga")
	fmt.Println(book)

	// menggunakan len untuk menghitung data
	fmt.Println(len(book))

	// menggunakan map[] untuk mengambil data spesifik
	fmt.Println(book["judul"])

	// menggunakan map[] = value untuk mengubah data 
	book["penulis"] = "Muhammad Ridho Febrian"
	fmt.Println(book)
	
}

