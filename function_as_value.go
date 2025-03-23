package main

import "fmt"

func getGoodbye(name string)string{
	return "Good Bye," +" "+ name
}
// kita bisa menyimpan sebuat function mejadi value di suatu variabel
func main(){
	// pastikan tidak menggunakan kurung jika ingin menyimpan function sebagai value pada variabel
	// karean ini ingin menyimpan, bukan memanggil function
	dadah := getGoodbye
	contoh := getGoodbye

	fmt.Println(dadah("Ridho"))
	fmt.Println(contoh("Febrian"))
}