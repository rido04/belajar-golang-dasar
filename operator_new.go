package main

import "fmt"

// OPERATOR NEW
// Sebelumnya jika ingin membuat pointer dengan menggunakan oprator &
// golang juga punya function new yang bisa di pakai untuk bikin pointer
// function new hanya return pointer ke data kosong, artinya tidak ada data awal

type Home struct{
	Warna string
	Nomor int
}
func main(){
	// buat dengan menggunakan syntax new, untuk mengarahkan ke value Home
	// kurang lebih sama kayak pointer dengan menggunakan syntax &
	rumah1 := new(Home)
	rumah2 := rumah1

	rumah2.Warna = "Hjjau"
	rumah2.Nomor = 20

	fmt.Println(rumah1)
	fmt.Println(rumah2)
}