package main

import "fmt"

func main(){
	nomor := 1

	for nomor <= 10{
		fmt.Println("Nomor Ke:", nomor)
		nomor ++
		// nomor ++ artinya naik atau nambah 1
	}
	fmt.Println("Selesai")

	// bisa juga menambahkan statement
	// ada init statement, yaitu statement sebelum for di eksekusi
	for number := 1; number <= 10; number ++{
		fmt.Println("in Numnber :", number)
	}
	fmt.Println("Done")
	// ada post statement, yaitu statement yang akan selalu di eksekusi di akhir tiap perulangan



	// ADA JUGA FOR RANGE
	// Contoh Manual 👇
	names := []string{"Muhammad", "Ridho", "Febrian"}
	for i := 0; i < len(names); i++{
		fmt.Println(names[i])
	}

	// ini contoh fo range, dengan output index
	for index, name := range names{
		fmt.Println("index", index, "=", name)
	}

	// contoh jika tidak perlu index
	for _,name := range names{
		fmt.Println(name)
	}
}