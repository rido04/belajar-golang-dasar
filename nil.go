package main

import "fmt"

// biasanya dalam bahasa pemrograman lain, object yang belum di inisiasikan maka secara otomatis nialinya null atau nil
// kalau di golang beda, karena secara otomatis di buat nilai defaultnya
// tapi di golang ada nil atau data kosong
// nil ini hanya bisa di gunakan di beberapa tipe data. seperti interface, function, map, slice, pointer, dan channel

// contoh dengan menggunakan tipe data map
func NewMap(name string) map[string]string{
	if name == ""{
		return nil
	} else {
		return map[string]string{
			"name" : name,
		}
	}
}

func main(){
	data := NewMap("")
	if data == nil{
		fmt.Println("Data Kosong")
	} else {
		fmt.Println(data["name"])
	}
}