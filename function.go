package main

import "fmt"

func sayHello(){
	fmt.Println("Hello World")
	// function yang telah di buat boleh di panggil lebih dari satu kali
}

func main(){
	sayHello()
	sayHello()
}