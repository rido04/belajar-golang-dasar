package main

func main() {
	var nilaiAkhir = 100
	var absensi = 80

	var lulusNilaiAkhir = nilaiAkhir > 80
	var lulusAbsensi = absensi > 80

	var lulus = lulusNilaiAkhir && lulusAbsensi
	println(lulus)

	var lulus2 = lulusNilaiAkhir || lulusAbsensi
	println(lulus2)

	var lulus3 = !lulusNilaiAkhir
	println(lulus3)

}

// ada beberapa operasi boolean yang bisa digunakan di golang

// && (and) Dan
// operasi ini akan mengembalikan nilai true jika kedua nilai yang diberikan bernilai true
// jika salah satu bernilai false, maka hasilnya akan false

// || (or) Atau
// operasi ini akan mengembalikan nilai true jika salah satu nilai yang diberikan bernilai true
// jika kedua nilai bernilai false, maka hasilnya akan false

// ! (not) Kebalikan dari nilai boolean yang diberikan
// operasi ini akan mengembalikan nilai true jika nilai yang diberikan bernilai false
// dan sebaliknya
