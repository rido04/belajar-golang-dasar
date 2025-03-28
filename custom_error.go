package main

import "fmt"

// MEMBUAT CUSTOM ERROR
// karena error adalah sebauh interface, kita bisa buat cutom error sendiri mengikuti kontrak atau ketentuan interface error itu sendiri

type validationError struct{
	Message string
}

type notFoundError struct{
	Pesan string
}

// masukkan ke variabel, disini saya pakai variabel v
// jadikan pointer dengan *
func(v *validationError) Error() string{
	return v.Message
}

// masukkan ke variabel, disini saya pakai variabel n
// jadika pointer dengan *
func (n *notFoundError) Error() string{
	return n.Pesan
}

// lalu kita gunakan
// contoh
func SaveData(id string, data any) error{ //balikan data error kalau terjadi error di akhir post statement
	if id ==""{
		return &validationError{Message: "Validasi Error"} //karena error interface, bisa balikan datanya dengan pointer tanda (&), karena pointer masuh mengikuti kontrak atau aturan dari si interface nya
	}

	if id != "Ridho" {
		return &notFoundError{Pesan: "User Tidak Ditemukan"}
	}

	// contoh kalau semuanya sukses saya kembalikan nil atau kosong
	return nil
}

func main(){
	// cara mengecek error nya
	err:= SaveData("", nil) //kita bisa konversi untuk mengecek error nya secara detail, apa yang error
	// bisa pakai switch case juga kalau mau, disini saya pakai if statement
	if err != nil{
		// masukkan detail error yang kita inginkan untuk mengecek error apa yang terjadi
		// bisa menggunakan konversi menggunakn type assertions
		//ok disini akan mengembalikan nilai boolean
		if validationErr, ok := err.(*validationError); ok { //konversi menggunakan type assertions
			// semisal errornya ini, akan mengembalikan pesan error di bawah berikut
			fmt.Println("Validasi Error", validationErr.Message)
		} else if notFoundErr, ok := err.(*notFoundError); ok{
			fmt.Println("Error not found", notFoundErr.Pesan)
		} else {
			fmt.Println("Erro tidak diketahui", err.Error())
		}

		// contoh pakai switch case
		// switch finalEror := err.(type){
		// case *validationError:
		// 	fmt.Println("Validasi Error", finalEror.Message)
		// case *notFoundError:
		// 	fmt.Println("Error not found", finalEror.Pesan)
		// default: fmt.Println("Error tidak diketahui", err.Error())
		// }
	} else {
		fmt.Print("Hore, tidak terjadi error")
	}
}