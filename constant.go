package main

import "fmt"

// variabel contstanta tidak bisa di isi ulang
// constanta tidak akan masalah atau terjadi error saat di run meskipun tidak di panggil atau di gunakan
// intinya tidak boleh di ubah lagi setelah di deklarasikan
func main(){
	const firstName string = "Muhammad"
	// bisa masukkan nilai langsung dan juga tipe datanya
	const middleName  = "Ridho"
	const lastName  = "Febrian"


	// Ini akan terjadi error jika di panggil ulang seperti di bawah
	// firstName = "Muhammad Ridho"
	// middleName = "Muhammad Ridho"
	// lastName = "Muhammad Ridho"


	// deklarasi multi constanta
	// bisa juga di isi langsung atau tidak
	const(  namaDepan string = "Muhammad"
			namaTengah = "Ridho"
	 		namaBelakang = "Febrian")
	fmt.Println(namaDepan)
	fmt.Println(namaTengah)
	fmt.Println(namaBelakang)
}