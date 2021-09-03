package main

import "fmt"

func main() {

	// switch juga untuk percabangan
	// digunakan saat kita ingin mencocokan sebuah variabel

	number := 2

	switch number {
	case 1:
		fmt.Println("Ini Satu")
	case 2:
		fmt.Println("Ini Dua")
	case 3:
		fmt.Println("Ini Tiga")
	default:
		fmt.Println("Angka Salah")
	}

	// kalau di ubah menjadi IF Control maka akan seperti dibawah ini
	// Dalam kasus ini penggunaan SWITCH lebih sederhana daripada IF ELSE
	if number == 2 {
		fmt.Println("Ini Dua")
	} else if number == 1 {
		fmt.Println("Ini Satu")
	} else if number == 3 {
		fmt.Println("Ini Tiga")
	} else {
		fmt.Println("Angka Salah")
	}
}
