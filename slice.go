package main

import "fmt"

func main() {
	names := [...]string{"John", "Paul", "George", "Ringo", "Ridho", "Jamal", "Kokom", "Riska"}
	// ini akan mengambil data array dari index ke 3 sampai ke 4
	slice1 := names[3:4]
	fmt.Println(slice1)

	// ini akan mengambil data array dari index paling awal sampai ke 4
	slice2 := names[:4]
	fmt.Println(slice2)

	// ini akan mengambil data array dari yang ke 5 sampai ke yang palin terakhir
	slice3 := names[5:]
	fmt.Println(slice3)

	// ini akan mengambil semua data di array
	slice4 := names[:]
	fmt.Println(slice4)

	// contoh jika di buat manual
	var slice5 []string = names[2:4]
	fmt.Println(slice5)

	// append slice
	days := [...]string{"Monday", "Tuesday", "Wednesday", "Thursday", "Friday", "Saturday", "Sunday"}
	daySlice1 := days[5:]
	fmt.Println(daySlice1)
	// akan menghasilkan saturday dan sunday dengan index 0 dan 1
	//  lalu ubah atau append data baru
	daySlice1[0] = "new Saturday"
	daySlice1[1] = "new Sunday"
	// ini kana mengubah data asli dalam array days atau array asli
	fmt.Println(daySlice1)
	fmt.Println(days)
	
	daySlice2 := append(daySlice1, "Libur")
	// bisa di asumsikan ketika append, dia akan menambah atau membuat array baru jika array asli sudah penuh
	// perumpaan nya seperti ini
	// daysBaru := [...]string{"Monday", "Tuesday", "Wednesday", "Thursday", "Friday", " new Saturday", "new Sunday", "Libur"}
	daySlice3 := append(daySlice2, "Hore")
	daySlice2[0] = "Old Saturday"
	fmt.Println(daySlice1)
	fmt.Println(daySlice2)
	fmt.Println(daySlice3)
	fmt.Println(days)


	// Make slice
	// buat dengan make, dengan parameter tipe data, panjangnya beraoa, dan kapasitas nya berapa
	newSlice := make([]string, 2, 5)
	// masukkan data array sesuai panjang yang telah ditentukan
	newSlice[0] = "Muhammad"
	newSlice[1] = "Ridho"

	fmt.Println(newSlice)
	fmt.Print(len(newSlice))
	fmt.Println(cap(newSlice))

	newSlice2 := append(newSlice, "Febrian")
	fmt.Println(newSlice2)
	// Ketika Di append, maka length array nya akan bertamabah
	fmt.Println(len(newSlice2))
	fmt.Println(cap(newSlice2))
	// Data Atau array yang di gunakan akan tetap memakai data asli
	// Kecuali kapasitas nya sudah tidak bisa menampung
	// otomatis akan terputus dari slice lama atau membuat data array yang baru
	newSlice2[0] = "Raden"
	fmt.Println(newSlice2)
	fmt.Println(newSlice)
	

	// Copy Slice

	fromSlice := days[:]
	// agar bisa di tampung semua datanya, perlu bikin slice
	toSlice := make([]string, len(fromSlice), cap(fromSlice))

	copy(toSlice, fromSlice)

	fmt.Println(fromSlice)
	fmt.Println(toSlice)

	fromSlice2 := newSlice2[:]
	toSlice2 := make([]string, len(fromSlice2), cap(fromSlice2))

	copy(toSlice2, fromSlice2)
	fmt.Println(fromSlice2)
	fmt.Println(toSlice2)

	fromSlice3 := daySlice2
	toSlice3 := make([]string, len(fromSlice3), cap(fromSlice3))

	copy(toSlice3, fromSlice3)
	fmt.Println(fromSlice3)
	fmt.Println(toSlice3)
	// Perlu diperhatikan, saat membuat slice, harusnya bikin array, malah bisa jadi bikin slice yang baru
	// [...] ini berarti array
	// [] ini berarti slice
	// bedanya ketika buat slice, itu tidak ada jumlah datanya

	iniArray := [...]int{1,2,3,4,5}
	iniSlice := []int{1,2,3,4,5}

	fmt.Println(iniArray)
	fmt.Println(iniSlice)

}

// function yang ada di slice
// len() => untuk mengetahui panjang dari slice pointer, bukan panjang dari array
// cap() => untuk mengetahui kapasitas dari slice pointer
// append() => untuk menambahkan data ke dalam slice pointer
// make() => untuk membuat slice baru
// copy() => untuk menyalin data dari slice ke slice lain
// new() => untuk membuat pointer baru 
// delete() => untuk menghapus data dari slice