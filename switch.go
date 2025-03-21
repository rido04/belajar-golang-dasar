package main

import (
	"fmt"
)

func main(){
	nama := "Ridho"

	// switch biasa di gunakan untuk mengecek 1 variable
	switch nama{
	case "Ridho":
		fmt.Println("Yoman Bro ",nama, ", balik lagi lu")
	case "Bagas":
		fmt.Println("Lu Siapa Bjir")
	// default bisa di bilang else punya switch case
	default:
		fmt.Println("Lah ini Lagi Siapa")
	}

	// bisa juga menggunakan short statement

	switch length := len(nama); length>5{
	case true:
		fmt.Println("Kepanjangan brader namanya")
	case false:
		fmt.Println("Udah pas nih namanya")
	}

	// switch tanpa expression
	length := len(nama)
	switch{
	case length > 10:
		fmt.Println("Nama sudah panjang")
	case length > 5:
		fmt.Println("Nama lumayan panjang")
	default:
		fmt.Println("Nama terlalu pendek")
	}
}