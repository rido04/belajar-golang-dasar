package main

import "fmt"

// POINTER DI FUNCTION
// saat membuat parameter di function, defaultnya adalah pass by value, artinya adalah data akan di copy lalu di kirim ke function tsb
// makanya, jika kita mengubah data di dalam function, data yang aslinya tidak akan berubah karena di duplikasi
// maka variabel jadi aman karena data nya tidak berubah
// namun kadang kita penegen membuat function yang mengubah data di parameter
// kalau mau kayak gini, bisa pakai pointer function
// untuk menjadikan sebuah parameter sebagai pointer, kit bisa pakai operator * di parameternya

type Animal struct{
	Name, Herbivore, Carnivore string
}

// contoh yang belum menggunakan pointer function
func ChangeAnimalStatus(animal Animal) {
	animal.Carnivore = "Karnvora"
}

// contoh menggunakan pointer
func AnimalStatus(animal *Animal){
	animal.Carnivore = ", Tapi dia Karnivora"
}


func main(){
	animal := Animal{"Hawimau", "bukan Herbivora", "" }
	// data yang dikirim ke function di bawah ini tidak berubah ke indonesia, karena datanya akan di copy, jadi belum memakai pointer
	// karena yang dikirim adalah duplikat dari si valuenya
	ChangeAnimalStatus(animal)
	// TIDAK BERUBAH
	fmt.Println(animal)
	// BERUBAH
	// karena sudah menggunakan pounter di function
	// gunakan operator dan (&) karena pointer
	AnimalStatus(&animal)
	fmt.Println(animal)
}