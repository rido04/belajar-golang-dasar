package main

import "fmt"

// interface adalah sebuah data abstrak, dia tidak bisa memiliki implementasi langsung
// biasanya berisikan definisi-definisi method
// biasanya digunakan sebahai kontrak
// kontrak artinya harus ada yang mengimplementasikannya

// buat dengan syntax interface
type HasName interface{
	// tentukan method apa saja yang harus dimiliki
	// di bawah ini adalah method, bukan field seperti yang ada di struct
	GetName() string
	GetAddress() string
}

func SayHello(value HasName){
	fmt.Println("Halo", value.GetName())
	fmt.Println("Alamat saya ada di", value.GetAddress())
}

// impelementasi interface
// setiap tipe data yang sesuai kontrak interface, secara otomatis dianggap sebagai interface tsb
// sehingga tidak perlu lagi implementasu interface secara manual 
// hal ini berbeda dengan bahasa pemrograman yang lain ketika membuat interface, kita harus menyebutkan secara eksplisit akan menggunakan interface dimana

type Person struct {
	// buat field nya sesuai dengan type interface yang sudah di buat
	Name,Address string
}

// jadi person di bawah ini adalah secara langsung implementasi dari interface
// gunakan func untuk inisialisasi method yang sudah di buat
// jadi buat sesuai kontrak interface yang sudah di buay, yaitu GetName dan GetAddress
func (person Person) GetName()string{
	return person.Name
}

// jadi semua implementasi dari hasName bisa kita kirim lewat parameter di atas
func (person Person) GetAddress()string{
	return person.Address
}

// adapun interface kosong
// biasanya di daam OOP ada satu data parent di puncak ynag bisa dianggap sebagai semua implementasi data yang ada di bahasa tersebut
// contoh ada di java.lang.object dan oop php
// dalam kasus ini, golang menggunakan interface kosong
// ini adalah interface yang ridak memiliki deklarasi method satupun, hal ini secara otomatis membuat semua tipe data akan menjadi implementasi nya
// interface kosong juga memiliki type alias bernama any


// contoh penggunaa interface kosong
// fmt.Println(a...interface{})
// panic(v...interface{})
// recover() interface{}
// dan lain lain
// dan sebenernya semua itu adalah interface kosong, artinya bisa dikirim data apapaun, benar benar bisa kirim data apapun bahkan sampai struct dan interface

// conoth interface kosong
func Ups() interface{}{
	// return 1
	// return true
	return "Ups"
	// bisa kembalikan nilai int, boolean, string, dlll. pokoknya bebas
}

// jika menggunakan any(sama aja, cuman lebih simpel)
func oops() any {
	return "Lagi belajar golang dasar nich, senggol dong"
}
func main(){
	person := Person{Name: "Ridho Febrian", Address: "Pakulonan Barat"}
	SayHello(person)
	kosong := Ups()
	fmt.Println(kosong)
	kosong2 := oops()
	fmt.Println(kosong2)
}