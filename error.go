package main

// ERROR
// selain panic dan recover, golang juga punya error interface lain, namanya adalah error

import (
	"errors"
	"fmt"
)

// aslinya golang udah punya error ini secara otomasti
// jadi golang punya libraryr error nya sendiri untuk membuat helper secara mudah, yang ada di "package errors"

func Pembagian(nilai int, pembagi int)(int, error) { //biasanya mengembalikan dua data, yaitu hasilnya, dan data kalau terjadi error, karena error itu interface
	if pembagi == 0{
		return 0, errors.New("Pembagian dengan NOL")
	} else {
		return nilai / pembagi, nil //error bisa dikembalikan dalam bentuk nil, karena error itu interface
	} //jadi kalau kita akses datanya, dia bisa mengembalikan data hasil, atauan data errornya
}

// contoh penggunaan error
func main(){
	hasil, err := Pembagian(100, 10)
	if err == nil{
		fmt.Println("Hasil", hasil)
	} else {
		fmt.Println("Error", err.Error()) //jadi error nya adalah error titik si method errornya (Error())
	}
}

// jadi kesimpulannya, cara menggunakan error di golang cukup menggunakan interface error pada saata membuat function
// lalu cara menggunakannya cukup dengan menggunakan errors.New()