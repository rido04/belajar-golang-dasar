package helper

// nama package harus disesuaikan dengan nama fodler

// PACKAGE
// adalah tempat yang bisa digunakan untuk mengatur kode program yang kita buat di golang
// dengan menggunakan package, kita bisa lebih merapihkan kode yang kita buat
// package sendiri sebenernya cuma direktori folder di sistem operasi kita
// contoh untuk package main
// berarti dia kode yang harus di jalankan karena MAIN

func SayHello(name string)string{
	// pastkan huruf depan pada functionnya adalah huruf besa, karena kita ingin akses function ini di package yang berbeda
	return "Hello" + name
}

var version = "1.0.0" //ini tidak bisa diakses di file lain, karena awal hurufnya huruf kecil
var Application = "Golang" //ini bisa di akses di file atau package lain, karena awalannya adalah huruf besar

func sayGoodbye(name string)string{
	return "Goodbye " + name
}

