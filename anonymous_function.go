package main

import (
	"fmt"
)

// anonyomus function adalah function yang langsung di buat di dalam variabel, tanpa membuat function itu sendiri

type BlackList func(string)bool
// kontraknya adalah function, param nya string, return nya boolean

func registerUser(name string, blackList BlackList){
	if blackList(name){
		fmt.Println("Kamu di blokir, ", name)
	} else {
		fmt.Println("Selamat Datang, ", name)
	}
}

func main(){
	// langsung buat anonymous fucntion di dalam variabel, tanpa di buat nama
	blackList:= func(name string)bool{
		return name == "anjing"
	}
	registerUser("Ridho", blackList)
	// bisa juga langsung di dalam values
	registerUser("anjing", func(name string)bool{
		return name == "anjing"
	})
	// initnya dia tidak punya nama function
}