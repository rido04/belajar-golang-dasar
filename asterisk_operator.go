package main

import "fmt"

// ASTERISK OPERATOR / OPERATOR *
// saat kita mengubah variabel pointer, maka yang berubah hanya variabel tsb
// semua variable yang mengacu ke data yang sama tidak akan berubah
// jika kita ingin mengubah seluruh variabel yang mengacu ke data tsb, kita bisa menggunakan operator *

type Jalan struct{
	Gang string
	Nomor int
}
func main(){
	jalan1 := Jalan{"H. Abdullah", 3}
	// di sini jalan2 masih mengacu ke jalan1
	jalan2 := &jalan1

	jalan2.Gang = "Abdul Ghani"
	fmt.Println(jalan1)
	fmt.Println(jalan2)

	// karena jalan 2 adalah pointer ke jalan1, jadi gunakan tanda & untuk membuat yang baru
	// jadi sekarang jalan2 sudah mengacu ke pointer yang baru, namun jalan1 tetap mengacu ke pointer yang lama

	jalan2 = &Jalan{"H. Sanusi", 4}
	fmt.Println(jalan1)
	fmt.Println(jalan2)
	// contoh menggunakan operator asterisk *
	*jalan2 = Jalan{"Gg. H. Sanusi", 10}
	// jalan 1 juga kaan berubah, karena yang di ubah adalah valuenya
	// jadi kasarnya, variabel jalan1 akan ikut isi data dari pointer baru variabel jalan2
	fmt.Println(jalan1)
	fmt.Println(jalan2)
	// OUTPUT :
	/*
	{Abdul Ghani 3}
	&{Abdul Ghani 3}
	{Gg. H. Sanusi 10}
	&{Gg. H. Sanusi 10}
	*/
}