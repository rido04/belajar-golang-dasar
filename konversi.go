package main

import "fmt"

// Konversi tipe data
func main() {
	var nilai32 = 32768
	// konversi nilai yang ada di variabel sebelumnya ke tipe data int64
	var nilai64 = int64(nilai32)
	// Perlu Diperhatikan jika nilai yang dikonversi melebihi batas tipe data maka akan terjadi overflow (Contoh = Nilai dari int 32 melebihi batas int 16)
	var nilai16 = int16(nilai32)
	var nilai8 = int8(nilai32)


	// jika nilai masih bisa di terima saat du konversi maka akan tetap sama nilainya

	// jika nilai melebihi batas maka akan terjadi overflow, oveverflow adalah balik lagi ke belakang, dari yang paling tinggi balik ke yang paling rendah (Contoh : int 16 melebihi batas dari 32767, maka akan balik ke -32768)
	fmt.Println(nilai32)
	fmt.Println(nilai64)
	fmt.Println(nilai16)
	fmt.Println(nilai8)


	// Konversi byte ke string
	// byte adalah tipe data yang sama dengan uint8
	var nama = "Muhammad Ridho Febrian"
	// mengambil byte atau index pertama dari variabel nama
	var r = nama[0]
	// gunakan string untuk mengkonversi byte ke string setelah memanggil nama[0]
	var rString = string(r)

	fmt.Println(nama)
	fmt.Println(r)
	fmt.Println(rString)
}

// maks tipe data 
// int 8 = -128 sampai 127
// int 16 = -32768 sampai 32767
// int 32 = -2147483648 sampai 2147483647
// int 64 = -9223372036854775808 sampai 9223372036854775807

