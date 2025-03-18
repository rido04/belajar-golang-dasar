package main

import "fmt"

func main() {
	// masukkan jumlah array maksimal ketika membuat variabel
	var names [3]string
	names[0] = "Muhammad"
	names[1] = "Ridho"
	names[2] = "Febrian"

	fmt.Println(names[0])
	fmt.Println(names[1])
	fmt.Println(names[2])

	// buat array secara langsung ketika membuat varuabel
	var values = [3]int{
		90,
		95,
		80,
	}

	fmt.Println(values)
	// output: [90 95 80]
	fmt.Println(values[0])
	// output: 90
	fmt.Println(values[1])
	// output: 95
	fmt.Println(values[2])
	// output: 80


	var values2 = [...]int{
		90,
		95,
		80,
	}
	// output: [90 95 80]
	fmt.Println(values2)
	fmt.Println(len(values2))
	values2[0] = 100
	fmt.Println(values2)
	// output: [100 95 80]
}

// array akan terjadi error ketika ada array baru yang di masukkan melebihi jumlah array yang di tentukan
// jika ingin menambahkan array baru, maka harus menambahkan jumlah array yang di tentukan

// di array juga bisa di isi langsung ketika membuat variabel
// ketika array yang di isi kurang dari jumlah array yang di tentukan, maka array yang tidak di isi akan di isi dengan nilai default (0)
// kalau string akan di isi dengan string kosong
// jika array yang di isi lebih dari jumlah array yang di tentukan, maka akan terjadi error
// jika tidak mau menentukan jumlah array, maka bisa menggunakan [...] dan tidak perlu menentukan jumlah array
// dan hanya bisa jika array langsung di isi nilainya ketika membuat variabel


// Adapaun Array Function
// len(array) => untuk mengetahui jumlah array
// array[index] => untuk mengakses array berdasarkan index
// array[index] = value => untuk mengubah nilai array berdasarkan index


// DI ARRAY TIDAK BISA MENGHAPUS DATA YANG ADA DI DALAMNYA