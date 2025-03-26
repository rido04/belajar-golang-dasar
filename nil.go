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

// contoh dengan menggunakan slice
func slice(condition bool) []int {
	if condition {
		return nil
	} else {
		return []int{1, 2, 3}
	}
}

// contoh menggunakan pointer
func pointer(condition bool) *int{
	if condition{
		return nil
	} else {
		value := 42
		return	&value
	}
}

// contoh menggunakan interface
func NewInterface(condition bool) any {
	if condition {
		return nil
	} else {
		return "Hello World!"
	}
}

// conoth menggunakan function
func NewFunction(condition bool) func(){
	if condition {
		return nil
	} else {
		return func() {
			fmt.Println("Haloo, ini dari function")
		}
	}
}

// contoh menggunakan channel
func NewChannel(condition bool) chan int{
	if condition {
		return nil
	} else {
		return make(chan int)
	}
}

func main(){
	// implementasi tipe data map
	data := NewMap("")
	if data == nil{
		fmt.Println("Data Kosong")
	} else {
		fmt.Println(data["name"])
	}

	// implementasi tipe data slice
	data2:= slice(true)
	if data2 == nil {
		fmt.Println("Data Slice Kosong")
	} else {
		fmt.Println(data2)
	}

	// implementasi tipe data pointer
	data3 := pointer(true) 
		if data3 == nil {
			fmt.Println("Pointernya kosong")
		} else {
			fmt.Println(*data3)
		}
	
	// implementasi tipe data channel
	data4 := NewChannel(true)
	if data4 == nil {
		fmt.Println("Channel kosong")
	} else {
		fmt.Println("channel dibuat")
	}

	// implementasi interface
	data5 := NewInterface(true)
	if data5 == nil {
		fmt.Println("Interface Kosong")
	} else {
		fmt.Println(data5)
	}

	// implementasi function
	data6 := NewFunction(true)
	if data6 == nil {
		fmt.Println("Function Kosong")
	} else {
		fmt.Print(data6)
	}
}