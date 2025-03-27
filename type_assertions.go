package main

import "fmt"

// type assertions adalah kemampuan merubah tupe data menjadi tipe data yang kita inginkan
// fitur ini seringkali digunakan ketika kita bertemu dengan tipe data interface kosong

// contoh
func random() interface{} {
	return "OK"
}
func main(){
	// ketika kita buat result yang di ambil dari random, maka hasil nya adalah tipe data interface kosong
	result := random()
	// contoh koversi menggunakan type assertions ke string
	resultString := result.(string)
	fmt.Println(resultString)
	
	// jika sudah melakukakn konversi lalu membuat konversi lagi, maka akan terjadi panic
	// resultInt := result.(int)
	// fmt.Println(resultInt)

	// NOTE: Pastikam tipe data yang akan di konversi sesuai dengan tipe data yang di return di atas

	// Adapun type assertion menggunakan switch
	// menggunakan switch lebih direkomendasikan
	result2 := random()
	// lalu ambil tyoe nya apa di dalam kurung, dia akan di simpan di variabel value
	switch value := result2.(type) {
	case string:
		fmt.Println("String", value)
	case int :
		fmt.Println("Int", value)
	default :
		fmt.Println("Unknown")
	 }
}