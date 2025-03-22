package main

import "fmt"

func main(){
	for i := 0; i <=10; i++{
		// contoh kalau misal i habis dibagi dua
		if i%2 == 0{
			continue
		}
		// ketika loop dari modulu 2 sudah habis di bagi, maka dia akan melanjutkan loop yang ada di post statement
		fmt.Println("perulangan ke", i)
		// maka output yang menhasilkan loop genap akan di alnjut atau di skip ke loop dengan hasil ganjil
	}
}