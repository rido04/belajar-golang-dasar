package main

import "fmt"
func getCompleteName()(firstName, middleName, lastName string, age int){
	firstName = "Muhammad"
	// bisa juga di buat jadi string kosong
	middleName = "Ridho"
	lastName = "Febrian"
	age = 21
	return firstName, middleName ,lastName, age
}

func main(){
	a, b, c, d := getCompleteName()
	fmt.Println(a,b,c,d)
}