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
}