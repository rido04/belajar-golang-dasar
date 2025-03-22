package main

import "fmt"

// post statementny adalah data yang dikembalikan atau kembalian data
func getHello(name string)string{
	// ketika post statement untuk mengembalikan data di buat, maka harus mengembalikan data dengab return
	return "Hallo" + " " + name
}

// bisa juga mengembalikan multiple value
func getFullName()(string, string){
	return "Ridho", "Febrian"
}

func main(){
	result := getHello("Muhammad Ridho Febrian")
	fmt.Println(result)
	firstName, lastName := getFullName()
	fmt.Println(firstName, lastName)
}