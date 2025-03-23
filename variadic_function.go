package main

import "fmt"

// jika menggunakan variabel argumen
// mirip dengan slice, bedanya tidak perlu membuat slice lagi di bawah
// di tandau dengan titik tiga (...)
func sumAll(numbers ...int)int{
	total := 0
	for _, number := range numbers{
		total += number
	}

	return total
}

func main(){
	fmt.Println(sumAll(10,10,10))
	fmt.Println(sumAll(10,10,10,10,10))
	fmt.Println(sumAll(10,10,10,10,10,10,10))

	// jika sudah memiliki slice bisa gunakan cara seperti ini
	total := sumAll(10,10,10,10,10)
	fmt.Println(total)

	nomors := []int{10,10,10,10,10,10}
	// pastikan tambah titik tiga (...) jika mneggunakan slice
	total = sumAll(nomors...)
	// maka slice akan di konversi menjadi variabel argumen
	fmt.Println(total)
}