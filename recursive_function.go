package main

import "fmt"

// recursive function adalah function yang memanggil dirinya sendiri
// contoh kasus nya adalah pada factorial
// misal punya nilai 10, lalu turun ke 9 (10, 9, 8, 7, dst)

// contoh ketika menggunakan loop

func factorialLoop(value int)int{
	result := 1
	for i := value; i >0; i--{
		result *= i
	}
	return result
}

// contoh menggunakan recursive function
func factorialRecursive(value int)int{
	if value == 1{
		return 1
	} else {
		return value * factorialRecursive(value-1)
	}
}

func recursive(hasil int) int{
	if hasil == 1{
		return 1
		// kalau hasil atau value yang di masukkan bukan 1, maka kalikan, lalu kurangi nilai yang akan di kalikan sebanyak 1, sampai hasil atau value kembali atau menghasilkan nilai 1
	} else {
		return hasil * recursive(hasil-1)
	}
}

func main(){
	fmt.Println(factorialLoop(10))
	fmt.Println(factorialRecursive(10))
}