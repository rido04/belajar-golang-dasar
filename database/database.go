package database

// PACKAGE INITIALIZATION
// sata kita membuat packagae, kita bisa membuat sebuah function yang akan diakses ketika package init diakses
// sanagat cocok ketika : package berisi function untuk berkomunikasi dengan database, kita membuat functon inisialisasi untuk membuka koneksi ke database
// untuk membuat function yang diakses secara otomatis ketika package diakses, cukup dengan membuat function dengan nama init

// contoh
var connection string

// function init akan otomastis dieksekusi ketiak package nya di panggil
func init() {
	connection = "MySQL"
}

// lalu buat function yang bisa diakses oleh package lain, karena dalam kasus ini adalah koneksi database
func GetDatabase() string {
	return connection
}
