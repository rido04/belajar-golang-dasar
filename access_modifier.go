package main

// ACCESS MODIFIER
// di golang, untuk menentukan access modifier, cukup dengan nama function atau varuiabel
// jika nama diawali dengan huruf besar, maka artinya bisa di akses di package lain, jika di mulai dengan huruf kecil, artinya tidak bisa diakses di package lain
// namun jika menggunakan huruf kecik di awal tetap bisa di akses oleh package yang sama

import (
	"belajar-golang-dasar/helper"
	"fmt"
)
func main(){
	// akses packahe lain yang ada di helper 
	fmt.Println(helper.Application) //huruf depannya besar

}