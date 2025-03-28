package main

import "fmt"

// POINTER DI METHOD
// walaupun method akan menempel di struct, sebenernya data struct yang di akses di method adalah pass by value
// sangat di rekomendasikan menggunakan pointer di method, sehingga tidak boros memory karena harus selalu diduplikasi ketika memanggil method

type Man struct{
	Name string
}

// contoh method yang belum menggunakan pointer
// karena di sini parameter nya adalah value, jadi dia akan di duplikat
func (man Man) Married(){
	man.Name = "Mr." + man.Name
}

// contoh yang menggunakan pointer
// gunakan operator * untuk menentukan apakah itu pointer
func (man *Man) Menikah(){
	man.Name = "Tuan " + man.Name
}
func main() {
	// contoh impelementasi yang belum menggunakan pointer
	ridho := Man{"Ridho Febrian"}
	ridho.Married()
	fmt.Println(ridho.Name)

	// contoh implpementasi yang menggunakan pointer
	ridho2 := Man{"Ridho Febrian"}
	ridho2.Menikah()
	fmt.Println(ridho2.Name)
}