package main

import "fmt"

// function dengan mengirim parameter
func sayHelloTo(firstName string, lastName string){
	fmt.Println("Haloo", firstName, lastName)
}

func kabar(kabar string, kondisi string){
	fmt.Println("kabarku", kabar, kondisi)
}

func nomor(nomor int){
	fmt.Println("Nomor", nomor)
}

func main(){
	// tipe data yang di kirim harus sama dengan yang di buat di awal fucntion
	sayHelloTo("Ridho", "Febrian")
	sayHelloTo("Joko", "Anwar")
	// bisa juga di panggil lebih dari satu kali dengan mengirim data yang berbeda
	kabar("Alhamdulillah", "Sehat")
	nomor(10)
}