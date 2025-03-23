package main

import "fmt"

// di sini saya menggunakan type declaration
// dimana function yang berfungsi sebagai parameter akan di buat alias supaya lebih efisien
type Filter func(string) string

func sayHelloWithFilter(name string, filter Filter) {
	filteredName := filter(name)
	fmt.Println("Hello, ", filteredName)
}

func spamFilter(name string) string{
	if name == "Anying"{
		return "Lagi puasa ga boleh kasar!"
	} else {
		return name
	}
}

func main(){
	// bisa di panggil langsuung ke dalam parameter
	// dalam kasus ini adalah spamFilter function
	sayHelloWithFilter("Ridho", spamFilter)
	
	// atau masukkan function nya ke dalam variable
	// dalam kasus ini adalah spamFilter function
	filter := spamFilter
	sayHelloWithFilter("Anying", filter)
}