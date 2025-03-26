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

func (person Person) GetAddress()string{
	return person.Address
}

// jadi semua implementasi dari hasName bisa kita kirim lewat parameter di atas
func main(){
	person := Person{Name: "Ridho Febrian", Address: "Pakulonan Barat"}
	SayHello(person)
}