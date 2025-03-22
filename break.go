package main

import "fmt"

// contoh jika menggunakan break
func main(){
	for i := 0; i <=10; i++{
		if i == 5{
			break
		}
		// loop akan dihentika ketika mencapai 5 kali
		fmt.Println("perulangan ke", i)
	}
}