package main

import "fmt"

func main() {

	// Bagaimana kalau kita belum tahu mau mengisi sebuah data ke suatau aray yang jumlahnya itu ada berapa

	// maksudnya ada sebuah Array yang kita tidak tau jumlahnya, dan kita ingin memasukkan sebuah data kedalam array tersebut

	// artinya kita hanya bisa mengisi 5 data ke array
	// var gamingConsole [5]string

	// nah kalau kita belum tahu panjang arraynya, maka kita bisa menambahkan data yang kurang atau lebih dari kapasitas array

	// kita bisa pake slice (kita gak butuh angka)

	// ini artinya kita buat sebuah slice
	var gamingConsole []string

	// untuk menambahkan data kedalam slice menggunakan keyword append
	gamingConsole = append(gamingConsole, "PlayStation4")

	// inisialisasi slice langsung
	city := []string{"Kediri", "Malang", "Surabaya", "Jakarta", "Bandung"}

	fmt.Println(city)
	fmt.Println(gamingConsole)

}
