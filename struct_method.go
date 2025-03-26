package main

import "fmt"

// jika kita ingin menambahkan method ke dalam struct, sehingga seakaan struct memiliki function
// method adalah function yang menempel di dalam si struct

type Pelanggan struct{
	Name,Address string
	Age	int
}

// sebelum nama function kita isi parameter struct nya
// contoh
func(pelanggan Pelanggan) sayHello(name string){
	fmt.Println("Hallo", name, "nama saya", pelanggan.Name)
}
// method say hello tsb hanya bisa di akses ketika kita sudah membuat objek dari si pelanggan

func main(){
	ridho := Pelanggan{Name: "Ridho"}
	// lalu panggil method nya
	ridho.sayHello("Ujang")
}